export type Currency = {
  code: string;
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
];
