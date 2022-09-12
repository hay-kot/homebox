import { describe, expect, test } from "vitest";
import { hasKey, parseDate } from "./base-api";

describe("hasKey works as expected", () => {
  test("hasKey returns true if the key exists", () => {
    const obj = { createdAt: "2021-01-01" };
    expect(hasKey(obj, "createdAt")).toBe(true);
  });

  test("hasKey returns false if the key does not exist", () => {
    const obj = { createdAt: "2021-01-01" };
    expect(hasKey(obj, "updatedAt")).toBe(false);
  });
});

describe("parseDate should work as expected", () => {
  test("parseDate should set defaults", () => {
    const obj = { createdAt: "2021-01-01", updatedAt: "2021-01-01" };
    const result = parseDate(obj);
    expect(result.createdAt).toBeInstanceOf(Date);
    expect(result.updatedAt).toBeInstanceOf(Date);
  });

  test("parseDate should set passed in types", () => {
    const obj = { key1: "2021-01-01", key2: "2021-01-01" };
    const result = parseDate(obj, ["key1", "key2"]);
    expect(result.key1).toBeInstanceOf(Date);
    expect(result.key2).toBeInstanceOf(Date);
  });
});
