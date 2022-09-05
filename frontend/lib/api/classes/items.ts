import { BaseAPI, route } from '../base';
import { Label } from './labels';
import { Location } from './locations';
import { Results } from './types';

export interface ItemCreate {
  name: string;
  description: string;
  locationId: string;
  labelIds: string[];
}

export interface Item {
  createdAt: string;
  description: string;
  id: string;
  labels: Label[];
  location: Location;
  manufacturer: string;
  modelNumber: string;
  name: string;
  notes: string;
  purchaseFrom: string;
  purchasePrice: number;
  purchaseTime: string;
  serialNumber: string;
  soldNotes: string;
  soldPrice: number;
  soldTime: string;
  soldTo: string;
  updatedAt: string;
}

export class ItemsApi extends BaseAPI {
  async getAll() {
    return this.http.get<Results<Item>>(route('/items'));
  }

  async create(item: ItemCreate) {
    return this.http.post<ItemCreate, Item>(route('/items'), item);
  }

  async get(id: string) {
    return this.http.get<Item>(route(`/items/${id}`));
  }

  async delete(id: string) {
    return this.http.delete<void>(route(`/items/${id}`));
  }

  async update(id: string, item: ItemCreate) {
    return this.http.put<ItemCreate, Item>(route(`/items/${id}`), item);
  }
}
