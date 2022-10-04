package oauth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/authhandler"
)

// TODO: rename


func AuthCodeURL(config *oauth2.Config, opts ...oauth2.AuthCodeOption) (string, string) {
	state := randomString(16)
	return config.AuthCodeURL(state, opts...), state
}

func StateAndPKCE() (string, *authhandler.PKCEParams) {
	verifier := randomString(64)
	verifierSha256 := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(verifierSha256[:])
	return randomString(16), &authhandler.PKCEParams{
		Challenge:       challenge,
		ChallengeMethod: "S256",
		Verifier:        "..",
	}
}

func AuthCodeURLWithPKCE(config *oauth2.Config,
	opts ...oauth2.AuthCodeOption) (string, string, string) { // TODO: this may not be needed
	state := randomString(16)
	verifier := randomString(64)
	verifierSha256 := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(verifierSha256[:])
	return config.AuthCodeURL(state, append(opts,
			oauth2.SetAuthURLParam("code_challenge", challenge),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"))...),
		state, verifier
}

// Exchange auth code for access token verified by a PKCE challenge
func ExchangeWithPKCE(ctx context.Context, config *oauth2.Config, code, verifier string) (*oauth2.Token, error) {
	return config.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", verifier))
}

func randomString(size int) string {
	rand.Seed(time.Now().UnixNano())
	raw := make([]byte, size)
	_, _ = rand.Read(raw)
	return base64.RawURLEncoding.EncodeToString(raw)
}


