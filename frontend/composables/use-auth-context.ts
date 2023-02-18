import { CookieRef } from "nuxt/dist/app/composables";
import { PublicApi } from "~~/lib/api/public";
import { UserOut } from "~~/lib/api/types/data-contracts";
import { UserClient } from "~~/lib/api/user";

export interface IAuthContext {
  self?: UserOut;
  get token(): string | null;
  get expiresAt(): string | null;
  get attachmentToken(): string | null;

  /**
   * The current user object for the session. This is undefined if the session is not authorized.
   */
  user?: UserOut;

  /**
   * Returns true if the session is expired.
   */
  isExpired(): boolean;

  /**
   * Returns true if the session is authorized.
   */
  isAuthorized(): boolean;

  /**
   * Invalidates the session by removing the token and the expiresAt.
   */
  invalidateSession(): void;

  /**
   * Logs out the user and calls the invalidateSession method.
   */
  logout(api: UserClient): ReturnType<UserClient["user"]["logout"]>;

  /**
   * Logs in the user and sets the authorization context via cookies
   */
  login(api: PublicApi, email: string, password: string): ReturnType<PublicApi["login"]>;
}

class AuthContext implements IAuthContext {
  user?: UserOut;
  private _token: CookieRef<string | null>;
  private _expiresAt: CookieRef<string | null>;
  private _attachmentToken: CookieRef<string | null>;

  get token() {
    return this._token.value;
  }

  get expiresAt() {
    return this._expiresAt.value;
  }

  get attachmentToken() {
    return this._attachmentToken.value;
  }

  constructor(
    token: CookieRef<string | null>,
    expiresAt: CookieRef<string | null>,
    attachmentToken: CookieRef<string | null>
  ) {
    this._token = token;
    this._expiresAt = expiresAt;
    this._attachmentToken = attachmentToken;
  }

  isExpired() {
    const expiresAt = this.expiresAt;
    if (expiresAt === null) {
      return true;
    }

    const expiresAtDate = new Date(expiresAt);
    const now = new Date();

    return now.getTime() > expiresAtDate.getTime();
  }

  isAuthorized() {
    return this._token.value !== null && !this.isExpired();
  }

  invalidateSession() {
    this.user = undefined;
    this._token.value = null;
    this._expiresAt.value = null;
    this._attachmentToken.value = null;
  }

  async login(api: PublicApi, email: string, password: string) {
    const r = await api.login(email, password);

    if (!r.error) {
      this._token.value = r.data.token;
      this._expiresAt.value = r.data.expiresAt as string;
      this._attachmentToken.value = r.data.attachmentToken;

      console.log({
        token: this._token.value,
        expiresAt: this._expiresAt.value,
        attachmentToken: this._attachmentToken.value,
      });
    }

    return r;
  }

  async logout(api: UserClient) {
    const r = await api.user.logout();

    if (!r.error) {
      this.invalidateSession();
    }

    return r;
  }
}

export function useAuthContext(): IAuthContext {
  const tokenCookie = useCookie("hb.auth.token");
  const expiresAtCookie = useCookie("hb.auth.expires_at");
  const attachmentTokenCookie = useCookie("hb.auth.attachment_token");

  return new AuthContext(tokenCookie, expiresAtCookie, attachmentTokenCookie);
}
