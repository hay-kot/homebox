import { describe, expect, test } from "vitest";

import { PublicApi } from "../../public";
import * as config from "../../../../test/config";
import { Requests } from "../../../requests";
import { overrideParts } from "../../base/urls";

describe("trusted header handling", () => {
  overrideParts(config.BASE_URL, "/api/v1");
  const requests = new Requests("");
  const pub = new PublicApi(requests);

  test("basic login using HTTP headers", async () => {
    const ssoHeaders = {
      "Remote-Email": "test@test.com",
      "Remote-Name": "Test User",
      "Remote-Groups": "admins,local",
    };

    const response = await pub.login_sso_header(ssoHeaders);
    expect(response.error).toBeFalsy();
  }, 20000);

  test("basic login using HTTP headers fails no headers", async () => {
    const ssoHeaders = {};

    const response = await pub.login_sso_header(ssoHeaders);
    expect(response.error).toBeTruthy();
  }, 20000);

  test("basic login using HTTP headers empty email header", async () => {
    const ssoHeaders = {
      "Remote-Email": "",
    };

    const response = await pub.login_sso_header(ssoHeaders);
    expect(response.error).toBeTruthy();
  }, 20000);
});
