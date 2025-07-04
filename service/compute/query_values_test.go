package compute

import (
	"fmt"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
)

func TestQueryValues(t *testing.T) {
	r := ListClustersRequest{
		FilterBy: &ListClustersFilterBy{
			IsPinned: true,
		},
		ForceSendFields: []string{"FilterBy"},
	}
	pb, err := listClustersRequestToPb(&r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", pb)
	params, err := query.Values(&pb)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", params)
	assert.Equal(t, len(params), 1)
	assert.Equal(t, params.Get("filter_by[is_pinned]"), "true")
}
