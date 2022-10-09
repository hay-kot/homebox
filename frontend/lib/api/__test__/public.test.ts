import { describe, test, expect } from "vitest";
import { factories } from "./factories";
import { client, sharedUserClient, userClient } from "./test-utils";

describe("[GET] /api/v1/status", () => {
  test("server should respond", async () => {
    const api = client();
    const { response, data } = await api.status();
    expect(response.status).toBe(200);
    expect(data.health).toBe(true);
  });
});

describe("first time user workflow (register, login, join group)", () => {
  const api = client();
  const userData = factories.user();

  test("user should be able to register", async () => {
    const { response } = await api.register(userData);
    expect(response.status).toBe(204);
  });

  test("user should be able to login", async () => {
    const { response, data } = await api.login(userData.email, userData.password);
    expect(response.status).toBe(200);
    expect(data.token).toBeTruthy();

    // Cleanup
    const userApi = userClient(data.token);
    {
      const { response } = await userApi.user.delete();
      expect(response.status).toBe(204);
    }
  });

  test("user should be able to join create join token and have user signup", async () => {
    // Setup User 1 Token

    const client = await sharedUserClient();
    const { data: user1 } = await client.user.self();

    const { response, data } = await client.group.createInvitation({
      expiresAt: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
      uses: 1,
    });

    expect(response.status).toBe(201);
    expect(data.token).toBeTruthy();

    // Create User 2 with token

    const duplicateUser = factories.user();
    duplicateUser.token = data.token;

    const { response: registerResp } = await api.register(duplicateUser);
    expect(registerResp.status).toBe(204);

    const { response: loginResp, data: loginData } = await api.login(duplicateUser.email, duplicateUser.password);
    expect(loginResp.status).toBe(200);

    // Get Self and Assert

    const client2 = userClient(loginData.token);

    const { data: user2 } = await client2.user.self();

    user2.item.groupName = user1.item.groupName;

    // Cleanup User 2

    const { response: deleteResp } = await client2.user.delete();
    expect(deleteResp.status).toBe(204);
  });
});
