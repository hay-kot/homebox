import { Requests } from "../../requests";

const ZERO_DATE = "0001-01-01T00:00:00Z";

type BaseApiType = {
  createdAt: string;
  updatedAt: string;

  [key: string]: any;
};

export function hasKey(obj: object, key: string): obj is Required<BaseApiType> {
  return typeof obj[key] === "string";
}

export function parseDate<T>(obj: T, keys: Array<keyof T> = []): T {
  const result = { ...obj };
  [...keys, "createdAt", "updatedAt"].forEach(key => {
    // @ts-ignore - TS doesn't know that we're checking for the key above
    if (hasKey(result, key)) {
      if (result[key] === ZERO_DATE) {
        const dt = new Date();
        dt.setFullYear(1);

        result[key] = dt;
        return;
      }

      // transform string to ensure dates are parsed as UTC dates instead of
      // localized time stamps
      const asStr = result[key] as string;
      const cleaned = asStr.replaceAll("-", "/").split("T")[0];
      result[key] = new Date(cleaned);
    }
  });

  return result;
}

export class BaseAPI {
  http: Requests;
  attachmentToken: string;

  constructor(requests: Requests, attachmentToken = "") {
    this.http = requests;
    this.attachmentToken = attachmentToken;
  }

  // if a attachmentToken is present it will be added to URL as a query param
  // this is done with a simple appending of the query param to the URL. If your
  // URL already has a query param, this will not work.
  authURL(url: string): string {
    if (this.attachmentToken) {
      return `/api/v1${url}?access_token=${this.attachmentToken}`;
    }
    return url;
  }

  /**
   * dropFields will remove any fields that are specified in the fields array
   * additionally, it will remove the `createdAt` and `updatedAt` fields if they
   * are present. This is useful for when you want to send a subset of fields to
   * the server like when performing an update.
   */
  protected dropFields<T>(obj: T, keys: Array<keyof T> = []): T {
    const result = { ...obj };
    [...keys, "createdAt", "updatedAt"].forEach(key => {
      // @ts-ignore - we are checking for the key above
      if (hasKey(result, key)) {
        // @ts-ignore - we are guarding against this above
        delete result[key];
      }
    });
    return result;
  }
}
