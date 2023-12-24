/* eslint-disable no-console */
import { prismaClient } from "../app/database.js";
import PakaiWA from "./wa.js";

const pakaiWASessions = {};

const check = await prismaClient.hardware.findMany({
	select: {
		name: true,
	},
	orderBy: {
		id: "asc",
	},
});

export function myFunction(item) {
	const device_id = item["name"];
	pakaiWASessions[device_id] = new PakaiWA(device_id);
	pakaiWASessions[device_id].init();
	pakaiWASessions[device_id].getQRCode();
}

check.forEach(myFunction);
