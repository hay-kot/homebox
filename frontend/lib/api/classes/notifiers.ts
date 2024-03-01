import { BaseAPI, route } from "../base";
import type { NotifierCreate, NotifierOut, NotifierUpdate } from "../types/data-contracts";

export class NotifiersAPI extends BaseAPI {
  getAll() {
    return this.http.get<NotifierOut[]>({ url: route("/notifiers") });
  }

  create(body: NotifierCreate) {
    return this.http.post<NotifierCreate, NotifierOut>({ url: route("/notifiers"), body });
  }

  update(id: string, body: NotifierUpdate) {
    if (body.url === "") {
      body.url = null;
    }

    return this.http.put<NotifierUpdate, NotifierOut>({ url: route(`/notifiers/${id}`), body });
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/notifiers/${id}`) });
  }

  test(url: string) {
    return this.http.post<{ url: string }, null>({ url: route(`/notifiers/test`), body: { url } });
  }
}
