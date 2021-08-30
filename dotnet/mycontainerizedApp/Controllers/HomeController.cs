using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Logging;
using mycontainerizedApp.Models;

namespace mycontainerizedApp.Controllers
{
    public class HomeController : Controller
    {
        private readonly ILogger<HomeController> _logger;
        private readonly IConfiguration configuration;

        public HomeController(ILogger<HomeController> logger, IConfiguration configuration)
        {
            _logger = logger;
            this.configuration = configuration;
        }

        public IActionResult Index()
        {
            
            ViewData["NodeName"] = configuration["NODE_NAME"];
            ViewData["Name"] = configuration["POD_NAME"];
            ViewData["Namespace"] = configuration["POD_NAMESPACE"];
            ViewData["PodIp"] = configuration["POD_IP"];
            ViewData["ServiceAccount"] = configuration["POD_SERVICE_ACCOUNT"];

            _logger.LogInformation("Home/Indes: ENV variables read");

            return View();
        }

        public IActionResult Privacy()
        {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}
