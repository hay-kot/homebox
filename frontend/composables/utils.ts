export function validDate(dt: Date | string | null | undefined): boolean {
  if (!dt) {
    return false;
  }

  // If it's a string, try to parse it
  if (typeof dt === "string") {
    const parsed = new Date(dt);
    if (isNaN(parsed.getTime())) {
      return false;
    }
  }

  // If it's a date, check if it's valid
  if (dt instanceof Date) {
    if (dt.getFullYear() < 1000) {
      return false;
    }
  }

  return true;
}

export function fmtCurrency(value: number | string, currency = "USD", locale = "en-Us"): string {
  if (typeof value === "string") {
    value = parseFloat(value);
  }

  const formatter = new Intl.NumberFormat(locale, {
    style: "currency",
    currency,
    minimumFractionDigits: 2,
  });
  return formatter.format(value);
}
