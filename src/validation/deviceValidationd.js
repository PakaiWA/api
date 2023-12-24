import Joi from "joi";

const addDeviceValidation = Joi.object({
	device_id: Joi.string().required(),
});

export { addDeviceValidation };
