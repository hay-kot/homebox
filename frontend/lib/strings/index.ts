export function titlecase(str: string) {
  return str
    .split(" ")
    .map(word => word[0].toUpperCase() + word.slice(1))
    .join(" ");
}

export function capitalize(str: string) {
  return str[0].toUpperCase() + str.slice(1);
}

export function truncate(str: string, length: number) {
  return str.length > length ? str.substring(0, length) + "..." : str;
}
