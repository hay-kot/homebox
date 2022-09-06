export function truncate(str: string, length: number) {
  return str.length > length ? str.substring(0, length) + '...' : str;
}

export function capitalize(str: string) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}
