import { CookieRef } from "nuxt/dist/app/composables";
import { PublicApi } from "~~/lib/api/public";
import { UserOut } from "~~/lib/api/types/data-contracts";
import { UserClient } from "~~/lib/api/user";

export interface IAuthContext {
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
  login(api: PublicApi, email: string, password: string, stayLoggedIn: boolean): ReturnType<PublicApi["login"]>;
}

class AuthContext implements IAuthContext {
  // eslint-disable-next-line no-use-before-define
  private static _instance?: AuthContext;

  private static readonly cookieTokenKey = "hb.auth.token";
  private static readonly cookieExpiresAtKey = "hb.auth.expires_at";
  private static readonly cookieAttachmentTokenKey = "hb.auth.attachment_token";

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

  private constructor(token: string, expiresAt: string, attachmentToken: string) {
    this._token = useCookie(token);
    this._expiresAt = useCookie(expiresAt);
    this._attachmentToken = useCookie(attachmentToken);
  }

  static get instance() {
    if (!this._instance) {
      this._instance = new AuthContext(
        AuthContext.cookieTokenKey,
        AuthContext.cookieExpiresAtKey,
        AuthContext.cookieAttachmentTokenKey
      );
    }

    return this._instance;
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
    return !!this._token.value && !this.isExpired();
  }

  invalidateSession() {
    this.user = undefined;

    // Delete the cookies
    this._token.value = null;
    this._expiresAt.value = null;
    this._attachmentToken.value = null;

    navigateTo("/");
    console.log("Session invalidated");
  }

  async login(api: PublicApi, email: string, password: string, stayLoggedIn: boolean) {
    const r = await api.login(email, password, stayLoggedIn);

    if (!r.error) {
      this._token.value = r.data.token;
      this._expiresAt.value = r.data.expiresAt as string;
      this._attachmentToken.value = r.data.attachmentToken;
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
  return AuthContext.instance;
}
