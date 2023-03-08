package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type serverResponse struct {
	Host    string            `json:"host"`
	Headers map[string]string `json:"headers"`
}

type ServerCredentials struct{}

func (c ServerCredentials) Name() string {
	return "auth-server"
}

func (c ServerCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.AuthServerUrl == "" {
		return nil, nil
	}

	resp, err := makeRequest(cfg.AuthServerUrl)
	if err != nil {
		return nil, nil
	}

	if resp == nil {
		return nil, nil
	}
	cfg.Host = resp.Host

	return func(r *http.Request) error {
		resp, err := makeRequest(cfg.AuthServerUrl)
		if err != nil {
			return err
		}
		if resp == nil {
			return fmt.Errorf("no response from auth server")
		}

		for k, v := range resp.Headers {
			r.Header.Set(k, v)
		}

		return nil
	}, nil
}

func makeRequest(serverUrl string) (*serverResponse, error) {
	resp, err := http.Get(serverUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response serverResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
