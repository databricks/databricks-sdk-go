package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
	"github.com/databricks/databricks-sdk-go/databricks/oauth"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"golang.org/x/oauth2"
)

var host string

const localServeAddr = "localhost:8020"

func main() {
	flag.StringVar(&host, "host", "", "Databricks Workspace Host")
	flag.Parse()

	host = ".."

	if host == "" {
		fmt.Printf("host is empty\n\n%s", flag.ErrHelp)
		os.Exit(1)
	}

	authConfig, err := (&databricks.Config{
		Host:        host,
		ClientID:    "databricks-cli",
		RedirectURL: "http://" + localServeAddr,
		Scopes:      []string{"jobs"},
	}).OAuthConfig()
	if err != nil {
		logger.Errorf("failed to init oauth helper: %s", err)
		return
	}

	tmpl := template.Must(template.ParseGlob("*.tmpl"))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("state") == "" {
			tmpl.ExecuteTemplate(w, "index.tmpl", nil)
			return
		}
		if r.FormValue("error_description") != "" {
			tmpl.ExecuteTemplate(w, "error.tmpl", r.FormValue("error_description"))
			return
		}
		stateCookie, _ := r.Cookie("DATABRICKS_OAUTH_STATE")
		if stateCookie == nil {
			tmpl.ExecuteTemplate(w, "error.tmpl", "missing state cookie")
			return
		}
		if r.FormValue("state") != stateCookie.Value {
			tmpl.ExecuteTemplate(w, "error.tmpl", "state mismatch")
			return
		}
		verifierCookie, _ := r.Cookie("DATABRICKS_OAUTH_VERIFIER")
		if verifierCookie == nil {
			tmpl.ExecuteTemplate(w, "error.tmpl", "missing verifier cookie")
			return
		}
		token, err := authConfig.Exchange(r.Context(), r.FormValue("code"), 
			oauth2.SetAuthURLParam("code_verifier", verifierCookie.Value))
		if err != nil {
			w.WriteHeader(400)
			tmpl.ExecuteTemplate(w, "error.tmpl", err)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "DATABRICKS_OAUTH_TOKEN",
			Value:   token.AccessToken,
			Expires: token.Expiry,
		})
	}))

	http.Handle("/oauth/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		state, pkce := oauth.StateAndPKCE()
		url := authConfig.AuthCodeURL(state, 
			oauth2.SetAuthURLParam("code_challenge", pkce.Challenge),
			oauth2.SetAuthURLParam("code_challenge_method", pkce.ChallengeMethod),
			oauth2.SetAuthURLParam("redirect_back", r.FormValue("redirect_back")))
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		http.SetCookie(w, &http.Cookie{
			Name:    "DATABRICKS_OAUTH_STATE",
			Value:   state,
			Expires: time.Now().Add(5 * time.Minute),
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "DATABRICKS_OAUTH_VERIFIER",
			Value:   pkce.Verifier,
			Expires: time.Now().Add(5 * time.Minute),
		})
	}))

	http.Handle("/jobs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("DATABRICKS_OAUTH_TOKEN")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/oauth/login?redirect_back=/jobs", http.StatusTemporaryRedirect)
			return
		}
		wsc := workspaces.New(&databricks.Config{
			Host:  host,
			Token: tokenCookie.Value, // TODO: make real oauth
		})
		allJobs, err := wsc.Jobs.List(r.Context(), jobs.ListRequest{})
		if err != nil {
			w.WriteHeader(500)
			tmpl.ExecuteTemplate(w, "error.tmpl", err)
			return
		}
		tmpl.ExecuteTemplate(w, "jobs.tmpl", allJobs)
	}))

	logger.Infof("listening on http://%s", localServeAddr)
	err = http.ListenAndServe(localServeAddr, http.DefaultServeMux)
	if err != nil {
		logger.Errorf("listen and serve: %s", err)
	}
}
