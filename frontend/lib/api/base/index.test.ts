import { describe, expect, it } from 'vitest';
import { UrlBuilder } from '.';

describe('UrlBuilder', () => {
	it('basic query parameter', () => {
		const result = UrlBuilder('/test', { a: 'b' });
		expect(result).toBe('/api/v1/test?a=b');
	});

	it('multiple query parameters', () => {
		const result = UrlBuilder('/test', { a: 'b', c: 'd' });
		expect(result).toBe('/api/v1/test?a=b&c=d');
	});

	it('no query parameters', () => {
		const result = UrlBuilder('/test');
		expect(result).toBe('/api/v1/test');
	});

	it('list-like query parameters', () => {
		const result = UrlBuilder('/test', { a: ['b', 'c'] });
		expect(result).toBe('/api/v1/test?a=b&a=c');
	});
});
