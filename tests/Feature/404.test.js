import supertest from "supertest";
import { api } from "../../src/app/api.js";

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
	it("404 Not Found User", async function () {
		for (let index = 0; index < 10; index++) {
			const result = await supertest(api)
				.get("/" + genrateURL())
				.set("Accept", "application/json")
				.set("Authorization", "Bearer 1121");
			expect(result.status).toBe(404);
			expect(result.body.message).toContain("Not Found");
		}
	});

	it("404 Not Found Member", async function () {
		for (let index = 0; index < 10; index++) {
			const result = await supertest(api)
				.get("/" + genrateURL())
				.set("Accept", "application/json")
				.set("Authorization", "Bearer 1122");
			expect(result.status).toBe(404);
			expect(result.body.message).toContain("Not Found");
		}
	});
	it("404 Not Found Premium", async function () {
		for (let index = 0; index < 10; index++) {
			const result = await supertest(api)
				.get("/" + genrateURL())
				.set("Accept", "application/json")
				.set("Authorization", "Bearer 1123");
			expect(result.status).toBe(404);
			expect(result.body.message).toContain("Not Found");
		}
	});
});
