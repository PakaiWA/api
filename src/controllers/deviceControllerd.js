import { logger } from "../app/logger.js";
import { prismaClient } from "../app/database.js";

const addDevice = async (req, res) => {
	const device_id = req.body.device_id;

	if (!device_id) {
		return res.status(400).json({
			success: false,
			message: "Missing device_id",
		});
	}

	const device_id_exists = await prismaClient.device.findFirst({
		where: {
			device_id: device_id,
			User: {
				token: global.token,
			},
		},
		select: {
			device_id: true,
		},
	});

	if (device_id_exists) {
		res.status(400).json({
			message: "Error: Device id already exists, choose another one.",
		});
	} else {
		const user = await prismaClient.user.findFirst({
			where: {
				token: global.token,
			},
			select: {
				email: true,
			},
		});
		try {
			const device = await prismaClient.device.create({
				data: {
					device_id: device_id,
					userEmail: user.email,
				},
				select: {
					device_id: true,
					status: true,
					created_at: true,
				},
			});
			res.status(201).json({
				id: device.device_id,
				status: device.status,
				created_at: device.created_at,
				meta: {
					location: `${global.host}/v1/devices/${device.device_id}`,
				},
			});
		} catch (error) {
			/* istanbul ignore next */
			res.status(500).json({ message: "Error: Internal server error" });
			/* istanbul ignore next */
			logger.error(error);
		}
	}
};

const getDeviceList = async (req, res) => {
	try {
		const device = await prismaClient.device.findMany({
			where: {
				User: {
					token: global.token,
				},
			},
			select: {
				device_id: true,
				status: true,
				created_at: true,
				connected_at: true,
				disconnected_at: true,
				disconnected_reason: true,
			},
		});

		if (device.length <= 0) {
			return res.status(404).json({
				message: "Device not found",
			});
		}
		const last_key = device.slice(-1);

		res.status(200).json({
			data: [
				...device.map((device) => ({
					id: device.device_id,
					status: device.status,
					created_at: device.created_at,
					connected_at: device.connected_at,
					disconnected_at: device.disconnected_at,
					disconnected_reason: device.disconnected_reason,
				})),
			],
			meta: {
				last_key: `${last_key[0].device_id}`,
			},
		});
	} catch (error) {
		/* istanbul ignore next */
		logger.error(error);
		/* istanbul ignore next */
		return res.status(500).json({ message: "Error: Internal server error" });
	}
};

const getDeviceStatus = async (req, res) => {
	const device_id = req.params.device_id;

	try {
		const device = await prismaClient.device.findFirst({
			where: {
				device_id: device_id,
				User: {
					token: global.token,
				},
			},
			select: {
				device_id: true,
				status: true,
				created_at: true,
				connected_at: true,
				disconnected_at: true,
				disconnected_reason: true,
			},
		});

		if (!device) {
			return res.status(404).json({
				message: "Device not found",
			});
		}

		res.status(200).json({
			data: {
				id: device.device_id,
				status: device.status,
				created_at: device.created_at,
				connected_at: device.connected_at,
				disconnected_at: device.disconnected_at,
				disconnected_reason: device.disconnected_reason,
			},
		});
	} catch (error) {
		/* istanbul ignore next */
		logger.error(error);
		/* istanbul ignore next */
		return res.status(500).json({ message: "Error: Internal server error" });
	}
};

const deleteDevice = async (req, res) => {
	const device_id = req.params.device_id;
	try {
		const device = await prismaClient.device.findFirst({
			where: {
				device_id: device_id,
				User: {
					token: global.token,
				},
			},
			select: {
				id: true,
				device_id: true,
				status: true,
				created_at: true,
				connected_at: true,
				disconnected_at: true,
				disconnected_reason: true,
			},
		});

		if (!device) {
			return res.status(404).json({
				message: "Device not found",
			});
		}

		await prismaClient.device.delete({
			where: {
				id: device.id,
			},
		});
		res.status(204).end();
	} catch (error) {
		/* istanbul ignore next */
		logger.error(error);
		/* istanbul ignore next */
		return res.status(500).json({ message: "Error: Internal server error" });
	}
};

export { addDevice, getDeviceList, getDeviceStatus, deleteDevice };
