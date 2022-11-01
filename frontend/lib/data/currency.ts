export type Codes = "USD" | "EUR" | "GBP" | "JPY" | "ZAR" | "AUD";

export type Currency = {
  code: Codes;
  local: string;
  symbol: string;
  name: string;
};

export const currencies: Currency[] = [
  {
    code: "USD",
    local: "en-US",
    symbol: "$",
    name: "US Dollar",
  },
  {
    code: "EUR",
    local: "de-DE",
    symbol: "€",
    name: "Euro",
  },
  {
    code: "GBP",
    local: "en-GB",
    symbol: "£",
    name: "British Pound",
  },
  {
    code: "JPY",
    local: "ja-JP",
    symbol: "¥",
    name: "Japanese Yen",
  },
  {
    code: "ZAR",
    local: "en-ZA",
    symbol: "R",
    name: "South African Rand",
  },
  {
    code: "AUD",
    local: "en-AU",
    symbol: "$",
    name: "Australian Dollar",
  },
];
