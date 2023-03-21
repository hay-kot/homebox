export type Codes =
  | "AUD"
  | "BGN"
  | "CHF"
  | "CZK"
  | "DKK"
  | "EUR"
  | "GBP"
  | "INR"
  | "JPY"
  | "NOK"
  | "NZD"
  | "PLN"
  | "RMB"
  | "RON"
  | "SEK"
  | "TRY"
  | "USD"
  | "ZAR";

export type Currency = {
  code: Codes;
  local: string;
  symbol: string;
  name: string;
};

export const currencies: Currency[] = [
  {
    code: "AUD",
    local: "en-AU",
    symbol: "$",
    name: "Australian Dollar",
  },
  {
    code: "GBP",
    local: "en-GB",
    symbol: "£",
    name: "British Pound",
  },
  {
    code: "RMB",
    local: "zh-CN",
    symbol: "¥",
    name: "Chinese Yuan",
  },
  {
    code: "DKK",
    local: "da-DK",
    symbol: "kr",
    name: "Danish Krone",
  },
  {
    code: "EUR",
    local: "de-DE",
    symbol: "€",
    name: "Euro",
  },
  {
    code: "INR",
    local: "en-IN",
    symbol: "₹",
    name: "Indian Rupee",
  },
  {
    code: "JPY",
    local: "ja-JP",
    symbol: "¥",
    name: "Japanese Yen",
  },
  {
    code: "NOK",
    local: "nb-NO",
    symbol: "kr",
    name: "Norwegian Krone",
  },
  {
    code: "NVD",
    local: "en-NZ",
    symbol: "NZ$",
    name: "New Zealand Dollar",
  },
  {
    code: "PLN",
    local: "pl-PL",
    symbol: "zł",
    name: "Polish Zloty",
  },
  {
    code: "RON",
    local: "ro-RO",
    symbol: "lei",
    name: "Romanian Leu",
  },
  {
    code: "ZAR",
    local: "en-ZA",
    symbol: "R",
    name: "South African Rand",
  },
  {
    code: "SEK",
    local: "sv-SE",
    symbol: "kr",
    name: "Swedish Krona",
  },
  {
    code: "TRY",
    local: "tr-TR",
    symbol: "₺",
    name: "Turkish Lira",
  },
  {
    code: "USD",
    local: "en-US",
    symbol: "$",
    name: "US Dollar",
  },
  {
    code: "BGN",
    local: "bg-BG",
    symbol: "lv",
    name: "Bulgarian lev",
  },
  {
    code: "CHF",
    local: "de-CH",
    symbol: "chf",
    name: "Swiss Francs",
  },
  {
    code: "CZK",
    local: "cs-CZ",
    symbol: "Kč",
    name: "Czech Koruna",
  },
];
