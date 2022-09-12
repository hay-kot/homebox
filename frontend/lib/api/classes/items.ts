import { BaseAPI, route } from "../base";
import { parseDate } from "../base/base-api";
import { ItemCreate, ItemOut } from "../types/data-contracts";
import { Results } from "./types";

export class ItemsApi extends BaseAPI {
  getAll() {
    return this.http.get<Results<ItemOut>>({ url: route("/items") });
  }

  create(item: ItemCreate) {
    return this.http.post<ItemCreate, ItemOut>({ url: route("/items"), body: item });
  }

  async get(id: string) {
    const payload = await this.http.get<ItemOut>({ url: route(`/items/${id}`) });

    if (!payload.data) {
      return payload;
    }

    // Parse Date Types
    payload.data = parseDate(payload.data, ["purchaseTime", "soldTime", "warrantyExpires"]);
    return payload;
  }

  delete(id: string) {
    return this.http.delete<void>({ url: route(`/items/${id}`) });
  }

  update(id: string, item: ItemCreate) {
    return this.http.put<ItemCreate, ItemOut>({ url: route(`/items/${id}`), body: item });
  }

  import(file: File) {
    const formData = new FormData();
    formData.append("csv", file);

    return this.http.post<FormData, void>({ url: route("/items/import"), data: formData });
  }
}
