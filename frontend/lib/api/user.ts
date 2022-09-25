import { BaseAPI, route } from "./base";
import { ItemsApi } from "./classes/items";
import { LabelsApi } from "./classes/labels";
import { LocationsApi } from "./classes/locations";
import { UserOut } from "./types/data-contracts";
import { Requests } from "~~/lib/requests";

export type Result<T> = {
  item: T;
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
    return this.http.get<Result<UserOut>>({ url: route("/users/self") });
  }

  public logout() {
    return this.http.post<object, void>({ url: route("/users/logout") });
  }

  public deleteAccount() {
    return this.http.delete<void>({ url: route("/users/self") });
  }
}
