/* eslint-disable no-console */
import { prismaClient } from "../app/database.js";
import sessionManager from "./whatsapp.js";

const clientSessionStore = {};
let clientInstance;

const check = await prismaClient.hardware.findMany({
	select: {
		name: true,
	},
	orderBy: {
		id: "asc",
	},
});

function myFunction(item) {
	const device_id = item["name"];
	clientInstance = sessionManager(device_id);
	clientSessionStore[device_id] = clientInstance;
}

check.forEach(myFunction);
