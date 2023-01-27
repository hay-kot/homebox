import { BaseAPI, route } from "../base";
import { LocationOutCount, LocationCreate, LocationOut, LocationUpdate, TreeItem } from "../types/data-contracts";
import { Results } from "../types/non-generated";

export type LocationsQuery = {
  filterChildren: boolean;
};

export class LocationsApi extends BaseAPI {
  getAll(q: LocationsQuery = { filterChildren: false }) {
    return this.http.get<Results<LocationOutCount>>({ url: route("/locations", q) });
  }

  getTree() {
    return this.http.get<Results<TreeItem>>({ url: route("/locations/tree") });
  }

  create(body: LocationCreate) {
    return this.http.post<LocationCreate, LocationOut>({ url: route("/locations"), body });
  }

  get(id: string) {
    return this.http.get<LocationOut>({ url: route(`/locations/${id}`) });
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/locations/${id}`) });
  }

  update(id: string, body: LocationUpdate) {
    return this.http.put<LocationUpdate, LocationOut>({ url: route(`/locations/${id}`), body });
  }
}
