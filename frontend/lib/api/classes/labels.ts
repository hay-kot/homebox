import { BaseAPI, route } from "../base";
import { LabelCreate, LabelOut } from "../types/data-contracts";
import { Results } from "./types";

export class LabelsApi extends BaseAPI {
  getAll() {
    return this.http.get<Results<LabelOut>>({ url: route("/labels") });
  }

  create(body: LabelCreate) {
    return this.http.post<LabelCreate, LabelOut>({ url: route("/labels"), body });
  }

  get(id: string) {
    return this.http.get<LabelOut>({ url: route(`/labels/${id}`) });
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/labels/${id}`) });
  }

  update(id: string, body: LabelCreate) {
    return this.http.put<LabelCreate, LabelOut>({ url: route(`/labels/${id}`), body });
  }
}
