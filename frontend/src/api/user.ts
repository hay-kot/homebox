import { BaseAPI, UrlBuilder } from './base';

export type Result<T> = {
  item: T;
};

export type User = {
  name: string;
  email: string;
  isSuperuser: boolean;
  id: number;
};

export class UserApi extends BaseAPI {
  public self() {
    return this.http.get<Result<User>>(UrlBuilder('/users/self'));
  }

  public logout() {
    return this.http.post<object, void>(UrlBuilder('/users/logout'), {});
  }
}
