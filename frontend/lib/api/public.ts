import { BaseAPI, route } from "./base";

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

export type StatusResult = {
  health: boolean;
  versions: string[];
  title: string;
  message: string;
};

export class PublicApi extends BaseAPI {
  public status() {
    return this.http.get<StatusResult>({ url: route("/status") });
  }

  public login(username: string, password: string) {
    return this.http.post<LoginPayload, LoginResult>({
      url: route("/users/login"),
      body: {
        username,
        password,
      },
    });
  }

  public register(body: RegisterPayload) {
    return this.http.post<RegisterPayload, LoginResult>({ url: route("/users/register"), body });
  }
}
