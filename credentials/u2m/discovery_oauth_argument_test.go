package u2m

import (
	"context"
	"testing"
)

func TestDiscoveryOAuthArgument_GetCacheKey(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("my-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	got := arg.GetCacheKey()
	if got != "my-profile" {
		t.Errorf("GetCacheKey(): want %q, got %q", "my-profile", got)
	}
}

func TestDiscoveryOAuthArgument_SetAndGetDiscoveredHost(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("test-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	arg.SetDiscoveredHost("https://adb-123.azuredatabricks.net")
	got := arg.GetDiscoveredHost()
	if got != "https://adb-123.azuredatabricks.net" {
		t.Errorf("GetDiscoveredHost(): want %q, got %q", "https://adb-123.azuredatabricks.net", got)
	}
}

func TestDiscoveryOAuthArgument_ZeroValueHasEmptyDiscoveredHost(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("test-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	got := arg.GetDiscoveredHost()
	if got != "" {
		t.Errorf("GetDiscoveredHost(): want empty string, got %q", got)
	}
}

func TestDiscoveryOAuthArgument_ConstructorRejectsEmptyProfile(t *testing.T) {
	_, err := NewBasicDiscoveryOAuthArgument("")
	if err == nil {
		t.Fatal("NewBasicDiscoveryOAuthArgument(): want error for empty profile, got nil")
	}
}

func TestDiscoveryOAuthArgument_NewPersistentAuthAcceptsIt(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("discovery-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(context.Background(), WithOAuthArgument(arg))
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()
}

type unsupportedOAuthArgument struct{}

func (u unsupportedOAuthArgument) GetCacheKey() string { return "unsupported" }

func TestDiscoveryOAuthArgument_ValidateArgRejectsUnknownTypes(t *testing.T) {
	_, err := NewPersistentAuth(context.Background(), WithOAuthArgument(unsupportedOAuthArgument{}))
	if err == nil {
		t.Fatal("NewPersistentAuth(): want error for unsupported OAuthArgument type, got nil")
	}
}
