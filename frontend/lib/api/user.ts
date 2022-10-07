import { BaseAPI } from "./base";
import { ItemsApi } from "./classes/items";
import { LabelsApi } from "./classes/labels";
import { LocationsApi } from "./classes/locations";
import { GroupApi } from "./classes/group";
import { UserApi } from "./classes/users";
import { Requests } from "~~/lib/requests";

export class UserClient extends BaseAPI {
  locations: LocationsApi;
  labels: LabelsApi;
  items: ItemsApi;
  group: GroupApi;
  user: UserApi;

  constructor(requests: Requests) {
    super(requests);

    this.locations = new LocationsApi(requests);
    this.labels = new LabelsApi(requests);
    this.items = new ItemsApi(requests);
    this.group = new GroupApi(requests);
    this.user = new UserApi(requests);

    Object.freeze(this);
  }

  /** @deprecated use this.user.self() */
  public self() {
    return this.user.self();
  }

  /** @deprecated use this.user.logout() */
  public logout() {
    return this.user.logout();
  }

  /** @deprecated use this.user.delete() */
  public deleteAccount() {
    return this.user.delete();
  }
}
