import { BaseAPI, route } from "../base";
import { parseDate } from "../base/base-api";
import {
  ItemAttachmentUpdate,
  ItemCreate,
  ItemOut,
  ItemSummary,
  ItemUpdate,
  MaintenanceEntry,
  MaintenanceEntryCreate,
  MaintenanceEntryUpdate,
  MaintenanceLog,
} from "../types/data-contracts";
import { AttachmentTypes, PaginationResult } from "../types/non-generated";
import { Requests } from "~~/lib/requests";

export type ItemsQuery = {
  includeArchived?: boolean;
  page?: number;
  pageSize?: number;
  locations?: string[];
  labels?: string[];
  q?: string;
};

export class AttachmentsAPI extends BaseAPI {
  add(id: string, file: File | Blob, filename: string, type: AttachmentTypes) {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("type", type);
    formData.append("name", filename);

    return this.http.post<FormData, ItemOut>({
      url: route(`/items/${id}/attachments`),
      data: formData,
    });
  }

  delete(id: string, attachmentId: string) {
    return this.http.delete<void>({ url: route(`/items/${id}/attachments/${attachmentId}`) });
  }

  update(id: string, attachmentId: string, data: ItemAttachmentUpdate) {
    return this.http.put<ItemAttachmentUpdate, ItemOut>({
      url: route(`/items/${id}/attachments/${attachmentId}`),
      body: data,
    });
  }
}

export class MaintenanceAPI extends BaseAPI {
  getLog(itemId: string) {
    return this.http.get<MaintenanceLog>({ url: route(`/items/${itemId}/maintenance`) });
  }

  create(itemId: string, data: MaintenanceEntryCreate) {
    return this.http.post<MaintenanceEntryCreate, MaintenanceEntry>({
      url: route(`/items/${itemId}/maintenance`),
      body: data,
    });
  }

  delete(itemId: string, entryId: string) {
    return this.http.delete<void>({ url: route(`/items/${itemId}/maintenance/${entryId}`) });
  }

  update(itemId: string, entryId: string, data: MaintenanceEntryUpdate) {
    return this.http.put<MaintenanceEntryUpdate, MaintenanceEntry>({
      url: route(`/items/${itemId}/maintenance/${entryId}`),
      body: data,
    });
  }
}

export class ItemsApi extends BaseAPI {
  attachments: AttachmentsAPI;
  maintenance: MaintenanceAPI;

  constructor(http: Requests, token = "") {
    super(http, token);
    this.attachments = new AttachmentsAPI(http);
    this.maintenance = new MaintenanceAPI(http);
  }

  getAll(q: ItemsQuery = {}) {
    return this.http.get<PaginationResult<ItemSummary>>({ url: route("/items", q) });
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

  import(file: File | Blob) {
    const formData = new FormData();
    formData.append("csv", file);

    return this.http.post<FormData, void>({
      url: route("/items/import"),
      data: formData,
    });
  }
}
