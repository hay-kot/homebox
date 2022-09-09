import { BaseAPI, route } from "../base";
import { Item } from "./items";
import { Details, OutType, Results } from "./types";

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
  getAll() {
    return this.http.get<Results<Label>>({ url: route("/labels") });
  }

  create(body: LabelCreate) {
    return this.http.post<LabelCreate, Label>({ url: route("/labels"), body });
  }

  get(id: string) {
    return this.http.get<Label>({ url: route(`/labels/${id}`) });
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/labels/${id}`) });
  }

  update(id: string, body: LabelUpdate) {
    return this.http.put<LabelUpdate, Label>({ url: route(`/labels/${id}`), body });
  }
}
