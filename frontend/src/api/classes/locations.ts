import { BaseAPI, UrlBuilder } from '../base';
import { type Results } from '../base/base-types';

export type LocationCreate = {
  name: string;
  description: string;
};

export type Location = LocationCreate & {
  id: string;
  groupId: string;
  createdAt: string;
  updatedAt: string;
};

export class LocationsApi extends BaseAPI {
  async getAll() {
    return this.http.get<Results<Location>>(UrlBuilder('/locations'));
  }

  async create(location: LocationCreate) {
    return this.http.post<LocationCreate, Location>(UrlBuilder('/locations'), location);
  }
}
