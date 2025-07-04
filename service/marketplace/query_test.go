package marketplace

import (
	"fmt"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
)

func TestQueryValuesArray(t *testing.T) {
	r := ListListingsRequest{
		Tags: []ListingTag{
			{
				TagName: "test",
			},
			{
				TagName: "test",
			},
		},
		ForceSendFields: []string{"FilterBy"},
	}
	pb, err := listListingsRequestToPb(&r)
	if err != nil {
		t.Fatal(err)
	}
	params, err := query.Values(&pb)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", params)
	assert.Equal(t, len(params), 1)
	assert.Equal(t, params.Get("tags"), "{test []}")
}
