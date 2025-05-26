// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func addExchangeForListingRequestToPb(st *AddExchangeForListingRequest) (*addExchangeForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addExchangeForListingRequestPb{}
	pb.ExchangeId = st.ExchangeId
	pb.ListingId = st.ListingId

	return pb, nil
}

type addExchangeForListingRequestPb struct {
	ExchangeId string `json:"exchange_id"`
	ListingId  string `json:"listing_id"`
}

func addExchangeForListingRequestFromPb(pb *addExchangeForListingRequestPb) (*AddExchangeForListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddExchangeForListingRequest{}
	st.ExchangeId = pb.ExchangeId
	st.ListingId = pb.ListingId

	return st, nil
}

func addExchangeForListingResponseToPb(st *AddExchangeForListingResponse) (*addExchangeForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addExchangeForListingResponsePb{}
	pb.ExchangeForListing = st.ExchangeForListing

	return pb, nil
}

type addExchangeForListingResponsePb struct {
	ExchangeForListing *ExchangeListing `json:"exchange_for_listing,omitempty"`
}

func addExchangeForListingResponseFromPb(pb *addExchangeForListingResponsePb) (*AddExchangeForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddExchangeForListingResponse{}
	st.ExchangeForListing = pb.ExchangeForListing

	return st, nil
}

func batchGetListingsRequestToPb(st *BatchGetListingsRequest) (*batchGetListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetListingsRequestPb{}
	pb.Ids = st.Ids

	return pb, nil
}

type batchGetListingsRequestPb struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

func batchGetListingsRequestFromPb(pb *batchGetListingsRequestPb) (*BatchGetListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetListingsRequest{}
	st.Ids = pb.Ids

	return st, nil
}

func batchGetListingsResponseToPb(st *BatchGetListingsResponse) (*batchGetListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetListingsResponsePb{}
	pb.Listings = st.Listings

	return pb, nil
}

type batchGetListingsResponsePb struct {
	Listings []Listing `json:"listings,omitempty"`
}

func batchGetListingsResponseFromPb(pb *batchGetListingsResponsePb) (*BatchGetListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetListingsResponse{}
	st.Listings = pb.Listings

	return st, nil
}

func batchGetProvidersRequestToPb(st *BatchGetProvidersRequest) (*batchGetProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetProvidersRequestPb{}
	pb.Ids = st.Ids

	return pb, nil
}

type batchGetProvidersRequestPb struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

func batchGetProvidersRequestFromPb(pb *batchGetProvidersRequestPb) (*BatchGetProvidersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetProvidersRequest{}
	st.Ids = pb.Ids

	return st, nil
}

func batchGetProvidersResponseToPb(st *BatchGetProvidersResponse) (*batchGetProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetProvidersResponsePb{}
	pb.Providers = st.Providers

	return pb, nil
}

type batchGetProvidersResponsePb struct {
	Providers []ProviderInfo `json:"providers,omitempty"`
}

func batchGetProvidersResponseFromPb(pb *batchGetProvidersResponsePb) (*BatchGetProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetProvidersResponse{}
	st.Providers = pb.Providers

	return st, nil
}

func consumerTermsToPb(st *ConsumerTerms) (*consumerTermsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &consumerTermsPb{}
	pb.Version = st.Version

	return pb, nil
}

type consumerTermsPb struct {
	Version string `json:"version"`
}

func consumerTermsFromPb(pb *consumerTermsPb) (*ConsumerTerms, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConsumerTerms{}
	st.Version = pb.Version

	return st, nil
}

func contactInfoToPb(st *ContactInfo) (*contactInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &contactInfoPb{}
	pb.Company = st.Company
	pb.Email = st.Email
	pb.FirstName = st.FirstName
	pb.LastName = st.LastName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type contactInfoPb struct {
	Company   string `json:"company,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func contactInfoFromPb(pb *contactInfoPb) (*ContactInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContactInfo{}
	st.Company = pb.Company
	st.Email = pb.Email
	st.FirstName = pb.FirstName
	st.LastName = pb.LastName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *contactInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st contactInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createExchangeFilterRequestToPb(st *CreateExchangeFilterRequest) (*createExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExchangeFilterRequestPb{}
	pb.Filter = st.Filter

	return pb, nil
}

type createExchangeFilterRequestPb struct {
	Filter ExchangeFilter `json:"filter"`
}

func createExchangeFilterRequestFromPb(pb *createExchangeFilterRequestPb) (*CreateExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeFilterRequest{}
	st.Filter = pb.Filter

	return st, nil
}

func createExchangeFilterResponseToPb(st *CreateExchangeFilterResponse) (*createExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExchangeFilterResponsePb{}
	pb.FilterId = st.FilterId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createExchangeFilterResponsePb struct {
	FilterId string `json:"filter_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExchangeFilterResponseFromPb(pb *createExchangeFilterResponsePb) (*CreateExchangeFilterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeFilterResponse{}
	st.FilterId = pb.FilterId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExchangeFilterResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExchangeFilterResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createExchangeRequestToPb(st *CreateExchangeRequest) (*createExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExchangeRequestPb{}
	pb.Exchange = st.Exchange

	return pb, nil
}

type createExchangeRequestPb struct {
	Exchange Exchange `json:"exchange"`
}

func createExchangeRequestFromPb(pb *createExchangeRequestPb) (*CreateExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeRequest{}
	st.Exchange = pb.Exchange

	return st, nil
}

func createExchangeResponseToPb(st *CreateExchangeResponse) (*createExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExchangeResponsePb{}
	pb.ExchangeId = st.ExchangeId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createExchangeResponsePb struct {
	ExchangeId string `json:"exchange_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExchangeResponseFromPb(pb *createExchangeResponsePb) (*CreateExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeResponse{}
	st.ExchangeId = pb.ExchangeId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExchangeResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExchangeResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createFileRequestToPb(st *CreateFileRequest) (*createFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFileRequestPb{}
	pb.DisplayName = st.DisplayName
	pb.FileParent = st.FileParent
	pb.MarketplaceFileType = st.MarketplaceFileType
	pb.MimeType = st.MimeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createFileRequestPb struct {
	DisplayName         string              `json:"display_name,omitempty"`
	FileParent          FileParent          `json:"file_parent"`
	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type"`
	MimeType            string              `json:"mime_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFileRequestFromPb(pb *createFileRequestPb) (*CreateFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFileRequest{}
	st.DisplayName = pb.DisplayName
	st.FileParent = pb.FileParent
	st.MarketplaceFileType = pb.MarketplaceFileType
	st.MimeType = pb.MimeType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createFileRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createFileRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createFileResponseToPb(st *CreateFileResponse) (*createFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFileResponsePb{}
	pb.FileInfo = st.FileInfo
	pb.UploadUrl = st.UploadUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createFileResponsePb struct {
	FileInfo  *FileInfo `json:"file_info,omitempty"`
	UploadUrl string    `json:"upload_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFileResponseFromPb(pb *createFileResponsePb) (*CreateFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFileResponse{}
	st.FileInfo = pb.FileInfo
	st.UploadUrl = pb.UploadUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createFileResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createFileResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createInstallationRequestToPb(st *CreateInstallationRequest) (*createInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createInstallationRequestPb{}
	pb.AcceptedConsumerTerms = st.AcceptedConsumerTerms
	pb.CatalogName = st.CatalogName
	pb.ListingId = st.ListingId
	pb.RecipientType = st.RecipientType
	pb.RepoDetail = st.RepoDetail
	pb.ShareName = st.ShareName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createInstallationRequestPb struct {
	AcceptedConsumerTerms *ConsumerTerms            `json:"accepted_consumer_terms,omitempty"`
	CatalogName           string                    `json:"catalog_name,omitempty"`
	ListingId             string                    `json:"-" url:"-"`
	RecipientType         DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	RepoDetail            *RepoInstallation         `json:"repo_detail,omitempty"`
	ShareName             string                    `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createInstallationRequestFromPb(pb *createInstallationRequestPb) (*CreateInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstallationRequest{}
	st.AcceptedConsumerTerms = pb.AcceptedConsumerTerms
	st.CatalogName = pb.CatalogName
	st.ListingId = pb.ListingId
	st.RecipientType = pb.RecipientType
	st.RepoDetail = pb.RepoDetail
	st.ShareName = pb.ShareName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createInstallationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createInstallationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createListingRequestToPb(st *CreateListingRequest) (*createListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createListingRequestPb{}
	pb.Listing = st.Listing

	return pb, nil
}

type createListingRequestPb struct {
	Listing Listing `json:"listing"`
}

func createListingRequestFromPb(pb *createListingRequestPb) (*CreateListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateListingRequest{}
	st.Listing = pb.Listing

	return st, nil
}

func createListingResponseToPb(st *CreateListingResponse) (*createListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createListingResponsePb{}
	pb.ListingId = st.ListingId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createListingResponsePb struct {
	ListingId string `json:"listing_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createListingResponseFromPb(pb *createListingResponsePb) (*CreateListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateListingResponse{}
	st.ListingId = pb.ListingId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createListingResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createListingResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPersonalizationRequestToPb(st *CreatePersonalizationRequest) (*createPersonalizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPersonalizationRequestPb{}
	pb.AcceptedConsumerTerms = st.AcceptedConsumerTerms
	pb.Comment = st.Comment
	pb.Company = st.Company
	pb.FirstName = st.FirstName
	pb.IntendedUse = st.IntendedUse
	pb.IsFromLighthouse = st.IsFromLighthouse
	pb.LastName = st.LastName
	pb.ListingId = st.ListingId
	pb.RecipientType = st.RecipientType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPersonalizationRequestPb struct {
	AcceptedConsumerTerms ConsumerTerms             `json:"accepted_consumer_terms"`
	Comment               string                    `json:"comment,omitempty"`
	Company               string                    `json:"company,omitempty"`
	FirstName             string                    `json:"first_name,omitempty"`
	IntendedUse           string                    `json:"intended_use"`
	IsFromLighthouse      bool                      `json:"is_from_lighthouse,omitempty"`
	LastName              string                    `json:"last_name,omitempty"`
	ListingId             string                    `json:"-" url:"-"`
	RecipientType         DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPersonalizationRequestFromPb(pb *createPersonalizationRequestPb) (*CreatePersonalizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePersonalizationRequest{}
	st.AcceptedConsumerTerms = pb.AcceptedConsumerTerms
	st.Comment = pb.Comment
	st.Company = pb.Company
	st.FirstName = pb.FirstName
	st.IntendedUse = pb.IntendedUse
	st.IsFromLighthouse = pb.IsFromLighthouse
	st.LastName = pb.LastName
	st.ListingId = pb.ListingId
	st.RecipientType = pb.RecipientType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPersonalizationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPersonalizationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPersonalizationRequestResponseToPb(st *CreatePersonalizationRequestResponse) (*createPersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPersonalizationRequestResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPersonalizationRequestResponsePb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPersonalizationRequestResponseFromPb(pb *createPersonalizationRequestResponsePb) (*CreatePersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePersonalizationRequestResponse{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPersonalizationRequestResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPersonalizationRequestResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createProviderRequestToPb(st *CreateProviderRequest) (*createProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createProviderRequestPb{}
	pb.Provider = st.Provider

	return pb, nil
}

type createProviderRequestPb struct {
	Provider ProviderInfo `json:"provider"`
}

func createProviderRequestFromPb(pb *createProviderRequestPb) (*CreateProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProviderRequest{}
	st.Provider = pb.Provider

	return st, nil
}

func createProviderResponseToPb(st *CreateProviderResponse) (*createProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createProviderResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createProviderResponsePb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createProviderResponseFromPb(pb *createProviderResponsePb) (*CreateProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProviderResponse{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createProviderResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createProviderResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dataRefreshInfoToPb(st *DataRefreshInfo) (*dataRefreshInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataRefreshInfoPb{}
	pb.Interval = st.Interval
	pb.Unit = st.Unit

	return pb, nil
}

type dataRefreshInfoPb struct {
	Interval int64       `json:"interval"`
	Unit     DataRefresh `json:"unit"`
}

func dataRefreshInfoFromPb(pb *dataRefreshInfoPb) (*DataRefreshInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataRefreshInfo{}
	st.Interval = pb.Interval
	st.Unit = pb.Unit

	return st, nil
}

func deleteExchangeFilterRequestToPb(st *DeleteExchangeFilterRequest) (*deleteExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeFilterRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteExchangeFilterRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteExchangeFilterRequestFromPb(pb *deleteExchangeFilterRequestPb) (*DeleteExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExchangeFilterRequest{}
	st.Id = pb.Id

	return st, nil
}

func deleteExchangeFilterResponseToPb(st *DeleteExchangeFilterResponse) (*deleteExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeFilterResponsePb{}

	return pb, nil
}

type deleteExchangeFilterResponsePb struct {
}

func deleteExchangeFilterResponseFromPb(pb *deleteExchangeFilterResponsePb) (*DeleteExchangeFilterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExchangeFilterResponse{}

	return st, nil
}

func deleteExchangeRequestToPb(st *DeleteExchangeRequest) (*deleteExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteExchangeRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteExchangeRequestFromPb(pb *deleteExchangeRequestPb) (*DeleteExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExchangeRequest{}
	st.Id = pb.Id

	return st, nil
}

func deleteExchangeResponseToPb(st *DeleteExchangeResponse) (*deleteExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeResponsePb{}

	return pb, nil
}

type deleteExchangeResponsePb struct {
}

func deleteExchangeResponseFromPb(pb *deleteExchangeResponsePb) (*DeleteExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExchangeResponse{}

	return st, nil
}

func deleteFileRequestToPb(st *DeleteFileRequest) (*deleteFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFileRequestPb{}
	pb.FileId = st.FileId

	return pb, nil
}

type deleteFileRequestPb struct {
	FileId string `json:"-" url:"-"`
}

func deleteFileRequestFromPb(pb *deleteFileRequestPb) (*DeleteFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFileRequest{}
	st.FileId = pb.FileId

	return st, nil
}

func deleteFileResponseToPb(st *DeleteFileResponse) (*deleteFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFileResponsePb{}

	return pb, nil
}

type deleteFileResponsePb struct {
}

func deleteFileResponseFromPb(pb *deleteFileResponsePb) (*DeleteFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFileResponse{}

	return st, nil
}

func deleteInstallationRequestToPb(st *DeleteInstallationRequest) (*deleteInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteInstallationRequestPb{}
	pb.InstallationId = st.InstallationId
	pb.ListingId = st.ListingId

	return pb, nil
}

type deleteInstallationRequestPb struct {
	InstallationId string `json:"-" url:"-"`
	ListingId      string `json:"-" url:"-"`
}

func deleteInstallationRequestFromPb(pb *deleteInstallationRequestPb) (*DeleteInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstallationRequest{}
	st.InstallationId = pb.InstallationId
	st.ListingId = pb.ListingId

	return st, nil
}

func deleteInstallationResponseToPb(st *DeleteInstallationResponse) (*deleteInstallationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteInstallationResponsePb{}

	return pb, nil
}

type deleteInstallationResponsePb struct {
}

func deleteInstallationResponseFromPb(pb *deleteInstallationResponsePb) (*DeleteInstallationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstallationResponse{}

	return st, nil
}

func deleteListingRequestToPb(st *DeleteListingRequest) (*deleteListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteListingRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteListingRequestFromPb(pb *deleteListingRequestPb) (*DeleteListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteListingRequest{}
	st.Id = pb.Id

	return st, nil
}

func deleteListingResponseToPb(st *DeleteListingResponse) (*deleteListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteListingResponsePb{}

	return pb, nil
}

type deleteListingResponsePb struct {
}

func deleteListingResponseFromPb(pb *deleteListingResponsePb) (*DeleteListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteListingResponse{}

	return st, nil
}

func deleteProviderRequestToPb(st *DeleteProviderRequest) (*deleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteProviderRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteProviderRequestFromPb(pb *deleteProviderRequestPb) (*DeleteProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteProviderRequest{}
	st.Id = pb.Id

	return st, nil
}

func deleteProviderResponseToPb(st *DeleteProviderResponse) (*deleteProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderResponsePb{}

	return pb, nil
}

type deleteProviderResponsePb struct {
}

func deleteProviderResponseFromPb(pb *deleteProviderResponsePb) (*DeleteProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteProviderResponse{}

	return st, nil
}

func exchangeToPb(st *Exchange) (*exchangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangePb{}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Filters = st.Filters
	pb.Id = st.Id
	pb.LinkedListings = st.LinkedListings
	pb.Name = st.Name
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type exchangePb struct {
	Comment        string            `json:"comment,omitempty"`
	CreatedAt      int64             `json:"created_at,omitempty"`
	CreatedBy      string            `json:"created_by,omitempty"`
	Filters        []ExchangeFilter  `json:"filters,omitempty"`
	Id             string            `json:"id,omitempty"`
	LinkedListings []ExchangeListing `json:"linked_listings,omitempty"`
	Name           string            `json:"name"`
	UpdatedAt      int64             `json:"updated_at,omitempty"`
	UpdatedBy      string            `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func exchangeFromPb(pb *exchangePb) (*Exchange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Exchange{}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Filters = pb.Filters
	st.Id = pb.Id
	st.LinkedListings = pb.LinkedListings
	st.Name = pb.Name
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *exchangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st exchangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func exchangeFilterToPb(st *ExchangeFilter) (*exchangeFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeFilterPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.ExchangeId = st.ExchangeId
	pb.FilterType = st.FilterType
	pb.FilterValue = st.FilterValue
	pb.Id = st.Id
	pb.Name = st.Name
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type exchangeFilterPb struct {
	CreatedAt   int64              `json:"created_at,omitempty"`
	CreatedBy   string             `json:"created_by,omitempty"`
	ExchangeId  string             `json:"exchange_id"`
	FilterType  ExchangeFilterType `json:"filter_type"`
	FilterValue string             `json:"filter_value"`
	Id          string             `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	UpdatedAt   int64              `json:"updated_at,omitempty"`
	UpdatedBy   string             `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func exchangeFilterFromPb(pb *exchangeFilterPb) (*ExchangeFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeFilter{}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.ExchangeId = pb.ExchangeId
	st.FilterType = pb.FilterType
	st.FilterValue = pb.FilterValue
	st.Id = pb.Id
	st.Name = pb.Name
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *exchangeFilterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st exchangeFilterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func exchangeListingToPb(st *ExchangeListing) (*exchangeListingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeListingPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.ExchangeId = st.ExchangeId
	pb.ExchangeName = st.ExchangeName
	pb.Id = st.Id
	pb.ListingId = st.ListingId
	pb.ListingName = st.ListingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type exchangeListingPb struct {
	CreatedAt    int64  `json:"created_at,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	ExchangeId   string `json:"exchange_id,omitempty"`
	ExchangeName string `json:"exchange_name,omitempty"`
	Id           string `json:"id,omitempty"`
	ListingId    string `json:"listing_id,omitempty"`
	ListingName  string `json:"listing_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func exchangeListingFromPb(pb *exchangeListingPb) (*ExchangeListing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeListing{}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.ExchangeId = pb.ExchangeId
	st.ExchangeName = pb.ExchangeName
	st.Id = pb.Id
	st.ListingId = pb.ListingId
	st.ListingName = pb.ListingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *exchangeListingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st exchangeListingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fileInfoToPb(st *FileInfo) (*fileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileInfoPb{}
	pb.CreatedAt = st.CreatedAt
	pb.DisplayName = st.DisplayName
	pb.DownloadLink = st.DownloadLink
	pb.FileParent = st.FileParent
	pb.Id = st.Id
	pb.MarketplaceFileType = st.MarketplaceFileType
	pb.MimeType = st.MimeType
	pb.Status = st.Status
	pb.StatusMessage = st.StatusMessage
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type fileInfoPb struct {
	CreatedAt           int64               `json:"created_at,omitempty"`
	DisplayName         string              `json:"display_name,omitempty"`
	DownloadLink        string              `json:"download_link,omitempty"`
	FileParent          *FileParent         `json:"file_parent,omitempty"`
	Id                  string              `json:"id,omitempty"`
	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type,omitempty"`
	MimeType            string              `json:"mime_type,omitempty"`
	Status              FileStatus          `json:"status,omitempty"`
	StatusMessage       string              `json:"status_message,omitempty"`
	UpdatedAt           int64               `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileInfoFromPb(pb *fileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.CreatedAt = pb.CreatedAt
	st.DisplayName = pb.DisplayName
	st.DownloadLink = pb.DownloadLink
	st.FileParent = pb.FileParent
	st.Id = pb.Id
	st.MarketplaceFileType = pb.MarketplaceFileType
	st.MimeType = pb.MimeType
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fileParentToPb(st *FileParent) (*fileParentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileParentPb{}
	pb.FileParentType = st.FileParentType
	pb.ParentId = st.ParentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type fileParentPb struct {
	FileParentType FileParentType `json:"file_parent_type,omitempty"`
	ParentId       string         `json:"parent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileParentFromPb(pb *fileParentPb) (*FileParent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileParent{}
	st.FileParentType = pb.FileParentType
	st.ParentId = pb.ParentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileParentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileParentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getExchangeRequestToPb(st *GetExchangeRequest) (*getExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExchangeRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getExchangeRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getExchangeRequestFromPb(pb *getExchangeRequestPb) (*GetExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExchangeRequest{}
	st.Id = pb.Id

	return st, nil
}

func getExchangeResponseToPb(st *GetExchangeResponse) (*getExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExchangeResponsePb{}
	pb.Exchange = st.Exchange

	return pb, nil
}

type getExchangeResponsePb struct {
	Exchange *Exchange `json:"exchange,omitempty"`
}

func getExchangeResponseFromPb(pb *getExchangeResponsePb) (*GetExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExchangeResponse{}
	st.Exchange = pb.Exchange

	return st, nil
}

func getFileRequestToPb(st *GetFileRequest) (*getFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFileRequestPb{}
	pb.FileId = st.FileId

	return pb, nil
}

type getFileRequestPb struct {
	FileId string `json:"-" url:"-"`
}

func getFileRequestFromPb(pb *getFileRequestPb) (*GetFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFileRequest{}
	st.FileId = pb.FileId

	return st, nil
}

func getFileResponseToPb(st *GetFileResponse) (*getFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFileResponsePb{}
	pb.FileInfo = st.FileInfo

	return pb, nil
}

type getFileResponsePb struct {
	FileInfo *FileInfo `json:"file_info,omitempty"`
}

func getFileResponseFromPb(pb *getFileResponsePb) (*GetFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFileResponse{}
	st.FileInfo = pb.FileInfo

	return st, nil
}

func getLatestVersionProviderAnalyticsDashboardResponseToPb(st *GetLatestVersionProviderAnalyticsDashboardResponse) (*getLatestVersionProviderAnalyticsDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLatestVersionProviderAnalyticsDashboardResponsePb{}
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getLatestVersionProviderAnalyticsDashboardResponsePb struct {
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getLatestVersionProviderAnalyticsDashboardResponseFromPb(pb *getLatestVersionProviderAnalyticsDashboardResponsePb) (*GetLatestVersionProviderAnalyticsDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionProviderAnalyticsDashboardResponse{}
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getLatestVersionProviderAnalyticsDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getLatestVersionProviderAnalyticsDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getListingContentMetadataRequestToPb(st *GetListingContentMetadataRequest) (*getListingContentMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingContentMetadataRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getListingContentMetadataRequestPb struct {
	ListingId string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getListingContentMetadataRequestFromPb(pb *getListingContentMetadataRequestPb) (*GetListingContentMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingContentMetadataRequest{}
	st.ListingId = pb.ListingId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getListingContentMetadataRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getListingContentMetadataRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getListingContentMetadataResponseToPb(st *GetListingContentMetadataResponse) (*getListingContentMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingContentMetadataResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.SharedDataObjects = st.SharedDataObjects

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getListingContentMetadataResponsePb struct {
	NextPageToken     string             `json:"next_page_token,omitempty"`
	SharedDataObjects []SharedDataObject `json:"shared_data_objects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getListingContentMetadataResponseFromPb(pb *getListingContentMetadataResponsePb) (*GetListingContentMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingContentMetadataResponse{}
	st.NextPageToken = pb.NextPageToken
	st.SharedDataObjects = pb.SharedDataObjects

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getListingContentMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getListingContentMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getListingRequestToPb(st *GetListingRequest) (*getListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getListingRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getListingRequestFromPb(pb *getListingRequestPb) (*GetListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingRequest{}
	st.Id = pb.Id

	return st, nil
}

func getListingResponseToPb(st *GetListingResponse) (*getListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingResponsePb{}
	pb.Listing = st.Listing

	return pb, nil
}

type getListingResponsePb struct {
	Listing *Listing `json:"listing,omitempty"`
}

func getListingResponseFromPb(pb *getListingResponsePb) (*GetListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingResponse{}
	st.Listing = pb.Listing

	return st, nil
}

func getListingsRequestToPb(st *GetListingsRequest) (*getListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getListingsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getListingsRequestFromPb(pb *getListingsRequestPb) (*GetListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getListingsResponseToPb(st *GetListingsResponse) (*getListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingsResponsePb{}
	pb.Listings = st.Listings
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getListingsResponsePb struct {
	Listings      []Listing `json:"listings,omitempty"`
	NextPageToken string    `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getListingsResponseFromPb(pb *getListingsResponsePb) (*GetListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingsResponse{}
	st.Listings = pb.Listings
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getListingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getListingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPersonalizationRequestRequestToPb(st *GetPersonalizationRequestRequest) (*getPersonalizationRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPersonalizationRequestRequestPb{}
	pb.ListingId = st.ListingId

	return pb, nil
}

type getPersonalizationRequestRequestPb struct {
	ListingId string `json:"-" url:"-"`
}

func getPersonalizationRequestRequestFromPb(pb *getPersonalizationRequestRequestPb) (*GetPersonalizationRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalizationRequestRequest{}
	st.ListingId = pb.ListingId

	return st, nil
}

func getPersonalizationRequestResponseToPb(st *GetPersonalizationRequestResponse) (*getPersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPersonalizationRequestResponsePb{}
	pb.PersonalizationRequests = st.PersonalizationRequests

	return pb, nil
}

type getPersonalizationRequestResponsePb struct {
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`
}

func getPersonalizationRequestResponseFromPb(pb *getPersonalizationRequestResponsePb) (*GetPersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalizationRequestResponse{}
	st.PersonalizationRequests = pb.PersonalizationRequests

	return st, nil
}

func getProviderRequestToPb(st *GetProviderRequest) (*getProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getProviderRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getProviderRequestFromPb(pb *getProviderRequestPb) (*GetProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderRequest{}
	st.Id = pb.Id

	return st, nil
}

func getProviderResponseToPb(st *GetProviderResponse) (*getProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderResponsePb{}
	pb.Provider = st.Provider

	return pb, nil
}

type getProviderResponsePb struct {
	Provider *ProviderInfo `json:"provider,omitempty"`
}

func getProviderResponseFromPb(pb *getProviderResponsePb) (*GetProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderResponse{}
	st.Provider = pb.Provider

	return st, nil
}

func installationToPb(st *Installation) (*installationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &installationPb{}
	pb.Installation = st.Installation

	return pb, nil
}

type installationPb struct {
	Installation *InstallationDetail `json:"installation,omitempty"`
}

func installationFromPb(pb *installationPb) (*Installation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Installation{}
	st.Installation = pb.Installation

	return st, nil
}

func installationDetailToPb(st *InstallationDetail) (*installationDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &installationDetailPb{}
	pb.CatalogName = st.CatalogName
	pb.ErrorMessage = st.ErrorMessage
	pb.Id = st.Id
	pb.InstalledOn = st.InstalledOn
	pb.ListingId = st.ListingId
	pb.ListingName = st.ListingName
	pb.RecipientType = st.RecipientType
	pb.RepoName = st.RepoName
	pb.RepoPath = st.RepoPath
	pb.ShareName = st.ShareName
	pb.Status = st.Status
	pb.TokenDetail = st.TokenDetail
	pb.Tokens = st.Tokens

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type installationDetailPb struct {
	CatalogName   string                    `json:"catalog_name,omitempty"`
	ErrorMessage  string                    `json:"error_message,omitempty"`
	Id            string                    `json:"id,omitempty"`
	InstalledOn   int64                     `json:"installed_on,omitempty"`
	ListingId     string                    `json:"listing_id,omitempty"`
	ListingName   string                    `json:"listing_name,omitempty"`
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	RepoName      string                    `json:"repo_name,omitempty"`
	RepoPath      string                    `json:"repo_path,omitempty"`
	ShareName     string                    `json:"share_name,omitempty"`
	Status        InstallationStatus        `json:"status,omitempty"`
	TokenDetail   *TokenDetail              `json:"token_detail,omitempty"`
	Tokens        []TokenInfo               `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func installationDetailFromPb(pb *installationDetailPb) (*InstallationDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstallationDetail{}
	st.CatalogName = pb.CatalogName
	st.ErrorMessage = pb.ErrorMessage
	st.Id = pb.Id
	st.InstalledOn = pb.InstalledOn
	st.ListingId = pb.ListingId
	st.ListingName = pb.ListingName
	st.RecipientType = pb.RecipientType
	st.RepoName = pb.RepoName
	st.RepoPath = pb.RepoPath
	st.ShareName = pb.ShareName
	st.Status = pb.Status
	st.TokenDetail = pb.TokenDetail
	st.Tokens = pb.Tokens

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *installationDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st installationDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAllInstallationsRequestToPb(st *ListAllInstallationsRequest) (*listAllInstallationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllInstallationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAllInstallationsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAllInstallationsRequestFromPb(pb *listAllInstallationsRequestPb) (*ListAllInstallationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllInstallationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAllInstallationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAllInstallationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAllInstallationsResponseToPb(st *ListAllInstallationsResponse) (*listAllInstallationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllInstallationsResponsePb{}
	pb.Installations = st.Installations
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAllInstallationsResponsePb struct {
	Installations []InstallationDetail `json:"installations,omitempty"`
	NextPageToken string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAllInstallationsResponseFromPb(pb *listAllInstallationsResponsePb) (*ListAllInstallationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllInstallationsResponse{}
	st.Installations = pb.Installations
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAllInstallationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAllInstallationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAllPersonalizationRequestsRequestToPb(st *ListAllPersonalizationRequestsRequest) (*listAllPersonalizationRequestsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllPersonalizationRequestsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAllPersonalizationRequestsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAllPersonalizationRequestsRequestFromPb(pb *listAllPersonalizationRequestsRequestPb) (*ListAllPersonalizationRequestsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllPersonalizationRequestsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAllPersonalizationRequestsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAllPersonalizationRequestsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAllPersonalizationRequestsResponseToPb(st *ListAllPersonalizationRequestsResponse) (*listAllPersonalizationRequestsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllPersonalizationRequestsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PersonalizationRequests = st.PersonalizationRequests

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAllPersonalizationRequestsResponsePb struct {
	NextPageToken           string                   `json:"next_page_token,omitempty"`
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAllPersonalizationRequestsResponseFromPb(pb *listAllPersonalizationRequestsResponsePb) (*ListAllPersonalizationRequestsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllPersonalizationRequestsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PersonalizationRequests = pb.PersonalizationRequests

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAllPersonalizationRequestsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAllPersonalizationRequestsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExchangeFiltersRequestToPb(st *ListExchangeFiltersRequest) (*listExchangeFiltersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangeFiltersRequestPb{}
	pb.ExchangeId = st.ExchangeId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExchangeFiltersRequestPb struct {
	ExchangeId string `json:"-" url:"exchange_id"`
	PageSize   int    `json:"-" url:"page_size,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangeFiltersRequestFromPb(pb *listExchangeFiltersRequestPb) (*ListExchangeFiltersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangeFiltersRequest{}
	st.ExchangeId = pb.ExchangeId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExchangeFiltersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExchangeFiltersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExchangeFiltersResponseToPb(st *ListExchangeFiltersResponse) (*listExchangeFiltersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangeFiltersResponsePb{}
	pb.Filters = st.Filters
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExchangeFiltersResponsePb struct {
	Filters       []ExchangeFilter `json:"filters,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangeFiltersResponseFromPb(pb *listExchangeFiltersResponsePb) (*ListExchangeFiltersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangeFiltersResponse{}
	st.Filters = pb.Filters
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExchangeFiltersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExchangeFiltersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExchangesForListingRequestToPb(st *ListExchangesForListingRequest) (*listExchangesForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangesForListingRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExchangesForListingRequestPb struct {
	ListingId string `json:"-" url:"listing_id"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangesForListingRequestFromPb(pb *listExchangesForListingRequestPb) (*ListExchangesForListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesForListingRequest{}
	st.ListingId = pb.ListingId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExchangesForListingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExchangesForListingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExchangesForListingResponseToPb(st *ListExchangesForListingResponse) (*listExchangesForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangesForListingResponsePb{}
	pb.ExchangeListing = st.ExchangeListing
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExchangesForListingResponsePb struct {
	ExchangeListing []ExchangeListing `json:"exchange_listing,omitempty"`
	NextPageToken   string            `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangesForListingResponseFromPb(pb *listExchangesForListingResponsePb) (*ListExchangesForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesForListingResponse{}
	st.ExchangeListing = pb.ExchangeListing
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExchangesForListingResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExchangesForListingResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExchangesRequestToPb(st *ListExchangesRequest) (*listExchangesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExchangesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangesRequestFromPb(pb *listExchangesRequestPb) (*ListExchangesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExchangesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExchangesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExchangesResponseToPb(st *ListExchangesResponse) (*listExchangesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangesResponsePb{}
	pb.Exchanges = st.Exchanges
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExchangesResponsePb struct {
	Exchanges     []Exchange `json:"exchanges,omitempty"`
	NextPageToken string     `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangesResponseFromPb(pb *listExchangesResponsePb) (*ListExchangesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesResponse{}
	st.Exchanges = pb.Exchanges
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExchangesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExchangesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFilesRequestToPb(st *ListFilesRequest) (*listFilesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFilesRequestPb{}
	pb.FileParent = st.FileParent
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFilesRequestPb struct {
	FileParent FileParent `json:"-" url:"file_parent"`
	PageSize   int        `json:"-" url:"page_size,omitempty"`
	PageToken  string     `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFilesRequestFromPb(pb *listFilesRequestPb) (*ListFilesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFilesRequest{}
	st.FileParent = pb.FileParent
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFilesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFilesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFilesResponseToPb(st *ListFilesResponse) (*listFilesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFilesResponsePb{}
	pb.FileInfos = st.FileInfos
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFilesResponsePb struct {
	FileInfos     []FileInfo `json:"file_infos,omitempty"`
	NextPageToken string     `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFilesResponseFromPb(pb *listFilesResponsePb) (*ListFilesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFilesResponse{}
	st.FileInfos = pb.FileInfos
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFilesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFilesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFulfillmentsRequestToPb(st *ListFulfillmentsRequest) (*listFulfillmentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFulfillmentsRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFulfillmentsRequestPb struct {
	ListingId string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFulfillmentsRequestFromPb(pb *listFulfillmentsRequestPb) (*ListFulfillmentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFulfillmentsRequest{}
	st.ListingId = pb.ListingId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFulfillmentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFulfillmentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFulfillmentsResponseToPb(st *ListFulfillmentsResponse) (*listFulfillmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFulfillmentsResponsePb{}
	pb.Fulfillments = st.Fulfillments
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFulfillmentsResponsePb struct {
	Fulfillments  []ListingFulfillment `json:"fulfillments,omitempty"`
	NextPageToken string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFulfillmentsResponseFromPb(pb *listFulfillmentsResponsePb) (*ListFulfillmentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFulfillmentsResponse{}
	st.Fulfillments = pb.Fulfillments
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFulfillmentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFulfillmentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listInstallationsRequestToPb(st *ListInstallationsRequest) (*listInstallationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstallationsRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listInstallationsRequestPb struct {
	ListingId string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listInstallationsRequestFromPb(pb *listInstallationsRequestPb) (*ListInstallationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstallationsRequest{}
	st.ListingId = pb.ListingId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listInstallationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listInstallationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listInstallationsResponseToPb(st *ListInstallationsResponse) (*listInstallationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstallationsResponsePb{}
	pb.Installations = st.Installations
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listInstallationsResponsePb struct {
	Installations []InstallationDetail `json:"installations,omitempty"`
	NextPageToken string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listInstallationsResponseFromPb(pb *listInstallationsResponsePb) (*ListInstallationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstallationsResponse{}
	st.Installations = pb.Installations
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listInstallationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listInstallationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listListingsForExchangeRequestToPb(st *ListListingsForExchangeRequest) (*listListingsForExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listListingsForExchangeRequestPb{}
	pb.ExchangeId = st.ExchangeId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listListingsForExchangeRequestPb struct {
	ExchangeId string `json:"-" url:"exchange_id"`
	PageSize   int    `json:"-" url:"page_size,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listListingsForExchangeRequestFromPb(pb *listListingsForExchangeRequestPb) (*ListListingsForExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsForExchangeRequest{}
	st.ExchangeId = pb.ExchangeId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listListingsForExchangeRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listListingsForExchangeRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listListingsForExchangeResponseToPb(st *ListListingsForExchangeResponse) (*listListingsForExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listListingsForExchangeResponsePb{}
	pb.ExchangeListings = st.ExchangeListings
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listListingsForExchangeResponsePb struct {
	ExchangeListings []ExchangeListing `json:"exchange_listings,omitempty"`
	NextPageToken    string            `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listListingsForExchangeResponseFromPb(pb *listListingsForExchangeResponsePb) (*ListListingsForExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsForExchangeResponse{}
	st.ExchangeListings = pb.ExchangeListings
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listListingsForExchangeResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listListingsForExchangeResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listListingsRequestToPb(st *ListListingsRequest) (*listListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listListingsRequestPb{}
	pb.Assets = st.Assets
	pb.Categories = st.Categories
	pb.IsFree = st.IsFree
	pb.IsPrivateExchange = st.IsPrivateExchange
	pb.IsStaffPick = st.IsStaffPick
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ProviderIds = st.ProviderIds
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listListingsRequestPb struct {
	Assets            []AssetType  `json:"-" url:"assets,omitempty"`
	Categories        []Category   `json:"-" url:"categories,omitempty"`
	IsFree            bool         `json:"-" url:"is_free,omitempty"`
	IsPrivateExchange bool         `json:"-" url:"is_private_exchange,omitempty"`
	IsStaffPick       bool         `json:"-" url:"is_staff_pick,omitempty"`
	PageSize          int          `json:"-" url:"page_size,omitempty"`
	PageToken         string       `json:"-" url:"page_token,omitempty"`
	ProviderIds       []string     `json:"-" url:"provider_ids,omitempty"`
	Tags              []ListingTag `json:"-" url:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listListingsRequestFromPb(pb *listListingsRequestPb) (*ListListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsRequest{}
	st.Assets = pb.Assets
	st.Categories = pb.Categories
	st.IsFree = pb.IsFree
	st.IsPrivateExchange = pb.IsPrivateExchange
	st.IsStaffPick = pb.IsStaffPick
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ProviderIds = pb.ProviderIds
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listListingsResponseToPb(st *ListListingsResponse) (*listListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listListingsResponsePb{}
	pb.Listings = st.Listings
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listListingsResponsePb struct {
	Listings      []Listing `json:"listings,omitempty"`
	NextPageToken string    `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listListingsResponseFromPb(pb *listListingsResponsePb) (*ListListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsResponse{}
	st.Listings = pb.Listings
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listListingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listListingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProviderAnalyticsDashboardResponseToPb(st *ListProviderAnalyticsDashboardResponse) (*listProviderAnalyticsDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderAnalyticsDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.Id = st.Id
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProviderAnalyticsDashboardResponsePb struct {
	DashboardId string `json:"dashboard_id"`
	Id          string `json:"id"`
	Version     int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProviderAnalyticsDashboardResponseFromPb(pb *listProviderAnalyticsDashboardResponsePb) (*ListProviderAnalyticsDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderAnalyticsDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.Id = pb.Id
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProviderAnalyticsDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProviderAnalyticsDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProvidersRequestToPb(st *ListProvidersRequest) (*listProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersRequestPb{}
	pb.IsFeatured = st.IsFeatured
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProvidersRequestPb struct {
	IsFeatured bool   `json:"-" url:"is_featured,omitempty"`
	PageSize   int    `json:"-" url:"page_size,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProvidersRequestFromPb(pb *listProvidersRequestPb) (*ListProvidersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersRequest{}
	st.IsFeatured = pb.IsFeatured
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProvidersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProvidersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProvidersResponseToPb(st *ListProvidersResponse) (*listProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Providers = st.Providers

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProvidersResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Providers     []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProvidersResponseFromPb(pb *listProvidersResponsePb) (*ListProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Providers = pb.Providers

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProvidersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProvidersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listingToPb(st *Listing) (*listingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingPb{}
	pb.Detail = st.Detail
	pb.Id = st.Id
	pb.Summary = st.Summary

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listingPb struct {
	Detail  *ListingDetail `json:"detail,omitempty"`
	Id      string         `json:"id,omitempty"`
	Summary ListingSummary `json:"summary"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listingFromPb(pb *listingPb) (*Listing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Listing{}
	st.Detail = pb.Detail
	st.Id = pb.Id
	st.Summary = pb.Summary

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listingDetailToPb(st *ListingDetail) (*listingDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingDetailPb{}
	pb.Assets = st.Assets
	pb.CollectionDateEnd = st.CollectionDateEnd
	pb.CollectionDateStart = st.CollectionDateStart
	pb.CollectionGranularity = st.CollectionGranularity
	pb.Cost = st.Cost
	pb.DataSource = st.DataSource
	pb.Description = st.Description
	pb.DocumentationLink = st.DocumentationLink
	pb.EmbeddedNotebookFileInfos = st.EmbeddedNotebookFileInfos
	pb.FileIds = st.FileIds
	pb.GeographicalCoverage = st.GeographicalCoverage
	pb.License = st.License
	pb.PricingModel = st.PricingModel
	pb.PrivacyPolicyLink = st.PrivacyPolicyLink
	pb.Size = st.Size
	pb.SupportLink = st.SupportLink
	pb.Tags = st.Tags
	pb.TermsOfService = st.TermsOfService
	pb.UpdateFrequency = st.UpdateFrequency

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listingDetailPb struct {
	Assets                    []AssetType      `json:"assets,omitempty"`
	CollectionDateEnd         int64            `json:"collection_date_end,omitempty"`
	CollectionDateStart       int64            `json:"collection_date_start,omitempty"`
	CollectionGranularity     *DataRefreshInfo `json:"collection_granularity,omitempty"`
	Cost                      Cost             `json:"cost,omitempty"`
	DataSource                string           `json:"data_source,omitempty"`
	Description               string           `json:"description,omitempty"`
	DocumentationLink         string           `json:"documentation_link,omitempty"`
	EmbeddedNotebookFileInfos []FileInfo       `json:"embedded_notebook_file_infos,omitempty"`
	FileIds                   []string         `json:"file_ids,omitempty"`
	GeographicalCoverage      string           `json:"geographical_coverage,omitempty"`
	License                   string           `json:"license,omitempty"`
	PricingModel              string           `json:"pricing_model,omitempty"`
	PrivacyPolicyLink         string           `json:"privacy_policy_link,omitempty"`
	Size                      float64          `json:"size,omitempty"`
	SupportLink               string           `json:"support_link,omitempty"`
	Tags                      []ListingTag     `json:"tags,omitempty"`
	TermsOfService            string           `json:"terms_of_service,omitempty"`
	UpdateFrequency           *DataRefreshInfo `json:"update_frequency,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listingDetailFromPb(pb *listingDetailPb) (*ListingDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingDetail{}
	st.Assets = pb.Assets
	st.CollectionDateEnd = pb.CollectionDateEnd
	st.CollectionDateStart = pb.CollectionDateStart
	st.CollectionGranularity = pb.CollectionGranularity
	st.Cost = pb.Cost
	st.DataSource = pb.DataSource
	st.Description = pb.Description
	st.DocumentationLink = pb.DocumentationLink
	st.EmbeddedNotebookFileInfos = pb.EmbeddedNotebookFileInfos
	st.FileIds = pb.FileIds
	st.GeographicalCoverage = pb.GeographicalCoverage
	st.License = pb.License
	st.PricingModel = pb.PricingModel
	st.PrivacyPolicyLink = pb.PrivacyPolicyLink
	st.Size = pb.Size
	st.SupportLink = pb.SupportLink
	st.Tags = pb.Tags
	st.TermsOfService = pb.TermsOfService
	st.UpdateFrequency = pb.UpdateFrequency

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listingDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listingDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listingFulfillmentToPb(st *ListingFulfillment) (*listingFulfillmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingFulfillmentPb{}
	pb.FulfillmentType = st.FulfillmentType
	pb.ListingId = st.ListingId
	pb.RecipientType = st.RecipientType
	pb.RepoInfo = st.RepoInfo
	pb.ShareInfo = st.ShareInfo

	return pb, nil
}

type listingFulfillmentPb struct {
	FulfillmentType FulfillmentType           `json:"fulfillment_type,omitempty"`
	ListingId       string                    `json:"listing_id"`
	RecipientType   DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	RepoInfo        *RepoInfo                 `json:"repo_info,omitempty"`
	ShareInfo       *ShareInfo                `json:"share_info,omitempty"`
}

func listingFulfillmentFromPb(pb *listingFulfillmentPb) (*ListingFulfillment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingFulfillment{}
	st.FulfillmentType = pb.FulfillmentType
	st.ListingId = pb.ListingId
	st.RecipientType = pb.RecipientType
	st.RepoInfo = pb.RepoInfo
	st.ShareInfo = pb.ShareInfo

	return st, nil
}

func listingSettingToPb(st *ListingSetting) (*listingSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingSettingPb{}
	pb.Visibility = st.Visibility

	return pb, nil
}

type listingSettingPb struct {
	Visibility Visibility `json:"visibility,omitempty"`
}

func listingSettingFromPb(pb *listingSettingPb) (*ListingSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingSetting{}
	st.Visibility = pb.Visibility

	return st, nil
}

func listingSummaryToPb(st *ListingSummary) (*listingSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingSummaryPb{}
	pb.Categories = st.Categories
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.CreatedById = st.CreatedById
	pb.ExchangeIds = st.ExchangeIds
	pb.GitRepo = st.GitRepo
	pb.ListingType = st.ListingType
	pb.Name = st.Name
	pb.ProviderId = st.ProviderId
	pb.ProviderRegion = st.ProviderRegion
	pb.PublishedAt = st.PublishedAt
	pb.PublishedBy = st.PublishedBy
	pb.Setting = st.Setting
	pb.Share = st.Share
	pb.Status = st.Status
	pb.Subtitle = st.Subtitle
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.UpdatedById = st.UpdatedById

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listingSummaryPb struct {
	Categories     []Category      `json:"categories,omitempty"`
	CreatedAt      int64           `json:"created_at,omitempty"`
	CreatedBy      string          `json:"created_by,omitempty"`
	CreatedById    int64           `json:"created_by_id,omitempty"`
	ExchangeIds    []string        `json:"exchange_ids,omitempty"`
	GitRepo        *RepoInfo       `json:"git_repo,omitempty"`
	ListingType    ListingType     `json:"listingType"`
	Name           string          `json:"name"`
	ProviderId     string          `json:"provider_id,omitempty"`
	ProviderRegion *RegionInfo     `json:"provider_region,omitempty"`
	PublishedAt    int64           `json:"published_at,omitempty"`
	PublishedBy    string          `json:"published_by,omitempty"`
	Setting        *ListingSetting `json:"setting,omitempty"`
	Share          *ShareInfo      `json:"share,omitempty"`
	Status         ListingStatus   `json:"status,omitempty"`
	Subtitle       string          `json:"subtitle,omitempty"`
	UpdatedAt      int64           `json:"updated_at,omitempty"`
	UpdatedBy      string          `json:"updated_by,omitempty"`
	UpdatedById    int64           `json:"updated_by_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listingSummaryFromPb(pb *listingSummaryPb) (*ListingSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingSummary{}
	st.Categories = pb.Categories
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CreatedById = pb.CreatedById
	st.ExchangeIds = pb.ExchangeIds
	st.GitRepo = pb.GitRepo
	st.ListingType = pb.ListingType
	st.Name = pb.Name
	st.ProviderId = pb.ProviderId
	st.ProviderRegion = pb.ProviderRegion
	st.PublishedAt = pb.PublishedAt
	st.PublishedBy = pb.PublishedBy
	st.Setting = pb.Setting
	st.Share = pb.Share
	st.Status = pb.Status
	st.Subtitle = pb.Subtitle
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UpdatedById = pb.UpdatedById

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listingSummaryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listingSummaryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listingTagToPb(st *ListingTag) (*listingTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingTagPb{}
	pb.TagName = st.TagName
	pb.TagValues = st.TagValues

	return pb, nil
}

type listingTagPb struct {
	TagName   ListingTagType `json:"tag_name,omitempty"`
	TagValues []string       `json:"tag_values,omitempty"`
}

func listingTagFromPb(pb *listingTagPb) (*ListingTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingTag{}
	st.TagName = pb.TagName
	st.TagValues = pb.TagValues

	return st, nil
}

func personalizationRequestToPb(st *PersonalizationRequest) (*personalizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &personalizationRequestPb{}
	pb.Comment = st.Comment
	pb.ConsumerRegion = st.ConsumerRegion
	pb.ContactInfo = st.ContactInfo
	pb.CreatedAt = st.CreatedAt
	pb.Id = st.Id
	pb.IntendedUse = st.IntendedUse
	pb.IsFromLighthouse = st.IsFromLighthouse
	pb.ListingId = st.ListingId
	pb.ListingName = st.ListingName
	pb.MetastoreId = st.MetastoreId
	pb.ProviderId = st.ProviderId
	pb.RecipientType = st.RecipientType
	pb.Share = st.Share
	pb.Status = st.Status
	pb.StatusMessage = st.StatusMessage
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type personalizationRequestPb struct {
	Comment          string                       `json:"comment,omitempty"`
	ConsumerRegion   RegionInfo                   `json:"consumer_region"`
	ContactInfo      *ContactInfo                 `json:"contact_info,omitempty"`
	CreatedAt        int64                        `json:"created_at,omitempty"`
	Id               string                       `json:"id,omitempty"`
	IntendedUse      string                       `json:"intended_use,omitempty"`
	IsFromLighthouse bool                         `json:"is_from_lighthouse,omitempty"`
	ListingId        string                       `json:"listing_id,omitempty"`
	ListingName      string                       `json:"listing_name,omitempty"`
	MetastoreId      string                       `json:"metastore_id,omitempty"`
	ProviderId       string                       `json:"provider_id,omitempty"`
	RecipientType    DeltaSharingRecipientType    `json:"recipient_type,omitempty"`
	Share            *ShareInfo                   `json:"share,omitempty"`
	Status           PersonalizationRequestStatus `json:"status,omitempty"`
	StatusMessage    string                       `json:"status_message,omitempty"`
	UpdatedAt        int64                        `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func personalizationRequestFromPb(pb *personalizationRequestPb) (*PersonalizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalizationRequest{}
	st.Comment = pb.Comment
	st.ConsumerRegion = pb.ConsumerRegion
	st.ContactInfo = pb.ContactInfo
	st.CreatedAt = pb.CreatedAt
	st.Id = pb.Id
	st.IntendedUse = pb.IntendedUse
	st.IsFromLighthouse = pb.IsFromLighthouse
	st.ListingId = pb.ListingId
	st.ListingName = pb.ListingName
	st.MetastoreId = pb.MetastoreId
	st.ProviderId = pb.ProviderId
	st.RecipientType = pb.RecipientType
	st.Share = pb.Share
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *personalizationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st personalizationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func providerAnalyticsDashboardToPb(st *ProviderAnalyticsDashboard) (*providerAnalyticsDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerAnalyticsDashboardPb{}
	pb.Id = st.Id

	return pb, nil
}

type providerAnalyticsDashboardPb struct {
	Id string `json:"id"`
}

func providerAnalyticsDashboardFromPb(pb *providerAnalyticsDashboardPb) (*ProviderAnalyticsDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderAnalyticsDashboard{}
	st.Id = pb.Id

	return st, nil
}

func providerInfoToPb(st *ProviderInfo) (*providerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerInfoPb{}
	pb.BusinessContactEmail = st.BusinessContactEmail
	pb.CompanyWebsiteLink = st.CompanyWebsiteLink
	pb.DarkModeIconFileId = st.DarkModeIconFileId
	pb.DarkModeIconFilePath = st.DarkModeIconFilePath
	pb.Description = st.Description
	pb.IconFileId = st.IconFileId
	pb.IconFilePath = st.IconFilePath
	pb.Id = st.Id
	pb.IsFeatured = st.IsFeatured
	pb.Name = st.Name
	pb.PrivacyPolicyLink = st.PrivacyPolicyLink
	pb.PublishedBy = st.PublishedBy
	pb.SupportContactEmail = st.SupportContactEmail
	pb.TermOfServiceLink = st.TermOfServiceLink

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type providerInfoPb struct {
	BusinessContactEmail string `json:"business_contact_email"`
	CompanyWebsiteLink   string `json:"company_website_link,omitempty"`
	DarkModeIconFileId   string `json:"dark_mode_icon_file_id,omitempty"`
	DarkModeIconFilePath string `json:"dark_mode_icon_file_path,omitempty"`
	Description          string `json:"description,omitempty"`
	IconFileId           string `json:"icon_file_id,omitempty"`
	IconFilePath         string `json:"icon_file_path,omitempty"`
	Id                   string `json:"id,omitempty"`
	IsFeatured           bool   `json:"is_featured,omitempty"`
	Name                 string `json:"name"`
	PrivacyPolicyLink    string `json:"privacy_policy_link"`
	PublishedBy          string `json:"published_by,omitempty"`
	SupportContactEmail  string `json:"support_contact_email,omitempty"`
	TermOfServiceLink    string `json:"term_of_service_link"`

	ForceSendFields []string `json:"-" url:"-"`
}

func providerInfoFromPb(pb *providerInfoPb) (*ProviderInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderInfo{}
	st.BusinessContactEmail = pb.BusinessContactEmail
	st.CompanyWebsiteLink = pb.CompanyWebsiteLink
	st.DarkModeIconFileId = pb.DarkModeIconFileId
	st.DarkModeIconFilePath = pb.DarkModeIconFilePath
	st.Description = pb.Description
	st.IconFileId = pb.IconFileId
	st.IconFilePath = pb.IconFilePath
	st.Id = pb.Id
	st.IsFeatured = pb.IsFeatured
	st.Name = pb.Name
	st.PrivacyPolicyLink = pb.PrivacyPolicyLink
	st.PublishedBy = pb.PublishedBy
	st.SupportContactEmail = pb.SupportContactEmail
	st.TermOfServiceLink = pb.TermOfServiceLink

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *providerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st providerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func regionInfoToPb(st *RegionInfo) (*regionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regionInfoPb{}
	pb.Cloud = st.Cloud
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type regionInfoPb struct {
	Cloud  string `json:"cloud,omitempty"`
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func regionInfoFromPb(pb *regionInfoPb) (*RegionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegionInfo{}
	st.Cloud = pb.Cloud
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *regionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st regionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func removeExchangeForListingRequestToPb(st *RemoveExchangeForListingRequest) (*removeExchangeForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeExchangeForListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type removeExchangeForListingRequestPb struct {
	Id string `json:"-" url:"-"`
}

func removeExchangeForListingRequestFromPb(pb *removeExchangeForListingRequestPb) (*RemoveExchangeForListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveExchangeForListingRequest{}
	st.Id = pb.Id

	return st, nil
}

func removeExchangeForListingResponseToPb(st *RemoveExchangeForListingResponse) (*removeExchangeForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeExchangeForListingResponsePb{}

	return pb, nil
}

type removeExchangeForListingResponsePb struct {
}

func removeExchangeForListingResponseFromPb(pb *removeExchangeForListingResponsePb) (*RemoveExchangeForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveExchangeForListingResponse{}

	return st, nil
}

func repoInfoToPb(st *RepoInfo) (*repoInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoInfoPb{}
	pb.GitRepoUrl = st.GitRepoUrl

	return pb, nil
}

type repoInfoPb struct {
	GitRepoUrl string `json:"git_repo_url"`
}

func repoInfoFromPb(pb *repoInfoPb) (*RepoInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoInfo{}
	st.GitRepoUrl = pb.GitRepoUrl

	return st, nil
}

func repoInstallationToPb(st *RepoInstallation) (*repoInstallationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoInstallationPb{}
	pb.RepoName = st.RepoName
	pb.RepoPath = st.RepoPath

	return pb, nil
}

type repoInstallationPb struct {
	RepoName string `json:"repo_name"`
	RepoPath string `json:"repo_path"`
}

func repoInstallationFromPb(pb *repoInstallationPb) (*RepoInstallation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoInstallation{}
	st.RepoName = pb.RepoName
	st.RepoPath = pb.RepoPath

	return st, nil
}

func searchListingsRequestToPb(st *SearchListingsRequest) (*searchListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchListingsRequestPb{}
	pb.Assets = st.Assets
	pb.Categories = st.Categories
	pb.IsFree = st.IsFree
	pb.IsPrivateExchange = st.IsPrivateExchange
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ProviderIds = st.ProviderIds
	pb.Query = st.Query

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchListingsRequestPb struct {
	Assets            []AssetType `json:"-" url:"assets,omitempty"`
	Categories        []Category  `json:"-" url:"categories,omitempty"`
	IsFree            bool        `json:"-" url:"is_free,omitempty"`
	IsPrivateExchange bool        `json:"-" url:"is_private_exchange,omitempty"`
	PageSize          int         `json:"-" url:"page_size,omitempty"`
	PageToken         string      `json:"-" url:"page_token,omitempty"`
	ProviderIds       []string    `json:"-" url:"provider_ids,omitempty"`
	Query             string      `json:"-" url:"query"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchListingsRequestFromPb(pb *searchListingsRequestPb) (*SearchListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchListingsRequest{}
	st.Assets = pb.Assets
	st.Categories = pb.Categories
	st.IsFree = pb.IsFree
	st.IsPrivateExchange = pb.IsPrivateExchange
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ProviderIds = pb.ProviderIds
	st.Query = pb.Query

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func searchListingsResponseToPb(st *SearchListingsResponse) (*searchListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchListingsResponsePb{}
	pb.Listings = st.Listings
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type searchListingsResponsePb struct {
	Listings      []Listing `json:"listings,omitempty"`
	NextPageToken string    `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchListingsResponseFromPb(pb *searchListingsResponsePb) (*SearchListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchListingsResponse{}
	st.Listings = pb.Listings
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchListingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchListingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func shareInfoToPb(st *ShareInfo) (*shareInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &shareInfoPb{}
	pb.Name = st.Name
	pb.Type = st.Type

	return pb, nil
}

type shareInfoPb struct {
	Name string           `json:"name"`
	Type ListingShareType `json:"type"`
}

func shareInfoFromPb(pb *shareInfoPb) (*ShareInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareInfo{}
	st.Name = pb.Name
	st.Type = pb.Type

	return st, nil
}

func sharedDataObjectToPb(st *SharedDataObject) (*sharedDataObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharedDataObjectPb{}
	pb.DataObjectType = st.DataObjectType
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sharedDataObjectPb struct {
	DataObjectType string `json:"data_object_type,omitempty"`
	Name           string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sharedDataObjectFromPb(pb *sharedDataObjectPb) (*SharedDataObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObject{}
	st.DataObjectType = pb.DataObjectType
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sharedDataObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sharedDataObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenDetailToPb(st *TokenDetail) (*tokenDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenDetailPb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ExpirationTime = st.ExpirationTime
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenDetailPb struct {
	BearerToken             string `json:"bearerToken,omitempty"`
	Endpoint                string `json:"endpoint,omitempty"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	ShareCredentialsVersion int    `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenDetailFromPb(pb *tokenDetailPb) (*TokenDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenDetail{}
	st.BearerToken = pb.BearerToken
	st.Endpoint = pb.Endpoint
	st.ExpirationTime = pb.ExpirationTime
	st.ShareCredentialsVersion = pb.ShareCredentialsVersion

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenInfoToPb(st *TokenInfo) (*tokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenInfoPb{}
	pb.ActivationUrl = st.ActivationUrl
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.ExpirationTime = st.ExpirationTime
	pb.Id = st.Id
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenInfoPb struct {
	ActivationUrl  string `json:"activation_url,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
	ExpirationTime int64  `json:"expiration_time,omitempty"`
	Id             string `json:"id,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	UpdatedBy      string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenInfoFromPb(pb *tokenInfoPb) (*TokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenInfo{}
	st.ActivationUrl = pb.ActivationUrl
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.ExpirationTime = pb.ExpirationTime
	st.Id = pb.Id
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateExchangeFilterRequestToPb(st *UpdateExchangeFilterRequest) (*updateExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeFilterRequestPb{}
	pb.Filter = st.Filter
	pb.Id = st.Id

	return pb, nil
}

type updateExchangeFilterRequestPb struct {
	Filter ExchangeFilter `json:"filter"`
	Id     string         `json:"-" url:"-"`
}

func updateExchangeFilterRequestFromPb(pb *updateExchangeFilterRequestPb) (*UpdateExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeFilterRequest{}
	st.Filter = pb.Filter
	st.Id = pb.Id

	return st, nil
}

func updateExchangeFilterResponseToPb(st *UpdateExchangeFilterResponse) (*updateExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeFilterResponsePb{}
	pb.Filter = st.Filter

	return pb, nil
}

type updateExchangeFilterResponsePb struct {
	Filter *ExchangeFilter `json:"filter,omitempty"`
}

func updateExchangeFilterResponseFromPb(pb *updateExchangeFilterResponsePb) (*UpdateExchangeFilterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeFilterResponse{}
	st.Filter = pb.Filter

	return st, nil
}

func updateExchangeRequestToPb(st *UpdateExchangeRequest) (*updateExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeRequestPb{}
	pb.Exchange = st.Exchange
	pb.Id = st.Id

	return pb, nil
}

type updateExchangeRequestPb struct {
	Exchange Exchange `json:"exchange"`
	Id       string   `json:"-" url:"-"`
}

func updateExchangeRequestFromPb(pb *updateExchangeRequestPb) (*UpdateExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeRequest{}
	st.Exchange = pb.Exchange
	st.Id = pb.Id

	return st, nil
}

func updateExchangeResponseToPb(st *UpdateExchangeResponse) (*updateExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeResponsePb{}
	pb.Exchange = st.Exchange

	return pb, nil
}

type updateExchangeResponsePb struct {
	Exchange *Exchange `json:"exchange,omitempty"`
}

func updateExchangeResponseFromPb(pb *updateExchangeResponsePb) (*UpdateExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeResponse{}
	st.Exchange = pb.Exchange

	return st, nil
}

func updateInstallationRequestToPb(st *UpdateInstallationRequest) (*updateInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInstallationRequestPb{}
	pb.Installation = st.Installation
	pb.InstallationId = st.InstallationId
	pb.ListingId = st.ListingId
	pb.RotateToken = st.RotateToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateInstallationRequestPb struct {
	Installation   InstallationDetail `json:"installation"`
	InstallationId string             `json:"-" url:"-"`
	ListingId      string             `json:"-" url:"-"`
	RotateToken    bool               `json:"rotate_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateInstallationRequestFromPb(pb *updateInstallationRequestPb) (*UpdateInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInstallationRequest{}
	st.Installation = pb.Installation
	st.InstallationId = pb.InstallationId
	st.ListingId = pb.ListingId
	st.RotateToken = pb.RotateToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateInstallationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateInstallationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateInstallationResponseToPb(st *UpdateInstallationResponse) (*updateInstallationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInstallationResponsePb{}
	pb.Installation = st.Installation

	return pb, nil
}

type updateInstallationResponsePb struct {
	Installation *InstallationDetail `json:"installation,omitempty"`
}

func updateInstallationResponseFromPb(pb *updateInstallationResponsePb) (*UpdateInstallationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInstallationResponse{}
	st.Installation = pb.Installation

	return st, nil
}

func updateListingRequestToPb(st *UpdateListingRequest) (*updateListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateListingRequestPb{}
	pb.Id = st.Id
	pb.Listing = st.Listing

	return pb, nil
}

type updateListingRequestPb struct {
	Id      string  `json:"-" url:"-"`
	Listing Listing `json:"listing"`
}

func updateListingRequestFromPb(pb *updateListingRequestPb) (*UpdateListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateListingRequest{}
	st.Id = pb.Id
	st.Listing = pb.Listing

	return st, nil
}

func updateListingResponseToPb(st *UpdateListingResponse) (*updateListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateListingResponsePb{}
	pb.Listing = st.Listing

	return pb, nil
}

type updateListingResponsePb struct {
	Listing *Listing `json:"listing,omitempty"`
}

func updateListingResponseFromPb(pb *updateListingResponsePb) (*UpdateListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateListingResponse{}
	st.Listing = pb.Listing

	return st, nil
}

func updatePersonalizationRequestRequestToPb(st *UpdatePersonalizationRequestRequest) (*updatePersonalizationRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePersonalizationRequestRequestPb{}
	pb.ListingId = st.ListingId
	pb.Reason = st.Reason
	pb.RequestId = st.RequestId
	pb.Share = st.Share
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updatePersonalizationRequestRequestPb struct {
	ListingId string                       `json:"-" url:"-"`
	Reason    string                       `json:"reason,omitempty"`
	RequestId string                       `json:"-" url:"-"`
	Share     *ShareInfo                   `json:"share,omitempty"`
	Status    PersonalizationRequestStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updatePersonalizationRequestRequestFromPb(pb *updatePersonalizationRequestRequestPb) (*UpdatePersonalizationRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalizationRequestRequest{}
	st.ListingId = pb.ListingId
	st.Reason = pb.Reason
	st.RequestId = pb.RequestId
	st.Share = pb.Share
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updatePersonalizationRequestRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updatePersonalizationRequestRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updatePersonalizationRequestResponseToPb(st *UpdatePersonalizationRequestResponse) (*updatePersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePersonalizationRequestResponsePb{}
	pb.Request = st.Request

	return pb, nil
}

type updatePersonalizationRequestResponsePb struct {
	Request *PersonalizationRequest `json:"request,omitempty"`
}

func updatePersonalizationRequestResponseFromPb(pb *updatePersonalizationRequestResponsePb) (*UpdatePersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalizationRequestResponse{}
	st.Request = pb.Request

	return st, nil
}

func updateProviderAnalyticsDashboardRequestToPb(st *UpdateProviderAnalyticsDashboardRequest) (*updateProviderAnalyticsDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderAnalyticsDashboardRequestPb{}
	pb.Id = st.Id
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateProviderAnalyticsDashboardRequestPb struct {
	Id      string `json:"-" url:"-"`
	Version int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateProviderAnalyticsDashboardRequestFromPb(pb *updateProviderAnalyticsDashboardRequestPb) (*UpdateProviderAnalyticsDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderAnalyticsDashboardRequest{}
	st.Id = pb.Id
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateProviderAnalyticsDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateProviderAnalyticsDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateProviderAnalyticsDashboardResponseToPb(st *UpdateProviderAnalyticsDashboardResponse) (*updateProviderAnalyticsDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderAnalyticsDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.Id = st.Id
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateProviderAnalyticsDashboardResponsePb struct {
	DashboardId string `json:"dashboard_id"`
	Id          string `json:"id"`
	Version     int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateProviderAnalyticsDashboardResponseFromPb(pb *updateProviderAnalyticsDashboardResponsePb) (*UpdateProviderAnalyticsDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderAnalyticsDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.Id = pb.Id
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateProviderAnalyticsDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateProviderAnalyticsDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateProviderRequestToPb(st *UpdateProviderRequest) (*updateProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderRequestPb{}
	pb.Id = st.Id
	pb.Provider = st.Provider

	return pb, nil
}

type updateProviderRequestPb struct {
	Id       string       `json:"-" url:"-"`
	Provider ProviderInfo `json:"provider"`
}

func updateProviderRequestFromPb(pb *updateProviderRequestPb) (*UpdateProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderRequest{}
	st.Id = pb.Id
	st.Provider = pb.Provider

	return st, nil
}

func updateProviderResponseToPb(st *UpdateProviderResponse) (*updateProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderResponsePb{}
	pb.Provider = st.Provider

	return pb, nil
}

type updateProviderResponsePb struct {
	Provider *ProviderInfo `json:"provider,omitempty"`
}

func updateProviderResponseFromPb(pb *updateProviderResponsePb) (*UpdateProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderResponse{}
	st.Provider = pb.Provider

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
