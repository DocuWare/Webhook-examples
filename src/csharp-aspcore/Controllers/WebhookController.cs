using System;
using System.Text.Json;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace csharp.Controllers
{
    public class WebhookController : Controller
    {
        [HttpPost]
        public async Task<IActionResult> Index()
        {
            Console.WriteLine("The request contains the following body:");

            try
            {
                JsonDocument jsonDocument = await JsonDocument.ParseAsync(Request.Body);

                JsonSerializerOptions jsonSerializerOptions = new JsonSerializerOptions()
                {
                    WriteIndented = true
                };

                Console.WriteLine(JsonSerializer.Serialize(jsonDocument, jsonSerializerOptions));
            }
            catch (Exception e)
            {
                Console.WriteLine("Can't handle the json format!");
                Console.WriteLine(e);
            }

            return Accepted();
        }
    }
}