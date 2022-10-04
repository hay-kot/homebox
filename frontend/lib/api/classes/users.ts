import { BaseAPI, route } from "../base";
import { UserOut } from "../types/data-contracts";
import { Result } from "../types/non-generated";

export class UserApi extends BaseAPI {
  public self() {
    return this.http.get<Result<UserOut>>({ url: route("/users/self") });
  }

  public logout() {
    return this.http.post<object, void>({ url: route("/users/logout") });
  }

  public delete() {
    return this.http.delete<void>({ url: route("/users/self") });
  }
}
