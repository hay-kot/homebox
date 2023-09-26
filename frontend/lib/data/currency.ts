export type Codes =
  | "AED"
  | "AUD"
  | "BGN"
  | "BRL"
  | "CAD"
  | "CHF"
  | "CZK"
  | "DKK"
  | "EUR"
  | "GBP"
  | "HKD"
  | "IDR"
  | "INR"
  | "JPY"
  | "KRW"
  | "MXN"
  | "MYR"
  | "NOK"
  | "NZD"
  | "PLN"
  | "RMB"
  | "RUB"
  | "RON"
  | "SAR"
  | "SEK"
  | "SGD"
  | "THB"
  | "TRY"
  | "USD"
  | "XAG"
  | "XAU"
  | "ZAR";

export type Currency = {
  code: Codes;
  local: string;
  symbol: string;
  name: string;
};

export const currencies: Currency[] = [
  { code: "AED", local: "United Arab Emirates", symbol: "د.إ", name: "United Arab Emirates Dirham" },
  { code: "AUD", local: "Australia", symbol: "A$", name: "Australian Dollar" },
  { code: "BGN", local: "bg-BG", symbol: "lv", name: "Bulgarian lev" },
  { code: "BRL", local: "Brazil", symbol: "R$", name: "Brazilian Real" },
  { code: "CAD", local: "Canada", symbol: "C$", name: "Canadian Dollar" },
  { code: "CHF", local: "Switzerland", symbol: "CHF", name: "Swiss Franc" },
  { code: "CZK", local: "cs-CZ", symbol: "Kč", name: "Czech Koruna" },
  { code: "DKK", local: "da-DK", symbol: "kr", name: "Danish Krone" },
  { code: "EUR", local: "Eurozone", symbol: "€", name: "Euro" },
  { code: "GBP", local: "United Kingdom", symbol: "£", name: "British Pound Sterling" },
  { code: "HKD", local: "Hong Kong", symbol: "HK$", name: "Hong Kong Dollar" },
  { code: "IDR", local: "Indonesia", symbol: "Rp", name: "Indonesian Rupiah" },
  { code: "INR", local: "India", symbol: "₹", name: "Indian Rupee" },
  { code: "JPY", local: "Japan", symbol: "¥", name: "Japanese Yen" },
  { code: "KRW", local: "South Korea", symbol: "₩", name: "South Korean Won" },
  { code: "MXN", local: "Mexico", symbol: "Mex$", name: "Mexican Peso" },
  { code: "MYR", local: "Malaysia", symbol: "RM", name: "Malaysian Ringgit" },  
  { code: "NOK", local: "Norway", symbol: "kr", name: "Norwegian Krone" },
  { code: "NZD", local: "New Zealand", symbol: "NZ$", name: "New Zealand Dollar" },
  { code: "PLN", local: "Poland", symbol: "zł", name: "Polish Zloty" },
  { code: "RMB", local: "zh-CN", symbol: "¥", name: "Chinese Yuan" },
  { code: "RON", local: "ro-RO", symbol: "lei", name: "Romanian Leu" },
  { code: "RUB", local: "Russia", symbol: "₽", name: "Russian Ruble" },
  { code: "SAR", local: "Saudi Arabia", symbol: "﷼", name: "Saudi Riyal" },
  { code: "SEK", local: "Sweden", symbol: "kr", name: "Swedish Krona" },
  { code: "SGD", local: "Singapore", symbol: "S$", name: "Singapore Dollar" },
  { code: "THB", local: "Thailand", symbol: "฿", name: "Thai Baht" },
  { code: "TRY", local: "Turkey", symbol: "₺", name: "Turkish Lira" },
  { code: "USD", local: "United States", symbol: "$", name: "United States Dollar" },
  { code: "XAG", local: "Global", symbol: "XAG", name: "Silver Troy Ounce" },
  { code: "XAU", local: "Global", symbol: "XAU", name: "Gold Troy Ounce" },
  { code: "ZAR", local: "South Africa", symbol: "R", name: "South African Rand" },
];
