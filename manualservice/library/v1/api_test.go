package v1

import (
	"testing"

	libraryv1 "github.com/databricks/databricks-sdk-go/library/v1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestCreateLibraryMarshalling(t *testing.T) {
	name := "test-library"
	library := &libraryv1.Library{
		Name: &name,
	}

	createLibraryRequest := &libraryv1.CreateLibraryRequest{
		Library: library,
	}
	fieldMask, err := fieldmaskpb.New(library, "name", "created_at")
	if err != nil {
		t.Fatalf("Failed to create field mask: %v", err)
	}
	updateLibraryRequest := &libraryv1.UpdateLibraryRequest{
		Library:    library,
		UpdateMask: fieldMask,
	}

	marshalledFieldMask, err := protojson.Marshal(fieldMask)
	if err != nil {
		t.Fatalf("Failed to marshal field mask: %v", err)
	}

	marshalled, err := protojson.Marshal(createLibraryRequest)
	if err != nil {
		t.Fatalf("Failed to marshal create library request: %v", err)
	}

	marshalledUpdate, err := protojson.Marshal(updateLibraryRequest)
	if err != nil {
		t.Fatalf("Failed to marshal update library request: %v", err)
	}

	if !cmp.Equal(`{"library":{"name":"test-library"}}`, string(marshalled)) {
		t.Fatalf("Create library request marshalling failed")
	}
	if !cmp.Equal("\"name,createdAt\"", string(marshalledFieldMask)) {
		t.Fatalf("Field mask marshalling failed")
	}
	if !cmp.Equal(`{"library":{"name":"test-library"},"update_mask":"\"name,createdAt\""}`, string(marshalledUpdate)) {
		t.Fatalf("Update library request marshalling failed want: %v, got: %v", `{"library":{"name":"test-library"},"update_mask":"name,createdAt"}`, string(marshalledUpdate))
	}
}
