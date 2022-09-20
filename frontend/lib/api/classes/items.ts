import { BaseAPI, route } from "../base";
import { parseDate } from "../base/base-api";
import { ItemAttachmentToken, ItemCreate, ItemOut, ItemSummary, ItemUpdate } from "../types/data-contracts";
import { AttachmentTypes } from "../types/non-generated";
import { Results } from "./types";

export class ItemsApi extends BaseAPI {
  getAll() {
    return this.http.get<Results<ItemSummary>>({ url: route("/items") });
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

  async update(id: string, item: ItemUpdate) {
    const payload = await this.http.put<ItemCreate, ItemOut>({
      url: route(`/items/${id}`),
      body: this.dropFields(item),
    });
    if (!payload.data) {
      return payload;
    }

    payload.data = parseDate(payload.data, ["purchaseTime", "soldTime", "warrantyExpires"]);
    return payload;
  }

  import(file: File) {
    const formData = new FormData();
    formData.append("csv", file);

    return this.http.post<FormData, void>({
      url: route("/items/import"),
      data: formData,
    });
  }

  addAttachment(id: string, file: File | Blob, filename: string, type: AttachmentTypes) {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("type", type);
    formData.append("name", filename);

    return this.http.post<FormData, ItemOut>({
      url: route(`/items/${id}/attachments`),
      data: formData,
    });
  }

  async getAttachmentUrl(id: string, attachmentId: string): Promise<string> {
    const payload = await this.http.get<ItemAttachmentToken>({
      url: route(`/items/${id}/attachments/${attachmentId}`),
    });

    if (!payload.data) {
      return "";
    }

    return route(`/items/${id}/attachments/download`, { token: payload.data.token });
  }

  deleteAttachment(id: string, attachmentId: string) {
    return this.http.delete<void>({ url: route(`/items/${id}/attachments/${attachmentId}`) });
  }
}
