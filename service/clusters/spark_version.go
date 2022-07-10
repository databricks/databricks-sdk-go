package clusters

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

// SparkVersionRequest - filtering request
type SparkVersionRequest struct {
	LongTermSupport bool   `json:"long_term_support,omitempty" tf:"optional,default:false"`
	Beta            bool   `json:"beta,omitempty" tf:"optional,default:false,conflicts:long_term_support"`
	Latest          bool   `json:"latest,omitempty" tf:"optional,default:true"`
	ML              bool   `json:"ml,omitempty" tf:"optional,default:false"`
	Genomics        bool   `json:"genomics,omitempty" tf:"optional,default:false"`
	GPU             bool   `json:"gpu,omitempty" tf:"optional,default:false"`
	Scala           string `json:"scala,omitempty" tf:"optional,default:2.12"`
	SparkVersion    string `json:"spark_version,omitempty" tf:"optional,default:"`
	Photon          bool   `json:"photon,omitempty" tf:"optional,default:false"`
	Graviton        bool   `json:"graviton,omitempty"`
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

// LatestSparkVersion returns latest version matching the request parameters
func (sv SparkVersionsList) Select(req SparkVersionRequest) (string, error) {
	var versions []string
	for _, version := range sv.SparkVersions {
		if strings.Contains(version.Version, "-scala"+req.Scala) {
			matches := ((!strings.Contains(version.Version, "apache-spark-")) &&
				(strings.Contains(version.Version, "-ml-") == req.ML) &&
				(strings.Contains(version.Version, "-hls-") == req.Genomics) &&
				(strings.Contains(version.Version, "-gpu-") == req.GPU) &&
				(strings.Contains(version.Version, "-photon-") == req.Photon) &&
				(strings.Contains(version.Version, "-aarch64-") == req.Graviton) &&
				(strings.Contains(version.Description, "Beta") == req.Beta))
			if matches && req.LongTermSupport {
				matches = (matches && (strings.Contains(version.Description, "LTS") || strings.Contains(version.Version, "-esr-")))
			}
			if matches && len(req.SparkVersion) > 0 {
				matches = (matches && strings.Contains(version.Description, "Apache Spark "+req.SparkVersion))
			}
			if matches {
				versions = append(versions, version.Version)
			}
		}
	}
	if len(versions) < 1 {
		return "", fmt.Errorf("spark versions query returned no results. Please change your search criteria and try again")
	} else if len(versions) > 1 {
		if req.Latest {
			sort.Sort(sparkVersionsType(versions))
		} else {
			return "", fmt.Errorf("spark versions query returned multiple results. Please change your search criteria and try again")
		}
	}
	return versions[0], nil
}
