import { BaseAPI } from "./base";
import { ItemsApi } from "./classes/items";
import { LabelsApi } from "./classes/labels";
import { LocationsApi } from "./classes/locations";
import { GroupApi } from "./classes/group";
import { UserApi } from "./classes/users";
import { ActionsAPI } from "./classes/actions";
import { StatsAPI } from "./classes/stats";
import { AssetsApi } from "./classes/assets";
import { Requests } from "~~/lib/requests";

export class UserClient extends BaseAPI {
  locations: LocationsApi;
  labels: LabelsApi;
  items: ItemsApi;
  group: GroupApi;
  user: UserApi;
  actions: ActionsAPI;
  stats: StatsAPI;
  assets: AssetsApi;

  constructor(requests: Requests, attachmentToken: string) {
    super(requests, attachmentToken);

    this.locations = new LocationsApi(requests);
    this.labels = new LabelsApi(requests);
    this.items = new ItemsApi(requests, attachmentToken);
    this.group = new GroupApi(requests);
    this.user = new UserApi(requests);
    this.actions = new ActionsAPI(requests);
    this.stats = new StatsAPI(requests);
    this.assets = new AssetsApi(requests);

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
