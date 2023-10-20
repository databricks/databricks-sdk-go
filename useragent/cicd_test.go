package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCiCdProviderDetect(t *testing.T) {
	cicdProvider := cicdProvider{
		name: "foo",
		envVars: []envVar{
			{"APPLE", "123"},
			{"BANANA", "456"},
			{"CHERRY", ""},
		},
	}

	// Set some of the env vars.
	t.Setenv("APPLE", "123")
	t.Setenv("CHERRY", "000")
	assert.False(t, cicdProvider.detect(), "should not detect when not all env vars are set")

	// Set the rest of the env vars.
	t.Setenv("BANANA", "456")
	assert.True(t, cicdProvider.detect(), "should detect when all env vars are set")
}

func TestCiCdProviderGithubActions(t *testing.T) {
	// No provider detected.
	assert.Equal(t, "", CiCdProvider())

	// Github Actions detected.
	t.Setenv("GITHUB_ACTIONS", "true")
	assert.Equal(t, "github", CiCdProvider())
}

func TestCiCdProviderGitlab(t *testing.T) {
	// Gitlab detected.
	t.Setenv("GITLAB_CI", "true")
	assert.Equal(t, "gitlab", CiCdProvider())
}

func TestCiCdProviderJenkins(t *testing.T) {
	// Jenkins detected.
	t.Setenv("JENKINS_URL", "https://jenkins.example.com")
	assert.Equal(t, "jenkins", CiCdProvider())
}

func TestCiCdProviderAzureDevops(t *testing.T) {
	// Azure Devops detected.
	t.Setenv("TF_BUILD", "True")
	assert.Equal(t, "azure-devops", CiCdProvider())
}

func TestCiCdProviderCircle(t *testing.T) {
	// Circle detected.
	t.Setenv("CIRCLECI", "true")
	assert.Equal(t, "circle", CiCdProvider())
}

func TestCiCdProviderTravis(t *testing.T) {
	// Travis detected.
	t.Setenv("TRAVIS", "true")
	assert.Equal(t, "travis", CiCdProvider())
}

func TestCiCdProviderBitbucket(t *testing.T) {
	// Bitbucket detected.
	t.Setenv("BITBUCKET_BUILD_NUMBER", "123")
	assert.Equal(t, "bitbucket", CiCdProvider())
}

func TestCiCdProviderGoogleCloudBuild(t *testing.T) {
	// Google Cloud Build detected.
	t.Setenv("PROJECT_ID", "foo")
	t.Setenv("BUILD_ID", "bar")
	t.Setenv("PROJECT_NUMBER", "baz")
	t.Setenv("LOCATION", "")
	assert.Equal(t, "google-cloud-build", CiCdProvider())
}

func TestCiCdProviderAwsCodeBuild(t *testing.T) {
	// AWS Code Build detected.
	t.Setenv("CODEBUILD_BUILD_ARN", "arn:aws:codebuild:us-east-1:123456789012:build/my-demo-project:b1e6deae-e4f2-4151-be79-3cc4e82a0bf0")
	assert.Equal(t, "aws-code-build", CiCdProvider())
}

func TestCiCdProviderTfCloud(t *testing.T) {
	// Terraform Cloud detected.
	t.Setenv("TFC_RUN_ID", "run-123")
	assert.Equal(t, "tf-cloud", CiCdProvider())
}

func TestCiCdProviderMultiple(t *testing.T) {
	// Multiple providers detected. The first one detected is set.
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITLAB_CI", "true")
	assert.Equal(t, "github", CiCdProvider())
}
