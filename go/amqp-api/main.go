package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/streadway/amqp"
)

func main() {
	serverUrl := os.Getenv("SERVER_URL")
	queueName := os.Getenv("QUEUE_NAME")

	amqpConnection, err := amqp.Dial(serverUrl)
	if err != nil {
		panic(err)
	}
	defer amqpConnection.Close()

	amqpChannel, err := amqpConnection.Channel()
	if err != nil {
		panic(err)
	}
	defer amqpChannel.Close()

	_, err = amqpChannel.QueueDeclare(
		queueName, true, false, false, false, nil,
	)
	if err != nil {
		panic(err)
	}

	//fiber app setup
	app := fiber.New()

	app.Use(
		logger.New(),
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "usage: curl -X POST --url 'http://ADDRESS/send?message=hello'",
		})
	})

	app.Post("/send", func(c *fiber.Ctx) error {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("message")),
		}

		if err := amqpChannel.Publish(
			"", queueName, false, false, message,
		); err != nil {
			return err
		}

		return nil
	})

	log.Fatal(app.Listen(":2500"))
}
