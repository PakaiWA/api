import supertest from "supertest";
import { app } from "../../src/app/api.js";

function genrateURL() {
	let chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
	let charLength = chars.length;
	let url = "";
	for (let i = 0; i < 10; i++) {
		url += chars.charAt(Math.floor(Math.random() * charLength));
	}
	return url;
}

describe("404 Test", function () {
	for (let index = 0; index < 10; index++) {
		it(index + ". 404 Not Found User", async function () {
			const result = await supertest(app)
				.get("/" + genrateURL())
				.set("Accept", "application/json;charset=utf-8")
				.set("Authorization", "Bearer 1121");
			expect(result.status).toBe(404);
			expect(result.body.message).not.toBeNull();
		});
	}
	for (let index = 0; index < 10; index++) {
		it(index + ". 404 Not Found Member", async function () {
			const result = await supertest(app)
				.get("/" + genrateURL())
				.set("Accept", "application/json;charset=utf-8")
				.set("Authorization", "Bearer 1122");
			expect(result.status).toBe(404);
			expect(result.body.message).not.toBeNull();
		});
	}
	for (let index = 0; index < 10; index++) {
		it(index + ". 404 Not Found Premium", async function () {
			const result = await supertest(app)
				.get("/" + genrateURL())
				.set("Accept", "application/json;charset=utf-8")
				.set("Authorization", "Bearer 1123");
			expect(result.status).toBe(404);
			expect(result.body.message).not.toBeNull();
		});
	}
});
