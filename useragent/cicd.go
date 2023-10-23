package useragent

import (
	"os"
	"sync"
)

type envVar struct {
	name          string
	expectedValue string
}

type cicdProvider struct {
	// The name of the CI/CD provider. This is the name included in the user
	// agent string.
	name string

	// The env vars that are expected to be set in the CI/CD provider's runner.
	envVars []envVar
}

// List of CI/CD providers and their env vars we can rely on to detect them.
func listCiCdProviders() []cicdProvider {
	return []cicdProvider{
		{"github", []envVar{{"GITHUB_ACTIONS", "true"}}},
		{"gitlab", []envVar{{"GITLAB_CI", "true"}}},
		{"jenkins", []envVar{{"JENKINS_URL", ""}}},
		{"azure-devops", []envVar{{"TF_BUILD", "True"}}},
		{"circle", []envVar{{"CIRCLECI", "true"}}},
		{"travis", []envVar{{"TRAVIS", "true"}}},
		{"bitbucket", []envVar{{"BITBUCKET_BUILD_NUMBER", ""}}},
		{"google-cloud-build", []envVar{{"PROJECT_ID", ""}, {"BUILD_ID", ""}, {"PROJECT_NUMBER", ""}, {"LOCATION", ""}}},
		{"aws-code-build", []envVar{{"CODEBUILD_BUILD_ARN", ""}}},
		{"tf-cloud", []envVar{{"TFC_RUN_ID", ""}}},
	}
}

// detect returns true if all env vars are set and have expected values.
func (p cicdProvider) detect() bool {
	os.Environ()
	for _, envVar := range p.envVars {
		v, ok := os.LookupEnv(envVar.name)
		if !ok {
			return false
		}
		if envVar.expectedValue != "" && v != envVar.expectedValue {
			return false
		}
	}
	return true
}

// lookupCiCdProvider returns the name of the CI/CD provider if detected. Returns the
// first one, if multiple are detected.
func lookupCiCdProvider() string {
	for _, p := range listCiCdProviders() {
		if p.detect() {
			return p.name
		}
	}
	return ""
}

var provider string

var providerOnce sync.Once

func CiCdProvider() string {
	providerOnce.Do(func() {
		provider = lookupCiCdProvider()
	})
	return provider
}
