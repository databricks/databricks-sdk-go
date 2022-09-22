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
