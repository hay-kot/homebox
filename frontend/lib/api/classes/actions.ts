import { BaseAPI, route } from "../base";
import { EnsureAssetIDResult } from "../types/data-contracts";

export class ActionsAPI extends BaseAPI {
  ensureAssetIDs() {
    return this.http.post<void, EnsureAssetIDResult>({
      url: route("/actions/ensure-asset-ids"),
    });
  }
}
