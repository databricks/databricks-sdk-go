package code

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#pull_request

func getTags(repo string) (tags []string, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/tags", repo))
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	var tagsResponse []struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(raw, &tagsResponse)
	if err != nil {
		return nil, fmt.Errorf("unmarshall: %w", err)
	}
	for _, v := range tagsResponse {
		tags = append(tags, v.Name)
	}
	return tags, nil
}

func compareCommits(repo, tag string) (lines []string, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/compare/%s...main", repo, tag))
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	var commitsResponse struct {
		Commits []struct {
			Commit struct {
				Message string `json:"message"`
			} `json:"commit"`
		} `json:"commits"`
	}
	err = json.Unmarshal(raw, &commitsResponse)
	if err != nil {
		return nil, fmt.Errorf("unmarshall: %w", err)
	}
	issueNre := regexp.MustCompile(`(?m)#(\d+)`)
	for _, v := range commitsResponse.Commits {
		msg := strings.Split(v.Commit.Message, "\n")[0]
		prLink := fmt.Sprintf("[#$1](https://github.com/%s/pull/$1)", repo)
		msg = issueNre.ReplaceAllString(msg, prLink)
		lines = append(lines, msg)
	}
	sort.Strings(lines)
	return lines, nil
}

func TestGithub(t *testing.T) {
	repo := "databricks/databricks-sdk-go"
	tags, err := getTags(repo)
	assert.NoError(t, err)

	_, err = compareCommits(repo, tags[0])
	assert.NoError(t, err)
}
