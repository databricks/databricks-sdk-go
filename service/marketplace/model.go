// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddExchangeForListingRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string

	// Wire name: 'listing_id'
	ListingId string
}

func addExchangeForListingRequestToPb(st *AddExchangeForListingRequest) (*addExchangeForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addExchangeForListingRequestPb{}
	pb.ExchangeId = st.ExchangeId

	pb.ListingId = st.ListingId

	return pb, nil
}

func (st *AddExchangeForListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &addExchangeForListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := addExchangeForListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AddExchangeForListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := addExchangeForListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type addExchangeForListingRequestPb struct {
	ExchangeId string `json:"exchange_id"`

	ListingId string `json:"listing_id"`
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

type AddExchangeForListingResponse struct {

	// Wire name: 'exchange_for_listing'
	ExchangeForListing *ExchangeListing
}

func addExchangeForListingResponseToPb(st *AddExchangeForListingResponse) (*addExchangeForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addExchangeForListingResponsePb{}
	exchangeForListingPb, err := exchangeListingToPb(st.ExchangeForListing)
	if err != nil {
		return nil, err
	}
	if exchangeForListingPb != nil {
		pb.ExchangeForListing = exchangeForListingPb
	}

	return pb, nil
}

func (st *AddExchangeForListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &addExchangeForListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := addExchangeForListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AddExchangeForListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := addExchangeForListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type addExchangeForListingResponsePb struct {
	ExchangeForListing *exchangeListingPb `json:"exchange_for_listing,omitempty"`
}

func addExchangeForListingResponseFromPb(pb *addExchangeForListingResponsePb) (*AddExchangeForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddExchangeForListingResponse{}
	exchangeForListingField, err := exchangeListingFromPb(pb.ExchangeForListing)
	if err != nil {
		return nil, err
	}
	if exchangeForListingField != nil {
		st.ExchangeForListing = exchangeForListingField
	}

	return st, nil
}

type AssetType string
type assetTypePb string

const AssetTypeAssetTypeApp AssetType = `ASSET_TYPE_APP`

const AssetTypeAssetTypeDataTable AssetType = `ASSET_TYPE_DATA_TABLE`

const AssetTypeAssetTypeGitRepo AssetType = `ASSET_TYPE_GIT_REPO`

const AssetTypeAssetTypeMedia AssetType = `ASSET_TYPE_MEDIA`

const AssetTypeAssetTypeModel AssetType = `ASSET_TYPE_MODEL`

const AssetTypeAssetTypeNotebook AssetType = `ASSET_TYPE_NOTEBOOK`

const AssetTypeAssetTypePartnerIntegration AssetType = `ASSET_TYPE_PARTNER_INTEGRATION`

// String representation for [fmt.Print]
func (f *AssetType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AssetType) Set(v string) error {
	switch v {
	case `ASSET_TYPE_APP`, `ASSET_TYPE_DATA_TABLE`, `ASSET_TYPE_GIT_REPO`, `ASSET_TYPE_MEDIA`, `ASSET_TYPE_MODEL`, `ASSET_TYPE_NOTEBOOK`, `ASSET_TYPE_PARTNER_INTEGRATION`:
		*f = AssetType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASSET_TYPE_APP", "ASSET_TYPE_DATA_TABLE", "ASSET_TYPE_GIT_REPO", "ASSET_TYPE_MEDIA", "ASSET_TYPE_MODEL", "ASSET_TYPE_NOTEBOOK", "ASSET_TYPE_PARTNER_INTEGRATION"`, v)
	}
}

// Type always returns AssetType to satisfy [pflag.Value] interface
func (f *AssetType) Type() string {
	return "AssetType"
}

func assetTypeToPb(st *AssetType) (*assetTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := assetTypePb(*st)
	return &pb, nil
}

func assetTypeFromPb(pb *assetTypePb) (*AssetType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AssetType(*pb)
	return &st, nil
}

// Get one batch of listings. One may specify up to 50 IDs per request.
type BatchGetListingsRequest struct {

	// Wire name: 'ids'
	Ids []string `tf:"-"`
}

func batchGetListingsRequestToPb(st *BatchGetListingsRequest) (*batchGetListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetListingsRequestPb{}
	pb.Ids = st.Ids

	return pb, nil
}

func (st *BatchGetListingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &batchGetListingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := batchGetListingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BatchGetListingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := batchGetListingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type BatchGetListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing
}

func batchGetListingsResponseToPb(st *BatchGetListingsResponse) (*batchGetListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetListingsResponsePb{}

	var listingsPb []listingPb
	for _, item := range st.Listings {
		itemPb, err := listingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			listingsPb = append(listingsPb, *itemPb)
		}
	}
	pb.Listings = listingsPb

	return pb, nil
}

func (st *BatchGetListingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &batchGetListingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := batchGetListingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BatchGetListingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := batchGetListingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type batchGetListingsResponsePb struct {
	Listings []listingPb `json:"listings,omitempty"`
}

func batchGetListingsResponseFromPb(pb *batchGetListingsResponsePb) (*BatchGetListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := listingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			listingsField = append(listingsField, *item)
		}
	}
	st.Listings = listingsField

	return st, nil
}

// Get one batch of providers. One may specify up to 50 IDs per request.
type BatchGetProvidersRequest struct {

	// Wire name: 'ids'
	Ids []string `tf:"-"`
}

func batchGetProvidersRequestToPb(st *BatchGetProvidersRequest) (*batchGetProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetProvidersRequestPb{}
	pb.Ids = st.Ids

	return pb, nil
}

func (st *BatchGetProvidersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &batchGetProvidersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := batchGetProvidersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BatchGetProvidersRequest) MarshalJSON() ([]byte, error) {
	pb, err := batchGetProvidersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type BatchGetProvidersResponse struct {

	// Wire name: 'providers'
	Providers []ProviderInfo
}

func batchGetProvidersResponseToPb(st *BatchGetProvidersResponse) (*batchGetProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &batchGetProvidersResponsePb{}

	var providersPb []providerInfoPb
	for _, item := range st.Providers {
		itemPb, err := providerInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			providersPb = append(providersPb, *itemPb)
		}
	}
	pb.Providers = providersPb

	return pb, nil
}

func (st *BatchGetProvidersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &batchGetProvidersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := batchGetProvidersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BatchGetProvidersResponse) MarshalJSON() ([]byte, error) {
	pb, err := batchGetProvidersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type batchGetProvidersResponsePb struct {
	Providers []providerInfoPb `json:"providers,omitempty"`
}

func batchGetProvidersResponseFromPb(pb *batchGetProvidersResponsePb) (*BatchGetProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetProvidersResponse{}

	var providersField []ProviderInfo
	for _, itemPb := range pb.Providers {
		item, err := providerInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			providersField = append(providersField, *item)
		}
	}
	st.Providers = providersField

	return st, nil
}

type Category string
type categoryPb string

const CategoryAdvertisingAndMarketing Category = `ADVERTISING_AND_MARKETING`

const CategoryClimateAndEnvironment Category = `CLIMATE_AND_ENVIRONMENT`

const CategoryCommerce Category = `COMMERCE`

const CategoryDemographics Category = `DEMOGRAPHICS`

const CategoryEconomics Category = `ECONOMICS`

const CategoryEducation Category = `EDUCATION`

const CategoryEnergy Category = `ENERGY`

const CategoryFinancial Category = `FINANCIAL`

const CategoryGaming Category = `GAMING`

const CategoryGeospatial Category = `GEOSPATIAL`

const CategoryHealth Category = `HEALTH`

const CategoryLookupTables Category = `LOOKUP_TABLES`

const CategoryManufacturing Category = `MANUFACTURING`

const CategoryMedia Category = `MEDIA`

const CategoryOther Category = `OTHER`

const CategoryPublicSector Category = `PUBLIC_SECTOR`

const CategoryRetail Category = `RETAIL`

const CategoryScienceAndResearch Category = `SCIENCE_AND_RESEARCH`

const CategorySecurity Category = `SECURITY`

const CategorySports Category = `SPORTS`

const CategoryTransportationAndLogistics Category = `TRANSPORTATION_AND_LOGISTICS`

const CategoryTravelAndTourism Category = `TRAVEL_AND_TOURISM`

// String representation for [fmt.Print]
func (f *Category) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Category) Set(v string) error {
	switch v {
	case `ADVERTISING_AND_MARKETING`, `CLIMATE_AND_ENVIRONMENT`, `COMMERCE`, `DEMOGRAPHICS`, `ECONOMICS`, `EDUCATION`, `ENERGY`, `FINANCIAL`, `GAMING`, `GEOSPATIAL`, `HEALTH`, `LOOKUP_TABLES`, `MANUFACTURING`, `MEDIA`, `OTHER`, `PUBLIC_SECTOR`, `RETAIL`, `SCIENCE_AND_RESEARCH`, `SECURITY`, `SPORTS`, `TRANSPORTATION_AND_LOGISTICS`, `TRAVEL_AND_TOURISM`:
		*f = Category(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADVERTISING_AND_MARKETING", "CLIMATE_AND_ENVIRONMENT", "COMMERCE", "DEMOGRAPHICS", "ECONOMICS", "EDUCATION", "ENERGY", "FINANCIAL", "GAMING", "GEOSPATIAL", "HEALTH", "LOOKUP_TABLES", "MANUFACTURING", "MEDIA", "OTHER", "PUBLIC_SECTOR", "RETAIL", "SCIENCE_AND_RESEARCH", "SECURITY", "SPORTS", "TRANSPORTATION_AND_LOGISTICS", "TRAVEL_AND_TOURISM"`, v)
	}
}

// Type always returns Category to satisfy [pflag.Value] interface
func (f *Category) Type() string {
	return "Category"
}

func categoryToPb(st *Category) (*categoryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := categoryPb(*st)
	return &pb, nil
}

func categoryFromPb(pb *categoryPb) (*Category, error) {
	if pb == nil {
		return nil, nil
	}
	st := Category(*pb)
	return &st, nil
}

type ConsumerTerms struct {

	// Wire name: 'version'
	Version string
}

func consumerTermsToPb(st *ConsumerTerms) (*consumerTermsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &consumerTermsPb{}
	pb.Version = st.Version

	return pb, nil
}

func (st *ConsumerTerms) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &consumerTermsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := consumerTermsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ConsumerTerms) MarshalJSON() ([]byte, error) {
	pb, err := consumerTermsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {

	// Wire name: 'company'
	Company string

	// Wire name: 'email'
	Email string

	// Wire name: 'first_name'
	FirstName string

	// Wire name: 'last_name'
	LastName string

	ForceSendFields []string `tf:"-"`
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

func (st *ContactInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &contactInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := contactInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ContactInfo) MarshalJSON() ([]byte, error) {
	pb, err := contactInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type contactInfoPb struct {
	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

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

type Cost string
type costPb string

const CostFree Cost = `FREE`

const CostPaid Cost = `PAID`

// String representation for [fmt.Print]
func (f *Cost) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Cost) Set(v string) error {
	switch v {
	case `FREE`, `PAID`:
		*f = Cost(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FREE", "PAID"`, v)
	}
}

// Type always returns Cost to satisfy [pflag.Value] interface
func (f *Cost) Type() string {
	return "Cost"
}

func costToPb(st *Cost) (*costPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := costPb(*st)
	return &pb, nil
}

func costFromPb(pb *costPb) (*Cost, error) {
	if pb == nil {
		return nil, nil
	}
	st := Cost(*pb)
	return &st, nil
}

type CreateExchangeFilterRequest struct {

	// Wire name: 'filter'
	Filter ExchangeFilter
}

func createExchangeFilterRequestToPb(st *CreateExchangeFilterRequest) (*createExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExchangeFilterRequestPb{}
	filterPb, err := exchangeFilterToPb(&st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = *filterPb
	}

	return pb, nil
}

func (st *CreateExchangeFilterRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExchangeFilterRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExchangeFilterRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExchangeFilterRequest) MarshalJSON() ([]byte, error) {
	pb, err := createExchangeFilterRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createExchangeFilterRequestPb struct {
	Filter exchangeFilterPb `json:"filter"`
}

func createExchangeFilterRequestFromPb(pb *createExchangeFilterRequestPb) (*CreateExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeFilterRequest{}
	filterField, err := exchangeFilterFromPb(&pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = *filterField
	}

	return st, nil
}

type CreateExchangeFilterResponse struct {

	// Wire name: 'filter_id'
	FilterId string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExchangeFilterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExchangeFilterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	pb, err := createExchangeFilterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type CreateExchangeRequest struct {

	// Wire name: 'exchange'
	Exchange Exchange
}

func createExchangeRequestToPb(st *CreateExchangeRequest) (*createExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExchangeRequestPb{}
	exchangePb, err := exchangeToPb(&st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = *exchangePb
	}

	return pb, nil
}

func (st *CreateExchangeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExchangeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExchangeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExchangeRequest) MarshalJSON() ([]byte, error) {
	pb, err := createExchangeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createExchangeRequestPb struct {
	Exchange exchangePb `json:"exchange"`
}

func createExchangeRequestFromPb(pb *createExchangeRequestPb) (*CreateExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeRequest{}
	exchangeField, err := exchangeFromPb(&pb.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangeField != nil {
		st.Exchange = *exchangeField
	}

	return st, nil
}

type CreateExchangeResponse struct {

	// Wire name: 'exchange_id'
	ExchangeId string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateExchangeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExchangeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExchangeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExchangeResponse) MarshalJSON() ([]byte, error) {
	pb, err := createExchangeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type CreateFileRequest struct {

	// Wire name: 'display_name'
	DisplayName string

	// Wire name: 'file_parent'
	FileParent FileParent

	// Wire name: 'marketplace_file_type'
	MarketplaceFileType MarketplaceFileType

	// Wire name: 'mime_type'
	MimeType string

	ForceSendFields []string `tf:"-"`
}

func createFileRequestToPb(st *CreateFileRequest) (*createFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFileRequestPb{}
	pb.DisplayName = st.DisplayName

	fileParentPb, err := fileParentToPb(&st.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentPb != nil {
		pb.FileParent = *fileParentPb
	}

	pb.MarketplaceFileType = st.MarketplaceFileType

	pb.MimeType = st.MimeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateFileRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createFileRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFileRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateFileRequest) MarshalJSON() ([]byte, error) {
	pb, err := createFileRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createFileRequestPb struct {
	DisplayName string `json:"display_name,omitempty"`

	FileParent fileParentPb `json:"file_parent"`

	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type"`

	MimeType string `json:"mime_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFileRequestFromPb(pb *createFileRequestPb) (*CreateFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFileRequest{}
	st.DisplayName = pb.DisplayName
	fileParentField, err := fileParentFromPb(&pb.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentField != nil {
		st.FileParent = *fileParentField
	}
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

type CreateFileResponse struct {

	// Wire name: 'file_info'
	FileInfo *FileInfo
	// Pre-signed POST URL to blob storage
	// Wire name: 'upload_url'
	UploadUrl string

	ForceSendFields []string `tf:"-"`
}

func createFileResponseToPb(st *CreateFileResponse) (*createFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFileResponsePb{}
	fileInfoPb, err := fileInfoToPb(st.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoPb != nil {
		pb.FileInfo = fileInfoPb
	}

	pb.UploadUrl = st.UploadUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateFileResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createFileResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFileResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateFileResponse) MarshalJSON() ([]byte, error) {
	pb, err := createFileResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createFileResponsePb struct {
	FileInfo *fileInfoPb `json:"file_info,omitempty"`
	// Pre-signed POST URL to blob storage
	UploadUrl string `json:"upload_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFileResponseFromPb(pb *createFileResponsePb) (*CreateFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFileResponse{}
	fileInfoField, err := fileInfoFromPb(pb.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoField != nil {
		st.FileInfo = fileInfoField
	}
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

type CreateInstallationRequest struct {

	// Wire name: 'accepted_consumer_terms'
	AcceptedConsumerTerms *ConsumerTerms

	// Wire name: 'catalog_name'
	CatalogName string

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType
	// for git repo installations
	// Wire name: 'repo_detail'
	RepoDetail *RepoInstallation

	// Wire name: 'share_name'
	ShareName string

	ForceSendFields []string `tf:"-"`
}

func createInstallationRequestToPb(st *CreateInstallationRequest) (*createInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createInstallationRequestPb{}
	acceptedConsumerTermsPb, err := consumerTermsToPb(st.AcceptedConsumerTerms)
	if err != nil {
		return nil, err
	}
	if acceptedConsumerTermsPb != nil {
		pb.AcceptedConsumerTerms = acceptedConsumerTermsPb
	}

	pb.CatalogName = st.CatalogName

	pb.ListingId = st.ListingId

	pb.RecipientType = st.RecipientType

	repoDetailPb, err := repoInstallationToPb(st.RepoDetail)
	if err != nil {
		return nil, err
	}
	if repoDetailPb != nil {
		pb.RepoDetail = repoDetailPb
	}

	pb.ShareName = st.ShareName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateInstallationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createInstallationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createInstallationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateInstallationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createInstallationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createInstallationRequestPb struct {
	AcceptedConsumerTerms *consumerTermsPb `json:"accepted_consumer_terms,omitempty"`

	CatalogName string `json:"catalog_name,omitempty"`

	ListingId string `json:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	// for git repo installations
	RepoDetail *repoInstallationPb `json:"repo_detail,omitempty"`

	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createInstallationRequestFromPb(pb *createInstallationRequestPb) (*CreateInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstallationRequest{}
	acceptedConsumerTermsField, err := consumerTermsFromPb(pb.AcceptedConsumerTerms)
	if err != nil {
		return nil, err
	}
	if acceptedConsumerTermsField != nil {
		st.AcceptedConsumerTerms = acceptedConsumerTermsField
	}
	st.CatalogName = pb.CatalogName
	st.ListingId = pb.ListingId
	st.RecipientType = pb.RecipientType
	repoDetailField, err := repoInstallationFromPb(pb.RepoDetail)
	if err != nil {
		return nil, err
	}
	if repoDetailField != nil {
		st.RepoDetail = repoDetailField
	}
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

type CreateListingRequest struct {

	// Wire name: 'listing'
	Listing Listing
}

func createListingRequestToPb(st *CreateListingRequest) (*createListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createListingRequestPb{}
	listingPb, err := listingToPb(&st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = *listingPb
	}

	return pb, nil
}

func (st *CreateListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := createListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createListingRequestPb struct {
	Listing listingPb `json:"listing"`
}

func createListingRequestFromPb(pb *createListingRequestPb) (*CreateListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateListingRequest{}
	listingField, err := listingFromPb(&pb.Listing)
	if err != nil {
		return nil, err
	}
	if listingField != nil {
		st.Listing = *listingField
	}

	return st, nil
}

type CreateListingResponse struct {

	// Wire name: 'listing_id'
	ListingId string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := createListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {

	// Wire name: 'accepted_consumer_terms'
	AcceptedConsumerTerms ConsumerTerms

	// Wire name: 'comment'
	Comment string

	// Wire name: 'company'
	Company string

	// Wire name: 'first_name'
	FirstName string

	// Wire name: 'intended_use'
	IntendedUse string

	// Wire name: 'is_from_lighthouse'
	IsFromLighthouse bool

	// Wire name: 'last_name'
	LastName string

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType

	ForceSendFields []string `tf:"-"`
}

func createPersonalizationRequestToPb(st *CreatePersonalizationRequest) (*createPersonalizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPersonalizationRequestPb{}
	acceptedConsumerTermsPb, err := consumerTermsToPb(&st.AcceptedConsumerTerms)
	if err != nil {
		return nil, err
	}
	if acceptedConsumerTermsPb != nil {
		pb.AcceptedConsumerTerms = *acceptedConsumerTermsPb
	}

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

func (st *CreatePersonalizationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPersonalizationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPersonalizationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePersonalizationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createPersonalizationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createPersonalizationRequestPb struct {
	AcceptedConsumerTerms consumerTermsPb `json:"accepted_consumer_terms"`

	Comment string `json:"comment,omitempty"`

	Company string `json:"company,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	IntendedUse string `json:"intended_use"`

	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`

	LastName string `json:"last_name,omitempty"`

	ListingId string `json:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPersonalizationRequestFromPb(pb *createPersonalizationRequestPb) (*CreatePersonalizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePersonalizationRequest{}
	acceptedConsumerTermsField, err := consumerTermsFromPb(&pb.AcceptedConsumerTerms)
	if err != nil {
		return nil, err
	}
	if acceptedConsumerTermsField != nil {
		st.AcceptedConsumerTerms = *acceptedConsumerTermsField
	}
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

type CreatePersonalizationRequestResponse struct {

	// Wire name: 'id'
	Id string

	ForceSendFields []string `tf:"-"`
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

func (st *CreatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPersonalizationRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPersonalizationRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := createPersonalizationRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type CreateProviderRequest struct {

	// Wire name: 'provider'
	Provider ProviderInfo
}

func createProviderRequestToPb(st *CreateProviderRequest) (*createProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createProviderRequestPb{}
	providerPb, err := providerInfoToPb(&st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = *providerPb
	}

	return pb, nil
}

func (st *CreateProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := createProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createProviderRequestPb struct {
	Provider providerInfoPb `json:"provider"`
}

func createProviderRequestFromPb(pb *createProviderRequestPb) (*CreateProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProviderRequest{}
	providerField, err := providerInfoFromPb(&pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = *providerField
	}

	return st, nil
}

type CreateProviderResponse struct {

	// Wire name: 'id'
	Id string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateProviderResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createProviderResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createProviderResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateProviderResponse) MarshalJSON() ([]byte, error) {
	pb, err := createProviderResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DataRefresh string
type dataRefreshPb string

const DataRefreshDaily DataRefresh = `DAILY`

const DataRefreshHourly DataRefresh = `HOURLY`

const DataRefreshMinute DataRefresh = `MINUTE`

const DataRefreshMonthly DataRefresh = `MONTHLY`

const DataRefreshNone DataRefresh = `NONE`

const DataRefreshQuarterly DataRefresh = `QUARTERLY`

const DataRefreshSecond DataRefresh = `SECOND`

const DataRefreshWeekly DataRefresh = `WEEKLY`

const DataRefreshYearly DataRefresh = `YEARLY`

// String representation for [fmt.Print]
func (f *DataRefresh) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataRefresh) Set(v string) error {
	switch v {
	case `DAILY`, `HOURLY`, `MINUTE`, `MONTHLY`, `NONE`, `QUARTERLY`, `SECOND`, `WEEKLY`, `YEARLY`:
		*f = DataRefresh(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAILY", "HOURLY", "MINUTE", "MONTHLY", "NONE", "QUARTERLY", "SECOND", "WEEKLY", "YEARLY"`, v)
	}
}

// Type always returns DataRefresh to satisfy [pflag.Value] interface
func (f *DataRefresh) Type() string {
	return "DataRefresh"
}

func dataRefreshToPb(st *DataRefresh) (*dataRefreshPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dataRefreshPb(*st)
	return &pb, nil
}

func dataRefreshFromPb(pb *dataRefreshPb) (*DataRefresh, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataRefresh(*pb)
	return &st, nil
}

type DataRefreshInfo struct {

	// Wire name: 'interval'
	Interval int64

	// Wire name: 'unit'
	Unit DataRefresh
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

func (st *DataRefreshInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dataRefreshInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dataRefreshInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DataRefreshInfo) MarshalJSON() ([]byte, error) {
	pb, err := dataRefreshInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dataRefreshInfoPb struct {
	Interval int64 `json:"interval"`

	Unit DataRefresh `json:"unit"`
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

// Delete an exchange filter
type DeleteExchangeFilterRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteExchangeFilterRequestToPb(st *DeleteExchangeFilterRequest) (*deleteExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeFilterRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *DeleteExchangeFilterRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExchangeFilterRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExchangeFilterRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExchangeFilterRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteExchangeFilterRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DeleteExchangeFilterResponse struct {
}

func deleteExchangeFilterResponseToPb(st *DeleteExchangeFilterResponse) (*deleteExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeFilterResponsePb{}

	return pb, nil
}

func (st *DeleteExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExchangeFilterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExchangeFilterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteExchangeFilterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Delete an exchange
type DeleteExchangeRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteExchangeRequestToPb(st *DeleteExchangeRequest) (*deleteExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *DeleteExchangeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExchangeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExchangeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExchangeRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteExchangeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DeleteExchangeResponse struct {
}

func deleteExchangeResponseToPb(st *DeleteExchangeResponse) (*deleteExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExchangeResponsePb{}

	return pb, nil
}

func (st *DeleteExchangeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExchangeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExchangeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExchangeResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteExchangeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Delete a file
type DeleteFileRequest struct {

	// Wire name: 'file_id'
	FileId string `tf:"-"`
}

func deleteFileRequestToPb(st *DeleteFileRequest) (*deleteFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFileRequestPb{}
	pb.FileId = st.FileId

	return pb, nil
}

func (st *DeleteFileRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFileRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFileRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFileRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteFileRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DeleteFileResponse struct {
}

func deleteFileResponseToPb(st *DeleteFileResponse) (*deleteFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFileResponsePb{}

	return pb, nil
}

func (st *DeleteFileResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFileResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFileResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFileResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteFileResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Uninstall from a listing
type DeleteInstallationRequest struct {

	// Wire name: 'installation_id'
	InstallationId string `tf:"-"`

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`
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

func (st *DeleteInstallationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteInstallationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteInstallationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteInstallationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteInstallationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteInstallationRequestPb struct {
	InstallationId string `json:"-" url:"-"`

	ListingId string `json:"-" url:"-"`
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

type DeleteInstallationResponse struct {
}

func deleteInstallationResponseToPb(st *DeleteInstallationResponse) (*deleteInstallationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteInstallationResponsePb{}

	return pb, nil
}

func (st *DeleteInstallationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteInstallationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteInstallationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteInstallationResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteInstallationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Delete a listing
type DeleteListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteListingRequestToPb(st *DeleteListingRequest) (*deleteListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *DeleteListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DeleteListingResponse struct {
}

func deleteListingResponseToPb(st *DeleteListingResponse) (*deleteListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteListingResponsePb{}

	return pb, nil
}

func (st *DeleteListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Delete provider
type DeleteProviderRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteProviderRequestToPb(st *DeleteProviderRequest) (*deleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *DeleteProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DeleteProviderResponse struct {
}

func deleteProviderResponseToPb(st *DeleteProviderResponse) (*deleteProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderResponsePb{}

	return pb, nil
}

func (st *DeleteProviderResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteProviderResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteProviderResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteProviderResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteProviderResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type DeltaSharingRecipientType string
type deltaSharingRecipientTypePb string

const DeltaSharingRecipientTypeDeltaSharingRecipientTypeDatabricks DeltaSharingRecipientType = `DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS`

const DeltaSharingRecipientTypeDeltaSharingRecipientTypeOpen DeltaSharingRecipientType = `DELTA_SHARING_RECIPIENT_TYPE_OPEN`

// String representation for [fmt.Print]
func (f *DeltaSharingRecipientType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeltaSharingRecipientType) Set(v string) error {
	switch v {
	case `DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS`, `DELTA_SHARING_RECIPIENT_TYPE_OPEN`:
		*f = DeltaSharingRecipientType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS", "DELTA_SHARING_RECIPIENT_TYPE_OPEN"`, v)
	}
}

// Type always returns DeltaSharingRecipientType to satisfy [pflag.Value] interface
func (f *DeltaSharingRecipientType) Type() string {
	return "DeltaSharingRecipientType"
}

func deltaSharingRecipientTypeToPb(st *DeltaSharingRecipientType) (*deltaSharingRecipientTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := deltaSharingRecipientTypePb(*st)
	return &pb, nil
}

func deltaSharingRecipientTypeFromPb(pb *deltaSharingRecipientTypePb) (*DeltaSharingRecipientType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeltaSharingRecipientType(*pb)
	return &st, nil
}

type Exchange struct {

	// Wire name: 'comment'
	Comment string

	// Wire name: 'created_at'
	CreatedAt int64

	// Wire name: 'created_by'
	CreatedBy string

	// Wire name: 'filters'
	Filters []ExchangeFilter

	// Wire name: 'id'
	Id string

	// Wire name: 'linked_listings'
	LinkedListings []ExchangeListing

	// Wire name: 'name'
	Name string

	// Wire name: 'updated_at'
	UpdatedAt int64

	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func exchangeToPb(st *Exchange) (*exchangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangePb{}
	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	var filtersPb []exchangeFilterPb
	for _, item := range st.Filters {
		itemPb, err := exchangeFilterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			filtersPb = append(filtersPb, *itemPb)
		}
	}
	pb.Filters = filtersPb

	pb.Id = st.Id

	var linkedListingsPb []exchangeListingPb
	for _, item := range st.LinkedListings {
		itemPb, err := exchangeListingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			linkedListingsPb = append(linkedListingsPb, *itemPb)
		}
	}
	pb.LinkedListings = linkedListingsPb

	pb.Name = st.Name

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Exchange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exchangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exchangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Exchange) MarshalJSON() ([]byte, error) {
	pb, err := exchangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type exchangePb struct {
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	Filters []exchangeFilterPb `json:"filters,omitempty"`

	Id string `json:"id,omitempty"`

	LinkedListings []exchangeListingPb `json:"linked_listings,omitempty"`

	Name string `json:"name"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

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

	var filtersField []ExchangeFilter
	for _, itemPb := range pb.Filters {
		item, err := exchangeFilterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			filtersField = append(filtersField, *item)
		}
	}
	st.Filters = filtersField
	st.Id = pb.Id

	var linkedListingsField []ExchangeListing
	for _, itemPb := range pb.LinkedListings {
		item, err := exchangeListingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			linkedListingsField = append(linkedListingsField, *item)
		}
	}
	st.LinkedListings = linkedListingsField
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

type ExchangeFilter struct {

	// Wire name: 'created_at'
	CreatedAt int64

	// Wire name: 'created_by'
	CreatedBy string

	// Wire name: 'exchange_id'
	ExchangeId string

	// Wire name: 'filter_type'
	FilterType ExchangeFilterType

	// Wire name: 'filter_value'
	FilterValue string

	// Wire name: 'id'
	Id string

	// Wire name: 'name'
	Name string

	// Wire name: 'updated_at'
	UpdatedAt int64

	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
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

func (st *ExchangeFilter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exchangeFilterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exchangeFilterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExchangeFilter) MarshalJSON() ([]byte, error) {
	pb, err := exchangeFilterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type exchangeFilterPb struct {
	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	ExchangeId string `json:"exchange_id"`

	FilterType ExchangeFilterType `json:"filter_type"`

	FilterValue string `json:"filter_value"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

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

type ExchangeFilterType string
type exchangeFilterTypePb string

const ExchangeFilterTypeGlobalMetastoreId ExchangeFilterType = `GLOBAL_METASTORE_ID`

// String representation for [fmt.Print]
func (f *ExchangeFilterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExchangeFilterType) Set(v string) error {
	switch v {
	case `GLOBAL_METASTORE_ID`:
		*f = ExchangeFilterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GLOBAL_METASTORE_ID"`, v)
	}
}

// Type always returns ExchangeFilterType to satisfy [pflag.Value] interface
func (f *ExchangeFilterType) Type() string {
	return "ExchangeFilterType"
}

func exchangeFilterTypeToPb(st *ExchangeFilterType) (*exchangeFilterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := exchangeFilterTypePb(*st)
	return &pb, nil
}

func exchangeFilterTypeFromPb(pb *exchangeFilterTypePb) (*ExchangeFilterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExchangeFilterType(*pb)
	return &st, nil
}

type ExchangeListing struct {

	// Wire name: 'created_at'
	CreatedAt int64

	// Wire name: 'created_by'
	CreatedBy string

	// Wire name: 'exchange_id'
	ExchangeId string

	// Wire name: 'exchange_name'
	ExchangeName string

	// Wire name: 'id'
	Id string

	// Wire name: 'listing_id'
	ListingId string

	// Wire name: 'listing_name'
	ListingName string

	ForceSendFields []string `tf:"-"`
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

func (st *ExchangeListing) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exchangeListingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exchangeListingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExchangeListing) MarshalJSON() ([]byte, error) {
	pb, err := exchangeListingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type exchangeListingPb struct {
	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	ExchangeId string `json:"exchange_id,omitempty"`

	ExchangeName string `json:"exchange_name,omitempty"`

	Id string `json:"id,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

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

type FileInfo struct {

	// Wire name: 'created_at'
	CreatedAt int64
	// Name displayed to users for applicable files, e.g. embedded notebooks
	// Wire name: 'display_name'
	DisplayName string

	// Wire name: 'download_link'
	DownloadLink string

	// Wire name: 'file_parent'
	FileParent *FileParent

	// Wire name: 'id'
	Id string

	// Wire name: 'marketplace_file_type'
	MarketplaceFileType MarketplaceFileType

	// Wire name: 'mime_type'
	MimeType string

	// Wire name: 'status'
	Status FileStatus
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	// Wire name: 'status_message'
	StatusMessage string

	// Wire name: 'updated_at'
	UpdatedAt int64

	ForceSendFields []string `tf:"-"`
}

func fileInfoToPb(st *FileInfo) (*fileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileInfoPb{}
	pb.CreatedAt = st.CreatedAt

	pb.DisplayName = st.DisplayName

	pb.DownloadLink = st.DownloadLink

	fileParentPb, err := fileParentToPb(st.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentPb != nil {
		pb.FileParent = fileParentPb
	}

	pb.Id = st.Id

	pb.MarketplaceFileType = st.MarketplaceFileType

	pb.MimeType = st.MimeType

	pb.Status = st.Status

	pb.StatusMessage = st.StatusMessage

	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileInfo) MarshalJSON() ([]byte, error) {
	pb, err := fileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type fileInfoPb struct {
	CreatedAt int64 `json:"created_at,omitempty"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName string `json:"display_name,omitempty"`

	DownloadLink string `json:"download_link,omitempty"`

	FileParent *fileParentPb `json:"file_parent,omitempty"`

	Id string `json:"id,omitempty"`

	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type,omitempty"`

	MimeType string `json:"mime_type,omitempty"`

	Status FileStatus `json:"status,omitempty"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	StatusMessage string `json:"status_message,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

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
	fileParentField, err := fileParentFromPb(pb.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentField != nil {
		st.FileParent = fileParentField
	}
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

type FileParent struct {

	// Wire name: 'file_parent_type'
	FileParentType FileParentType
	// TODO make the following fields required
	// Wire name: 'parent_id'
	ParentId string

	ForceSendFields []string `tf:"-"`
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

func (st *FileParent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileParentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileParentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileParent) MarshalJSON() ([]byte, error) {
	pb, err := fileParentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type fileParentPb struct {
	FileParentType FileParentType `json:"file_parent_type,omitempty"`
	// TODO make the following fields required
	ParentId string `json:"parent_id,omitempty"`

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

type FileParentType string
type fileParentTypePb string

const FileParentTypeListing FileParentType = `LISTING`

const FileParentTypeListingResource FileParentType = `LISTING_RESOURCE`

const FileParentTypeProvider FileParentType = `PROVIDER`

// String representation for [fmt.Print]
func (f *FileParentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileParentType) Set(v string) error {
	switch v {
	case `LISTING`, `LISTING_RESOURCE`, `PROVIDER`:
		*f = FileParentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING", "LISTING_RESOURCE", "PROVIDER"`, v)
	}
}

// Type always returns FileParentType to satisfy [pflag.Value] interface
func (f *FileParentType) Type() string {
	return "FileParentType"
}

func fileParentTypeToPb(st *FileParentType) (*fileParentTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := fileParentTypePb(*st)
	return &pb, nil
}

func fileParentTypeFromPb(pb *fileParentTypePb) (*FileParentType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FileParentType(*pb)
	return &st, nil
}

type FileStatus string
type fileStatusPb string

const FileStatusFileStatusPublished FileStatus = `FILE_STATUS_PUBLISHED`

const FileStatusFileStatusSanitizationFailed FileStatus = `FILE_STATUS_SANITIZATION_FAILED`

const FileStatusFileStatusSanitizing FileStatus = `FILE_STATUS_SANITIZING`

const FileStatusFileStatusStaging FileStatus = `FILE_STATUS_STAGING`

// String representation for [fmt.Print]
func (f *FileStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileStatus) Set(v string) error {
	switch v {
	case `FILE_STATUS_PUBLISHED`, `FILE_STATUS_SANITIZATION_FAILED`, `FILE_STATUS_SANITIZING`, `FILE_STATUS_STAGING`:
		*f = FileStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_STATUS_PUBLISHED", "FILE_STATUS_SANITIZATION_FAILED", "FILE_STATUS_SANITIZING", "FILE_STATUS_STAGING"`, v)
	}
}

// Type always returns FileStatus to satisfy [pflag.Value] interface
func (f *FileStatus) Type() string {
	return "FileStatus"
}

func fileStatusToPb(st *FileStatus) (*fileStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := fileStatusPb(*st)
	return &pb, nil
}

func fileStatusFromPb(pb *fileStatusPb) (*FileStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := FileStatus(*pb)
	return &st, nil
}

type FulfillmentType string
type fulfillmentTypePb string

const FulfillmentTypeInstall FulfillmentType = `INSTALL`

const FulfillmentTypeRequestAccess FulfillmentType = `REQUEST_ACCESS`

// String representation for [fmt.Print]
func (f *FulfillmentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FulfillmentType) Set(v string) error {
	switch v {
	case `INSTALL`, `REQUEST_ACCESS`:
		*f = FulfillmentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INSTALL", "REQUEST_ACCESS"`, v)
	}
}

// Type always returns FulfillmentType to satisfy [pflag.Value] interface
func (f *FulfillmentType) Type() string {
	return "FulfillmentType"
}

func fulfillmentTypeToPb(st *FulfillmentType) (*fulfillmentTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := fulfillmentTypePb(*st)
	return &pb, nil
}

func fulfillmentTypeFromPb(pb *fulfillmentTypePb) (*FulfillmentType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FulfillmentType(*pb)
	return &st, nil
}

// Get an exchange
type GetExchangeRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func getExchangeRequestToPb(st *GetExchangeRequest) (*getExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExchangeRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *GetExchangeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExchangeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExchangeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExchangeRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExchangeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type GetExchangeResponse struct {

	// Wire name: 'exchange'
	Exchange *Exchange
}

func getExchangeResponseToPb(st *GetExchangeResponse) (*getExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExchangeResponsePb{}
	exchangePb, err := exchangeToPb(st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = exchangePb
	}

	return pb, nil
}

func (st *GetExchangeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExchangeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExchangeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExchangeResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExchangeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExchangeResponsePb struct {
	Exchange *exchangePb `json:"exchange,omitempty"`
}

func getExchangeResponseFromPb(pb *getExchangeResponsePb) (*GetExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExchangeResponse{}
	exchangeField, err := exchangeFromPb(pb.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangeField != nil {
		st.Exchange = exchangeField
	}

	return st, nil
}

// Get a file
type GetFileRequest struct {

	// Wire name: 'file_id'
	FileId string `tf:"-"`
}

func getFileRequestToPb(st *GetFileRequest) (*getFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFileRequestPb{}
	pb.FileId = st.FileId

	return pb, nil
}

func (st *GetFileRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getFileRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getFileRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetFileRequest) MarshalJSON() ([]byte, error) {
	pb, err := getFileRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type GetFileResponse struct {

	// Wire name: 'file_info'
	FileInfo *FileInfo
}

func getFileResponseToPb(st *GetFileResponse) (*getFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFileResponsePb{}
	fileInfoPb, err := fileInfoToPb(st.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoPb != nil {
		pb.FileInfo = fileInfoPb
	}

	return pb, nil
}

func (st *GetFileResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getFileResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getFileResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetFileResponse) MarshalJSON() ([]byte, error) {
	pb, err := getFileResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getFileResponsePb struct {
	FileInfo *fileInfoPb `json:"file_info,omitempty"`
}

func getFileResponseFromPb(pb *getFileResponsePb) (*GetFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFileResponse{}
	fileInfoField, err := fileInfoFromPb(pb.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoField != nil {
		st.FileInfo = fileInfoField
	}

	return st, nil
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	// Wire name: 'version'
	Version int64

	ForceSendFields []string `tf:"-"`
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

func (st *GetLatestVersionProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLatestVersionProviderAnalyticsDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLatestVersionProviderAnalyticsDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLatestVersionProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := getLatestVersionProviderAnalyticsDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getLatestVersionProviderAnalyticsDashboardResponsePb struct {
	// version here is latest logical version of the dashboard template
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

// Get listing content metadata
type GetListingContentMetadataRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *GetListingContentMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getListingContentMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getListingContentMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetListingContentMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := getListingContentMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getListingContentMetadataRequestPb struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

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

type GetListingContentMetadataResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'shared_data_objects'
	SharedDataObjects []SharedDataObject

	ForceSendFields []string `tf:"-"`
}

func getListingContentMetadataResponseToPb(st *GetListingContentMetadataResponse) (*getListingContentMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingContentMetadataResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var sharedDataObjectsPb []sharedDataObjectPb
	for _, item := range st.SharedDataObjects {
		itemPb, err := sharedDataObjectToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sharedDataObjectsPb = append(sharedDataObjectsPb, *itemPb)
		}
	}
	pb.SharedDataObjects = sharedDataObjectsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetListingContentMetadataResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getListingContentMetadataResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getListingContentMetadataResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetListingContentMetadataResponse) MarshalJSON() ([]byte, error) {
	pb, err := getListingContentMetadataResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getListingContentMetadataResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	SharedDataObjects []sharedDataObjectPb `json:"shared_data_objects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getListingContentMetadataResponseFromPb(pb *getListingContentMetadataResponsePb) (*GetListingContentMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingContentMetadataResponse{}
	st.NextPageToken = pb.NextPageToken

	var sharedDataObjectsField []SharedDataObject
	for _, itemPb := range pb.SharedDataObjects {
		item, err := sharedDataObjectFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sharedDataObjectsField = append(sharedDataObjectsField, *item)
		}
	}
	st.SharedDataObjects = sharedDataObjectsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getListingContentMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getListingContentMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get listing
type GetListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func getListingRequestToPb(st *GetListingRequest) (*getListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *GetListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type GetListingResponse struct {

	// Wire name: 'listing'
	Listing *Listing
}

func getListingResponseToPb(st *GetListingResponse) (*getListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingResponsePb{}
	listingPb, err := listingToPb(st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = listingPb
	}

	return pb, nil
}

func (st *GetListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := getListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getListingResponsePb struct {
	Listing *listingPb `json:"listing,omitempty"`
}

func getListingResponseFromPb(pb *getListingResponsePb) (*GetListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingResponse{}
	listingField, err := listingFromPb(pb.Listing)
	if err != nil {
		return nil, err
	}
	if listingField != nil {
		st.Listing = listingField
	}

	return st, nil
}

// List listings
type GetListingsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *GetListingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getListingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getListingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetListingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getListingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getListingsRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

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

type GetListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func getListingsResponseToPb(st *GetListingsResponse) (*getListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getListingsResponsePb{}

	var listingsPb []listingPb
	for _, item := range st.Listings {
		itemPb, err := listingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			listingsPb = append(listingsPb, *itemPb)
		}
	}
	pb.Listings = listingsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetListingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getListingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getListingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetListingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getListingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getListingsResponsePb struct {
	Listings []listingPb `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getListingsResponseFromPb(pb *getListingsResponsePb) (*GetListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := listingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			listingsField = append(listingsField, *item)
		}
	}
	st.Listings = listingsField
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

// Get the personalization request for a listing
type GetPersonalizationRequestRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`
}

func getPersonalizationRequestRequestToPb(st *GetPersonalizationRequestRequest) (*getPersonalizationRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPersonalizationRequestRequestPb{}
	pb.ListingId = st.ListingId

	return pb, nil
}

func (st *GetPersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPersonalizationRequestRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPersonalizationRequestRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPersonalizationRequestRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPersonalizationRequestRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type GetPersonalizationRequestResponse struct {

	// Wire name: 'personalization_requests'
	PersonalizationRequests []PersonalizationRequest
}

func getPersonalizationRequestResponseToPb(st *GetPersonalizationRequestResponse) (*getPersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPersonalizationRequestResponsePb{}

	var personalizationRequestsPb []personalizationRequestPb
	for _, item := range st.PersonalizationRequests {
		itemPb, err := personalizationRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			personalizationRequestsPb = append(personalizationRequestsPb, *itemPb)
		}
	}
	pb.PersonalizationRequests = personalizationRequestsPb

	return pb, nil
}

func (st *GetPersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPersonalizationRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPersonalizationRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPersonalizationRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPersonalizationRequestResponsePb struct {
	PersonalizationRequests []personalizationRequestPb `json:"personalization_requests,omitempty"`
}

func getPersonalizationRequestResponseFromPb(pb *getPersonalizationRequestResponsePb) (*GetPersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalizationRequestResponse{}

	var personalizationRequestsField []PersonalizationRequest
	for _, itemPb := range pb.PersonalizationRequests {
		item, err := personalizationRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			personalizationRequestsField = append(personalizationRequestsField, *item)
		}
	}
	st.PersonalizationRequests = personalizationRequestsField

	return st, nil
}

// Get a provider
type GetProviderRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func getProviderRequestToPb(st *GetProviderRequest) (*getProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *GetProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := getProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type GetProviderResponse struct {

	// Wire name: 'provider'
	Provider *ProviderInfo
}

func getProviderResponseToPb(st *GetProviderResponse) (*getProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderResponsePb{}
	providerPb, err := providerInfoToPb(st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = providerPb
	}

	return pb, nil
}

func (st *GetProviderResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getProviderResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getProviderResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetProviderResponse) MarshalJSON() ([]byte, error) {
	pb, err := getProviderResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getProviderResponsePb struct {
	Provider *providerInfoPb `json:"provider,omitempty"`
}

func getProviderResponseFromPb(pb *getProviderResponsePb) (*GetProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderResponse{}
	providerField, err := providerInfoFromPb(pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = providerField
	}

	return st, nil
}

type Installation struct {

	// Wire name: 'installation'
	Installation *InstallationDetail
}

func installationToPb(st *Installation) (*installationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &installationPb{}
	installationPb, err := installationDetailToPb(st.Installation)
	if err != nil {
		return nil, err
	}
	if installationPb != nil {
		pb.Installation = installationPb
	}

	return pb, nil
}

func (st *Installation) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &installationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := installationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Installation) MarshalJSON() ([]byte, error) {
	pb, err := installationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type installationPb struct {
	Installation *installationDetailPb `json:"installation,omitempty"`
}

func installationFromPb(pb *installationPb) (*Installation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Installation{}
	installationField, err := installationDetailFromPb(pb.Installation)
	if err != nil {
		return nil, err
	}
	if installationField != nil {
		st.Installation = installationField
	}

	return st, nil
}

type InstallationDetail struct {

	// Wire name: 'catalog_name'
	CatalogName string

	// Wire name: 'error_message'
	ErrorMessage string

	// Wire name: 'id'
	Id string

	// Wire name: 'installed_on'
	InstalledOn int64

	// Wire name: 'listing_id'
	ListingId string

	// Wire name: 'listing_name'
	ListingName string

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType

	// Wire name: 'repo_name'
	RepoName string

	// Wire name: 'repo_path'
	RepoPath string

	// Wire name: 'share_name'
	ShareName string

	// Wire name: 'status'
	Status InstallationStatus

	// Wire name: 'token_detail'
	TokenDetail *TokenDetail

	// Wire name: 'tokens'
	Tokens []TokenInfo

	ForceSendFields []string `tf:"-"`
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

	tokenDetailPb, err := tokenDetailToPb(st.TokenDetail)
	if err != nil {
		return nil, err
	}
	if tokenDetailPb != nil {
		pb.TokenDetail = tokenDetailPb
	}

	var tokensPb []tokenInfoPb
	for _, item := range st.Tokens {
		itemPb, err := tokenInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokensPb = append(tokensPb, *itemPb)
		}
	}
	pb.Tokens = tokensPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *InstallationDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &installationDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := installationDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstallationDetail) MarshalJSON() ([]byte, error) {
	pb, err := installationDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type installationDetailPb struct {
	CatalogName string `json:"catalog_name,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`

	Id string `json:"id,omitempty"`

	InstalledOn int64 `json:"installed_on,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	RepoName string `json:"repo_name,omitempty"`

	RepoPath string `json:"repo_path,omitempty"`

	ShareName string `json:"share_name,omitempty"`

	Status InstallationStatus `json:"status,omitempty"`

	TokenDetail *tokenDetailPb `json:"token_detail,omitempty"`

	Tokens []tokenInfoPb `json:"tokens,omitempty"`

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
	tokenDetailField, err := tokenDetailFromPb(pb.TokenDetail)
	if err != nil {
		return nil, err
	}
	if tokenDetailField != nil {
		st.TokenDetail = tokenDetailField
	}

	var tokensField []TokenInfo
	for _, itemPb := range pb.Tokens {
		item, err := tokenInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokensField = append(tokensField, *item)
		}
	}
	st.Tokens = tokensField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *installationDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st installationDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstallationStatus string
type installationStatusPb string

const InstallationStatusFailed InstallationStatus = `FAILED`

const InstallationStatusInstalled InstallationStatus = `INSTALLED`

// String representation for [fmt.Print]
func (f *InstallationStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstallationStatus) Set(v string) error {
	switch v {
	case `FAILED`, `INSTALLED`:
		*f = InstallationStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "INSTALLED"`, v)
	}
}

// Type always returns InstallationStatus to satisfy [pflag.Value] interface
func (f *InstallationStatus) Type() string {
	return "InstallationStatus"
}

func installationStatusToPb(st *InstallationStatus) (*installationStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := installationStatusPb(*st)
	return &pb, nil
}

func installationStatusFromPb(pb *installationStatusPb) (*InstallationStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstallationStatus(*pb)
	return &st, nil
}

// List all installations
type ListAllInstallationsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListAllInstallationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAllInstallationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAllInstallationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAllInstallationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAllInstallationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAllInstallationsRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

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

type ListAllInstallationsResponse struct {

	// Wire name: 'installations'
	Installations []InstallationDetail

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listAllInstallationsResponseToPb(st *ListAllInstallationsResponse) (*listAllInstallationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllInstallationsResponsePb{}

	var installationsPb []installationDetailPb
	for _, item := range st.Installations {
		itemPb, err := installationDetailToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			installationsPb = append(installationsPb, *itemPb)
		}
	}
	pb.Installations = installationsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListAllInstallationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAllInstallationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAllInstallationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAllInstallationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAllInstallationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAllInstallationsResponsePb struct {
	Installations []installationDetailPb `json:"installations,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAllInstallationsResponseFromPb(pb *listAllInstallationsResponsePb) (*ListAllInstallationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllInstallationsResponse{}

	var installationsField []InstallationDetail
	for _, itemPb := range pb.Installations {
		item, err := installationDetailFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			installationsField = append(installationsField, *item)
		}
	}
	st.Installations = installationsField
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

// List all personalization requests
type ListAllPersonalizationRequestsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListAllPersonalizationRequestsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAllPersonalizationRequestsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAllPersonalizationRequestsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAllPersonalizationRequestsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAllPersonalizationRequestsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAllPersonalizationRequestsRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

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

type ListAllPersonalizationRequestsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'personalization_requests'
	PersonalizationRequests []PersonalizationRequest

	ForceSendFields []string `tf:"-"`
}

func listAllPersonalizationRequestsResponseToPb(st *ListAllPersonalizationRequestsResponse) (*listAllPersonalizationRequestsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllPersonalizationRequestsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var personalizationRequestsPb []personalizationRequestPb
	for _, item := range st.PersonalizationRequests {
		itemPb, err := personalizationRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			personalizationRequestsPb = append(personalizationRequestsPb, *itemPb)
		}
	}
	pb.PersonalizationRequests = personalizationRequestsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListAllPersonalizationRequestsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAllPersonalizationRequestsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAllPersonalizationRequestsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAllPersonalizationRequestsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAllPersonalizationRequestsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAllPersonalizationRequestsResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	PersonalizationRequests []personalizationRequestPb `json:"personalization_requests,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAllPersonalizationRequestsResponseFromPb(pb *listAllPersonalizationRequestsResponsePb) (*ListAllPersonalizationRequestsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllPersonalizationRequestsResponse{}
	st.NextPageToken = pb.NextPageToken

	var personalizationRequestsField []PersonalizationRequest
	for _, itemPb := range pb.PersonalizationRequests {
		item, err := personalizationRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			personalizationRequestsField = append(personalizationRequestsField, *item)
		}
	}
	st.PersonalizationRequests = personalizationRequestsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAllPersonalizationRequestsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAllPersonalizationRequestsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List exchange filters
type ListExchangeFiltersRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListExchangeFiltersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExchangeFiltersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExchangeFiltersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExchangeFiltersRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExchangeFiltersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExchangeFiltersRequestPb struct {
	ExchangeId string `json:"-" url:"exchange_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ListExchangeFiltersResponse struct {

	// Wire name: 'filters'
	Filters []ExchangeFilter

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listExchangeFiltersResponseToPb(st *ListExchangeFiltersResponse) (*listExchangeFiltersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangeFiltersResponsePb{}

	var filtersPb []exchangeFilterPb
	for _, item := range st.Filters {
		itemPb, err := exchangeFilterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			filtersPb = append(filtersPb, *itemPb)
		}
	}
	pb.Filters = filtersPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExchangeFiltersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExchangeFiltersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExchangeFiltersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExchangeFiltersResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExchangeFiltersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExchangeFiltersResponsePb struct {
	Filters []exchangeFilterPb `json:"filters,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangeFiltersResponseFromPb(pb *listExchangeFiltersResponsePb) (*ListExchangeFiltersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangeFiltersResponse{}

	var filtersField []ExchangeFilter
	for _, itemPb := range pb.Filters {
		item, err := exchangeFilterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			filtersField = append(filtersField, *item)
		}
	}
	st.Filters = filtersField
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

// List exchanges for listing
type ListExchangesForListingRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListExchangesForListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExchangesForListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExchangesForListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExchangesForListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExchangesForListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExchangesForListingRequestPb struct {
	ListingId string `json:"-" url:"listing_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

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

type ListExchangesForListingResponse struct {

	// Wire name: 'exchange_listing'
	ExchangeListing []ExchangeListing

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listExchangesForListingResponseToPb(st *ListExchangesForListingResponse) (*listExchangesForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangesForListingResponsePb{}

	var exchangeListingPb []exchangeListingPb
	for _, item := range st.ExchangeListing {
		itemPb, err := exchangeListingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			exchangeListingPb = append(exchangeListingPb, *itemPb)
		}
	}
	pb.ExchangeListing = exchangeListingPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExchangesForListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExchangesForListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExchangesForListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExchangesForListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExchangesForListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExchangesForListingResponsePb struct {
	ExchangeListing []exchangeListingPb `json:"exchange_listing,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangesForListingResponseFromPb(pb *listExchangesForListingResponsePb) (*ListExchangesForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesForListingResponse{}

	var exchangeListingField []ExchangeListing
	for _, itemPb := range pb.ExchangeListing {
		item, err := exchangeListingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			exchangeListingField = append(exchangeListingField, *item)
		}
	}
	st.ExchangeListing = exchangeListingField
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

// List exchanges
type ListExchangesRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListExchangesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExchangesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExchangesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExchangesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExchangesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExchangesRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

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

type ListExchangesResponse struct {

	// Wire name: 'exchanges'
	Exchanges []Exchange

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listExchangesResponseToPb(st *ListExchangesResponse) (*listExchangesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExchangesResponsePb{}

	var exchangesPb []exchangePb
	for _, item := range st.Exchanges {
		itemPb, err := exchangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			exchangesPb = append(exchangesPb, *itemPb)
		}
	}
	pb.Exchanges = exchangesPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExchangesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExchangesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExchangesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExchangesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExchangesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExchangesResponsePb struct {
	Exchanges []exchangePb `json:"exchanges,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExchangesResponseFromPb(pb *listExchangesResponsePb) (*ListExchangesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesResponse{}

	var exchangesField []Exchange
	for _, itemPb := range pb.Exchanges {
		item, err := exchangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			exchangesField = append(exchangesField, *item)
		}
	}
	st.Exchanges = exchangesField
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

// List files
type ListFilesRequest struct {

	// Wire name: 'file_parent'
	FileParent FileParent `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listFilesRequestToPb(st *ListFilesRequest) (*listFilesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFilesRequestPb{}
	fileParentPb, err := fileParentToPb(&st.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentPb != nil {
		pb.FileParent = *fileParentPb
	}

	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListFilesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFilesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFilesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFilesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listFilesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFilesRequestPb struct {
	FileParent fileParentPb `json:"-" url:"file_parent"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFilesRequestFromPb(pb *listFilesRequestPb) (*ListFilesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFilesRequest{}
	fileParentField, err := fileParentFromPb(&pb.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentField != nil {
		st.FileParent = *fileParentField
	}
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

type ListFilesResponse struct {

	// Wire name: 'file_infos'
	FileInfos []FileInfo

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listFilesResponseToPb(st *ListFilesResponse) (*listFilesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFilesResponsePb{}

	var fileInfosPb []fileInfoPb
	for _, item := range st.FileInfos {
		itemPb, err := fileInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fileInfosPb = append(fileInfosPb, *itemPb)
		}
	}
	pb.FileInfos = fileInfosPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListFilesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFilesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFilesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFilesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listFilesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFilesResponsePb struct {
	FileInfos []fileInfoPb `json:"file_infos,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFilesResponseFromPb(pb *listFilesResponsePb) (*ListFilesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFilesResponse{}

	var fileInfosField []FileInfo
	for _, itemPb := range pb.FileInfos {
		item, err := fileInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fileInfosField = append(fileInfosField, *item)
		}
	}
	st.FileInfos = fileInfosField
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

// List all listing fulfillments
type ListFulfillmentsRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListFulfillmentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFulfillmentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFulfillmentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listFulfillmentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFulfillmentsRequestPb struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

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

type ListFulfillmentsResponse struct {

	// Wire name: 'fulfillments'
	Fulfillments []ListingFulfillment

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listFulfillmentsResponseToPb(st *ListFulfillmentsResponse) (*listFulfillmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFulfillmentsResponsePb{}

	var fulfillmentsPb []listingFulfillmentPb
	for _, item := range st.Fulfillments {
		itemPb, err := listingFulfillmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fulfillmentsPb = append(fulfillmentsPb, *itemPb)
		}
	}
	pb.Fulfillments = fulfillmentsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListFulfillmentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFulfillmentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFulfillmentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFulfillmentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listFulfillmentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFulfillmentsResponsePb struct {
	Fulfillments []listingFulfillmentPb `json:"fulfillments,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFulfillmentsResponseFromPb(pb *listFulfillmentsResponsePb) (*ListFulfillmentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFulfillmentsResponse{}

	var fulfillmentsField []ListingFulfillment
	for _, itemPb := range pb.Fulfillments {
		item, err := listingFulfillmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fulfillmentsField = append(fulfillmentsField, *item)
		}
	}
	st.Fulfillments = fulfillmentsField
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

// List installations for a listing
type ListInstallationsRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListInstallationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listInstallationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listInstallationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListInstallationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listInstallationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listInstallationsRequestPb struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

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

type ListInstallationsResponse struct {

	// Wire name: 'installations'
	Installations []InstallationDetail

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listInstallationsResponseToPb(st *ListInstallationsResponse) (*listInstallationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstallationsResponsePb{}

	var installationsPb []installationDetailPb
	for _, item := range st.Installations {
		itemPb, err := installationDetailToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			installationsPb = append(installationsPb, *itemPb)
		}
	}
	pb.Installations = installationsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListInstallationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listInstallationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listInstallationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListInstallationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listInstallationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listInstallationsResponsePb struct {
	Installations []installationDetailPb `json:"installations,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listInstallationsResponseFromPb(pb *listInstallationsResponsePb) (*ListInstallationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstallationsResponse{}

	var installationsField []InstallationDetail
	for _, itemPb := range pb.Installations {
		item, err := installationDetailFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			installationsField = append(installationsField, *item)
		}
	}
	st.Installations = installationsField
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

// List listings for exchange
type ListListingsForExchangeRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListListingsForExchangeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listListingsForExchangeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listListingsForExchangeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListListingsForExchangeRequest) MarshalJSON() ([]byte, error) {
	pb, err := listListingsForExchangeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listListingsForExchangeRequestPb struct {
	ExchangeId string `json:"-" url:"exchange_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ListListingsForExchangeResponse struct {

	// Wire name: 'exchange_listings'
	ExchangeListings []ExchangeListing

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listListingsForExchangeResponseToPb(st *ListListingsForExchangeResponse) (*listListingsForExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listListingsForExchangeResponsePb{}

	var exchangeListingsPb []exchangeListingPb
	for _, item := range st.ExchangeListings {
		itemPb, err := exchangeListingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			exchangeListingsPb = append(exchangeListingsPb, *itemPb)
		}
	}
	pb.ExchangeListings = exchangeListingsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListListingsForExchangeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listListingsForExchangeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listListingsForExchangeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListListingsForExchangeResponse) MarshalJSON() ([]byte, error) {
	pb, err := listListingsForExchangeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listListingsForExchangeResponsePb struct {
	ExchangeListings []exchangeListingPb `json:"exchange_listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listListingsForExchangeResponseFromPb(pb *listListingsForExchangeResponsePb) (*ListListingsForExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsForExchangeResponse{}

	var exchangeListingsField []ExchangeListing
	for _, itemPb := range pb.ExchangeListings {
		item, err := exchangeListingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			exchangeListingsField = append(exchangeListingsField, *item)
		}
	}
	st.ExchangeListings = exchangeListingsField
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

// List listings
type ListListingsRequest struct {
	// Matches any of the following asset types
	// Wire name: 'assets'
	Assets []AssetType `tf:"-"`
	// Matches any of the following categories
	// Wire name: 'categories'
	Categories []Category `tf:"-"`
	// Filters each listing based on if it is free.
	// Wire name: 'is_free'
	IsFree bool `tf:"-"`
	// Filters each listing based on if it is a private exchange.
	// Wire name: 'is_private_exchange'
	IsPrivateExchange bool `tf:"-"`
	// Filters each listing based on whether it is a staff pick.
	// Wire name: 'is_staff_pick'
	IsStaffPick bool `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Matches any of the following provider ids
	// Wire name: 'provider_ids'
	ProviderIds []string `tf:"-"`
	// Matches any of the following tags
	// Wire name: 'tags'
	Tags []ListingTag `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

	var tagsPb []listingTagPb
	for _, item := range st.Tags {
		itemPb, err := listingTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListListingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listListingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listListingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListListingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listListingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listListingsRequestPb struct {
	// Matches any of the following asset types
	Assets []AssetType `json:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `json:"-" url:"categories,omitempty"`
	// Filters each listing based on if it is free.
	IsFree bool `json:"-" url:"is_free,omitempty"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange bool `json:"-" url:"is_private_exchange,omitempty"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick bool `json:"-" url:"is_staff_pick,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []string `json:"-" url:"provider_ids,omitempty"`
	// Matches any of the following tags
	Tags []listingTagPb `json:"-" url:"tags,omitempty"`

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

	var tagsField []ListingTag
	for _, itemPb := range pb.Tags {
		item, err := listingTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listListingsResponseToPb(st *ListListingsResponse) (*listListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listListingsResponsePb{}

	var listingsPb []listingPb
	for _, item := range st.Listings {
		itemPb, err := listingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			listingsPb = append(listingsPb, *itemPb)
		}
	}
	pb.Listings = listingsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListListingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listListingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listListingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListListingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listListingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listListingsResponsePb struct {
	Listings []listingPb `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listListingsResponseFromPb(pb *listListingsResponsePb) (*ListListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := listingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			listingsField = append(listingsField, *item)
		}
	}
	st.Listings = listingsField
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

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string

	// Wire name: 'id'
	Id string

	// Wire name: 'version'
	Version int64

	ForceSendFields []string `tf:"-"`
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

func (st *ListProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProviderAnalyticsDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProviderAnalyticsDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := listProviderAnalyticsDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProviderAnalyticsDashboardResponsePb struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId string `json:"dashboard_id"`

	Id string `json:"id"`

	Version int64 `json:"version,omitempty"`

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

// List providers
type ListProvidersRequest struct {

	// Wire name: 'is_featured'
	IsFeatured bool `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProvidersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProvidersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProvidersRequest) MarshalJSON() ([]byte, error) {
	pb, err := listProvidersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProvidersRequestPb struct {
	IsFeatured bool `json:"-" url:"is_featured,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ListProvidersResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'providers'
	Providers []ProviderInfo

	ForceSendFields []string `tf:"-"`
}

func listProvidersResponseToPb(st *ListProvidersResponse) (*listProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var providersPb []providerInfoPb
	for _, item := range st.Providers {
		itemPb, err := providerInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			providersPb = append(providersPb, *itemPb)
		}
	}
	pb.Providers = providersPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProvidersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProvidersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProvidersResponse) MarshalJSON() ([]byte, error) {
	pb, err := listProvidersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProvidersResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Providers []providerInfoPb `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProvidersResponseFromPb(pb *listProvidersResponsePb) (*ListProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersResponse{}
	st.NextPageToken = pb.NextPageToken

	var providersField []ProviderInfo
	for _, itemPb := range pb.Providers {
		item, err := providerInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			providersField = append(providersField, *item)
		}
	}
	st.Providers = providersField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProvidersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProvidersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Listing struct {

	// Wire name: 'detail'
	Detail *ListingDetail

	// Wire name: 'id'
	Id string

	// Wire name: 'summary'
	Summary ListingSummary

	ForceSendFields []string `tf:"-"`
}

func listingToPb(st *Listing) (*listingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingPb{}
	detailPb, err := listingDetailToPb(st.Detail)
	if err != nil {
		return nil, err
	}
	if detailPb != nil {
		pb.Detail = detailPb
	}

	pb.Id = st.Id

	summaryPb, err := listingSummaryToPb(&st.Summary)
	if err != nil {
		return nil, err
	}
	if summaryPb != nil {
		pb.Summary = *summaryPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Listing) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Listing) MarshalJSON() ([]byte, error) {
	pb, err := listingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listingPb struct {
	Detail *listingDetailPb `json:"detail,omitempty"`

	Id string `json:"id,omitempty"`

	Summary listingSummaryPb `json:"summary"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listingFromPb(pb *listingPb) (*Listing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Listing{}
	detailField, err := listingDetailFromPb(pb.Detail)
	if err != nil {
		return nil, err
	}
	if detailField != nil {
		st.Detail = detailField
	}
	st.Id = pb.Id
	summaryField, err := listingSummaryFromPb(&pb.Summary)
	if err != nil {
		return nil, err
	}
	if summaryField != nil {
		st.Summary = *summaryField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	// Wire name: 'assets'
	Assets []AssetType
	// The ending date timestamp for when the data spans
	// Wire name: 'collection_date_end'
	CollectionDateEnd int64
	// The starting date timestamp for when the data spans
	// Wire name: 'collection_date_start'
	CollectionDateStart int64
	// Smallest unit of time in the dataset
	// Wire name: 'collection_granularity'
	CollectionGranularity *DataRefreshInfo
	// Whether the dataset is free or paid
	// Wire name: 'cost'
	Cost Cost
	// Where/how the data is sourced
	// Wire name: 'data_source'
	DataSource string

	// Wire name: 'description'
	Description string

	// Wire name: 'documentation_link'
	DocumentationLink string

	// Wire name: 'embedded_notebook_file_infos'
	EmbeddedNotebookFileInfos []FileInfo

	// Wire name: 'file_ids'
	FileIds []string
	// Which geo region the listing data is collected from
	// Wire name: 'geographical_coverage'
	GeographicalCoverage string
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	// Wire name: 'license'
	License string
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	// Wire name: 'pricing_model'
	PricingModel string

	// Wire name: 'privacy_policy_link'
	PrivacyPolicyLink string
	// size of the dataset in GB
	// Wire name: 'size'
	Size float64

	// Wire name: 'support_link'
	SupportLink string
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	// Wire name: 'tags'
	Tags []ListingTag

	// Wire name: 'terms_of_service'
	TermsOfService string
	// How often data is updated
	// Wire name: 'update_frequency'
	UpdateFrequency *DataRefreshInfo

	ForceSendFields []string `tf:"-"`
}

func listingDetailToPb(st *ListingDetail) (*listingDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingDetailPb{}
	pb.Assets = st.Assets

	pb.CollectionDateEnd = st.CollectionDateEnd

	pb.CollectionDateStart = st.CollectionDateStart

	collectionGranularityPb, err := dataRefreshInfoToPb(st.CollectionGranularity)
	if err != nil {
		return nil, err
	}
	if collectionGranularityPb != nil {
		pb.CollectionGranularity = collectionGranularityPb
	}

	pb.Cost = st.Cost

	pb.DataSource = st.DataSource

	pb.Description = st.Description

	pb.DocumentationLink = st.DocumentationLink

	var embeddedNotebookFileInfosPb []fileInfoPb
	for _, item := range st.EmbeddedNotebookFileInfos {
		itemPb, err := fileInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddedNotebookFileInfosPb = append(embeddedNotebookFileInfosPb, *itemPb)
		}
	}
	pb.EmbeddedNotebookFileInfos = embeddedNotebookFileInfosPb

	pb.FileIds = st.FileIds

	pb.GeographicalCoverage = st.GeographicalCoverage

	pb.License = st.License

	pb.PricingModel = st.PricingModel

	pb.PrivacyPolicyLink = st.PrivacyPolicyLink

	pb.Size = st.Size

	pb.SupportLink = st.SupportLink

	var tagsPb []listingTagPb
	for _, item := range st.Tags {
		itemPb, err := listingTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.TermsOfService = st.TermsOfService

	updateFrequencyPb, err := dataRefreshInfoToPb(st.UpdateFrequency)
	if err != nil {
		return nil, err
	}
	if updateFrequencyPb != nil {
		pb.UpdateFrequency = updateFrequencyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListingDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listingDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listingDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListingDetail) MarshalJSON() ([]byte, error) {
	pb, err := listingDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listingDetailPb struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets []AssetType `json:"assets,omitempty"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd int64 `json:"collection_date_end,omitempty"`
	// The starting date timestamp for when the data spans
	CollectionDateStart int64 `json:"collection_date_start,omitempty"`
	// Smallest unit of time in the dataset
	CollectionGranularity *dataRefreshInfoPb `json:"collection_granularity,omitempty"`
	// Whether the dataset is free or paid
	Cost Cost `json:"cost,omitempty"`
	// Where/how the data is sourced
	DataSource string `json:"data_source,omitempty"`

	Description string `json:"description,omitempty"`

	DocumentationLink string `json:"documentation_link,omitempty"`

	EmbeddedNotebookFileInfos []fileInfoPb `json:"embedded_notebook_file_infos,omitempty"`

	FileIds []string `json:"file_ids,omitempty"`
	// Which geo region the listing data is collected from
	GeographicalCoverage string `json:"geographical_coverage,omitempty"`
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	License string `json:"license,omitempty"`
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	PricingModel string `json:"pricing_model,omitempty"`

	PrivacyPolicyLink string `json:"privacy_policy_link,omitempty"`
	// size of the dataset in GB
	Size float64 `json:"size,omitempty"`

	SupportLink string `json:"support_link,omitempty"`
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	Tags []listingTagPb `json:"tags,omitempty"`

	TermsOfService string `json:"terms_of_service,omitempty"`
	// How often data is updated
	UpdateFrequency *dataRefreshInfoPb `json:"update_frequency,omitempty"`

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
	collectionGranularityField, err := dataRefreshInfoFromPb(pb.CollectionGranularity)
	if err != nil {
		return nil, err
	}
	if collectionGranularityField != nil {
		st.CollectionGranularity = collectionGranularityField
	}
	st.Cost = pb.Cost
	st.DataSource = pb.DataSource
	st.Description = pb.Description
	st.DocumentationLink = pb.DocumentationLink

	var embeddedNotebookFileInfosField []FileInfo
	for _, itemPb := range pb.EmbeddedNotebookFileInfos {
		item, err := fileInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddedNotebookFileInfosField = append(embeddedNotebookFileInfosField, *item)
		}
	}
	st.EmbeddedNotebookFileInfos = embeddedNotebookFileInfosField
	st.FileIds = pb.FileIds
	st.GeographicalCoverage = pb.GeographicalCoverage
	st.License = pb.License
	st.PricingModel = pb.PricingModel
	st.PrivacyPolicyLink = pb.PrivacyPolicyLink
	st.Size = pb.Size
	st.SupportLink = pb.SupportLink

	var tagsField []ListingTag
	for _, itemPb := range pb.Tags {
		item, err := listingTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.TermsOfService = pb.TermsOfService
	updateFrequencyField, err := dataRefreshInfoFromPb(pb.UpdateFrequency)
	if err != nil {
		return nil, err
	}
	if updateFrequencyField != nil {
		st.UpdateFrequency = updateFrequencyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listingDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listingDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListingFulfillment struct {

	// Wire name: 'fulfillment_type'
	FulfillmentType FulfillmentType

	// Wire name: 'listing_id'
	ListingId string

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType

	// Wire name: 'repo_info'
	RepoInfo *RepoInfo

	// Wire name: 'share_info'
	ShareInfo *ShareInfo
}

func listingFulfillmentToPb(st *ListingFulfillment) (*listingFulfillmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingFulfillmentPb{}
	pb.FulfillmentType = st.FulfillmentType

	pb.ListingId = st.ListingId

	pb.RecipientType = st.RecipientType

	repoInfoPb, err := repoInfoToPb(st.RepoInfo)
	if err != nil {
		return nil, err
	}
	if repoInfoPb != nil {
		pb.RepoInfo = repoInfoPb
	}

	shareInfoPb, err := shareInfoToPb(st.ShareInfo)
	if err != nil {
		return nil, err
	}
	if shareInfoPb != nil {
		pb.ShareInfo = shareInfoPb
	}

	return pb, nil
}

func (st *ListingFulfillment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listingFulfillmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listingFulfillmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListingFulfillment) MarshalJSON() ([]byte, error) {
	pb, err := listingFulfillmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listingFulfillmentPb struct {
	FulfillmentType FulfillmentType `json:"fulfillment_type,omitempty"`

	ListingId string `json:"listing_id"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	RepoInfo *repoInfoPb `json:"repo_info,omitempty"`

	ShareInfo *shareInfoPb `json:"share_info,omitempty"`
}

func listingFulfillmentFromPb(pb *listingFulfillmentPb) (*ListingFulfillment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingFulfillment{}
	st.FulfillmentType = pb.FulfillmentType
	st.ListingId = pb.ListingId
	st.RecipientType = pb.RecipientType
	repoInfoField, err := repoInfoFromPb(pb.RepoInfo)
	if err != nil {
		return nil, err
	}
	if repoInfoField != nil {
		st.RepoInfo = repoInfoField
	}
	shareInfoField, err := shareInfoFromPb(pb.ShareInfo)
	if err != nil {
		return nil, err
	}
	if shareInfoField != nil {
		st.ShareInfo = shareInfoField
	}

	return st, nil
}

type ListingSetting struct {

	// Wire name: 'visibility'
	Visibility Visibility
}

func listingSettingToPb(st *ListingSetting) (*listingSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listingSettingPb{}
	pb.Visibility = st.Visibility

	return pb, nil
}

func (st *ListingSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listingSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listingSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListingSetting) MarshalJSON() ([]byte, error) {
	pb, err := listingSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type ListingShareType string
type listingShareTypePb string

const ListingShareTypeFull ListingShareType = `FULL`

const ListingShareTypeSample ListingShareType = `SAMPLE`

// String representation for [fmt.Print]
func (f *ListingShareType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingShareType) Set(v string) error {
	switch v {
	case `FULL`, `SAMPLE`:
		*f = ListingShareType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL", "SAMPLE"`, v)
	}
}

// Type always returns ListingShareType to satisfy [pflag.Value] interface
func (f *ListingShareType) Type() string {
	return "ListingShareType"
}

func listingShareTypeToPb(st *ListingShareType) (*listingShareTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listingShareTypePb(*st)
	return &pb, nil
}

func listingShareTypeFromPb(pb *listingShareTypePb) (*ListingShareType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingShareType(*pb)
	return &st, nil
}

// Enums
type ListingStatus string
type listingStatusPb string

const ListingStatusDraft ListingStatus = `DRAFT`

const ListingStatusPending ListingStatus = `PENDING`

const ListingStatusPublished ListingStatus = `PUBLISHED`

const ListingStatusSuspended ListingStatus = `SUSPENDED`

// String representation for [fmt.Print]
func (f *ListingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingStatus) Set(v string) error {
	switch v {
	case `DRAFT`, `PENDING`, `PUBLISHED`, `SUSPENDED`:
		*f = ListingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DRAFT", "PENDING", "PUBLISHED", "SUSPENDED"`, v)
	}
}

// Type always returns ListingStatus to satisfy [pflag.Value] interface
func (f *ListingStatus) Type() string {
	return "ListingStatus"
}

func listingStatusToPb(st *ListingStatus) (*listingStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listingStatusPb(*st)
	return &pb, nil
}

func listingStatusFromPb(pb *listingStatusPb) (*ListingStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingStatus(*pb)
	return &st, nil
}

type ListingSummary struct {

	// Wire name: 'categories'
	Categories []Category

	// Wire name: 'created_at'
	CreatedAt int64

	// Wire name: 'created_by'
	CreatedBy string

	// Wire name: 'created_by_id'
	CreatedById int64

	// Wire name: 'exchange_ids'
	ExchangeIds []string
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	// Wire name: 'git_repo'
	GitRepo *RepoInfo

	// Wire name: 'listingType'
	ListingType ListingType

	// Wire name: 'name'
	Name string

	// Wire name: 'provider_id'
	ProviderId string

	// Wire name: 'provider_region'
	ProviderRegion *RegionInfo

	// Wire name: 'published_at'
	PublishedAt int64

	// Wire name: 'published_by'
	PublishedBy string

	// Wire name: 'setting'
	Setting *ListingSetting

	// Wire name: 'share'
	Share *ShareInfo
	// Enums
	// Wire name: 'status'
	Status ListingStatus

	// Wire name: 'subtitle'
	Subtitle string

	// Wire name: 'updated_at'
	UpdatedAt int64

	// Wire name: 'updated_by'
	UpdatedBy string

	// Wire name: 'updated_by_id'
	UpdatedById int64

	ForceSendFields []string `tf:"-"`
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

	gitRepoPb, err := repoInfoToPb(st.GitRepo)
	if err != nil {
		return nil, err
	}
	if gitRepoPb != nil {
		pb.GitRepo = gitRepoPb
	}

	pb.ListingType = st.ListingType

	pb.Name = st.Name

	pb.ProviderId = st.ProviderId

	providerRegionPb, err := regionInfoToPb(st.ProviderRegion)
	if err != nil {
		return nil, err
	}
	if providerRegionPb != nil {
		pb.ProviderRegion = providerRegionPb
	}

	pb.PublishedAt = st.PublishedAt

	pb.PublishedBy = st.PublishedBy

	settingPb, err := listingSettingToPb(st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = settingPb
	}

	sharePb, err := shareInfoToPb(st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = sharePb
	}

	pb.Status = st.Status

	pb.Subtitle = st.Subtitle

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.UpdatedById = st.UpdatedById

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListingSummary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listingSummaryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listingSummaryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListingSummary) MarshalJSON() ([]byte, error) {
	pb, err := listingSummaryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listingSummaryPb struct {
	Categories []Category `json:"categories,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	CreatedById int64 `json:"created_by_id,omitempty"`

	ExchangeIds []string `json:"exchange_ids,omitempty"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo *repoInfoPb `json:"git_repo,omitempty"`

	ListingType ListingType `json:"listingType"`

	Name string `json:"name"`

	ProviderId string `json:"provider_id,omitempty"`

	ProviderRegion *regionInfoPb `json:"provider_region,omitempty"`

	PublishedAt int64 `json:"published_at,omitempty"`

	PublishedBy string `json:"published_by,omitempty"`

	Setting *listingSettingPb `json:"setting,omitempty"`

	Share *shareInfoPb `json:"share,omitempty"`
	// Enums
	Status ListingStatus `json:"status,omitempty"`

	Subtitle string `json:"subtitle,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

	UpdatedById int64 `json:"updated_by_id,omitempty"`

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
	gitRepoField, err := repoInfoFromPb(pb.GitRepo)
	if err != nil {
		return nil, err
	}
	if gitRepoField != nil {
		st.GitRepo = gitRepoField
	}
	st.ListingType = pb.ListingType
	st.Name = pb.Name
	st.ProviderId = pb.ProviderId
	providerRegionField, err := regionInfoFromPb(pb.ProviderRegion)
	if err != nil {
		return nil, err
	}
	if providerRegionField != nil {
		st.ProviderRegion = providerRegionField
	}
	st.PublishedAt = pb.PublishedAt
	st.PublishedBy = pb.PublishedBy
	settingField, err := listingSettingFromPb(pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = settingField
	}
	shareField, err := shareInfoFromPb(pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = shareField
	}
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

type ListingTag struct {
	// Tag name (enum)
	// Wire name: 'tag_name'
	TagName ListingTagType
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	// Wire name: 'tag_values'
	TagValues []string
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

func (st *ListingTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listingTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listingTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListingTag) MarshalJSON() ([]byte, error) {
	pb, err := listingTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listingTagPb struct {
	// Tag name (enum)
	TagName ListingTagType `json:"tag_name,omitempty"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues []string `json:"tag_values,omitempty"`
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

type ListingTagType string
type listingTagTypePb string

const ListingTagTypeListingTagTypeLanguage ListingTagType = `LISTING_TAG_TYPE_LANGUAGE`

const ListingTagTypeListingTagTypeTask ListingTagType = `LISTING_TAG_TYPE_TASK`

// String representation for [fmt.Print]
func (f *ListingTagType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingTagType) Set(v string) error {
	switch v {
	case `LISTING_TAG_TYPE_LANGUAGE`, `LISTING_TAG_TYPE_TASK`:
		*f = ListingTagType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING_TAG_TYPE_LANGUAGE", "LISTING_TAG_TYPE_TASK"`, v)
	}
}

// Type always returns ListingTagType to satisfy [pflag.Value] interface
func (f *ListingTagType) Type() string {
	return "ListingTagType"
}

func listingTagTypeToPb(st *ListingTagType) (*listingTagTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listingTagTypePb(*st)
	return &pb, nil
}

func listingTagTypeFromPb(pb *listingTagTypePb) (*ListingTagType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingTagType(*pb)
	return &st, nil
}

type ListingType string
type listingTypePb string

const ListingTypePersonalized ListingType = `PERSONALIZED`

const ListingTypeStandard ListingType = `STANDARD`

// String representation for [fmt.Print]
func (f *ListingType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListingType) Set(v string) error {
	switch v {
	case `PERSONALIZED`, `STANDARD`:
		*f = ListingType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PERSONALIZED", "STANDARD"`, v)
	}
}

// Type always returns ListingType to satisfy [pflag.Value] interface
func (f *ListingType) Type() string {
	return "ListingType"
}

func listingTypeToPb(st *ListingType) (*listingTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listingTypePb(*st)
	return &pb, nil
}

func listingTypeFromPb(pb *listingTypePb) (*ListingType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingType(*pb)
	return &st, nil
}

type MarketplaceFileType string
type marketplaceFileTypePb string

const MarketplaceFileTypeApp MarketplaceFileType = `APP`

const MarketplaceFileTypeEmbeddedNotebook MarketplaceFileType = `EMBEDDED_NOTEBOOK`

const MarketplaceFileTypeProviderIcon MarketplaceFileType = `PROVIDER_ICON`

// String representation for [fmt.Print]
func (f *MarketplaceFileType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MarketplaceFileType) Set(v string) error {
	switch v {
	case `APP`, `EMBEDDED_NOTEBOOK`, `PROVIDER_ICON`:
		*f = MarketplaceFileType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APP", "EMBEDDED_NOTEBOOK", "PROVIDER_ICON"`, v)
	}
}

// Type always returns MarketplaceFileType to satisfy [pflag.Value] interface
func (f *MarketplaceFileType) Type() string {
	return "MarketplaceFileType"
}

func marketplaceFileTypeToPb(st *MarketplaceFileType) (*marketplaceFileTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplaceFileTypePb(*st)
	return &pb, nil
}

func marketplaceFileTypeFromPb(pb *marketplaceFileTypePb) (*MarketplaceFileType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MarketplaceFileType(*pb)
	return &st, nil
}

type PersonalizationRequest struct {

	// Wire name: 'comment'
	Comment string

	// Wire name: 'consumer_region'
	ConsumerRegion RegionInfo
	// contact info for the consumer requesting data or performing a listing
	// installation
	// Wire name: 'contact_info'
	ContactInfo *ContactInfo

	// Wire name: 'created_at'
	CreatedAt int64

	// Wire name: 'id'
	Id string

	// Wire name: 'intended_use'
	IntendedUse string

	// Wire name: 'is_from_lighthouse'
	IsFromLighthouse bool

	// Wire name: 'listing_id'
	ListingId string

	// Wire name: 'listing_name'
	ListingName string

	// Wire name: 'metastore_id'
	MetastoreId string

	// Wire name: 'provider_id'
	ProviderId string

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType

	// Wire name: 'share'
	Share *ShareInfo

	// Wire name: 'status'
	Status PersonalizationRequestStatus

	// Wire name: 'status_message'
	StatusMessage string

	// Wire name: 'updated_at'
	UpdatedAt int64

	ForceSendFields []string `tf:"-"`
}

func personalizationRequestToPb(st *PersonalizationRequest) (*personalizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &personalizationRequestPb{}
	pb.Comment = st.Comment

	consumerRegionPb, err := regionInfoToPb(&st.ConsumerRegion)
	if err != nil {
		return nil, err
	}
	if consumerRegionPb != nil {
		pb.ConsumerRegion = *consumerRegionPb
	}

	contactInfoPb, err := contactInfoToPb(st.ContactInfo)
	if err != nil {
		return nil, err
	}
	if contactInfoPb != nil {
		pb.ContactInfo = contactInfoPb
	}

	pb.CreatedAt = st.CreatedAt

	pb.Id = st.Id

	pb.IntendedUse = st.IntendedUse

	pb.IsFromLighthouse = st.IsFromLighthouse

	pb.ListingId = st.ListingId

	pb.ListingName = st.ListingName

	pb.MetastoreId = st.MetastoreId

	pb.ProviderId = st.ProviderId

	pb.RecipientType = st.RecipientType

	sharePb, err := shareInfoToPb(st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = sharePb
	}

	pb.Status = st.Status

	pb.StatusMessage = st.StatusMessage

	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PersonalizationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &personalizationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := personalizationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PersonalizationRequest) MarshalJSON() ([]byte, error) {
	pb, err := personalizationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type personalizationRequestPb struct {
	Comment string `json:"comment,omitempty"`

	ConsumerRegion regionInfoPb `json:"consumer_region"`
	// contact info for the consumer requesting data or performing a listing
	// installation
	ContactInfo *contactInfoPb `json:"contact_info,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	Id string `json:"id,omitempty"`

	IntendedUse string `json:"intended_use,omitempty"`

	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

	MetastoreId string `json:"metastore_id,omitempty"`

	ProviderId string `json:"provider_id,omitempty"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	Share *shareInfoPb `json:"share,omitempty"`

	Status PersonalizationRequestStatus `json:"status,omitempty"`

	StatusMessage string `json:"status_message,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func personalizationRequestFromPb(pb *personalizationRequestPb) (*PersonalizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalizationRequest{}
	st.Comment = pb.Comment
	consumerRegionField, err := regionInfoFromPb(&pb.ConsumerRegion)
	if err != nil {
		return nil, err
	}
	if consumerRegionField != nil {
		st.ConsumerRegion = *consumerRegionField
	}
	contactInfoField, err := contactInfoFromPb(pb.ContactInfo)
	if err != nil {
		return nil, err
	}
	if contactInfoField != nil {
		st.ContactInfo = contactInfoField
	}
	st.CreatedAt = pb.CreatedAt
	st.Id = pb.Id
	st.IntendedUse = pb.IntendedUse
	st.IsFromLighthouse = pb.IsFromLighthouse
	st.ListingId = pb.ListingId
	st.ListingName = pb.ListingName
	st.MetastoreId = pb.MetastoreId
	st.ProviderId = pb.ProviderId
	st.RecipientType = pb.RecipientType
	shareField, err := shareInfoFromPb(pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = shareField
	}
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

type PersonalizationRequestStatus string
type personalizationRequestStatusPb string

const PersonalizationRequestStatusDenied PersonalizationRequestStatus = `DENIED`

const PersonalizationRequestStatusFulfilled PersonalizationRequestStatus = `FULFILLED`

const PersonalizationRequestStatusNew PersonalizationRequestStatus = `NEW`

const PersonalizationRequestStatusRequestPending PersonalizationRequestStatus = `REQUEST_PENDING`

// String representation for [fmt.Print]
func (f *PersonalizationRequestStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalizationRequestStatus) Set(v string) error {
	switch v {
	case `DENIED`, `FULFILLED`, `NEW`, `REQUEST_PENDING`:
		*f = PersonalizationRequestStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DENIED", "FULFILLED", "NEW", "REQUEST_PENDING"`, v)
	}
}

// Type always returns PersonalizationRequestStatus to satisfy [pflag.Value] interface
func (f *PersonalizationRequestStatus) Type() string {
	return "PersonalizationRequestStatus"
}

func personalizationRequestStatusToPb(st *PersonalizationRequestStatus) (*personalizationRequestStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := personalizationRequestStatusPb(*st)
	return &pb, nil
}

func personalizationRequestStatusFromPb(pb *personalizationRequestStatusPb) (*PersonalizationRequestStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := PersonalizationRequestStatus(*pb)
	return &st, nil
}

type ProviderAnalyticsDashboard struct {

	// Wire name: 'id'
	Id string
}

func providerAnalyticsDashboardToPb(st *ProviderAnalyticsDashboard) (*providerAnalyticsDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerAnalyticsDashboardPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *ProviderAnalyticsDashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &providerAnalyticsDashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := providerAnalyticsDashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ProviderAnalyticsDashboard) MarshalJSON() ([]byte, error) {
	pb, err := providerAnalyticsDashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type ProviderInfo struct {

	// Wire name: 'business_contact_email'
	BusinessContactEmail string

	// Wire name: 'company_website_link'
	CompanyWebsiteLink string

	// Wire name: 'dark_mode_icon_file_id'
	DarkModeIconFileId string

	// Wire name: 'dark_mode_icon_file_path'
	DarkModeIconFilePath string

	// Wire name: 'description'
	Description string

	// Wire name: 'icon_file_id'
	IconFileId string

	// Wire name: 'icon_file_path'
	IconFilePath string

	// Wire name: 'id'
	Id string
	// is_featured is accessible by consumers only
	// Wire name: 'is_featured'
	IsFeatured bool

	// Wire name: 'name'
	Name string

	// Wire name: 'privacy_policy_link'
	PrivacyPolicyLink string
	// published_by is only applicable to data aggregators (e.g. Crux)
	// Wire name: 'published_by'
	PublishedBy string

	// Wire name: 'support_contact_email'
	SupportContactEmail string

	// Wire name: 'term_of_service_link'
	TermOfServiceLink string

	ForceSendFields []string `tf:"-"`
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

func (st *ProviderInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &providerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := providerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ProviderInfo) MarshalJSON() ([]byte, error) {
	pb, err := providerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type providerInfoPb struct {
	BusinessContactEmail string `json:"business_contact_email"`

	CompanyWebsiteLink string `json:"company_website_link,omitempty"`

	DarkModeIconFileId string `json:"dark_mode_icon_file_id,omitempty"`

	DarkModeIconFilePath string `json:"dark_mode_icon_file_path,omitempty"`

	Description string `json:"description,omitempty"`

	IconFileId string `json:"icon_file_id,omitempty"`

	IconFilePath string `json:"icon_file_path,omitempty"`

	Id string `json:"id,omitempty"`
	// is_featured is accessible by consumers only
	IsFeatured bool `json:"is_featured,omitempty"`

	Name string `json:"name"`

	PrivacyPolicyLink string `json:"privacy_policy_link"`
	// published_by is only applicable to data aggregators (e.g. Crux)
	PublishedBy string `json:"published_by,omitempty"`

	SupportContactEmail string `json:"support_contact_email,omitempty"`

	TermOfServiceLink string `json:"term_of_service_link"`

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

type RegionInfo struct {

	// Wire name: 'cloud'
	Cloud string

	// Wire name: 'region'
	Region string

	ForceSendFields []string `tf:"-"`
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

func (st *RegionInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &regionInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := regionInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegionInfo) MarshalJSON() ([]byte, error) {
	pb, err := regionInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type regionInfoPb struct {
	Cloud string `json:"cloud,omitempty"`

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

// Remove an exchange for listing
type RemoveExchangeForListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func removeExchangeForListingRequestToPb(st *RemoveExchangeForListingRequest) (*removeExchangeForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeExchangeForListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *RemoveExchangeForListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &removeExchangeForListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := removeExchangeForListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RemoveExchangeForListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := removeExchangeForListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type RemoveExchangeForListingResponse struct {
}

func removeExchangeForListingResponseToPb(st *RemoveExchangeForListingResponse) (*removeExchangeForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeExchangeForListingResponsePb{}

	return pb, nil
}

func (st *RemoveExchangeForListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &removeExchangeForListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := removeExchangeForListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RemoveExchangeForListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := removeExchangeForListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	// Wire name: 'git_repo_url'
	GitRepoUrl string
}

func repoInfoToPb(st *RepoInfo) (*repoInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoInfoPb{}
	pb.GitRepoUrl = st.GitRepoUrl

	return pb, nil
}

func (st *RepoInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoInfo) MarshalJSON() ([]byte, error) {
	pb, err := repoInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type repoInfoPb struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
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

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	// Wire name: 'repo_name'
	RepoName string
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	// Wire name: 'repo_path'
	RepoPath string
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

func (st *RepoInstallation) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoInstallationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoInstallationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoInstallation) MarshalJSON() ([]byte, error) {
	pb, err := repoInstallationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type repoInstallationPb struct {
	// the user-specified repo name for their installed git repo listing
	RepoName string `json:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
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

// Search listings
type SearchListingsRequest struct {
	// Matches any of the following asset types
	// Wire name: 'assets'
	Assets []AssetType `tf:"-"`
	// Matches any of the following categories
	// Wire name: 'categories'
	Categories []Category `tf:"-"`

	// Wire name: 'is_free'
	IsFree bool `tf:"-"`

	// Wire name: 'is_private_exchange'
	IsPrivateExchange bool `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Matches any of the following provider ids
	// Wire name: 'provider_ids'
	ProviderIds []string `tf:"-"`
	// Fuzzy matches query
	// Wire name: 'query'
	Query string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *SearchListingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchListingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchListingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchListingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := searchListingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchListingsRequestPb struct {
	// Matches any of the following asset types
	Assets []AssetType `json:"-" url:"assets,omitempty"`
	// Matches any of the following categories
	Categories []Category `json:"-" url:"categories,omitempty"`

	IsFree bool `json:"-" url:"is_free,omitempty"`

	IsPrivateExchange bool `json:"-" url:"is_private_exchange,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Matches any of the following provider ids
	ProviderIds []string `json:"-" url:"provider_ids,omitempty"`
	// Fuzzy matches query
	Query string `json:"-" url:"query"`

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

type SearchListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing

	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func searchListingsResponseToPb(st *SearchListingsResponse) (*searchListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchListingsResponsePb{}

	var listingsPb []listingPb
	for _, item := range st.Listings {
		itemPb, err := listingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			listingsPb = append(listingsPb, *itemPb)
		}
	}
	pb.Listings = listingsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchListingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchListingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchListingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchListingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchListingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchListingsResponsePb struct {
	Listings []listingPb `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchListingsResponseFromPb(pb *searchListingsResponsePb) (*SearchListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := listingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			listingsField = append(listingsField, *item)
		}
	}
	st.Listings = listingsField
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

type ShareInfo struct {

	// Wire name: 'name'
	Name string

	// Wire name: 'type'
	Type ListingShareType
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

func (st *ShareInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &shareInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := shareInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ShareInfo) MarshalJSON() ([]byte, error) {
	pb, err := shareInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type shareInfoPb struct {
	Name string `json:"name"`

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

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	// Wire name: 'data_object_type'
	DataObjectType string
	// Name of the shared object
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
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

func (st *SharedDataObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharedDataObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sharedDataObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SharedDataObject) MarshalJSON() ([]byte, error) {
	pb, err := sharedDataObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sharedDataObjectPb struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType string `json:"data_object_type,omitempty"`
	// Name of the shared object
	Name string `json:"name,omitempty"`

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

type TokenDetail struct {

	// Wire name: 'bearerToken'
	BearerToken string

	// Wire name: 'endpoint'
	Endpoint string

	// Wire name: 'expirationTime'
	ExpirationTime string
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	// Wire name: 'shareCredentialsVersion'
	ShareCredentialsVersion int

	ForceSendFields []string `tf:"-"`
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

func (st *TokenDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenDetail) MarshalJSON() ([]byte, error) {
	pb, err := tokenDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tokenDetailPb struct {
	BearerToken string `json:"bearerToken,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

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

type TokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string
	// Time at which this Recipient Token was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of Recipient Token creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// Unique id of the Recipient Token.
	// Wire name: 'id'
	Id string
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of Recipient Token updater.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
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

func (st *TokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := tokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tokenInfoPb struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Recipient Token creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique id of the Recipient Token.
	Id string `json:"id,omitempty"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of Recipient Token updater.
	UpdatedBy string `json:"updated_by,omitempty"`

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

type UpdateExchangeFilterRequest struct {

	// Wire name: 'filter'
	Filter ExchangeFilter

	// Wire name: 'id'
	Id string `tf:"-"`
}

func updateExchangeFilterRequestToPb(st *UpdateExchangeFilterRequest) (*updateExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeFilterRequestPb{}
	filterPb, err := exchangeFilterToPb(&st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = *filterPb
	}

	pb.Id = st.Id

	return pb, nil
}

func (st *UpdateExchangeFilterRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExchangeFilterRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExchangeFilterRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExchangeFilterRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateExchangeFilterRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExchangeFilterRequestPb struct {
	Filter exchangeFilterPb `json:"filter"`

	Id string `json:"-" url:"-"`
}

func updateExchangeFilterRequestFromPb(pb *updateExchangeFilterRequestPb) (*UpdateExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeFilterRequest{}
	filterField, err := exchangeFilterFromPb(&pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = *filterField
	}
	st.Id = pb.Id

	return st, nil
}

type UpdateExchangeFilterResponse struct {

	// Wire name: 'filter'
	Filter *ExchangeFilter
}

func updateExchangeFilterResponseToPb(st *UpdateExchangeFilterResponse) (*updateExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeFilterResponsePb{}
	filterPb, err := exchangeFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	return pb, nil
}

func (st *UpdateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExchangeFilterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExchangeFilterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateExchangeFilterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExchangeFilterResponsePb struct {
	Filter *exchangeFilterPb `json:"filter,omitempty"`
}

func updateExchangeFilterResponseFromPb(pb *updateExchangeFilterResponsePb) (*UpdateExchangeFilterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeFilterResponse{}
	filterField, err := exchangeFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}

	return st, nil
}

type UpdateExchangeRequest struct {

	// Wire name: 'exchange'
	Exchange Exchange

	// Wire name: 'id'
	Id string `tf:"-"`
}

func updateExchangeRequestToPb(st *UpdateExchangeRequest) (*updateExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeRequestPb{}
	exchangePb, err := exchangeToPb(&st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = *exchangePb
	}

	pb.Id = st.Id

	return pb, nil
}

func (st *UpdateExchangeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExchangeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExchangeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExchangeRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateExchangeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExchangeRequestPb struct {
	Exchange exchangePb `json:"exchange"`

	Id string `json:"-" url:"-"`
}

func updateExchangeRequestFromPb(pb *updateExchangeRequestPb) (*UpdateExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeRequest{}
	exchangeField, err := exchangeFromPb(&pb.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangeField != nil {
		st.Exchange = *exchangeField
	}
	st.Id = pb.Id

	return st, nil
}

type UpdateExchangeResponse struct {

	// Wire name: 'exchange'
	Exchange *Exchange
}

func updateExchangeResponseToPb(st *UpdateExchangeResponse) (*updateExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExchangeResponsePb{}
	exchangePb, err := exchangeToPb(st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = exchangePb
	}

	return pb, nil
}

func (st *UpdateExchangeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExchangeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExchangeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExchangeResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateExchangeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExchangeResponsePb struct {
	Exchange *exchangePb `json:"exchange,omitempty"`
}

func updateExchangeResponseFromPb(pb *updateExchangeResponsePb) (*UpdateExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeResponse{}
	exchangeField, err := exchangeFromPb(pb.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangeField != nil {
		st.Exchange = exchangeField
	}

	return st, nil
}

type UpdateInstallationRequest struct {

	// Wire name: 'installation'
	Installation InstallationDetail

	// Wire name: 'installation_id'
	InstallationId string `tf:"-"`

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'rotate_token'
	RotateToken bool

	ForceSendFields []string `tf:"-"`
}

func updateInstallationRequestToPb(st *UpdateInstallationRequest) (*updateInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInstallationRequestPb{}
	installationPb, err := installationDetailToPb(&st.Installation)
	if err != nil {
		return nil, err
	}
	if installationPb != nil {
		pb.Installation = *installationPb
	}

	pb.InstallationId = st.InstallationId

	pb.ListingId = st.ListingId

	pb.RotateToken = st.RotateToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateInstallationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateInstallationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateInstallationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateInstallationRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateInstallationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateInstallationRequestPb struct {
	Installation installationDetailPb `json:"installation"`

	InstallationId string `json:"-" url:"-"`

	ListingId string `json:"-" url:"-"`

	RotateToken bool `json:"rotate_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateInstallationRequestFromPb(pb *updateInstallationRequestPb) (*UpdateInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInstallationRequest{}
	installationField, err := installationDetailFromPb(&pb.Installation)
	if err != nil {
		return nil, err
	}
	if installationField != nil {
		st.Installation = *installationField
	}
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

type UpdateInstallationResponse struct {

	// Wire name: 'installation'
	Installation *InstallationDetail
}

func updateInstallationResponseToPb(st *UpdateInstallationResponse) (*updateInstallationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInstallationResponsePb{}
	installationPb, err := installationDetailToPb(st.Installation)
	if err != nil {
		return nil, err
	}
	if installationPb != nil {
		pb.Installation = installationPb
	}

	return pb, nil
}

func (st *UpdateInstallationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateInstallationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateInstallationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateInstallationResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateInstallationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateInstallationResponsePb struct {
	Installation *installationDetailPb `json:"installation,omitempty"`
}

func updateInstallationResponseFromPb(pb *updateInstallationResponsePb) (*UpdateInstallationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInstallationResponse{}
	installationField, err := installationDetailFromPb(pb.Installation)
	if err != nil {
		return nil, err
	}
	if installationField != nil {
		st.Installation = installationField
	}

	return st, nil
}

type UpdateListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'listing'
	Listing Listing
}

func updateListingRequestToPb(st *UpdateListingRequest) (*updateListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateListingRequestPb{}
	pb.Id = st.Id

	listingPb, err := listingToPb(&st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = *listingPb
	}

	return pb, nil
}

func (st *UpdateListingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateListingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateListingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateListingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateListingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateListingRequestPb struct {
	Id string `json:"-" url:"-"`

	Listing listingPb `json:"listing"`
}

func updateListingRequestFromPb(pb *updateListingRequestPb) (*UpdateListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateListingRequest{}
	st.Id = pb.Id
	listingField, err := listingFromPb(&pb.Listing)
	if err != nil {
		return nil, err
	}
	if listingField != nil {
		st.Listing = *listingField
	}

	return st, nil
}

type UpdateListingResponse struct {

	// Wire name: 'listing'
	Listing *Listing
}

func updateListingResponseToPb(st *UpdateListingResponse) (*updateListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateListingResponsePb{}
	listingPb, err := listingToPb(st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = listingPb
	}

	return pb, nil
}

func (st *UpdateListingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateListingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateListingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateListingResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateListingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateListingResponsePb struct {
	Listing *listingPb `json:"listing,omitempty"`
}

func updateListingResponseFromPb(pb *updateListingResponsePb) (*UpdateListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateListingResponse{}
	listingField, err := listingFromPb(pb.Listing)
	if err != nil {
		return nil, err
	}
	if listingField != nil {
		st.Listing = listingField
	}

	return st, nil
}

type UpdatePersonalizationRequestRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'reason'
	Reason string

	// Wire name: 'request_id'
	RequestId string `tf:"-"`

	// Wire name: 'share'
	Share *ShareInfo

	// Wire name: 'status'
	Status PersonalizationRequestStatus

	ForceSendFields []string `tf:"-"`
}

func updatePersonalizationRequestRequestToPb(st *UpdatePersonalizationRequestRequest) (*updatePersonalizationRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePersonalizationRequestRequestPb{}
	pb.ListingId = st.ListingId

	pb.Reason = st.Reason

	pb.RequestId = st.RequestId

	sharePb, err := shareInfoToPb(st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = sharePb
	}

	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdatePersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePersonalizationRequestRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePersonalizationRequestRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePersonalizationRequestRequest) MarshalJSON() ([]byte, error) {
	pb, err := updatePersonalizationRequestRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updatePersonalizationRequestRequestPb struct {
	ListingId string `json:"-" url:"-"`

	Reason string `json:"reason,omitempty"`

	RequestId string `json:"-" url:"-"`

	Share *shareInfoPb `json:"share,omitempty"`

	Status PersonalizationRequestStatus `json:"status"`

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
	shareField, err := shareInfoFromPb(pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = shareField
	}
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

type UpdatePersonalizationRequestResponse struct {

	// Wire name: 'request'
	Request *PersonalizationRequest
}

func updatePersonalizationRequestResponseToPb(st *UpdatePersonalizationRequestResponse) (*updatePersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePersonalizationRequestResponsePb{}
	requestPb, err := personalizationRequestToPb(st.Request)
	if err != nil {
		return nil, err
	}
	if requestPb != nil {
		pb.Request = requestPb
	}

	return pb, nil
}

func (st *UpdatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePersonalizationRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePersonalizationRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := updatePersonalizationRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updatePersonalizationRequestResponsePb struct {
	Request *personalizationRequestPb `json:"request,omitempty"`
}

func updatePersonalizationRequestResponseFromPb(pb *updatePersonalizationRequestResponsePb) (*UpdatePersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalizationRequestResponse{}
	requestField, err := personalizationRequestFromPb(pb.Request)
	if err != nil {
		return nil, err
	}
	if requestField != nil {
		st.Request = requestField
	}

	return st, nil
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	// Wire name: 'id'
	Id string `tf:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	// Wire name: 'version'
	Version int64

	ForceSendFields []string `tf:"-"`
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

func (st *UpdateProviderAnalyticsDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateProviderAnalyticsDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateProviderAnalyticsDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateProviderAnalyticsDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateProviderAnalyticsDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateProviderAnalyticsDashboardRequestPb struct {
	// id is immutable property and can't be updated.
	Id string `json:"-" url:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version int64 `json:"version,omitempty"`

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

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	// Wire name: 'dashboard_id'
	DashboardId string
	// id & version should be the same as the request
	// Wire name: 'id'
	Id string

	// Wire name: 'version'
	Version int64

	ForceSendFields []string `tf:"-"`
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

func (st *UpdateProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateProviderAnalyticsDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateProviderAnalyticsDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateProviderAnalyticsDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateProviderAnalyticsDashboardResponsePb struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId string `json:"dashboard_id"`
	// id & version should be the same as the request
	Id string `json:"id"`

	Version int64 `json:"version,omitempty"`

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

type UpdateProviderRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'provider'
	Provider ProviderInfo
}

func updateProviderRequestToPb(st *UpdateProviderRequest) (*updateProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderRequestPb{}
	pb.Id = st.Id

	providerPb, err := providerInfoToPb(&st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = *providerPb
	}

	return pb, nil
}

func (st *UpdateProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateProviderRequestPb struct {
	Id string `json:"-" url:"-"`

	Provider providerInfoPb `json:"provider"`
}

func updateProviderRequestFromPb(pb *updateProviderRequestPb) (*UpdateProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderRequest{}
	st.Id = pb.Id
	providerField, err := providerInfoFromPb(&pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = *providerField
	}

	return st, nil
}

type UpdateProviderResponse struct {

	// Wire name: 'provider'
	Provider *ProviderInfo
}

func updateProviderResponseToPb(st *UpdateProviderResponse) (*updateProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderResponsePb{}
	providerPb, err := providerInfoToPb(st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = providerPb
	}

	return pb, nil
}

func (st *UpdateProviderResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateProviderResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateProviderResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateProviderResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateProviderResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateProviderResponsePb struct {
	Provider *providerInfoPb `json:"provider,omitempty"`
}

func updateProviderResponseFromPb(pb *updateProviderResponsePb) (*UpdateProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderResponse{}
	providerField, err := providerInfoFromPb(pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = providerField
	}

	return st, nil
}

type Visibility string
type visibilityPb string

const VisibilityPrivate Visibility = `PRIVATE`

const VisibilityPublic Visibility = `PUBLIC`

// String representation for [fmt.Print]
func (f *Visibility) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Visibility) Set(v string) error {
	switch v {
	case `PRIVATE`, `PUBLIC`:
		*f = Visibility(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PRIVATE", "PUBLIC"`, v)
	}
}

// Type always returns Visibility to satisfy [pflag.Value] interface
func (f *Visibility) Type() string {
	return "Visibility"
}

func visibilityToPb(st *Visibility) (*visibilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := visibilityPb(*st)
	return &pb, nil
}

func visibilityFromPb(pb *visibilityPb) (*Visibility, error) {
	if pb == nil {
		return nil, nil
	}
	st := Visibility(*pb)
	return &st, nil
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
