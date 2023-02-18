import { BaseAPI, route } from "../base";

export class ReportsAPI extends BaseAPI {
  billOfMaterialsURL(): string {
    return route("/reporting/bill-of-materials");
  }
}
