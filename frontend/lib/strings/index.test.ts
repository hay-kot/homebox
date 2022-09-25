import { describe, test, expect } from "vitest";
import { titlecase, capitalize, truncate } from ".";

describe("title case tests", () => {
  test("should return the same string if it's already title case", () => {
    expect(titlecase("Hello World")).toBe("Hello World");
  });

  test("should title case a lower case word", () => {
    expect(titlecase("hello")).toBe("Hello");
  });

  test("should title case a sentence", () => {
    expect(titlecase("hello world")).toBe("Hello World");
  });

  test("should title case a sentence with multiple words", () => {
    expect(titlecase("hello world this is a test")).toBe("Hello World This Is A Test");
  });
});

describe("capitilize tests", () => {
  test("should return the same string if it's already capitalized", () => {
    expect(capitalize("Hello")).toBe("Hello");
  });

  test("should capitalize a lower case word", () => {
    expect(capitalize("hello")).toBe("Hello");
  });

  test("should capitalize a sentence", () => {
    expect(capitalize("hello world")).toBe("Hello world");
  });

  test("should capitalize a sentence with multiple words", () => {
    expect(capitalize("hello world this is a test")).toBe("Hello world this is a test");
  });
});

describe("truncase tests", () => {
  test("should return the same string if it's already truncated", () => {
    expect(truncate("Hello", 5)).toBe("Hello");
  });

  test("should truncate a lower case word", () => {
    expect(truncate("hello", 3)).toBe("hel...");
  });

  test("should truncate a sentence", () => {
    expect(truncate("hello world", 5)).toBe("hello...");
  });

  test("should truncate a sentence with multiple words", () => {
    expect(truncate("hello world this is a test", 10)).toBe("hello worl...");
  });
});
