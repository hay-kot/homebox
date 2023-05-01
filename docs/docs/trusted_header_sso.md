# Trusted Header SSO

## Overview
As an additional login method, authentication via special HTTP headers is possible. When using this method then a user is logged in based on data provided by HTTP headers. No login screen will be presented to enter email and password to log in.
This is useful when you plan to integrate Homebox into an environment where a single sign-on service (like Authelia or Authentik) is running.

### References:
Authelia: https://www.authelia.com/integration/trusted-header-sso/introduction/


## Usage
By default, this method is disabled. To use it, enable the variable `HBOX_OPTIONS_HEADER_SSO_ENABLED`.

When this is enabled and the login page is accessed, then the HTTP header named `Remote-Email` is checked. If this exists then the user is automatically logged in if a Homebox user with this email address already exists. If such a user does not exists and `HBOX_OPTIONS_HEADER_SSO_AUTOREGISTER` is enabled, then a user will automatically be registered with the name taken from HTTP header `Remote-Name`.

When if trusted header SSO is enabled but the HTTP header `Remote-Email` is not found or is empty, then the regular login page will be presented to the user and the regular password-base login can be used. So, basically the login flow is not altered in any way when enabling trusted header SSO as long as no special headers are found.

The name of the HTTP headers to use can be altered by setting the variables `HBOX_OPTIONS_HEADER_SSO_HEADER_EMAIL` and `HBOX_OPTIONS_HEADER_SSO_HEADER_NAME`. They default to `Remote-Email` and `Remote-Name`.

**Note:**
For security reasons, an IP address has to be configured that is allowed to send the trusted SSO headers. This is usually the IP address of your SSO proxy. This should make sure that not random clients can login by just sending those headers. The IP has to be set to the variable `HBOX_OPTIONS_HEADER_SSO_ALLOWED_IP`. If the SSO headers are received from a different IP than this configured one, then the authorization will be rejected.