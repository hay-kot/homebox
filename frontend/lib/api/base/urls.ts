const parts = {
  host: 'http://localhost.com',
  prefix: '/api/v1',
};

export function OverrideParts(host: string, prefix: string) {
  parts.host = host;
  parts.prefix = prefix;
}

export type QueryValue = string | string[] | number | number[] | boolean | null | undefined;

export function UrlBuilder(rest: string, params: Record<string, QueryValue> = {}): string {
  const url = new URL(parts.prefix + rest, parts.host);

  for (const [key, value] of Object.entries(params)) {
    if (Array.isArray(value)) {
      for (const item of value) {
        url.searchParams.append(key, String(item));
      }
    } else {
      url.searchParams.append(key, String(value));
    }
  }

  // we return the path only, without the base URL
  return url.toString().replace('http://localhost.com', '');
}
