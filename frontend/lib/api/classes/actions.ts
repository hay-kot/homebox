import { BaseAPI, route } from "../base";
import { ActionAmountResult } from "../types/data-contracts";

export class ActionsAPI extends BaseAPI {
  ensureAssetIDs() {
    return this.http.post<void, ActionAmountResult>({
      url: route("/actions/ensure-asset-ids"),
    });
  }

  resetItemDateTimes() {
    return this.http.post<void, ActionAmountResult>({
      url: route("/actions/zero-item-time-fields"),
    });
  }

  ensureImportRefs() {
    return this.http.post<void, ActionAmountResult>({
      url: route("/actions/ensure-import-refs"),
    });
  }
}
