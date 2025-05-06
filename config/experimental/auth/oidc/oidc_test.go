package oidc

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIDTokenSourceFn(t *testing.T) {
	wantToken := &IDToken{Value: "from-func"}
	wantErr := fmt.Errorf("test error")
	wantAud := "func-audience"
	wantCtx := context.Background()

	ts := IDTokenSourceFn(func(gotCtx context.Context, gotAud string) (*IDToken, error) {
		if gotCtx != wantCtx {
			t.Errorf("unexpected context: got %v, want %v", gotCtx, wantCtx)
		}
		if gotAud != wantAud {
			t.Errorf("unexpected audience: got %q, want %q", gotAud, wantAud)
		}
		return wantToken, wantErr
	})

	gotToken, gotErr := ts.IDToken(wantCtx, wantAud)

	if gotErr != wantErr {
		t.Errorf("IDToken() want error: %v, got error: %v", wantErr, gotErr)
	}
	if !cmp.Equal(gotToken, wantToken) {
		t.Errorf("IDToken() token = %v, want %v", gotToken, wantToken)
	}
}

func TestNewEnvIDTokenSource(t *testing.T) {
	testCases := []struct {
		desc     string
		envName  string
		envValue string
		audience string
		want     *IDToken
		wantErr  bool
	}{
		{
			desc:     "Success - variable set",
			envName:  "OIDC_TEST_TOKEN_SUCCESS",
			envValue: "test-token-123",
			audience: "test-audience-1",
			want:     &IDToken{Value: "test-token-123"},
			wantErr:  false,
		},
		{
			desc:     "Failure - variable not set",
			envName:  "OIDC_TEST_TOKEN_MISSING",
			envValue: "",
			audience: "test-audience-2",
			want:     nil,
			wantErr:  true,
		},
		{
			desc:     "Failure - variable set to empty string",
			envName:  "OIDC_TEST_TOKEN_EMPTY",
			envValue: "",
			audience: "test-audience-3",
			want:     nil,
			wantErr:  true,
		},
		{
			desc:     "Success - different variable name",
			envName:  "ANOTHER_OIDC_TOKEN",
			envValue: "another-token-456",
			audience: "test-audience-4",
			want:     &IDToken{Value: "another-token-456"},
			wantErr:  false,
		},
		{
			desc:     "Success - empty audience string",
			envName:  "OIDC_TEST_TOKEN_NO_AUDIENCE",
			envValue: "token-no-audience",
			audience: "",
			want:     &IDToken{Value: "token-no-audience"},
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Setenv(tc.envName, tc.envValue)

			ts := NewEnvIDTokenSource(tc.envName)
			got, gotErr := ts.IDToken(context.Background(), tc.audience)

			if tc.wantErr && gotErr == nil {
				t.Fatalf("IDToken() want error, got none")
			}
			if !tc.wantErr && gotErr != nil {
				t.Fatalf("IDToken() want no error, got error: %v", gotErr)
			}
			if !cmp.Equal(got, tc.want) {
				t.Errorf("IDToken() token = %v, want %v", got, tc.want)
			}
		})
	}
}
