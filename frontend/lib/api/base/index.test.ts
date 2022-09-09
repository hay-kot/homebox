import { describe, expect, it } from "vitest";
import { route } from ".";

describe("UrlBuilder", () => {
  it("basic query parameter", () => {
    const result = route("/test", { a: "b" });
    expect(result).toBe("/api/v1/test?a=b");
  });

  it("multiple query parameters", () => {
    const result = route("/test", { a: "b", c: "d" });
    expect(result).toBe("/api/v1/test?a=b&c=d");
  });

  it("no query parameters", () => {
    const result = route("/test");
    expect(result).toBe("/api/v1/test");
  });

  it("list-like query parameters", () => {
    const result = route("/test", { a: ["b", "c"] });
    expect(result).toBe("/api/v1/test?a=b&a=c");
  });
});
