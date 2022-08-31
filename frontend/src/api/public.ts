import { BaseAPI, UrlBuilder } from './base';

export type LoginResult = {
	token: string;
	expiresAt: string;
};

export type LoginPayload = {
	username: string;
	password: string;
};

export type RegisterPayload = {
	user: {
		email: string;
		password: string;
		name: string;
	};
	groupName: string;
};

export class PublicApi extends BaseAPI {
	public login(username: string, password: string) {
		return this.http.post<LoginPayload, LoginResult>(
			UrlBuilder('/users/login'),
			{
				username,
				password,
			}
		);
	}

	public register(payload: RegisterPayload) {
		return this.http.post<RegisterPayload, LoginResult>(
			UrlBuilder('/users/register'),
			payload
		);
	}
}
