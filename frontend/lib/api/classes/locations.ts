import { BaseAPI, UrlBuilder } from '../base';
import { Details, OutType, Results } from './types';

export type LocationCreate = Details;

export type Location = LocationCreate &
  OutType & {
    groupId: string;
  };

export type LocationUpdate = LocationCreate;

export class LocationsApi extends BaseAPI {
  async getAll() {
    return this.http.get<Results<Location>>(UrlBuilder('/locations'));
  }

  async create(location: LocationCreate) {
    return this.http.post<LocationCreate, Location>(UrlBuilder('/locations'), location);
  }

  async get(id: string) {
    return this.http.get<Location>(UrlBuilder(`/locations/${id}`));
  }
  async delete(id: string) {
    return this.http.delete<void>(UrlBuilder(`/locations/${id}`));
  }

  async update(id: string, location: LocationUpdate) {
    return this.http.put<LocationUpdate, Location>(UrlBuilder(`/locations/${id}`), location);
  }
}
