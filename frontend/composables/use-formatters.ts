const cache = {
  currency: "",
};

export function resetCurrency() {
  cache.currency = "";
}

export async function useFormatCurrency() {
  if (cache.currency === "") {
    const client = useUserApi();

    const { data: group } = await client.group.get();

    if (group) {
      cache.currency = group.currency;
    }
  }

  return (value: number | string) => fmtCurrency(value, cache.currency);
}

export type DateTimeFormat = "relative" | "long" | "short" | "human";
export type DateTimeType = "date" | "time" | "datetime";

function ordinalIndicator(num: number) {
  if (num > 3 && num < 21) return "th";
  switch (num % 10) {
    case 1:
      return "st";
    case 2:
      return "nd";
    case 3:
      return "rd";
    default:
      return "th";
  }
}

export function fmtDate(value: string | Date, fmt: DateTimeFormat = "human"): string {
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  const dt = typeof value === "string" ? new Date(value) : value;
  if (!dt) {
    return "";
  }

  if (!validDate(dt)) {
    return "";
  }

  switch (fmt) {
    case "relative":
      return useTimeAgo(dt).value + useDateFormat(dt, " (YYYY-MM-DD)").value;
    case "long":
      return useDateFormat(dt, "YYYY-MM-DD (dddd)").value;
    case "short":
      return useDateFormat(dt, "YYYY-MM-DD").value;
    case "human":
      // January 1st, 2021
      return `${months[dt.getMonth()]} ${dt.getDate()}${ordinalIndicator(dt.getDate())}, ${dt.getFullYear()}`;
    default:
      return "";
  }
}
