import { describe, it, expect } from "vitest";
import { titlecase, capitalize, truncate } from ".";

describe("title case tests", () => {
  it("should return the same string if it's already title case", () => {
    expect(titlecase("Hello World")).toBe("Hello World");
  });

  it("should title case a lower case word", () => {
    expect(titlecase("hello")).toBe("Hello");
  });

  it("should title case a sentence", () => {
    expect(titlecase("hello world")).toBe("Hello World");
  });

  it("should title case a sentence with multiple words", () => {
    expect(titlecase("hello world this is a test")).toBe("Hello World This Is A Test");
  });
});

describe("capitilize tests", () => {
  it("should return the same string if it's already capitalized", () => {
    expect(capitalize("Hello")).toBe("Hello");
  });

  it("should capitalize a lower case word", () => {
    expect(capitalize("hello")).toBe("Hello");
  });

  it("should capitalize a sentence", () => {
    expect(capitalize("hello world")).toBe("Hello world");
  });

  it("should capitalize a sentence with multiple words", () => {
    expect(capitalize("hello world this is a test")).toBe("Hello world this is a test");
  });
});

describe("truncase tests", () => {
  it("should return the same string if it's already truncated", () => {
    expect(truncate("Hello", 5)).toBe("Hello");
  });

  it("should truncate a lower case word", () => {
    expect(truncate("hello", 3)).toBe("hel...");
  });

  it("should truncate a sentence", () => {
    expect(truncate("hello world", 5)).toBe("hello...");
  });

  it("should truncate a sentence with multiple words", () => {
    expect(truncate("hello world this is a test", 10)).toBe("hello worl...");
  });
});
