import { faker } from "@faker-js/faker";
import { describe, expect, test } from "vitest";
import { factories } from "../factories";

describe("basic user workflows", () => {
  test("user should be able to change password", async () => {
    const { client, user } = await factories.client.singleUse();
    const password = faker.internet.password();

    // Change Password
    {
      const response = await client.user.changePassword(user.password, password);
      expect(response.error).toBeFalsy();
      expect(response.status).toBe(204);
    }

    // Ensure New Login is Valid
    {
      const pub = factories.client.public();
      const response = await pub.login(user.email, password);
      expect(response.error).toBeFalsy();
      expect(response.status).toBe(200);
    }

    await client.user.delete();
  }, 20000);
});
