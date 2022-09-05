import { BaseAPI, route } from '../base';
import { Item } from './items';
import { Details, OutType, Results } from './types';

export type LocationCreate = Details;

export type Location = LocationCreate &
  OutType & {
    groupId: string;
    items: Item[];
    itemCount: number;
  };

export type LocationUpdate = LocationCreate;

export class LocationsApi extends BaseAPI {
  async getAll() {
    return this.http.get<Results<Location>>(route('/locations'));
  }

  async create(location: LocationCreate) {
    return this.http.post<LocationCreate, Location>(route('/locations'), location);
  }

  async get(id: string) {
    return this.http.get<Location>(route(`/locations/${id}`));
  }
  async delete(id: string) {
    return this.http.delete<void>(route(`/locations/${id}`));
  }

  async update(id: string, location: LocationUpdate) {
    return this.http.put<LocationUpdate, Location>(route(`/locations/${id}`), location);
  }
}
