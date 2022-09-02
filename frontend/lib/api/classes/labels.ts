import { BaseAPI, UrlBuilder } from '../base';
import { Details, OutType, Results } from './types';

export type LabelCreate = Details & {
  color: string;
};

export type LabelUpdate = LabelCreate;

export type Label = LabelCreate &
  OutType & {
    groupId: string;
  };

export class LabelsApi extends BaseAPI {
  async getAll() {
    return this.http.get<Results<Label>>(UrlBuilder('/labels'));
  }

  async create(label: LabelCreate) {
    return this.http.post<LabelCreate, Label>(UrlBuilder('/labels'), label);
  }

  async get(id: string) {
    return this.http.get<Label>(UrlBuilder(`/labels/${id}`));
  }

  async delete(id: string) {
    return this.http.delete<void>(UrlBuilder(`/labels/${id}`));
  }

  async update(id: string, label: LabelUpdate) {
    return this.http.put<LabelUpdate, Label>(UrlBuilder(`/labels/${id}`), label);
  }
}
