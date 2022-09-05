import { beforeAll, expect } from 'vitest';
import { Requests } from '../../requests';
import { overrideParts } from '../base/urls';
import { PublicApi } from '../public';
import * as config from '../../../test/config';
import { UserApi } from '../user';

export function client() {
  overrideParts(config.BASE_URL, '/api/v1');
  const requests = new Requests('');
  return new PublicApi(requests);
}

export function userClient(token: string) {
  overrideParts(config.BASE_URL, '/api/v1');
  const requests = new Requests('', token);
  return new UserApi(requests);
}

const cache = {
  token: '',
};

/*
 * Shared UserApi token for tests where the creation of a user is _not_ import
 * to the test. This is useful for tests that are testing the user API itself.
 */
export async function sharedUserClient(): Promise<UserApi> {
  if (cache.token) {
    return userClient(cache.token);
  }
  const testUser = {
    groupName: 'test-group',
    user: {
      email: '__test__@__test__.com',
      name: '__test__',
      password: '__test__',
    },
  };

  const api = client();
  const { response: tryLoginResp, data } = await api.login(testUser.user.email, testUser.user.password);

  if (tryLoginResp.status === 200) {
    cache.token = data.token;
    return userClient(cache.token);
  }

  const { response: registerResp } = await api.register(testUser);
  expect(registerResp.status).toBe(204);

  const { response: loginResp, data: loginData } = await api.login(testUser.user.email, testUser.user.password);
  expect(loginResp.status).toBe(200);

  cache.token = loginData.token;
  return userClient(data.token);
}

beforeAll(async () => {
  await sharedUserClient();
});
