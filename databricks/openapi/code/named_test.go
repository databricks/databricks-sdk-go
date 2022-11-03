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
	} {
		assert.Equal(t, expected, actual)
	}
}
