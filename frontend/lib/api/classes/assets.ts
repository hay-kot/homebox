import { BaseAPI, route } from "../base";
import { ItemSummary } from "../types/data-contracts";
import { PaginationResult } from "../types/non-generated";

export class AssetsApi extends BaseAPI {
  async get(id: string, page = 1, pageSize = 50) {
    return await this.http.get<PaginationResult<ItemSummary>>({
      url: route(`/asset/${id}`, { page, pageSize }),
    });
  }
}
