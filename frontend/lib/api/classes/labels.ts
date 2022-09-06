import { BaseAPI, route } from '../base';
import { Item } from './items';
import { Details, OutType, Results } from './types';

export type LabelCreate = Details & {
  color: string;
};

export type LabelUpdate = LabelCreate;

export type Label = LabelCreate &
  OutType & {
    groupId: string;
    items: Item[];
  };

export class LabelsApi extends BaseAPI {
  async getAll() {
    return this.http.get<Results<Label>>(route('/labels'));
  }

  async create(label: LabelCreate) {
    return this.http.post<LabelCreate, Label>(route('/labels'), label);
  }

  async get(id: string) {
    return this.http.get<Label>(route(`/labels/${id}`));
  }

  async delete(id: string) {
    return this.http.delete<void>(route(`/labels/${id}`));
  }

  async update(id: string, label: LabelUpdate) {
    return this.http.put<LabelUpdate, Label>(route(`/labels/${id}`), label);
  }
}
