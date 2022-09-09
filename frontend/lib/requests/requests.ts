export enum Method {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
}

export type RequestInterceptor = (r: Response) => void;
export type ResponseInterceptor = (r: Response) => void;

export interface TResponse<T> {
  status: number;
  error: boolean;
  data: T;
  response: Response;
}

export type RequestArgs<T> = {
  url: string;
  body?: T;
  data?: FormData;
  headers?: Record<string, string>;
};

export class Requests {
  private baseUrl: string;
  private token: () => string;
  private headers: Record<string, string> = {};
  private responseInterceptors: ResponseInterceptor[] = [];

  addResponseInterceptor(interceptor: ResponseInterceptor) {
    this.responseInterceptors.push(interceptor);
  }

  private callResponseInterceptors(response: Response) {
    this.responseInterceptors.forEach(i => i(response));
  }

  private url(rest: string): string {
    return this.baseUrl + rest;
  }

  constructor(baseUrl: string, token: string | (() => string) = '', headers: Record<string, string> = {}) {
    this.baseUrl = baseUrl;
    this.token = typeof token === 'string' ? () => token : token;
    this.headers = headers;
  }

  public get<T>(args: RequestArgs<T>): Promise<TResponse<T>> {
    return this.do<T>(Method.GET, args);
  }

  public post<T, U>(args: RequestArgs<T>): Promise<TResponse<U>> {
    return this.do<U>(Method.POST, args);
  }

  public put<T, U>(args: RequestArgs<T>): Promise<TResponse<U>> {
    return this.do<U>(Method.PUT, args);
  }

  public delete<T>(args: RequestArgs<T>): Promise<TResponse<T>> {
    return this.do<T>(Method.DELETE, args);
  }

  private methodSupportsBody(method: Method): boolean {
    return method === Method.POST || method === Method.PUT;
  }

  private async do<T>(method: Method, rargs: RequestArgs<unknown>): Promise<TResponse<T>> {
    const payload: RequestInit = {
      method,
      headers: {
        ...rargs.headers,
        ...this.headers,
      },
    };

    const token = this.token();
    if (token !== '' && payload.headers !== undefined) {
      payload.headers['Authorization'] = token;
    }

    if (this.methodSupportsBody(method)) {
      if (rargs.data) {
        payload.body = rargs.data;
      } else {
        payload.headers['Content-Type'] = 'application/json';
        payload.body = JSON.stringify(rargs.body);
      }
    }

    const response = await fetch(this.url(rargs.url), payload);
    this.callResponseInterceptors(response);

    const data: T = await (async () => {
      if (response.status === 204) {
        return {} as T;
      }

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
