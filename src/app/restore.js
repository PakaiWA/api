import "dotenv/config";
import QRCode from "qrcode";
import Whatsapp from "whatsapp-web.js";
const { Client, LocalAuth } = Whatsapp;
import { prismaClient } from "./database.js";
import { logger } from "./logger.js";

const clientSessionStore = {};
let clientInstance;

const sessionManager = function (device_id) {
	let timer = 0;
	const client = new Client({
		puppeteer: {
			args: ["--no-sandbox"],
		},
		authStrategy: new LocalAuth({
			dataPath: "./sessions",
			clientId: device_id,
		}),
	});

	client.on("qr", (qr) => {
		QRCode.toString(
			qr,
			{
				type: "terminal",
				small: true,
			},
			function (err, url) {
				logger.info(url);
			}
		);
		if (timer > 5) {
			client.destroy();
			logger.info("Destroy");
		} else {
			logger.info(device_id + ": " + timer);
			timer++;
		}
	});

	client.on("authenticated", () => {
		logger.info(`Client: ${device_id} is Authenticated`);
	});

	client.on("auth_failure", (msg) => {
		logger.info("auth_failure", msg);
	});

	client.on("ready", () => {
		logger.info(`Client : ${device_id} is ready.`);
	});

	client.on("message", async (msg) => {
		logger.info(`Client: ${device_id} received a mesage`, msg.body);
	});

	client.on("disconnected", (reason) => {
		logger.info("disconnected", reason);
		client.destroy();
	});

	client.initialize();

	return client;
};

const check = await prismaClient.hardware.findMany({
	select: {
		name: true,
	},
});

function myFunction(item) {
	const device_id = item["name"];
	clientInstance = sessionManager(device_id);
	clientSessionStore[device_id] = clientInstance;
}

check.forEach(myFunction);
