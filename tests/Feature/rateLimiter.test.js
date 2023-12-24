import supertest from "supertest";
import { api } from "../../src/app/api.js";

describe("GET /", function () {
	it("should return 200 OK", async function () {
		let result;
		for (let i = 0; i < 100; i++) {
			result = await supertest(api)
				.get("/")
				.set("Accept", "application/json")
				.set("Authorization", "Bearer 1121");
		}
		expect(result.status).toBe(200);
		expect(result.body.message).toContain("PakaiWA.my.id");
		expect(result.body.version).not.toBeNull();
		expect(result.body.stability).not.toBeNull();
	});
});
