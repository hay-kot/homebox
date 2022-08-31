export enum Method {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
}

export interface TResponse<T> {
  status: number;
  error: boolean;
  data: T;
  response: Response;
}

export class Requests {
  private baseUrl: string;
  private token: () => string;
  private headers: Record<string, string> = {};
  private logger?: (response: Response) => void;

  private url(rest: string): string {
    return this.baseUrl + rest;
  }

  constructor(
    baseUrl: string,
    token: string | (() => string) = '',
    headers: Record<string, string> = {},
    logger?: (response: Response) => void
  ) {
    this.baseUrl = baseUrl;
    this.token = typeof token === 'string' ? () => token : token;
    this.headers = headers;
    this.logger = logger;
  }

  public get<T>(url: string): Promise<TResponse<T>> {
    return this.do<T>(Method.GET, url);
  }

  public post<T, U>(url: string, payload: T): Promise<TResponse<U>> {
    return this.do<U>(Method.POST, url, payload);
  }

  public put<T, U>(url: string, payload: T): Promise<TResponse<U>> {
    return this.do<U>(Method.PUT, url, payload);
  }

  public delete<T>(url: string): Promise<TResponse<T>> {
    return this.do<T>(Method.DELETE, url);
  }

  private methodSupportsBody(method: Method): boolean {
    return method === Method.POST || method === Method.PUT;
  }

  private async do<T>(method: Method, url: string, payload: Object = {}): Promise<TResponse<T>> {
    const args: RequestInit = {
      method,
      headers: {
        'Content-Type': 'application/json',
        ...this.headers,
      },
    };

    const token = this.token();
    if (token !== '' && args.headers !== undefined) {
      // @ts-expect-error -- headers is always defined at this point
      args.headers['Authorization'] = token;
    }

    if (this.methodSupportsBody(method)) {
      args.body = JSON.stringify(payload);
    }

    const response = await fetch(this.url(url), args);

    if (this.logger) {
      this.logger(response);
    }

    const data: T = await (async () => {
      try {
        return await response.json();
      } catch (e) {
        return {} as T;
      }
    })();

    return {
      status: response.status,
      error: !response.ok,
      data,
      response,
    };
  }
}
