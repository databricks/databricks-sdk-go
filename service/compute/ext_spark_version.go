package compute

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

// SparkVersionRequest - filtering request
type SparkVersionRequest struct {
	Id              string `json:"id,omitempty"`
	LongTermSupport bool   `json:"long_term_support,omitempty" tf:"optional,default:false"`
	Beta            bool   `json:"beta,omitempty" tf:"optional,default:false,conflicts:long_term_support"`
	Latest          bool   `json:"latest,omitempty" tf:"optional,default:true"`
	ML              bool   `json:"ml,omitempty" tf:"optional,default:false"`
	Genomics        bool   `json:"genomics,omitempty" tf:"optional,default:false"`
	GPU             bool   `json:"gpu,omitempty" tf:"optional,default:false"`
	Scala           string `json:"scala,omitempty" tf:"optional,default:2.12"`
	SparkVersion    string `json:"spark_version,omitempty" tf:"optional,default:"`
	Photon          bool   `json:"photon,omitempty" tf:"optional,default:false"`
	Graviton        bool   `json:"graviton,omitempty" tf:"optional,default:false"`
}

type sparkVersionsType []string

func (s sparkVersionsType) Len() int {
	return len(s)
}
func (s sparkVersionsType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var dbrVersionRegex = regexp.MustCompile(`^(\d+\.\d+)\.x-.*`)

func extractDbrVersions(s string) string {
	m := dbrVersionRegex.FindStringSubmatch(s)
	if len(m) > 1 {
		return m[1]
	}
	return s
}

func (s sparkVersionsType) Less(i, j int) bool {
	return semver.Compare("v"+extractDbrVersions(s[i]), "v"+extractDbrVersions(s[j])) > 0
}

func (sv GetSparkVersionsResponse) Select(req SparkVersionRequest) (string, error) {
	var versions []string
	for _, version := range sv.Versions {
		if strings.Contains(version.Key, "-scala"+req.Scala) {
			matches := ((!strings.Contains(version.Key, "apache-spark-")) &&
				(strings.Contains(version.Key, "-ml-") == req.ML) &&
				(strings.Contains(version.Key, "-hls-") == req.Genomics) &&
				(strings.Contains(version.Key, "-gpu-") == req.GPU) &&
				(strings.Contains(version.Key, "-photon-") == req.Photon) &&
				(strings.Contains(version.Key, "-aarch64-") == req.Graviton) &&
				(strings.Contains(version.Name, "Beta") == req.Beta))
			if matches && req.LongTermSupport {
				matches = (matches && (strings.Contains(version.Name, "LTS") || strings.Contains(version.Key, "-esr-")))
			}
			if matches && len(req.SparkVersion) > 0 {
				matches = (matches && strings.Contains(version.Name, "Apache Spark "+req.SparkVersion))
			}
			if matches {
				versions = append(versions, version.Key)
			}
		}
	}
	if len(versions) < 1 {
		return "", fmt.Errorf("spark versions query returned no results. Please change your search criteria and try again")
	} else if len(versions) > 1 {
		if req.Latest {
			sort.Sort(sparkVersionsType(versions))
		} else {
			return "", fmt.Errorf("spark versions query returned multiple results %#v. Please change your search criteria and try again", versions)
		}
	}
	return versions[0], nil
}

// SelectSparkVersion returns latest DBR version matching the request parameters.
// If there are multiple versions matching the request, it will error (if latest = false)
// or return the latest version.
// Possible parameters are:
// - LongTermSupport: LTS versions only
// - Beta: Beta versions only
// - ML: ML versions only
// - Genomics: Genomics versions only
// - GPU: GPU versions only
// - Scala: Scala version
// - SparkVersion: Apache Spark version
// - Photon: Photon versions only (deprecated)
// - Graviton: Graviton versions only (deprecated)
func (a *ClustersAPI) SelectSparkVersion(ctx context.Context, r SparkVersionRequest) (string, error) {
	sv, err := a.SparkVersions(ctx)
	if err != nil {
		return "", err
	}
	return sv.Select(r)
}
