import { BaseAPI, route } from "./base";
import { ApiSummary, TokenResponse, UserRegistration } from "./types/data-contracts";

export type LoginPayload = {
  username: string;
  password: string;
};

export type StatusResult = {
  health: boolean;
  versions: string[];
  title: string;
  message: string;
};

export class PublicApi extends BaseAPI {
  public status() {
    return this.http.get<ApiSummary>({ url: route("/status") });
  }

  public login(username: string, password: string) {
    return this.http.post<LoginPayload, TokenResponse>({
      url: route("/users/login"),
      body: {
        username,
        password,
      },
    });
  }

  public register(body: UserRegistration) {
    return this.http.post<UserRegistration, TokenResponse>({ url: route("/users/register"), body });
  }
}
