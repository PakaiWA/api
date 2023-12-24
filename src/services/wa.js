import { logger } from "../app/logger.js";
import "dotenv/config";
import QRCode from "qrcode";
import WAWebJS from "whatsapp-web.js";
const { Client, LocalAuth } = WAWebJS;

class PakaiWA {
	constructor(device_id, msg) {
		this.device_id = device_id;
		this.msg = msg;
		this.cooldown = 0;
		this.client = new Client({
			puppeteer: {
				args: ["--no-sandbox"],
			},
			authStrategy: new LocalAuth({
				dataPath: "./sessions",
				clientId: device_id,
			}),
			info: {},
		});
	}

	init() {
		this.client.initialize();
	}

	getQRCode() {
		logger.info("Request QR Code: " + this.device_id);
		this.client.on("qr", (qr) => {
			logger.info("QR for " + this.device_id + "\n\n");
			if (this.cooldown < 2) {
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
				this.cooldown++;
			} else {
				this.client.destroy();
				logger.info("Destroy");
			}
		});

		this.client.on("authenticated", () => {
			logger.info(`Client: ${this.device_id} is Authenticated`);
		});

		this.client.on("auth_failure", (msg) => {
			logger.info("auth_failure", msg);
		});

		this.client.on("ready", () => {
			logger.info(`Client : ${this.device_id} is ready.`);
		});

		this.client.on("message", async (msg) => {
			if (!msg.isStatus) {
				logger.info(msg);
			}
			logger.info(
				`Client: ${this.device_id} received a mesage`,
				msg.body + " from " + msg.from
			);
			if (msg.body === "!ping") {
				this.client.sendMessage(msg.from, "pong");
			}
		});

		this.client.on("disconnected", (reason) => {
			logger.info("disconnected", reason);
			this.client.destroy();
		});
	}

	sendMessage(msg) {
		logger.info(msg);
		this.client.sendMessage(msg);
	}
}
export default PakaiWA;
