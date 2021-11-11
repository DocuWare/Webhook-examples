# C# ASP .Net Core webhook example
This example is done with ASP .Net Core 3.1 and written in C#.

## Prerequisites
Be familiar with C# and ASP .Net Core:
- [C# documentation](https://docs.microsoft.com/en-us/dotnet/csharp/)
- [ASP .Net Core](https://docs.microsoft.com/en-us/aspnet/core/introduction-to-aspnet-core?view=aspnetcore-3.1)

## How to use the project
1. Download and install [Visual Studio](https://visualstudio.microsoft.com/vs/) or [Visual Studio Code](https://code.visualstudio.com/).
2. Open the solution in [Visual Studio](https://visualstudio.microsoft.com/vs/) or [Visual Studio Code](https://code.visualstudio.com/).
3. Start it in Debug mode or Run it.
4. Send a POST request to http://localhost:5000/webhook with a JSON body e.g. { "Test": "test-value" }.
5. See the console output of the body.