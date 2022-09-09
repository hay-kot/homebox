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
  purchaseTime: Date;
  serialNumber: string;
  soldNotes: string;
  soldPrice: number;
  soldTime: Date;
  soldTo: string;
  updatedAt: string;
  lifetimeWarranty: boolean;
  warrantyExpires: Date;
  warrantyDetails: string;
}

export class ItemsApi extends BaseAPI {
  getAll() {
    return this.http.get<Results<Item>>({ url: route('/items') });
  }

  create(item: ItemCreate) {
    return this.http.post<ItemCreate, Item>({ url: route('/items'), body: item });
  }

  async get(id: string) {
    const payload = await this.http.get<Item>({ url: route(`/items/${id}`) });

    if (!payload.data) {
      return payload;
    }

    // Parse Date Types
    payload.data.purchaseTime = new Date(payload.data.purchaseTime);
    payload.data.soldTime = new Date(payload.data.soldTime);
    payload.data.warrantyExpires = new Date(payload.data.warrantyExpires);

    return payload;
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/items/${id}`) });
  }

  update(id: string, item: ItemCreate) {
    return this.http.put<ItemCreate, Item>({ url: route(`/items/${id}`), body: item });
  }

  import(file: File) {
    const formData = new FormData();
    formData.append('csv', file);

    return this.http.post<FormData, void>({ url: route('/items/import'), data: formData });
  }
}
