import { BaseAPI, route } from "../base";
import type { ItemSummary } from "../types/data-contracts";
import type { PaginationResult } from "../types/non-generated";

export class AssetsApi extends BaseAPI {
  async get(id: string, page = 1, pageSize = 50) {
    return await this.http.get<PaginationResult<ItemSummary>>({
      url: route(`/assets/${id}`, { page, pageSize }),
    });
  }
}
