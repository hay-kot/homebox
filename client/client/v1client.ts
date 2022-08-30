import axios, { Axios } from "axios";

interface Wrap<T> {
  item: T;
}

interface Status {
  status: string;
  message: string;
}

export interface ApiSummary {
  health: boolean;
  versions: string[];
  title: string;
  message: string;
}

export interface ApiToken {
  token: string;
  expiresAt: string;
}

export interface UserSelf {
  id: string;
  name: string;
  email: string;
  isSuperuser: boolean;
}

export class v1ApiClient {
  version: string;
  baseUrl: string;
  requests: Axios;

  token: string;
  expires: Date;

  constructor(baseUrl: string, version = "v1") {
    this.version = version;
    this.baseUrl = baseUrl;
    this.requests = axios.create({
      baseURL: `${this.baseUrl}/${this.version}`,
    });
  }

  v1(url: string) {
    return `${this.baseUrl}/api/v1${url}`;
  }

  api(url: string) {
    return `${this.baseUrl}/api${url}`;
  }

  setToken(token: string, expires: Date) {
    this.token = token;
    this.expires = expires;

    this.requests.defaults.headers.common["Authorization"] = token;
  }

  async login(username: string, password: string) {
    const response = await this.requests.post<ApiToken>(
      this.v1("/users/login"),
      {
        username,
        password,
      }
    );

    this.setToken(response.data.token, new Date(response.data.expiresAt));

    return response;
  }

  async logout() {
    const response = await this.requests.post<any>(this.v1("/users/logout"));

    if (response.status === 200) {
      this.setToken("", new Date());
    }

    return response;
  }

  async self() {
    return this.requests.get<Wrap<UserSelf>>(this.v1("/users/self"));
  }

  async status() {
    return this.requests.get<Wrap<Status>>(this.api("/status"));
  }
}
