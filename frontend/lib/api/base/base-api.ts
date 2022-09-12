import { Requests } from "../../requests";
//  <
// 	TGetResult,
// 	TPostData,
// 	TPostResult,
// 	TPutData = TPostData,
// 	TPutResult = TPostResult,
// 	TDeleteResult = void
// >

type BaseApiType = {
  createdAt: string;
  updatedAt: string;
};

export function hasKey(obj: object, key: string): obj is Required<BaseApiType> {
  return typeof obj[key] === "string";
}

export function parseDate<T>(obj: T, keys: Array<keyof T> = []): T {
  const result = { ...obj };
  [...keys, "createdAt", "updatedAt"].forEach(key => {
    // @ts-ignore - we are checking for the key above
    if (hasKey(result, key)) {
      // @ts-ignore - we are guarding against this above
      result[key] = new Date(result[key]);
    }
  });

  return result;
}

export class BaseAPI {
  http: Requests;

  constructor(requests: Requests) {
    this.http = requests;
  }

  /**
   * dropFields will remove any fields that are specified in the fields array
   * additionally, it will remove the `createdAt` and `updatedAt` fields if they
   * are present. This is useful for when you want to send a subset of fields to
   * the server like when performing an update.
   */
  dropFields<T>(obj: T, keys: Array<keyof T> = []): T {
    const result = { ...obj };
    console.log("dropFields", result);
    [...keys, "createdAt", "updatedAt"].forEach(key => {
      console.log(key);
      // @ts-ignore - we are checking for the key above
      if (hasKey(result, key)) {
        // @ts-ignore - we are guarding against this above
        delete result[key];
        console.log("dropping", key);
      }
    });
    console.log("dropFields", result);
    return result;
  }
}
