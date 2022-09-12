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
