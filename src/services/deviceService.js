import { addDeviceValidation } from "../validation/deviceValidation";
import { validate } from "../validation/validation.js";
import { prismaClient } from "../app/database.js";
import { ResponseError } from "../errors/response-error.js";

const addDevice = async (req) => {
	const device = validate(addDeviceValidation, req);

	const checkDevice = await prismaClient.device.count({
		where: {
			device_id: device.device_id,
		},
	});

	if (checkDevice === 1) {
		throw new ResponseError(400, "Device already exists");
	}

	return prismaClient.device.create({
		data: {
			device_id: device.device_id,
			userEmail: "a@b.c",
		},
		select: {
			device_id: true,
			status: true,
			created_at: true,
		},
	});
};

export { addDevice };
