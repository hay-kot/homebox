import { BaseAPI, route } from "../base";
import { Item } from "./items";
import { Details, OutType, Results } from "./types";

export type LocationCreate = Details;

export type Location = LocationCreate &
  OutType & {
    groupId: string;
    items: Item[];
    itemCount: number;
  };

export type LocationUpdate = LocationCreate;

export class LocationsApi extends BaseAPI {
  getAll() {
    return this.http.get<Results<Location>>({ url: route("/locations") });
  }

  create(body: LocationCreate) {
    return this.http.post<LocationCreate, Location>({ url: route("/locations"), body });
  }

  get(id: string) {
    return this.http.get<Location>({ url: route(`/locations/${id}`) });
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/locations/${id}`) });
  }

  update(id: string, body: LocationUpdate) {
    return this.http.put<LocationUpdate, Location>({ url: route(`/locations/${id}`), body });
  }
}
