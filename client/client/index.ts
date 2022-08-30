import { v1ApiClient } from "./v1client";

export function getClientV1(baseUrl: string): v1ApiClient {
  return new v1ApiClient(baseUrl, "v1");
}
