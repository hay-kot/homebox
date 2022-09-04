import { describe, it, expect } from 'vitest';
import { Requests } from '../../requests';
import { OverrideParts } from '../base/urls';
import { PublicApi } from '../public';
import * as config from '../../../test/config';
import { UserApi } from '../user';

function client() {
  OverrideParts(config.BASE_URL, '/api/v1');
  const requests = new Requests('');
  return new PublicApi(requests);
}

function userClient(token: string) {
  OverrideParts(config.BASE_URL, '/api/v1');
  const requests = new Requests('', token);
  return new UserApi(requests);
}

describe('[GET] /api/v1/status', () => {
  it('basic query parameter', async () => {
    const api = client();
    const { response, data } = await api.status();
    expect(response.status).toBe(200);
    expect(data.health).toBe(true);
  });
});

describe('first time user workflow (register, login)', () => {
  const api = client();
  const userData = {
    groupName: 'test-group',
    user: {
      email: 'test-user@email.com',
      name: 'test-user',
      password: 'test-password',
    },
  };

  it('user should be able to register', async () => {
    const { response } = await api.register(userData);
    expect(response.status).toBe(204);
  });

  it('user should be able to login', async () => {
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
