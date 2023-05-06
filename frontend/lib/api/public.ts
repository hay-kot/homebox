import { BaseAPI, route } from "./base";
import { ApiSummary, LoginForm, TokenResponse, UserRegistration } from "./types/data-contracts";

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

  // headers parameter only here for unit testing
  public login_sso_header(headers = {}) {
    const testHeaders = {
      /** TODO: remove headers here. Only for testing. Usually the SSO servie will add this */
      // "Remote-Email": "demo3@example.com",
      // "Remote-Name": "Fritz3",
      // "Remote-Groups": "admins,local",
    };
    const queryHeaders = { ...headers, ...testHeaders };
    return this.http.post<string, TokenResponse>({
      url: route("/users/login-sso-header"),
      headers: queryHeaders,
    });
  }

  public register(body: UserRegistration) {
    return this.http.post<UserRegistration, TokenResponse>({ url: route("/users/register"), body });
  }
}
