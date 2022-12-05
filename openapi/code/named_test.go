package code

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentFromDescription(t *testing.T) {
	n := Named{Description: `A data lakehouse unifies the best of data 
		warehouses and data lakes in one simple platform to handle all your 
		data, analytics and AI use cases. It's built on an open and reliable
		data foundation that efficiently handles all data types and applies 
		one common security and governance approach across all of your data 
		and cloud platforms.`}
	assert.Equal(t, strings.TrimLeft(`
    // A data lakehouse unifies the best of data warehouses and data lakes in
    // one simple platform to handle all your data, analytics and AI use cases.
    // It's built on an open and reliable data foundation that efficiently
    // handles all data types and applies one common security and governance
    // approach across all of your data and cloud platforms.`,
		"\n\t "), n.Comment("    // ", 80))
}

func TestCommentFromDescriptionWithSummaryAndBlankLines(t *testing.T) {
	n := Named{Description: strings.Join([]string{
		"Databricks Lakehouse",
		"",
		`A data lakehouse unifies the best of data warehouses and data lakes 
		in one simple platform to handle all your data, analytics and AI use 
		cases.`,
		"",
		`It's built on an open and reliable data foundation that efficiently 
		handles all data types and applies one common security and governance 
		approach across all of your data and cloud platforms.`,
	}, "\n")}
	assert.Equal(t, strings.TrimLeft(`
    // Databricks Lakehouse
    // 
    // A data lakehouse unifies the best of data warehouses and data lakes in
    // one simple platform to handle all your data, analytics and AI use cases.
    // 
    // It's built on an open and reliable data foundation that efficiently
    // handles all data types and applies one common security and governance
    // approach across all of your data and cloud platforms.`,
		"\n\t "), n.Comment("    // ", 80))
}

func TestSentencesFromDescription(t *testing.T) {
	n := Named{Description: strings.Join([]string{
		"Databricks Lakehouse.",
		"",
		`A data lakehouse unifies the best of data warehouses and data lakes 
		in one simple platform to handle all your data, analytics and AI use 
		cases.`,
	}, "\n")}
	assert.Equal(t, "Databricks Lakehouse.", n.Summary())
	assert.Equal(t, "A data lakehouse unifies the best of data "+
		"warehouses and data lakes in one simple platform to "+
		"handle all your data, analytics and AI use cases.",
		n.DescriptionWithoutSummary())

	assert.Equal(t, "", (&Named{}).Summary())
	assert.Equal(t, "", (&Named{}).DescriptionWithoutSummary())
	assert.Equal(t, "Foo.", (&Named{Description: "Foo"}).Summary())
	assert.Equal(t, "", (&Named{Description: "Foo"}).DescriptionWithoutSummary())
}

func TestCommentWithLinks(t *testing.T) {
	n := Named{Description: `This is an [example](https://example.com) of
		linking to other web pages. Follows this 
		[convention](https://tip.golang.org/doc/comment#links).`}
	assert.Equal(t, strings.TrimLeft(`
    // This is an [example] of linking to other web pages. Follows this
    // [convention].
    // 
    // [convention]: https://tip.golang.org/doc/comment#links
    // [example]: https://example.com`,
		"\n\t "), n.Comment("    // ", 80))
}

func TestNonConflictingCamelNames(t *testing.T) {
	n := Named{Name: "Import"}
	assert.True(t, n.IsNameReserved())
}

func TestNamedDecamel(t *testing.T) {
	for in, out := range map[string][]string{
		"BigHTMLParser": {"big", "html", "parser"},
		"parseHTML":     {"parse", "html"},
		"parse HTML":    {"parse", "html"},
		"parse-HTML":    {"parse", "html"},
		"parse--HTML":   {"parse", "html"},
		"parse_HTML":    {"parse", "html"},
		"parseHTMLNow":  {"parse", "html", "now"},
		"parseHtml":     {"parse", "html"},
		"ParseHtml":     {"parse", "html"},
		"clusterID":     {"cluster", "id"},
		"positionX":     {"position", "x"},
		"parseHtmlNow":  {"parse", "html", "now"},
		"HTMLParser":    {"html", "parser"},
		"BigO":          {"big", "o"},
		"OCaml":         {"o", "caml"},
	} {
		assert.Equal(t, out, (&Named{Name: in}).splitASCII())
	}
}

func TestNamedTransforms(t *testing.T) {
	n := Named{Name: "bigBrownFOX"}
	for actual, expected := range map[string]string{
		n.CamelName():    "bigBrownFox",
		n.PascalName():   "BigBrownFox",
		n.ConstantName(): "BIG_BROWN_FOX",
		n.SnakeName():    "big_brown_fox",
		n.KebabName():    "big-brown-fox",
		n.TitleName():    "Big Brown Fox",
	} {
		assert.Equal(t, expected, actual)
	}
}

func TestNamedSingular(t *testing.T) {
	for in, out := range map[string]string{
		"buses":             "bus",
		"boxes":             "box",
		"branches":          "branch",
		"blitzes":           "blitz",
		"cluster-policies":  "cluster-policy",
		"clusters":          "cluster",
		"dbfs":              "dbfs",
		"alerts":            "alert",
		"dashboards":        "dashboard",
		"data-sources":      "data-source",
		"dbsql-permissions": "dbsql-permission",
		"queries":           "query",
		"delta-pipelines":   "delta-pipeline",
		"repos":             "repo",
		"metastores":        "metastore",
		"tables":            "table",
		"workspace":         "workspace",
		"warehouses":        "warehouse",
	} {
		n := &Named{Name: in}
		assert.Equal(t, out, n.Singular().Name)
	}
}
