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
}
