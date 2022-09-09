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
    return this.http.get<Results<Location>>({ url: route('/locations') });
  }

  async create(body: LocationCreate) {
    return this.http.post<LocationCreate, Location>({ url: route('/locations'), body });
  }

  async get(id: string) {
    return this.http.get<Location>({ url: route(`/locations/${id}`) });
  }
  async delete(id: string) {
    return this.http.delete<void>({ url: route(`/locations/${id}`) });
  }

  async update(id: string, body: LocationUpdate) {
    return this.http.put<LocationUpdate, Location>({ url: route(`/locations/${id}`), body });
  }
}
