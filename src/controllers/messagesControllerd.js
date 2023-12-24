import { logger } from "../app/logger.js";
const sendMessage = async (req, res) => {
	const body = req.body;

	const text = body.message;
	const receiver = body.phone_number;
	const device_id = body.device_id;
	const message_type = body.message_type;

	const message = {
		message: text,
		phone_number: receiver,
		message_type: message_type,
		device_id: device_id,
	};

	logger.info(message);

	try {
		res.status(200).json(message);
	} catch (error) {
		logger.error("Something went wrong");
		res.status(400).json(message);
	}
};

export default sendMessage;
