package u2m

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const loginDatabricksHost = "https://login.databricks.com"

// DeriveHostFromIssuer extracts the workspace host (scheme + host) from an
// issuer URL returned in the OAuth callback iss parameter.
//
// Examples:
//
//	"https://adb-xxx.azuredatabricks.net/oidc"           -> "https://adb-xxx.azuredatabricks.net"
//	"https://nike.databricks.com/oidc/accounts/xxx"      -> "https://nike.databricks.com"
func DeriveHostFromIssuer(issuer string) (string, error) {
	if issuer == "" {
		return "", fmt.Errorf("issuer must not be empty")
	}
	u, err := url.Parse(issuer)
	if err != nil {
		return "", fmt.Errorf("parsing issuer URL %q: %w", issuer, err)
	}
	// Allow http for localhost (consistent with validateHost in workspace_oauth_argument.go).
	if u.Scheme == "http" && strings.HasPrefix(u.Host, "127.0.0.1") {
		// ok
	} else if u.Scheme != "https" {
		return "", fmt.Errorf("issuer must use https scheme: %q", issuer)
	}
	if u.Host == "" {
		return "", fmt.Errorf("issuer must have a non-empty host: %q", issuer)
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host), nil
}

// DeriveTokenEndpoint derives the token endpoint from an issuer URL by
// appending /v1/token to the issuer path.
//
// Example:
//
//	"https://adb-xxx.net/oidc" -> "https://adb-xxx.net/oidc/v1/token"
func DeriveTokenEndpoint(issuer string) string {
	return strings.TrimRight(issuer, "/") + "/v1/token"
}

// BuildDiscoveryAuthorizeURL builds the login.databricks.com URL that initiates
// the discovery OAuth flow. The OIDC authorize path with all OAuth query params
// is URL-encoded as the destination_url parameter.
func BuildDiscoveryAuthorizeURL(redirectAddr, state string, pkce PKCEParams, scopes []string) string {
	// Build the nested OIDC authorize path with query parameters.
	authParams := url.Values{}
	authParams.Set("client_id", appClientID)
	authParams.Set("redirect_uri", "http://"+redirectAddr)
	authParams.Set("response_type", "code")
	authParams.Set("scope", strings.Join(scopes, " "))
	authParams.Set("state", state)
	authParams.Set("code_challenge", pkce.Challenge)
	authParams.Set("code_challenge_method", pkce.ChallengeMethod)
	destinationURL := "/oidc/v1/authorize?" + authParams.Encode()

	// Wrap the authorize path as the destination_url query parameter on
	// login.databricks.com.
	topParams := url.Values{}
	topParams.Set("destination_url", destinationURL)
	return loginDatabricksHost + "/?" + topParams.Encode()
}

// PKCEParams holds the PKCE challenge parameters used to build the discovery
// authorize URL. This mirrors authhandler.PKCEParams but is used directly so
// the caller does not need to import authhandler.
type PKCEParams struct {
	Challenge       string
	ChallengeMethod string
	Verifier        string
}

// discoveryTokenSource handles the OAuth PKCE flow for login.databricks.com
// discovery login. Unlike the standard flow, the token endpoint is not known
// until the callback provides the iss parameter identifying the workspace.
type discoveryTokenSource struct {
	pa *PersistentAuth
}

// challenge initiates the discovery OAuth flow through login.databricks.com.
// It builds a custom authorize URL, opens the browser, waits for the callback,
// derives the workspace host and token endpoint from the iss parameter, and
// exchanges the authorization code for tokens.
func (d *discoveryTokenSource) challenge() error {
	cb, err := d.pa.newCallbackServer()
	if err != nil {
		return fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()

	state, authPKCE, err := d.pa.stateAndPKCE()
	if err != nil {
		return fmt.Errorf("state and pkce: %w", err)
	}

	// Determine scopes: same logic as oauth2Config().
	scopes := d.pa.scopes
	if len(scopes) == 0 {
		scopes = []string{"all-apis"}
	}
	if !d.pa.disableOfflineAccess {
		scopes = append([]string{"offline_access"}, scopes...)
	}

	pkce := PKCEParams{
		Challenge:       authPKCE.Challenge,
		ChallengeMethod: authPKCE.ChallengeMethod,
		Verifier:        authPKCE.Verifier,
	}
	authorizeURL := BuildDiscoveryAuthorizeURL(d.pa.redirectAddr, state, pkce, scopes)

	// Use cb.Handler to open the browser and wait for the callback.
	code, returnedState, err := cb.Handler(authorizeURL)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	issuer := cb.Issuer()

	// Validate state matches what we generated.
	if returnedState != state {
		return fmt.Errorf("state mismatch: expected %q, got %q", state, returnedState)
	}

	// Derive host and token endpoint from the issuer.
	host, err := DeriveHostFromIssuer(issuer)
	if err != nil {
		return fmt.Errorf("deriving host from issuer: %w", err)
	}
	tokenEndpoint := DeriveTokenEndpoint(issuer)

	// Exchange authorization code for tokens.
	cfg := &oauth2.Config{
		ClientID: appClientID,
		Endpoint: oauth2.Endpoint{
			TokenURL:  tokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: "http://" + d.pa.redirectAddr,
	}
	ctx := d.pa.setOAuthContext(d.pa.ctx)
	token, err := cfg.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", pkce.Verifier))
	if err != nil {
		return fmt.Errorf("token exchange: %w", err)
	}

	// Store discovered host on the argument.
	if discoveryArg, ok := d.pa.oAuthArgument.(DiscoveryOAuthArgument); ok {
		discoveryArg.SetDiscoveredHost(host)
	}

	// Cache the token.
	if err := d.pa.cache.Store(d.pa.oAuthArgument.GetCacheKey(), token); err != nil {
		return fmt.Errorf("storing token: %w", err)
	}
	return nil
}
