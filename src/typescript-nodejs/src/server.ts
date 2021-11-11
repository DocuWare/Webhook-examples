import express from "express";
import * as http from "http";

import * as winston from "winston";
import * as expressWinston from "express-winston";
import cors from "cors";
import debug from "debug";

const app: express.Application = express();
const server: http.Server = http.createServer(app);
const port = 5000;
const logServerStatus: debug.Debugger = debug("server");

// Parse all incoming requests as JSON.
app.use(express.json());

// Enable CORS.
app.use(cors());

// Configure logger options.
const loggerOptions: expressWinston.LoggerOptions = {
    transports: [new winston.transports.Console()],
    format: winston.format.combine(
        winston.format.json(),
        winston.format.prettyPrint(),
        winston.format.colorize({ all: true})
    ),
    requestWhitelist: [
        "headers",
        "query",
        "body"
    ],
    responseWhitelist: [
        "header",
        "body"
    ]
};

if (process.env.DEBUG === "false") {
    loggerOptions.meta = false;
}

// Initialize the logger.
app.use(expressWinston.logger(loggerOptions));

// Add webhook route.
app.post( "/webhook", (req: express.Request, res: express.Response) => {
    logServerStatus("Webhook request received!");
    // check console for logging output!
    res.sendStatus(202);
});

// Start server.
server.listen(port, () => {
    logServerStatus(`Webhook server is running on port: ${port}!`);
});