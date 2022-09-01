export const prefix = '/api/v1';

export type QueryValue =
	| string
	| string[]
	| number
	| number[]
	| boolean
	| null
	| undefined;

export function UrlBuilder(
	rest: string,
	params: Record<string, QueryValue> = {}
): string {
	// we use a stub base URL to leverage the URL class
	const url = new URL(prefix + rest, 'http://localhost.com');

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
