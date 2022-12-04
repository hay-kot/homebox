import { BaseAPI, route } from "../base";
import { GroupStatistics, TotalsByOrganizer, ValueOverTime } from "../types/data-contracts";

function YYYY_DD_MM(date?: Date): string {
  if (!date) {
    return "";
  }
  // with leading zeros
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const day = date.getDate().toString().padStart(2, "0");
  return `${year}-${month}-${day}`;
}
export class StatsAPI extends BaseAPI {
  totalPriceOverTime(start?: Date, end?: Date) {
    return this.http.get<ValueOverTime>({
      url: route("/groups/statistics/purchase-price", { start: YYYY_DD_MM(start), end: YYYY_DD_MM(end) }),
    });
  }

  /**
   * Returns ths general statistics for the group. This mostly just
   * includes the totals for various group properties.
   */
  group() {
    return this.http.get<GroupStatistics>({
      url: route("/groups/statistics"),
    });
  }

  labels() {
    return this.http.get<TotalsByOrganizer[]>({
      url: route("/groups/statistics/labels"),
    });
  }

  locations() {
    return this.http.get<TotalsByOrganizer[]>({
      url: route("/groups/statistics/locations"),
    });
  }
}
