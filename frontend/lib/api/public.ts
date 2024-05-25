import { BaseAPI, route } from "./base";
import type {
  APISummary,
  LoginForm,
  PasswordResetRequest,
  TokenResponse,
  UserRegistration,
} from "./types/data-contracts";

export type StatusResult = {
  health: boolean;
  versions: string[];
  title: string;
  message: string;
};

export class PublicApi extends BaseAPI {
  public status() {
    return this.http.get<APISummary>({ url: route("/status") });
  }

  public login(username: string, password: string, stayLoggedIn = false) {
    return this.http.post<LoginForm, TokenResponse>({
      url: route("/users/login"),
      body: {
        username,
        password,
        stayLoggedIn,
      },
    });
  }

  public register(body: UserRegistration) {
    return this.http.post<UserRegistration, TokenResponse>({ url: route("/users/register"), body });
  }

  public resetPasseord(email: string) {
    return this.http.post<PasswordResetRequest, void>({
      url: route("/users/request-password-reset"),
      body: { email },
    });
  }
}
