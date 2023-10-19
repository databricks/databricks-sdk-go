package useragent

import "os"

type envVar struct {
	name          string
	expectedValue string
}

type cicdProvider struct {
	name    string
	envVars []envVar
}

// TODO: add bamboo to this?
var cicdProviders = []cicdProvider{
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

func (p cicdProvider) detect() bool {
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

func CiCdProvider() string {
	for _, p := range cicdProviders {
		if p.detect() {
			return p.name
		}
	}
	return ""
}
