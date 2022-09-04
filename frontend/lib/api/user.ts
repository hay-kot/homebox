import { Requests } from '~~/lib/requests';
import { BaseAPI, UrlBuilder } from './base';
import { ItemsApi } from './classes/items';
import { LabelsApi } from './classes/labels';
import { LocationsApi } from './classes/locations';

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
  locations: LocationsApi;
  labels: LabelsApi;
  items: ItemsApi;
  constructor(requests: Requests) {
    super(requests);

    this.locations = new LocationsApi(requests);
    this.labels = new LabelsApi(requests);
    this.items = new ItemsApi(requests);

    Object.freeze(this);
  }

  public self() {
    return this.http.get<Result<User>>(UrlBuilder('/users/self'));
  }

  public logout() {
    return this.http.post<object, void>(UrlBuilder('/users/logout'), {});
  }

  public deleteAccount() {
    return this.http.delete<void>(UrlBuilder('/users/self'));
  }
}
