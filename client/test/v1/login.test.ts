import { getClientV1 } from "../../client";
import { describe, it, expect } from "vitest";
import * as config from "../config";
import axios, { AxiosError } from "axios";

const client = getClientV1(config.BASE_URL);

describe("POST /api/v1/login", function () {
  it("user can login", async function (done) {
    try {
      const res = await client.login("admin@admin.com", "admin");
      expect(res.status).toBe(200);
      expect(res.statusText).toBe("OK");

      expect(res.data.expiresAt).exist;
      expect(res.data.token).exist;

      done();
    } catch (err) {
      done(err);
    }
  });
});

describe("POST /api/v1/users/logout", function () {
  it("user can logout", async function (done) {
    try {
      const myclient = getClientV1(config.BASE_URL);

      const res = await myclient.login("admin@admin.com", "admin");
      expect(res.status).toBe(200);
      expect(res.statusText).toBe("OK");

      const res2 = await myclient.logout();
      expect(res2.status).toBe(204);
      expect(res2.statusText).toBe("No Content");

      // Try to get self again
      try {
        const res3 = await myclient.self();
        expect(res3.status).toBe(401);
        expect(res3.statusText).toBe("Unauthorized");
      } catch (e) {
        if (axios.isAxiosError(e)) {
          expect(e.response.status).toBe(401);
          done();
        } else {
          done(e);
        }
      }

      done();
    } catch (err) {
      done(err);
    }
  });
});

describe("GET /api/v1/users/self", function () {
  it("user can access basic self details", async function (done) {
    try {
      const res = await client.self();
      expect(res.status).toBe(200);
      expect(res.statusText).toBe("OK");

      expect(res.data.item.id).exist;
      expect(res.data.item.name).toBe("Admin");
      expect(res.data.item.email).toBe("admin@admin.com");

      done();
    } catch (err) {
      done(err);
    }
  });
});
