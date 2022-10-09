import { beforeAll, expect } from "vitest";
import { UserClient } from "../user";
import { factories } from "./factories";

const cache = {
  token: "",
};

/*
 * Shared UserApi token for tests where the creation of a user is _not_ import
 * to the test. This is useful for tests that are testing the user API itself.
 */
export async function sharedUserClient(): Promise<UserClient> {
  if (cache.token) {
    return factories.client.user(cache.token);
  }
  const testUser = {
    email: "__test__@__test__.com",
    name: "__test__",
    password: "__test__",
    token: "",
  };

  const api = factories.client.public();
  const { response: tryLoginResp, data } = await api.login(testUser.email, testUser.password);

  if (tryLoginResp.status === 200) {
    cache.token = data.token;
    return factories.client.user(cache.token);
  }

  const { response: registerResp } = await api.register(testUser);
  expect(registerResp.status).toBe(204);

  const { response: loginResp, data: loginData } = await api.login(testUser.email, testUser.password);
  expect(loginResp.status).toBe(200);

  cache.token = loginData.token;
  return factories.client.user(data.token);
}

beforeAll(async () => {
  await sharedUserClient();
});
