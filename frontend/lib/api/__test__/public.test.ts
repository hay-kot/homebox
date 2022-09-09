import { describe, test, expect } from "vitest";
import { client, userClient } from "./test-utils";

describe("[GET] /api/v1/status", () => {
  test("server should respond", async () => {
    const api = client();
    const { response, data } = await api.status();
    expect(response.status).toBe(200);
    expect(data.health).toBe(true);
  });
});

describe("first time user workflow (register, login)", () => {
  const api = client();
  const userData = {
    groupName: "test-group",
    user: {
      email: "test-user@email.com",
      name: "test-user",
      password: "test-password",
    },
  };

  test("user should be able to register", async () => {
    const { response } = await api.register(userData);
    expect(response.status).toBe(204);
  });

  test("user should be able to login", async () => {
    const { response, data } = await api.login(userData.user.email, userData.user.password);
    expect(response.status).toBe(200);
    expect(data.token).toBeTruthy();

    // Cleanup
    const userApi = userClient(data.token);
    {
      const { response } = await userApi.deleteAccount();
      expect(response.status).toBe(204);
    }
  });
});
