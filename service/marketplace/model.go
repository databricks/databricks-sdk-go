// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/marketplace/marketplacepb"
)

type AddExchangeForListingRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string ``

	// Wire name: 'listing_id'
	ListingId string ``
}

func AddExchangeForListingRequestToPb(st *AddExchangeForListingRequest) (*marketplacepb.AddExchangeForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.AddExchangeForListingRequestPb{}
	pb.ExchangeId = st.ExchangeId
	pb.ListingId = st.ListingId

	return pb, nil
}

func AddExchangeForListingRequestFromPb(pb *marketplacepb.AddExchangeForListingRequestPb) (*AddExchangeForListingRequest, error) {
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
	ExchangeForListing *ExchangeListing ``
}

func AddExchangeForListingResponseToPb(st *AddExchangeForListingResponse) (*marketplacepb.AddExchangeForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.AddExchangeForListingResponsePb{}
	exchangeForListingPb, err := ExchangeListingToPb(st.ExchangeForListing)
	if err != nil {
		return nil, err
	}
	if exchangeForListingPb != nil {
		pb.ExchangeForListing = exchangeForListingPb
	}

	return pb, nil
}

func AddExchangeForListingResponseFromPb(pb *marketplacepb.AddExchangeForListingResponsePb) (*AddExchangeForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddExchangeForListingResponse{}
	exchangeForListingField, err := ExchangeListingFromPb(pb.ExchangeForListing)
	if err != nil {
		return nil, err
	}
	if exchangeForListingField != nil {
		st.ExchangeForListing = exchangeForListingField
	}

	return st, nil
}

type AssetType string

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

// Values returns all possible values for AssetType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AssetType) Values() []AssetType {
	return []AssetType{
		AssetTypeAssetTypeApp,
		AssetTypeAssetTypeDataTable,
		AssetTypeAssetTypeGitRepo,
		AssetTypeAssetTypeMedia,
		AssetTypeAssetTypeModel,
		AssetTypeAssetTypeNotebook,
		AssetTypeAssetTypePartnerIntegration,
	}
}

// Type always returns AssetType to satisfy [pflag.Value] interface
func (f *AssetType) Type() string {
	return "AssetType"
}

func AssetTypeToPb(st *AssetType) (*marketplacepb.AssetTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.AssetTypePb(*st)
	return &pb, nil
}

func AssetTypeFromPb(pb *marketplacepb.AssetTypePb) (*AssetType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AssetType(*pb)
	return &st, nil
}

type BatchGetListingsRequest struct {

	// Wire name: 'ids'
	Ids []string `tf:"-"`
}

func BatchGetListingsRequestToPb(st *BatchGetListingsRequest) (*marketplacepb.BatchGetListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.BatchGetListingsRequestPb{}
	pb.Ids = st.Ids

	return pb, nil
}

func BatchGetListingsRequestFromPb(pb *marketplacepb.BatchGetListingsRequestPb) (*BatchGetListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetListingsRequest{}
	st.Ids = pb.Ids

	return st, nil
}

type BatchGetListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing ``
}

func BatchGetListingsResponseToPb(st *BatchGetListingsResponse) (*marketplacepb.BatchGetListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.BatchGetListingsResponsePb{}

	var listingsPb []marketplacepb.ListingPb
	for _, item := range st.Listings {
		itemPb, err := ListingToPb(&item)
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

func BatchGetListingsResponseFromPb(pb *marketplacepb.BatchGetListingsResponsePb) (*BatchGetListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := ListingFromPb(&itemPb)
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

type BatchGetProvidersRequest struct {

	// Wire name: 'ids'
	Ids []string `tf:"-"`
}

func BatchGetProvidersRequestToPb(st *BatchGetProvidersRequest) (*marketplacepb.BatchGetProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.BatchGetProvidersRequestPb{}
	pb.Ids = st.Ids

	return pb, nil
}

func BatchGetProvidersRequestFromPb(pb *marketplacepb.BatchGetProvidersRequestPb) (*BatchGetProvidersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetProvidersRequest{}
	st.Ids = pb.Ids

	return st, nil
}

type BatchGetProvidersResponse struct {

	// Wire name: 'providers'
	Providers []ProviderInfo ``
}

func BatchGetProvidersResponseToPb(st *BatchGetProvidersResponse) (*marketplacepb.BatchGetProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.BatchGetProvidersResponsePb{}

	var providersPb []marketplacepb.ProviderInfoPb
	for _, item := range st.Providers {
		itemPb, err := ProviderInfoToPb(&item)
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

func BatchGetProvidersResponseFromPb(pb *marketplacepb.BatchGetProvidersResponsePb) (*BatchGetProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BatchGetProvidersResponse{}

	var providersField []ProviderInfo
	for _, itemPb := range pb.Providers {
		item, err := ProviderInfoFromPb(&itemPb)
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

// Values returns all possible values for Category.
//
// There is no guarantee on the order of the values in the slice.
func (f *Category) Values() []Category {
	return []Category{
		CategoryAdvertisingAndMarketing,
		CategoryClimateAndEnvironment,
		CategoryCommerce,
		CategoryDemographics,
		CategoryEconomics,
		CategoryEducation,
		CategoryEnergy,
		CategoryFinancial,
		CategoryGaming,
		CategoryGeospatial,
		CategoryHealth,
		CategoryLookupTables,
		CategoryManufacturing,
		CategoryMedia,
		CategoryOther,
		CategoryPublicSector,
		CategoryRetail,
		CategoryScienceAndResearch,
		CategorySecurity,
		CategorySports,
		CategoryTransportationAndLogistics,
		CategoryTravelAndTourism,
	}
}

// Type always returns Category to satisfy [pflag.Value] interface
func (f *Category) Type() string {
	return "Category"
}

func CategoryToPb(st *Category) (*marketplacepb.CategoryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.CategoryPb(*st)
	return &pb, nil
}

func CategoryFromPb(pb *marketplacepb.CategoryPb) (*Category, error) {
	if pb == nil {
		return nil, nil
	}
	st := Category(*pb)
	return &st, nil
}

type ConsumerTerms struct {

	// Wire name: 'version'
	Version string ``
}

func ConsumerTermsToPb(st *ConsumerTerms) (*marketplacepb.ConsumerTermsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ConsumerTermsPb{}
	pb.Version = st.Version

	return pb, nil
}

func ConsumerTermsFromPb(pb *marketplacepb.ConsumerTermsPb) (*ConsumerTerms, error) {
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
	Company string ``

	// Wire name: 'email'
	Email string ``

	// Wire name: 'first_name'
	FirstName string ``

	// Wire name: 'last_name'
	LastName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ContactInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContactInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ContactInfoToPb(st *ContactInfo) (*marketplacepb.ContactInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ContactInfoPb{}
	pb.Company = st.Company
	pb.Email = st.Email
	pb.FirstName = st.FirstName
	pb.LastName = st.LastName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ContactInfoFromPb(pb *marketplacepb.ContactInfoPb) (*ContactInfo, error) {
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

type Cost string

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

// Values returns all possible values for Cost.
//
// There is no guarantee on the order of the values in the slice.
func (f *Cost) Values() []Cost {
	return []Cost{
		CostFree,
		CostPaid,
	}
}

// Type always returns Cost to satisfy [pflag.Value] interface
func (f *Cost) Type() string {
	return "Cost"
}

func CostToPb(st *Cost) (*marketplacepb.CostPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.CostPb(*st)
	return &pb, nil
}

func CostFromPb(pb *marketplacepb.CostPb) (*Cost, error) {
	if pb == nil {
		return nil, nil
	}
	st := Cost(*pb)
	return &st, nil
}

type CreateExchangeFilterRequest struct {

	// Wire name: 'filter'
	Filter ExchangeFilter ``
}

func CreateExchangeFilterRequestToPb(st *CreateExchangeFilterRequest) (*marketplacepb.CreateExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateExchangeFilterRequestPb{}
	filterPb, err := ExchangeFilterToPb(&st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = *filterPb
	}

	return pb, nil
}

func CreateExchangeFilterRequestFromPb(pb *marketplacepb.CreateExchangeFilterRequestPb) (*CreateExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeFilterRequest{}
	filterField, err := ExchangeFilterFromPb(&pb.Filter)
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
	FilterId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateExchangeFilterResponseToPb(st *CreateExchangeFilterResponse) (*marketplacepb.CreateExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateExchangeFilterResponsePb{}
	pb.FilterId = st.FilterId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateExchangeFilterResponseFromPb(pb *marketplacepb.CreateExchangeFilterResponsePb) (*CreateExchangeFilterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeFilterResponse{}
	st.FilterId = pb.FilterId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateExchangeRequest struct {

	// Wire name: 'exchange'
	Exchange Exchange ``
}

func CreateExchangeRequestToPb(st *CreateExchangeRequest) (*marketplacepb.CreateExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateExchangeRequestPb{}
	exchangePb, err := ExchangeToPb(&st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = *exchangePb
	}

	return pb, nil
}

func CreateExchangeRequestFromPb(pb *marketplacepb.CreateExchangeRequestPb) (*CreateExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeRequest{}
	exchangeField, err := ExchangeFromPb(&pb.Exchange)
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
	ExchangeId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateExchangeResponseToPb(st *CreateExchangeResponse) (*marketplacepb.CreateExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateExchangeResponsePb{}
	pb.ExchangeId = st.ExchangeId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateExchangeResponseFromPb(pb *marketplacepb.CreateExchangeResponsePb) (*CreateExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExchangeResponse{}
	st.ExchangeId = pb.ExchangeId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateFileRequest struct {

	// Wire name: 'display_name'
	DisplayName string ``

	// Wire name: 'file_parent'
	FileParent FileParent ``

	// Wire name: 'marketplace_file_type'
	MarketplaceFileType MarketplaceFileType ``

	// Wire name: 'mime_type'
	MimeType        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateFileRequestToPb(st *CreateFileRequest) (*marketplacepb.CreateFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateFileRequestPb{}
	pb.DisplayName = st.DisplayName
	fileParentPb, err := FileParentToPb(&st.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentPb != nil {
		pb.FileParent = *fileParentPb
	}
	marketplaceFileTypePb, err := MarketplaceFileTypeToPb(&st.MarketplaceFileType)
	if err != nil {
		return nil, err
	}
	if marketplaceFileTypePb != nil {
		pb.MarketplaceFileType = *marketplaceFileTypePb
	}
	pb.MimeType = st.MimeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateFileRequestFromPb(pb *marketplacepb.CreateFileRequestPb) (*CreateFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFileRequest{}
	st.DisplayName = pb.DisplayName
	fileParentField, err := FileParentFromPb(&pb.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentField != nil {
		st.FileParent = *fileParentField
	}
	marketplaceFileTypeField, err := MarketplaceFileTypeFromPb(&pb.MarketplaceFileType)
	if err != nil {
		return nil, err
	}
	if marketplaceFileTypeField != nil {
		st.MarketplaceFileType = *marketplaceFileTypeField
	}
	st.MimeType = pb.MimeType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateFileResponse struct {

	// Wire name: 'file_info'
	FileInfo *FileInfo ``
	// Pre-signed POST URL to blob storage
	// Wire name: 'upload_url'
	UploadUrl       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateFileResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateFileResponseToPb(st *CreateFileResponse) (*marketplacepb.CreateFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateFileResponsePb{}
	fileInfoPb, err := FileInfoToPb(st.FileInfo)
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

func CreateFileResponseFromPb(pb *marketplacepb.CreateFileResponsePb) (*CreateFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFileResponse{}
	fileInfoField, err := FileInfoFromPb(pb.FileInfo)
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

type CreateInstallationRequest struct {

	// Wire name: 'accepted_consumer_terms'
	AcceptedConsumerTerms *ConsumerTerms ``

	// Wire name: 'catalog_name'
	CatalogName string ``

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType ``
	// for git repo installations
	// Wire name: 'repo_detail'
	RepoDetail *RepoInstallation ``

	// Wire name: 'share_name'
	ShareName       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateInstallationRequestToPb(st *CreateInstallationRequest) (*marketplacepb.CreateInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateInstallationRequestPb{}
	acceptedConsumerTermsPb, err := ConsumerTermsToPb(st.AcceptedConsumerTerms)
	if err != nil {
		return nil, err
	}
	if acceptedConsumerTermsPb != nil {
		pb.AcceptedConsumerTerms = acceptedConsumerTermsPb
	}
	pb.CatalogName = st.CatalogName
	pb.ListingId = st.ListingId
	recipientTypePb, err := DeltaSharingRecipientTypeToPb(&st.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypePb != nil {
		pb.RecipientType = *recipientTypePb
	}
	repoDetailPb, err := RepoInstallationToPb(st.RepoDetail)
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

func CreateInstallationRequestFromPb(pb *marketplacepb.CreateInstallationRequestPb) (*CreateInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstallationRequest{}
	acceptedConsumerTermsField, err := ConsumerTermsFromPb(pb.AcceptedConsumerTerms)
	if err != nil {
		return nil, err
	}
	if acceptedConsumerTermsField != nil {
		st.AcceptedConsumerTerms = acceptedConsumerTermsField
	}
	st.CatalogName = pb.CatalogName
	st.ListingId = pb.ListingId
	recipientTypeField, err := DeltaSharingRecipientTypeFromPb(&pb.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypeField != nil {
		st.RecipientType = *recipientTypeField
	}
	repoDetailField, err := RepoInstallationFromPb(pb.RepoDetail)
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

type CreateListingRequest struct {

	// Wire name: 'listing'
	Listing Listing ``
}

func CreateListingRequestToPb(st *CreateListingRequest) (*marketplacepb.CreateListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateListingRequestPb{}
	listingPb, err := ListingToPb(&st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = *listingPb
	}

	return pb, nil
}

func CreateListingRequestFromPb(pb *marketplacepb.CreateListingRequestPb) (*CreateListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateListingRequest{}
	listingField, err := ListingFromPb(&pb.Listing)
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
	ListingId       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateListingResponseToPb(st *CreateListingResponse) (*marketplacepb.CreateListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateListingResponsePb{}
	pb.ListingId = st.ListingId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateListingResponseFromPb(pb *marketplacepb.CreateListingResponsePb) (*CreateListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateListingResponse{}
	st.ListingId = pb.ListingId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {

	// Wire name: 'accepted_consumer_terms'
	AcceptedConsumerTerms ConsumerTerms ``

	// Wire name: 'comment'
	Comment string ``

	// Wire name: 'company'
	Company string ``

	// Wire name: 'first_name'
	FirstName string ``

	// Wire name: 'intended_use'
	IntendedUse string ``

	// Wire name: 'is_from_lighthouse'
	IsFromLighthouse bool ``

	// Wire name: 'last_name'
	LastName string ``

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'recipient_type'
	RecipientType   DeltaSharingRecipientType ``
	ForceSendFields []string                  `tf:"-"`
}

func (s *CreatePersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreatePersonalizationRequestToPb(st *CreatePersonalizationRequest) (*marketplacepb.CreatePersonalizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreatePersonalizationRequestPb{}
	acceptedConsumerTermsPb, err := ConsumerTermsToPb(&st.AcceptedConsumerTerms)
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
	recipientTypePb, err := DeltaSharingRecipientTypeToPb(&st.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypePb != nil {
		pb.RecipientType = *recipientTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreatePersonalizationRequestFromPb(pb *marketplacepb.CreatePersonalizationRequestPb) (*CreatePersonalizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePersonalizationRequest{}
	acceptedConsumerTermsField, err := ConsumerTermsFromPb(&pb.AcceptedConsumerTerms)
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
	recipientTypeField, err := DeltaSharingRecipientTypeFromPb(&pb.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypeField != nil {
		st.RecipientType = *recipientTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreatePersonalizationRequestResponse struct {

	// Wire name: 'id'
	Id              string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreatePersonalizationRequestResponseToPb(st *CreatePersonalizationRequestResponse) (*marketplacepb.CreatePersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreatePersonalizationRequestResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreatePersonalizationRequestResponseFromPb(pb *marketplacepb.CreatePersonalizationRequestResponsePb) (*CreatePersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePersonalizationRequestResponse{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateProviderRequest struct {

	// Wire name: 'provider'
	Provider ProviderInfo ``
}

func CreateProviderRequestToPb(st *CreateProviderRequest) (*marketplacepb.CreateProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateProviderRequestPb{}
	providerPb, err := ProviderInfoToPb(&st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = *providerPb
	}

	return pb, nil
}

func CreateProviderRequestFromPb(pb *marketplacepb.CreateProviderRequestPb) (*CreateProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProviderRequest{}
	providerField, err := ProviderInfoFromPb(&pb.Provider)
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
	Id              string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateProviderResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProviderResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateProviderResponseToPb(st *CreateProviderResponse) (*marketplacepb.CreateProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.CreateProviderResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateProviderResponseFromPb(pb *marketplacepb.CreateProviderResponsePb) (*CreateProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProviderResponse{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DataRefresh string

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

// Values returns all possible values for DataRefresh.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataRefresh) Values() []DataRefresh {
	return []DataRefresh{
		DataRefreshDaily,
		DataRefreshHourly,
		DataRefreshMinute,
		DataRefreshMonthly,
		DataRefreshNone,
		DataRefreshQuarterly,
		DataRefreshSecond,
		DataRefreshWeekly,
		DataRefreshYearly,
	}
}

// Type always returns DataRefresh to satisfy [pflag.Value] interface
func (f *DataRefresh) Type() string {
	return "DataRefresh"
}

func DataRefreshToPb(st *DataRefresh) (*marketplacepb.DataRefreshPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.DataRefreshPb(*st)
	return &pb, nil
}

func DataRefreshFromPb(pb *marketplacepb.DataRefreshPb) (*DataRefresh, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataRefresh(*pb)
	return &st, nil
}

type DataRefreshInfo struct {

	// Wire name: 'interval'
	Interval int64 ``

	// Wire name: 'unit'
	Unit DataRefresh ``
}

func DataRefreshInfoToPb(st *DataRefreshInfo) (*marketplacepb.DataRefreshInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DataRefreshInfoPb{}
	pb.Interval = st.Interval
	unitPb, err := DataRefreshToPb(&st.Unit)
	if err != nil {
		return nil, err
	}
	if unitPb != nil {
		pb.Unit = *unitPb
	}

	return pb, nil
}

func DataRefreshInfoFromPb(pb *marketplacepb.DataRefreshInfoPb) (*DataRefreshInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataRefreshInfo{}
	st.Interval = pb.Interval
	unitField, err := DataRefreshFromPb(&pb.Unit)
	if err != nil {
		return nil, err
	}
	if unitField != nil {
		st.Unit = *unitField
	}

	return st, nil
}

type DeleteExchangeFilterRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteExchangeFilterRequestToPb(st *DeleteExchangeFilterRequest) (*marketplacepb.DeleteExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DeleteExchangeFilterRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteExchangeFilterRequestFromPb(pb *marketplacepb.DeleteExchangeFilterRequestPb) (*DeleteExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExchangeFilterRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteExchangeRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteExchangeRequestToPb(st *DeleteExchangeRequest) (*marketplacepb.DeleteExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DeleteExchangeRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteExchangeRequestFromPb(pb *marketplacepb.DeleteExchangeRequestPb) (*DeleteExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExchangeRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteFileRequest struct {

	// Wire name: 'file_id'
	FileId string `tf:"-"`
}

func DeleteFileRequestToPb(st *DeleteFileRequest) (*marketplacepb.DeleteFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DeleteFileRequestPb{}
	pb.FileId = st.FileId

	return pb, nil
}

func DeleteFileRequestFromPb(pb *marketplacepb.DeleteFileRequestPb) (*DeleteFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFileRequest{}
	st.FileId = pb.FileId

	return st, nil
}

type DeleteInstallationRequest struct {

	// Wire name: 'installation_id'
	InstallationId string `tf:"-"`

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`
}

func DeleteInstallationRequestToPb(st *DeleteInstallationRequest) (*marketplacepb.DeleteInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DeleteInstallationRequestPb{}
	pb.InstallationId = st.InstallationId
	pb.ListingId = st.ListingId

	return pb, nil
}

func DeleteInstallationRequestFromPb(pb *marketplacepb.DeleteInstallationRequestPb) (*DeleteInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstallationRequest{}
	st.InstallationId = pb.InstallationId
	st.ListingId = pb.ListingId

	return st, nil
}

type DeleteListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteListingRequestToPb(st *DeleteListingRequest) (*marketplacepb.DeleteListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DeleteListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteListingRequestFromPb(pb *marketplacepb.DeleteListingRequestPb) (*DeleteListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteListingRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteProviderRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteProviderRequestToPb(st *DeleteProviderRequest) (*marketplacepb.DeleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.DeleteProviderRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteProviderRequestFromPb(pb *marketplacepb.DeleteProviderRequestPb) (*DeleteProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteProviderRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeltaSharingRecipientType string

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

// Values returns all possible values for DeltaSharingRecipientType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeltaSharingRecipientType) Values() []DeltaSharingRecipientType {
	return []DeltaSharingRecipientType{
		DeltaSharingRecipientTypeDeltaSharingRecipientTypeDatabricks,
		DeltaSharingRecipientTypeDeltaSharingRecipientTypeOpen,
	}
}

// Type always returns DeltaSharingRecipientType to satisfy [pflag.Value] interface
func (f *DeltaSharingRecipientType) Type() string {
	return "DeltaSharingRecipientType"
}

func DeltaSharingRecipientTypeToPb(st *DeltaSharingRecipientType) (*marketplacepb.DeltaSharingRecipientTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.DeltaSharingRecipientTypePb(*st)
	return &pb, nil
}

func DeltaSharingRecipientTypeFromPb(pb *marketplacepb.DeltaSharingRecipientTypePb) (*DeltaSharingRecipientType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeltaSharingRecipientType(*pb)
	return &st, nil
}

type Exchange struct {

	// Wire name: 'comment'
	Comment string ``

	// Wire name: 'created_at'
	CreatedAt int64 ``

	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'filters'
	Filters []ExchangeFilter ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'linked_listings'
	LinkedListings []ExchangeListing ``

	// Wire name: 'name'
	Name string ``

	// Wire name: 'updated_at'
	UpdatedAt int64 ``

	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Exchange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Exchange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExchangeToPb(st *Exchange) (*marketplacepb.ExchangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ExchangePb{}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy

	var filtersPb []marketplacepb.ExchangeFilterPb
	for _, item := range st.Filters {
		itemPb, err := ExchangeFilterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			filtersPb = append(filtersPb, *itemPb)
		}
	}
	pb.Filters = filtersPb
	pb.Id = st.Id

	var linkedListingsPb []marketplacepb.ExchangeListingPb
	for _, item := range st.LinkedListings {
		itemPb, err := ExchangeListingToPb(&item)
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

func ExchangeFromPb(pb *marketplacepb.ExchangePb) (*Exchange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Exchange{}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy

	var filtersField []ExchangeFilter
	for _, itemPb := range pb.Filters {
		item, err := ExchangeFilterFromPb(&itemPb)
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
		item, err := ExchangeListingFromPb(&itemPb)
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

type ExchangeFilter struct {

	// Wire name: 'created_at'
	CreatedAt int64 ``

	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'exchange_id'
	ExchangeId string ``

	// Wire name: 'filter_type'
	FilterType ExchangeFilterType ``

	// Wire name: 'filter_value'
	FilterValue string ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'name'
	Name string ``

	// Wire name: 'updated_at'
	UpdatedAt int64 ``

	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExchangeFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExchangeFilterToPb(st *ExchangeFilter) (*marketplacepb.ExchangeFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ExchangeFilterPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.ExchangeId = st.ExchangeId
	filterTypePb, err := ExchangeFilterTypeToPb(&st.FilterType)
	if err != nil {
		return nil, err
	}
	if filterTypePb != nil {
		pb.FilterType = *filterTypePb
	}
	pb.FilterValue = st.FilterValue
	pb.Id = st.Id
	pb.Name = st.Name
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExchangeFilterFromPb(pb *marketplacepb.ExchangeFilterPb) (*ExchangeFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeFilter{}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.ExchangeId = pb.ExchangeId
	filterTypeField, err := ExchangeFilterTypeFromPb(&pb.FilterType)
	if err != nil {
		return nil, err
	}
	if filterTypeField != nil {
		st.FilterType = *filterTypeField
	}
	st.FilterValue = pb.FilterValue
	st.Id = pb.Id
	st.Name = pb.Name
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExchangeFilterType string

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

// Values returns all possible values for ExchangeFilterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExchangeFilterType) Values() []ExchangeFilterType {
	return []ExchangeFilterType{
		ExchangeFilterTypeGlobalMetastoreId,
	}
}

// Type always returns ExchangeFilterType to satisfy [pflag.Value] interface
func (f *ExchangeFilterType) Type() string {
	return "ExchangeFilterType"
}

func ExchangeFilterTypeToPb(st *ExchangeFilterType) (*marketplacepb.ExchangeFilterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.ExchangeFilterTypePb(*st)
	return &pb, nil
}

func ExchangeFilterTypeFromPb(pb *marketplacepb.ExchangeFilterTypePb) (*ExchangeFilterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExchangeFilterType(*pb)
	return &st, nil
}

type ExchangeListing struct {

	// Wire name: 'created_at'
	CreatedAt int64 ``

	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'exchange_id'
	ExchangeId string ``

	// Wire name: 'exchange_name'
	ExchangeName string ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'listing_id'
	ListingId string ``

	// Wire name: 'listing_name'
	ListingName     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExchangeListing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeListing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExchangeListingToPb(st *ExchangeListing) (*marketplacepb.ExchangeListingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ExchangeListingPb{}
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

func ExchangeListingFromPb(pb *marketplacepb.ExchangeListingPb) (*ExchangeListing, error) {
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

type FileInfo struct {

	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Name displayed to users for applicable files, e.g. embedded notebooks
	// Wire name: 'display_name'
	DisplayName string ``

	// Wire name: 'download_link'
	DownloadLink string ``

	// Wire name: 'file_parent'
	FileParent *FileParent ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'marketplace_file_type'
	MarketplaceFileType MarketplaceFileType ``

	// Wire name: 'mime_type'
	MimeType string ``

	// Wire name: 'status'
	Status FileStatus ``
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	// Wire name: 'status_message'
	StatusMessage string ``

	// Wire name: 'updated_at'
	UpdatedAt       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FileInfoToPb(st *FileInfo) (*marketplacepb.FileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.FileInfoPb{}
	pb.CreatedAt = st.CreatedAt
	pb.DisplayName = st.DisplayName
	pb.DownloadLink = st.DownloadLink
	fileParentPb, err := FileParentToPb(st.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentPb != nil {
		pb.FileParent = fileParentPb
	}
	pb.Id = st.Id
	marketplaceFileTypePb, err := MarketplaceFileTypeToPb(&st.MarketplaceFileType)
	if err != nil {
		return nil, err
	}
	if marketplaceFileTypePb != nil {
		pb.MarketplaceFileType = *marketplaceFileTypePb
	}
	pb.MimeType = st.MimeType
	statusPb, err := FileStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StatusMessage = st.StatusMessage
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FileInfoFromPb(pb *marketplacepb.FileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.CreatedAt = pb.CreatedAt
	st.DisplayName = pb.DisplayName
	st.DownloadLink = pb.DownloadLink
	fileParentField, err := FileParentFromPb(pb.FileParent)
	if err != nil {
		return nil, err
	}
	if fileParentField != nil {
		st.FileParent = fileParentField
	}
	st.Id = pb.Id
	marketplaceFileTypeField, err := MarketplaceFileTypeFromPb(&pb.MarketplaceFileType)
	if err != nil {
		return nil, err
	}
	if marketplaceFileTypeField != nil {
		st.MarketplaceFileType = *marketplaceFileTypeField
	}
	st.MimeType = pb.MimeType
	statusField, err := FileStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StatusMessage = pb.StatusMessage
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type FileParent struct {

	// Wire name: 'file_parent_type'
	FileParentType FileParentType ``
	// TODO make the following fields required
	// Wire name: 'parent_id'
	ParentId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *FileParent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileParent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FileParentToPb(st *FileParent) (*marketplacepb.FileParentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.FileParentPb{}
	fileParentTypePb, err := FileParentTypeToPb(&st.FileParentType)
	if err != nil {
		return nil, err
	}
	if fileParentTypePb != nil {
		pb.FileParentType = *fileParentTypePb
	}
	pb.ParentId = st.ParentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FileParentFromPb(pb *marketplacepb.FileParentPb) (*FileParent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileParent{}
	fileParentTypeField, err := FileParentTypeFromPb(&pb.FileParentType)
	if err != nil {
		return nil, err
	}
	if fileParentTypeField != nil {
		st.FileParentType = *fileParentTypeField
	}
	st.ParentId = pb.ParentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type FileParentType string

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

// Values returns all possible values for FileParentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileParentType) Values() []FileParentType {
	return []FileParentType{
		FileParentTypeListing,
		FileParentTypeListingResource,
		FileParentTypeProvider,
	}
}

// Type always returns FileParentType to satisfy [pflag.Value] interface
func (f *FileParentType) Type() string {
	return "FileParentType"
}

func FileParentTypeToPb(st *FileParentType) (*marketplacepb.FileParentTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.FileParentTypePb(*st)
	return &pb, nil
}

func FileParentTypeFromPb(pb *marketplacepb.FileParentTypePb) (*FileParentType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FileParentType(*pb)
	return &st, nil
}

type FileStatus string

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

// Values returns all possible values for FileStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *FileStatus) Values() []FileStatus {
	return []FileStatus{
		FileStatusFileStatusPublished,
		FileStatusFileStatusSanitizationFailed,
		FileStatusFileStatusSanitizing,
		FileStatusFileStatusStaging,
	}
}

// Type always returns FileStatus to satisfy [pflag.Value] interface
func (f *FileStatus) Type() string {
	return "FileStatus"
}

func FileStatusToPb(st *FileStatus) (*marketplacepb.FileStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.FileStatusPb(*st)
	return &pb, nil
}

func FileStatusFromPb(pb *marketplacepb.FileStatusPb) (*FileStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := FileStatus(*pb)
	return &st, nil
}

type FulfillmentType string

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

// Values returns all possible values for FulfillmentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FulfillmentType) Values() []FulfillmentType {
	return []FulfillmentType{
		FulfillmentTypeInstall,
		FulfillmentTypeRequestAccess,
	}
}

// Type always returns FulfillmentType to satisfy [pflag.Value] interface
func (f *FulfillmentType) Type() string {
	return "FulfillmentType"
}

func FulfillmentTypeToPb(st *FulfillmentType) (*marketplacepb.FulfillmentTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.FulfillmentTypePb(*st)
	return &pb, nil
}

func FulfillmentTypeFromPb(pb *marketplacepb.FulfillmentTypePb) (*FulfillmentType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FulfillmentType(*pb)
	return &st, nil
}

type GetExchangeRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetExchangeRequestToPb(st *GetExchangeRequest) (*marketplacepb.GetExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetExchangeRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetExchangeRequestFromPb(pb *marketplacepb.GetExchangeRequestPb) (*GetExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExchangeRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetExchangeResponse struct {

	// Wire name: 'exchange'
	Exchange *Exchange ``
}

func GetExchangeResponseToPb(st *GetExchangeResponse) (*marketplacepb.GetExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetExchangeResponsePb{}
	exchangePb, err := ExchangeToPb(st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = exchangePb
	}

	return pb, nil
}

func GetExchangeResponseFromPb(pb *marketplacepb.GetExchangeResponsePb) (*GetExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExchangeResponse{}
	exchangeField, err := ExchangeFromPb(pb.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangeField != nil {
		st.Exchange = exchangeField
	}

	return st, nil
}

type GetFileRequest struct {

	// Wire name: 'file_id'
	FileId string `tf:"-"`
}

func GetFileRequestToPb(st *GetFileRequest) (*marketplacepb.GetFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetFileRequestPb{}
	pb.FileId = st.FileId

	return pb, nil
}

func GetFileRequestFromPb(pb *marketplacepb.GetFileRequestPb) (*GetFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFileRequest{}
	st.FileId = pb.FileId

	return st, nil
}

type GetFileResponse struct {

	// Wire name: 'file_info'
	FileInfo *FileInfo ``
}

func GetFileResponseToPb(st *GetFileResponse) (*marketplacepb.GetFileResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetFileResponsePb{}
	fileInfoPb, err := FileInfoToPb(st.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoPb != nil {
		pb.FileInfo = fileInfoPb
	}

	return pb, nil
}

func GetFileResponseFromPb(pb *marketplacepb.GetFileResponsePb) (*GetFileResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFileResponse{}
	fileInfoField, err := FileInfoFromPb(pb.FileInfo)
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
	Version         int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetLatestVersionProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLatestVersionProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetLatestVersionProviderAnalyticsDashboardResponseToPb(st *GetLatestVersionProviderAnalyticsDashboardResponse) (*marketplacepb.GetLatestVersionProviderAnalyticsDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetLatestVersionProviderAnalyticsDashboardResponsePb{}
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetLatestVersionProviderAnalyticsDashboardResponseFromPb(pb *marketplacepb.GetLatestVersionProviderAnalyticsDashboardResponsePb) (*GetLatestVersionProviderAnalyticsDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionProviderAnalyticsDashboardResponse{}
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetListingContentMetadataRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetListingContentMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetListingContentMetadataRequestToPb(st *GetListingContentMetadataRequest) (*marketplacepb.GetListingContentMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetListingContentMetadataRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetListingContentMetadataRequestFromPb(pb *marketplacepb.GetListingContentMetadataRequestPb) (*GetListingContentMetadataRequest, error) {
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

type GetListingContentMetadataResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'shared_data_objects'
	SharedDataObjects []SharedDataObject ``
	ForceSendFields   []string           `tf:"-"`
}

func (s *GetListingContentMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetListingContentMetadataResponseToPb(st *GetListingContentMetadataResponse) (*marketplacepb.GetListingContentMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetListingContentMetadataResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var sharedDataObjectsPb []marketplacepb.SharedDataObjectPb
	for _, item := range st.SharedDataObjects {
		itemPb, err := SharedDataObjectToPb(&item)
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

func GetListingContentMetadataResponseFromPb(pb *marketplacepb.GetListingContentMetadataResponsePb) (*GetListingContentMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingContentMetadataResponse{}
	st.NextPageToken = pb.NextPageToken

	var sharedDataObjectsField []SharedDataObject
	for _, itemPb := range pb.SharedDataObjects {
		item, err := SharedDataObjectFromPb(&itemPb)
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

type GetListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetListingRequestToPb(st *GetListingRequest) (*marketplacepb.GetListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetListingRequestFromPb(pb *marketplacepb.GetListingRequestPb) (*GetListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetListingResponse struct {

	// Wire name: 'listing'
	Listing *Listing ``
}

func GetListingResponseToPb(st *GetListingResponse) (*marketplacepb.GetListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetListingResponsePb{}
	listingPb, err := ListingToPb(st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = listingPb
	}

	return pb, nil
}

func GetListingResponseFromPb(pb *marketplacepb.GetListingResponsePb) (*GetListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingResponse{}
	listingField, err := ListingFromPb(pb.Listing)
	if err != nil {
		return nil, err
	}
	if listingField != nil {
		st.Listing = listingField
	}

	return st, nil
}

type GetListingsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetListingsRequestToPb(st *GetListingsRequest) (*marketplacepb.GetListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetListingsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetListingsRequestFromPb(pb *marketplacepb.GetListingsRequestPb) (*GetListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetListingsResponseToPb(st *GetListingsResponse) (*marketplacepb.GetListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetListingsResponsePb{}

	var listingsPb []marketplacepb.ListingPb
	for _, item := range st.Listings {
		itemPb, err := ListingToPb(&item)
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

func GetListingsResponseFromPb(pb *marketplacepb.GetListingsResponsePb) (*GetListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := ListingFromPb(&itemPb)
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

type GetPersonalizationRequestRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`
}

func GetPersonalizationRequestRequestToPb(st *GetPersonalizationRequestRequest) (*marketplacepb.GetPersonalizationRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetPersonalizationRequestRequestPb{}
	pb.ListingId = st.ListingId

	return pb, nil
}

func GetPersonalizationRequestRequestFromPb(pb *marketplacepb.GetPersonalizationRequestRequestPb) (*GetPersonalizationRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalizationRequestRequest{}
	st.ListingId = pb.ListingId

	return st, nil
}

type GetPersonalizationRequestResponse struct {

	// Wire name: 'personalization_requests'
	PersonalizationRequests []PersonalizationRequest ``
}

func GetPersonalizationRequestResponseToPb(st *GetPersonalizationRequestResponse) (*marketplacepb.GetPersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetPersonalizationRequestResponsePb{}

	var personalizationRequestsPb []marketplacepb.PersonalizationRequestPb
	for _, item := range st.PersonalizationRequests {
		itemPb, err := PersonalizationRequestToPb(&item)
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

func GetPersonalizationRequestResponseFromPb(pb *marketplacepb.GetPersonalizationRequestResponsePb) (*GetPersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalizationRequestResponse{}

	var personalizationRequestsField []PersonalizationRequest
	for _, itemPb := range pb.PersonalizationRequests {
		item, err := PersonalizationRequestFromPb(&itemPb)
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

type GetProviderRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetProviderRequestToPb(st *GetProviderRequest) (*marketplacepb.GetProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetProviderRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetProviderRequestFromPb(pb *marketplacepb.GetProviderRequestPb) (*GetProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetProviderResponse struct {

	// Wire name: 'provider'
	Provider *ProviderInfo ``
}

func GetProviderResponseToPb(st *GetProviderResponse) (*marketplacepb.GetProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.GetProviderResponsePb{}
	providerPb, err := ProviderInfoToPb(st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = providerPb
	}

	return pb, nil
}

func GetProviderResponseFromPb(pb *marketplacepb.GetProviderResponsePb) (*GetProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderResponse{}
	providerField, err := ProviderInfoFromPb(pb.Provider)
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
	Installation *InstallationDetail ``
}

func InstallationToPb(st *Installation) (*marketplacepb.InstallationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.InstallationPb{}
	installationPb, err := InstallationDetailToPb(st.Installation)
	if err != nil {
		return nil, err
	}
	if installationPb != nil {
		pb.Installation = installationPb
	}

	return pb, nil
}

func InstallationFromPb(pb *marketplacepb.InstallationPb) (*Installation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Installation{}
	installationField, err := InstallationDetailFromPb(pb.Installation)
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
	CatalogName string ``

	// Wire name: 'error_message'
	ErrorMessage string ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'installed_on'
	InstalledOn int64 ``

	// Wire name: 'listing_id'
	ListingId string ``

	// Wire name: 'listing_name'
	ListingName string ``

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType ``

	// Wire name: 'repo_name'
	RepoName string ``

	// Wire name: 'repo_path'
	RepoPath string ``

	// Wire name: 'share_name'
	ShareName string ``

	// Wire name: 'status'
	Status InstallationStatus ``

	// Wire name: 'token_detail'
	TokenDetail *TokenDetail ``

	// Wire name: 'tokens'
	Tokens          []TokenInfo ``
	ForceSendFields []string    `tf:"-"`
}

func (s *InstallationDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstallationDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func InstallationDetailToPb(st *InstallationDetail) (*marketplacepb.InstallationDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.InstallationDetailPb{}
	pb.CatalogName = st.CatalogName
	pb.ErrorMessage = st.ErrorMessage
	pb.Id = st.Id
	pb.InstalledOn = st.InstalledOn
	pb.ListingId = st.ListingId
	pb.ListingName = st.ListingName
	recipientTypePb, err := DeltaSharingRecipientTypeToPb(&st.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypePb != nil {
		pb.RecipientType = *recipientTypePb
	}
	pb.RepoName = st.RepoName
	pb.RepoPath = st.RepoPath
	pb.ShareName = st.ShareName
	statusPb, err := InstallationStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	tokenDetailPb, err := TokenDetailToPb(st.TokenDetail)
	if err != nil {
		return nil, err
	}
	if tokenDetailPb != nil {
		pb.TokenDetail = tokenDetailPb
	}

	var tokensPb []marketplacepb.TokenInfoPb
	for _, item := range st.Tokens {
		itemPb, err := TokenInfoToPb(&item)
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

func InstallationDetailFromPb(pb *marketplacepb.InstallationDetailPb) (*InstallationDetail, error) {
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
	recipientTypeField, err := DeltaSharingRecipientTypeFromPb(&pb.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypeField != nil {
		st.RecipientType = *recipientTypeField
	}
	st.RepoName = pb.RepoName
	st.RepoPath = pb.RepoPath
	st.ShareName = pb.ShareName
	statusField, err := InstallationStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	tokenDetailField, err := TokenDetailFromPb(pb.TokenDetail)
	if err != nil {
		return nil, err
	}
	if tokenDetailField != nil {
		st.TokenDetail = tokenDetailField
	}

	var tokensField []TokenInfo
	for _, itemPb := range pb.Tokens {
		item, err := TokenInfoFromPb(&itemPb)
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

type InstallationStatus string

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

// Values returns all possible values for InstallationStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *InstallationStatus) Values() []InstallationStatus {
	return []InstallationStatus{
		InstallationStatusFailed,
		InstallationStatusInstalled,
	}
}

// Type always returns InstallationStatus to satisfy [pflag.Value] interface
func (f *InstallationStatus) Type() string {
	return "InstallationStatus"
}

func InstallationStatusToPb(st *InstallationStatus) (*marketplacepb.InstallationStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.InstallationStatusPb(*st)
	return &pb, nil
}

func InstallationStatusFromPb(pb *marketplacepb.InstallationStatusPb) (*InstallationStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstallationStatus(*pb)
	return &st, nil
}

type ListAllInstallationsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListAllInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAllInstallationsRequestToPb(st *ListAllInstallationsRequest) (*marketplacepb.ListAllInstallationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListAllInstallationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListAllInstallationsRequestFromPb(pb *marketplacepb.ListAllInstallationsRequestPb) (*ListAllInstallationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllInstallationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAllInstallationsResponse struct {

	// Wire name: 'installations'
	Installations []InstallationDetail ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListAllInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAllInstallationsResponseToPb(st *ListAllInstallationsResponse) (*marketplacepb.ListAllInstallationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListAllInstallationsResponsePb{}

	var installationsPb []marketplacepb.InstallationDetailPb
	for _, item := range st.Installations {
		itemPb, err := InstallationDetailToPb(&item)
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

func ListAllInstallationsResponseFromPb(pb *marketplacepb.ListAllInstallationsResponsePb) (*ListAllInstallationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllInstallationsResponse{}

	var installationsField []InstallationDetail
	for _, itemPb := range pb.Installations {
		item, err := InstallationDetailFromPb(&itemPb)
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

type ListAllPersonalizationRequestsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListAllPersonalizationRequestsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAllPersonalizationRequestsRequestToPb(st *ListAllPersonalizationRequestsRequest) (*marketplacepb.ListAllPersonalizationRequestsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListAllPersonalizationRequestsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListAllPersonalizationRequestsRequestFromPb(pb *marketplacepb.ListAllPersonalizationRequestsRequestPb) (*ListAllPersonalizationRequestsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllPersonalizationRequestsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAllPersonalizationRequestsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'personalization_requests'
	PersonalizationRequests []PersonalizationRequest ``
	ForceSendFields         []string                 `tf:"-"`
}

func (s *ListAllPersonalizationRequestsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAllPersonalizationRequestsResponseToPb(st *ListAllPersonalizationRequestsResponse) (*marketplacepb.ListAllPersonalizationRequestsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListAllPersonalizationRequestsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var personalizationRequestsPb []marketplacepb.PersonalizationRequestPb
	for _, item := range st.PersonalizationRequests {
		itemPb, err := PersonalizationRequestToPb(&item)
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

func ListAllPersonalizationRequestsResponseFromPb(pb *marketplacepb.ListAllPersonalizationRequestsResponsePb) (*ListAllPersonalizationRequestsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllPersonalizationRequestsResponse{}
	st.NextPageToken = pb.NextPageToken

	var personalizationRequestsField []PersonalizationRequest
	for _, itemPb := range pb.PersonalizationRequests {
		item, err := PersonalizationRequestFromPb(&itemPb)
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

type ListExchangeFiltersRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListExchangeFiltersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExchangeFiltersRequestToPb(st *ListExchangeFiltersRequest) (*marketplacepb.ListExchangeFiltersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListExchangeFiltersRequestPb{}
	pb.ExchangeId = st.ExchangeId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExchangeFiltersRequestFromPb(pb *marketplacepb.ListExchangeFiltersRequestPb) (*ListExchangeFiltersRequest, error) {
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

type ListExchangeFiltersResponse struct {

	// Wire name: 'filters'
	Filters []ExchangeFilter ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListExchangeFiltersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExchangeFiltersResponseToPb(st *ListExchangeFiltersResponse) (*marketplacepb.ListExchangeFiltersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListExchangeFiltersResponsePb{}

	var filtersPb []marketplacepb.ExchangeFilterPb
	for _, item := range st.Filters {
		itemPb, err := ExchangeFilterToPb(&item)
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

func ListExchangeFiltersResponseFromPb(pb *marketplacepb.ListExchangeFiltersResponsePb) (*ListExchangeFiltersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangeFiltersResponse{}

	var filtersField []ExchangeFilter
	for _, itemPb := range pb.Filters {
		item, err := ExchangeFilterFromPb(&itemPb)
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

type ListExchangesForListingRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListExchangesForListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExchangesForListingRequestToPb(st *ListExchangesForListingRequest) (*marketplacepb.ListExchangesForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListExchangesForListingRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExchangesForListingRequestFromPb(pb *marketplacepb.ListExchangesForListingRequestPb) (*ListExchangesForListingRequest, error) {
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

type ListExchangesForListingResponse struct {

	// Wire name: 'exchange_listing'
	ExchangeListing []ExchangeListing ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListExchangesForListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExchangesForListingResponseToPb(st *ListExchangesForListingResponse) (*marketplacepb.ListExchangesForListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListExchangesForListingResponsePb{}

	var exchangeListingPb []marketplacepb.ExchangeListingPb
	for _, item := range st.ExchangeListing {
		itemPb, err := ExchangeListingToPb(&item)
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

func ListExchangesForListingResponseFromPb(pb *marketplacepb.ListExchangesForListingResponsePb) (*ListExchangesForListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesForListingResponse{}

	var exchangeListingField []ExchangeListing
	for _, itemPb := range pb.ExchangeListing {
		item, err := ExchangeListingFromPb(&itemPb)
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

type ListExchangesRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListExchangesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExchangesRequestToPb(st *ListExchangesRequest) (*marketplacepb.ListExchangesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListExchangesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExchangesRequestFromPb(pb *marketplacepb.ListExchangesRequestPb) (*ListExchangesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListExchangesResponse struct {

	// Wire name: 'exchanges'
	Exchanges []Exchange ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListExchangesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExchangesResponseToPb(st *ListExchangesResponse) (*marketplacepb.ListExchangesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListExchangesResponsePb{}

	var exchangesPb []marketplacepb.ExchangePb
	for _, item := range st.Exchanges {
		itemPb, err := ExchangeToPb(&item)
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

func ListExchangesResponseFromPb(pb *marketplacepb.ListExchangesResponsePb) (*ListExchangesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExchangesResponse{}

	var exchangesField []Exchange
	for _, itemPb := range pb.Exchanges {
		item, err := ExchangeFromPb(&itemPb)
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

type ListFilesRequest struct {

	// Wire name: 'file_parent'
	FileParent FileParent `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFilesRequestToPb(st *ListFilesRequest) (*marketplacepb.ListFilesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListFilesRequestPb{}
	fileParentPb, err := FileParentToPb(&st.FileParent)
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

func ListFilesRequestFromPb(pb *marketplacepb.ListFilesRequestPb) (*ListFilesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFilesRequest{}
	fileParentField, err := FileParentFromPb(&pb.FileParent)
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

type ListFilesResponse struct {

	// Wire name: 'file_infos'
	FileInfos []FileInfo ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFilesResponseToPb(st *ListFilesResponse) (*marketplacepb.ListFilesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListFilesResponsePb{}

	var fileInfosPb []marketplacepb.FileInfoPb
	for _, item := range st.FileInfos {
		itemPb, err := FileInfoToPb(&item)
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

func ListFilesResponseFromPb(pb *marketplacepb.ListFilesResponsePb) (*ListFilesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFilesResponse{}

	var fileInfosField []FileInfo
	for _, itemPb := range pb.FileInfos {
		item, err := FileInfoFromPb(&itemPb)
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

type ListFulfillmentsRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListFulfillmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFulfillmentsRequestToPb(st *ListFulfillmentsRequest) (*marketplacepb.ListFulfillmentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListFulfillmentsRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListFulfillmentsRequestFromPb(pb *marketplacepb.ListFulfillmentsRequestPb) (*ListFulfillmentsRequest, error) {
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

type ListFulfillmentsResponse struct {

	// Wire name: 'fulfillments'
	Fulfillments []ListingFulfillment ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListFulfillmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFulfillmentsResponseToPb(st *ListFulfillmentsResponse) (*marketplacepb.ListFulfillmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListFulfillmentsResponsePb{}

	var fulfillmentsPb []marketplacepb.ListingFulfillmentPb
	for _, item := range st.Fulfillments {
		itemPb, err := ListingFulfillmentToPb(&item)
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

func ListFulfillmentsResponseFromPb(pb *marketplacepb.ListFulfillmentsResponsePb) (*ListFulfillmentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFulfillmentsResponse{}

	var fulfillmentsField []ListingFulfillment
	for _, itemPb := range pb.Fulfillments {
		item, err := ListingFulfillmentFromPb(&itemPb)
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

type ListInstallationsRequest struct {

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListInstallationsRequestToPb(st *ListInstallationsRequest) (*marketplacepb.ListInstallationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListInstallationsRequestPb{}
	pb.ListingId = st.ListingId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListInstallationsRequestFromPb(pb *marketplacepb.ListInstallationsRequestPb) (*ListInstallationsRequest, error) {
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

type ListInstallationsResponse struct {

	// Wire name: 'installations'
	Installations []InstallationDetail ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListInstallationsResponseToPb(st *ListInstallationsResponse) (*marketplacepb.ListInstallationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListInstallationsResponsePb{}

	var installationsPb []marketplacepb.InstallationDetailPb
	for _, item := range st.Installations {
		itemPb, err := InstallationDetailToPb(&item)
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

func ListInstallationsResponseFromPb(pb *marketplacepb.ListInstallationsResponsePb) (*ListInstallationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstallationsResponse{}

	var installationsField []InstallationDetail
	for _, itemPb := range pb.Installations {
		item, err := InstallationDetailFromPb(&itemPb)
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

type ListListingsForExchangeRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListListingsForExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListListingsForExchangeRequestToPb(st *ListListingsForExchangeRequest) (*marketplacepb.ListListingsForExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListListingsForExchangeRequestPb{}
	pb.ExchangeId = st.ExchangeId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListListingsForExchangeRequestFromPb(pb *marketplacepb.ListListingsForExchangeRequestPb) (*ListListingsForExchangeRequest, error) {
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

type ListListingsForExchangeResponse struct {

	// Wire name: 'exchange_listings'
	ExchangeListings []ExchangeListing ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListListingsForExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListListingsForExchangeResponseToPb(st *ListListingsForExchangeResponse) (*marketplacepb.ListListingsForExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListListingsForExchangeResponsePb{}

	var exchangeListingsPb []marketplacepb.ExchangeListingPb
	for _, item := range st.ExchangeListings {
		itemPb, err := ExchangeListingToPb(&item)
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

func ListListingsForExchangeResponseFromPb(pb *marketplacepb.ListListingsForExchangeResponsePb) (*ListListingsForExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsForExchangeResponse{}

	var exchangeListingsField []ExchangeListing
	for _, itemPb := range pb.ExchangeListings {
		item, err := ExchangeListingFromPb(&itemPb)
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
	Tags            []ListingTag `tf:"-"`
	ForceSendFields []string     `tf:"-"`
}

func (s *ListListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListListingsRequestToPb(st *ListListingsRequest) (*marketplacepb.ListListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListListingsRequestPb{}

	var assetsPb []marketplacepb.AssetTypePb
	for _, item := range st.Assets {
		itemPb, err := AssetTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			assetsPb = append(assetsPb, *itemPb)
		}
	}
	pb.Assets = assetsPb

	var categoriesPb []marketplacepb.CategoryPb
	for _, item := range st.Categories {
		itemPb, err := CategoryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			categoriesPb = append(categoriesPb, *itemPb)
		}
	}
	pb.Categories = categoriesPb
	pb.IsFree = st.IsFree
	pb.IsPrivateExchange = st.IsPrivateExchange
	pb.IsStaffPick = st.IsStaffPick
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ProviderIds = st.ProviderIds

	var tagsPb []marketplacepb.ListingTagPb
	for _, item := range st.Tags {
		itemPb, err := ListingTagToPb(&item)
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

func ListListingsRequestFromPb(pb *marketplacepb.ListListingsRequestPb) (*ListListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsRequest{}

	var assetsField []AssetType
	for _, itemPb := range pb.Assets {
		item, err := AssetTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			assetsField = append(assetsField, *item)
		}
	}
	st.Assets = assetsField

	var categoriesField []Category
	for _, itemPb := range pb.Categories {
		item, err := CategoryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			categoriesField = append(categoriesField, *item)
		}
	}
	st.Categories = categoriesField
	st.IsFree = pb.IsFree
	st.IsPrivateExchange = pb.IsPrivateExchange
	st.IsStaffPick = pb.IsStaffPick
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ProviderIds = pb.ProviderIds

	var tagsField []ListingTag
	for _, itemPb := range pb.Tags {
		item, err := ListingTagFromPb(&itemPb)
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

type ListListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListListingsResponseToPb(st *ListListingsResponse) (*marketplacepb.ListListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListListingsResponsePb{}

	var listingsPb []marketplacepb.ListingPb
	for _, item := range st.Listings {
		itemPb, err := ListingToPb(&item)
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

func ListListingsResponseFromPb(pb *marketplacepb.ListListingsResponsePb) (*ListListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := ListingFromPb(&itemPb)
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

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'version'
	Version         int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProviderAnalyticsDashboardResponseToPb(st *ListProviderAnalyticsDashboardResponse) (*marketplacepb.ListProviderAnalyticsDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListProviderAnalyticsDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.Id = st.Id
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListProviderAnalyticsDashboardResponseFromPb(pb *marketplacepb.ListProviderAnalyticsDashboardResponsePb) (*ListProviderAnalyticsDashboardResponse, error) {
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

type ListProvidersRequest struct {

	// Wire name: 'is_featured'
	IsFeatured bool `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProvidersRequestToPb(st *ListProvidersRequest) (*marketplacepb.ListProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListProvidersRequestPb{}
	pb.IsFeatured = st.IsFeatured
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListProvidersRequestFromPb(pb *marketplacepb.ListProvidersRequestPb) (*ListProvidersRequest, error) {
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

type ListProvidersResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'providers'
	Providers       []ProviderInfo ``
	ForceSendFields []string       `tf:"-"`
}

func (s *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProvidersResponseToPb(st *ListProvidersResponse) (*marketplacepb.ListProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListProvidersResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var providersPb []marketplacepb.ProviderInfoPb
	for _, item := range st.Providers {
		itemPb, err := ProviderInfoToPb(&item)
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

func ListProvidersResponseFromPb(pb *marketplacepb.ListProvidersResponsePb) (*ListProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersResponse{}
	st.NextPageToken = pb.NextPageToken

	var providersField []ProviderInfo
	for _, itemPb := range pb.Providers {
		item, err := ProviderInfoFromPb(&itemPb)
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

type Listing struct {

	// Wire name: 'detail'
	Detail *ListingDetail ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'summary'
	Summary         ListingSummary ``
	ForceSendFields []string       `tf:"-"`
}

func (s *Listing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Listing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListingToPb(st *Listing) (*marketplacepb.ListingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListingPb{}
	detailPb, err := ListingDetailToPb(st.Detail)
	if err != nil {
		return nil, err
	}
	if detailPb != nil {
		pb.Detail = detailPb
	}
	pb.Id = st.Id
	summaryPb, err := ListingSummaryToPb(&st.Summary)
	if err != nil {
		return nil, err
	}
	if summaryPb != nil {
		pb.Summary = *summaryPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListingFromPb(pb *marketplacepb.ListingPb) (*Listing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Listing{}
	detailField, err := ListingDetailFromPb(pb.Detail)
	if err != nil {
		return nil, err
	}
	if detailField != nil {
		st.Detail = detailField
	}
	st.Id = pb.Id
	summaryField, err := ListingSummaryFromPb(&pb.Summary)
	if err != nil {
		return nil, err
	}
	if summaryField != nil {
		st.Summary = *summaryField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	// Wire name: 'assets'
	Assets []AssetType ``
	// The ending date timestamp for when the data spans
	// Wire name: 'collection_date_end'
	CollectionDateEnd int64 ``
	// The starting date timestamp for when the data spans
	// Wire name: 'collection_date_start'
	CollectionDateStart int64 ``
	// Smallest unit of time in the dataset
	// Wire name: 'collection_granularity'
	CollectionGranularity *DataRefreshInfo ``
	// Whether the dataset is free or paid
	// Wire name: 'cost'
	Cost Cost ``
	// Where/how the data is sourced
	// Wire name: 'data_source'
	DataSource string ``

	// Wire name: 'description'
	Description string ``

	// Wire name: 'documentation_link'
	DocumentationLink string ``

	// Wire name: 'embedded_notebook_file_infos'
	EmbeddedNotebookFileInfos []FileInfo ``

	// Wire name: 'file_ids'
	FileIds []string ``
	// Which geo region the listing data is collected from
	// Wire name: 'geographical_coverage'
	GeographicalCoverage string ``
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	// Wire name: 'license'
	License string ``
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	// Wire name: 'pricing_model'
	PricingModel string ``

	// Wire name: 'privacy_policy_link'
	PrivacyPolicyLink string ``
	// size of the dataset in GB
	// Wire name: 'size'
	Size float64 ``

	// Wire name: 'support_link'
	SupportLink string ``
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	// Wire name: 'tags'
	Tags []ListingTag ``

	// Wire name: 'terms_of_service'
	TermsOfService string ``
	// How often data is updated
	// Wire name: 'update_frequency'
	UpdateFrequency *DataRefreshInfo ``
	ForceSendFields []string         `tf:"-"`
}

func (s *ListingDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListingDetailToPb(st *ListingDetail) (*marketplacepb.ListingDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListingDetailPb{}

	var assetsPb []marketplacepb.AssetTypePb
	for _, item := range st.Assets {
		itemPb, err := AssetTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			assetsPb = append(assetsPb, *itemPb)
		}
	}
	pb.Assets = assetsPb
	pb.CollectionDateEnd = st.CollectionDateEnd
	pb.CollectionDateStart = st.CollectionDateStart
	collectionGranularityPb, err := DataRefreshInfoToPb(st.CollectionGranularity)
	if err != nil {
		return nil, err
	}
	if collectionGranularityPb != nil {
		pb.CollectionGranularity = collectionGranularityPb
	}
	costPb, err := CostToPb(&st.Cost)
	if err != nil {
		return nil, err
	}
	if costPb != nil {
		pb.Cost = *costPb
	}
	pb.DataSource = st.DataSource
	pb.Description = st.Description
	pb.DocumentationLink = st.DocumentationLink

	var embeddedNotebookFileInfosPb []marketplacepb.FileInfoPb
	for _, item := range st.EmbeddedNotebookFileInfos {
		itemPb, err := FileInfoToPb(&item)
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

	var tagsPb []marketplacepb.ListingTagPb
	for _, item := range st.Tags {
		itemPb, err := ListingTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	pb.TermsOfService = st.TermsOfService
	updateFrequencyPb, err := DataRefreshInfoToPb(st.UpdateFrequency)
	if err != nil {
		return nil, err
	}
	if updateFrequencyPb != nil {
		pb.UpdateFrequency = updateFrequencyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListingDetailFromPb(pb *marketplacepb.ListingDetailPb) (*ListingDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingDetail{}

	var assetsField []AssetType
	for _, itemPb := range pb.Assets {
		item, err := AssetTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			assetsField = append(assetsField, *item)
		}
	}
	st.Assets = assetsField
	st.CollectionDateEnd = pb.CollectionDateEnd
	st.CollectionDateStart = pb.CollectionDateStart
	collectionGranularityField, err := DataRefreshInfoFromPb(pb.CollectionGranularity)
	if err != nil {
		return nil, err
	}
	if collectionGranularityField != nil {
		st.CollectionGranularity = collectionGranularityField
	}
	costField, err := CostFromPb(&pb.Cost)
	if err != nil {
		return nil, err
	}
	if costField != nil {
		st.Cost = *costField
	}
	st.DataSource = pb.DataSource
	st.Description = pb.Description
	st.DocumentationLink = pb.DocumentationLink

	var embeddedNotebookFileInfosField []FileInfo
	for _, itemPb := range pb.EmbeddedNotebookFileInfos {
		item, err := FileInfoFromPb(&itemPb)
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
		item, err := ListingTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.TermsOfService = pb.TermsOfService
	updateFrequencyField, err := DataRefreshInfoFromPb(pb.UpdateFrequency)
	if err != nil {
		return nil, err
	}
	if updateFrequencyField != nil {
		st.UpdateFrequency = updateFrequencyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListingFulfillment struct {

	// Wire name: 'fulfillment_type'
	FulfillmentType FulfillmentType ``

	// Wire name: 'listing_id'
	ListingId string ``

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType ``

	// Wire name: 'repo_info'
	RepoInfo *RepoInfo ``

	// Wire name: 'share_info'
	ShareInfo *ShareInfo ``
}

func ListingFulfillmentToPb(st *ListingFulfillment) (*marketplacepb.ListingFulfillmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListingFulfillmentPb{}
	fulfillmentTypePb, err := FulfillmentTypeToPb(&st.FulfillmentType)
	if err != nil {
		return nil, err
	}
	if fulfillmentTypePb != nil {
		pb.FulfillmentType = *fulfillmentTypePb
	}
	pb.ListingId = st.ListingId
	recipientTypePb, err := DeltaSharingRecipientTypeToPb(&st.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypePb != nil {
		pb.RecipientType = *recipientTypePb
	}
	repoInfoPb, err := RepoInfoToPb(st.RepoInfo)
	if err != nil {
		return nil, err
	}
	if repoInfoPb != nil {
		pb.RepoInfo = repoInfoPb
	}
	shareInfoPb, err := ShareInfoToPb(st.ShareInfo)
	if err != nil {
		return nil, err
	}
	if shareInfoPb != nil {
		pb.ShareInfo = shareInfoPb
	}

	return pb, nil
}

func ListingFulfillmentFromPb(pb *marketplacepb.ListingFulfillmentPb) (*ListingFulfillment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingFulfillment{}
	fulfillmentTypeField, err := FulfillmentTypeFromPb(&pb.FulfillmentType)
	if err != nil {
		return nil, err
	}
	if fulfillmentTypeField != nil {
		st.FulfillmentType = *fulfillmentTypeField
	}
	st.ListingId = pb.ListingId
	recipientTypeField, err := DeltaSharingRecipientTypeFromPb(&pb.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypeField != nil {
		st.RecipientType = *recipientTypeField
	}
	repoInfoField, err := RepoInfoFromPb(pb.RepoInfo)
	if err != nil {
		return nil, err
	}
	if repoInfoField != nil {
		st.RepoInfo = repoInfoField
	}
	shareInfoField, err := ShareInfoFromPb(pb.ShareInfo)
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
	Visibility Visibility ``
}

func ListingSettingToPb(st *ListingSetting) (*marketplacepb.ListingSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListingSettingPb{}
	visibilityPb, err := VisibilityToPb(&st.Visibility)
	if err != nil {
		return nil, err
	}
	if visibilityPb != nil {
		pb.Visibility = *visibilityPb
	}

	return pb, nil
}

func ListingSettingFromPb(pb *marketplacepb.ListingSettingPb) (*ListingSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingSetting{}
	visibilityField, err := VisibilityFromPb(&pb.Visibility)
	if err != nil {
		return nil, err
	}
	if visibilityField != nil {
		st.Visibility = *visibilityField
	}

	return st, nil
}

type ListingShareType string

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

// Values returns all possible values for ListingShareType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingShareType) Values() []ListingShareType {
	return []ListingShareType{
		ListingShareTypeFull,
		ListingShareTypeSample,
	}
}

// Type always returns ListingShareType to satisfy [pflag.Value] interface
func (f *ListingShareType) Type() string {
	return "ListingShareType"
}

func ListingShareTypeToPb(st *ListingShareType) (*marketplacepb.ListingShareTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.ListingShareTypePb(*st)
	return &pb, nil
}

func ListingShareTypeFromPb(pb *marketplacepb.ListingShareTypePb) (*ListingShareType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingShareType(*pb)
	return &st, nil
}

// Enums
type ListingStatus string

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

// Values returns all possible values for ListingStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingStatus) Values() []ListingStatus {
	return []ListingStatus{
		ListingStatusDraft,
		ListingStatusPending,
		ListingStatusPublished,
		ListingStatusSuspended,
	}
}

// Type always returns ListingStatus to satisfy [pflag.Value] interface
func (f *ListingStatus) Type() string {
	return "ListingStatus"
}

func ListingStatusToPb(st *ListingStatus) (*marketplacepb.ListingStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.ListingStatusPb(*st)
	return &pb, nil
}

func ListingStatusFromPb(pb *marketplacepb.ListingStatusPb) (*ListingStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingStatus(*pb)
	return &st, nil
}

type ListingSummary struct {

	// Wire name: 'categories'
	Categories []Category ``

	// Wire name: 'created_at'
	CreatedAt int64 ``

	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'created_by_id'
	CreatedById int64 ``

	// Wire name: 'exchange_ids'
	ExchangeIds []string ``
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	// Wire name: 'git_repo'
	GitRepo *RepoInfo ``

	// Wire name: 'listingType'
	ListingType ListingType ``

	// Wire name: 'name'
	Name string ``

	// Wire name: 'provider_id'
	ProviderId string ``

	// Wire name: 'provider_region'
	ProviderRegion *RegionInfo ``

	// Wire name: 'published_at'
	PublishedAt int64 ``

	// Wire name: 'published_by'
	PublishedBy string ``

	// Wire name: 'setting'
	Setting *ListingSetting ``

	// Wire name: 'share'
	Share *ShareInfo ``

	// Wire name: 'status'
	Status ListingStatus ``

	// Wire name: 'subtitle'
	Subtitle string ``

	// Wire name: 'updated_at'
	UpdatedAt int64 ``

	// Wire name: 'updated_by'
	UpdatedBy string ``

	// Wire name: 'updated_by_id'
	UpdatedById     int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListingSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListingSummaryToPb(st *ListingSummary) (*marketplacepb.ListingSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListingSummaryPb{}

	var categoriesPb []marketplacepb.CategoryPb
	for _, item := range st.Categories {
		itemPb, err := CategoryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			categoriesPb = append(categoriesPb, *itemPb)
		}
	}
	pb.Categories = categoriesPb
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.CreatedById = st.CreatedById
	pb.ExchangeIds = st.ExchangeIds
	gitRepoPb, err := RepoInfoToPb(st.GitRepo)
	if err != nil {
		return nil, err
	}
	if gitRepoPb != nil {
		pb.GitRepo = gitRepoPb
	}
	listingTypePb, err := ListingTypeToPb(&st.ListingType)
	if err != nil {
		return nil, err
	}
	if listingTypePb != nil {
		pb.ListingType = *listingTypePb
	}
	pb.Name = st.Name
	pb.ProviderId = st.ProviderId
	providerRegionPb, err := RegionInfoToPb(st.ProviderRegion)
	if err != nil {
		return nil, err
	}
	if providerRegionPb != nil {
		pb.ProviderRegion = providerRegionPb
	}
	pb.PublishedAt = st.PublishedAt
	pb.PublishedBy = st.PublishedBy
	settingPb, err := ListingSettingToPb(st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = settingPb
	}
	sharePb, err := ShareInfoToPb(st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = sharePb
	}
	statusPb, err := ListingStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.Subtitle = st.Subtitle
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.UpdatedById = st.UpdatedById

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListingSummaryFromPb(pb *marketplacepb.ListingSummaryPb) (*ListingSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingSummary{}

	var categoriesField []Category
	for _, itemPb := range pb.Categories {
		item, err := CategoryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			categoriesField = append(categoriesField, *item)
		}
	}
	st.Categories = categoriesField
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CreatedById = pb.CreatedById
	st.ExchangeIds = pb.ExchangeIds
	gitRepoField, err := RepoInfoFromPb(pb.GitRepo)
	if err != nil {
		return nil, err
	}
	if gitRepoField != nil {
		st.GitRepo = gitRepoField
	}
	listingTypeField, err := ListingTypeFromPb(&pb.ListingType)
	if err != nil {
		return nil, err
	}
	if listingTypeField != nil {
		st.ListingType = *listingTypeField
	}
	st.Name = pb.Name
	st.ProviderId = pb.ProviderId
	providerRegionField, err := RegionInfoFromPb(pb.ProviderRegion)
	if err != nil {
		return nil, err
	}
	if providerRegionField != nil {
		st.ProviderRegion = providerRegionField
	}
	st.PublishedAt = pb.PublishedAt
	st.PublishedBy = pb.PublishedBy
	settingField, err := ListingSettingFromPb(pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = settingField
	}
	shareField, err := ShareInfoFromPb(pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = shareField
	}
	statusField, err := ListingStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.Subtitle = pb.Subtitle
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UpdatedById = pb.UpdatedById

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListingTag struct {
	// Tag name (enum)
	// Wire name: 'tag_name'
	TagName ListingTagType ``
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	// Wire name: 'tag_values'
	TagValues []string ``
}

func ListingTagToPb(st *ListingTag) (*marketplacepb.ListingTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ListingTagPb{}
	tagNamePb, err := ListingTagTypeToPb(&st.TagName)
	if err != nil {
		return nil, err
	}
	if tagNamePb != nil {
		pb.TagName = *tagNamePb
	}
	pb.TagValues = st.TagValues

	return pb, nil
}

func ListingTagFromPb(pb *marketplacepb.ListingTagPb) (*ListingTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListingTag{}
	tagNameField, err := ListingTagTypeFromPb(&pb.TagName)
	if err != nil {
		return nil, err
	}
	if tagNameField != nil {
		st.TagName = *tagNameField
	}
	st.TagValues = pb.TagValues

	return st, nil
}

type ListingTagType string

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

// Values returns all possible values for ListingTagType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingTagType) Values() []ListingTagType {
	return []ListingTagType{
		ListingTagTypeListingTagTypeLanguage,
		ListingTagTypeListingTagTypeTask,
	}
}

// Type always returns ListingTagType to satisfy [pflag.Value] interface
func (f *ListingTagType) Type() string {
	return "ListingTagType"
}

func ListingTagTypeToPb(st *ListingTagType) (*marketplacepb.ListingTagTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.ListingTagTypePb(*st)
	return &pb, nil
}

func ListingTagTypeFromPb(pb *marketplacepb.ListingTagTypePb) (*ListingTagType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingTagType(*pb)
	return &st, nil
}

type ListingType string

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

// Values returns all possible values for ListingType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListingType) Values() []ListingType {
	return []ListingType{
		ListingTypePersonalized,
		ListingTypeStandard,
	}
}

// Type always returns ListingType to satisfy [pflag.Value] interface
func (f *ListingType) Type() string {
	return "ListingType"
}

func ListingTypeToPb(st *ListingType) (*marketplacepb.ListingTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.ListingTypePb(*st)
	return &pb, nil
}

func ListingTypeFromPb(pb *marketplacepb.ListingTypePb) (*ListingType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListingType(*pb)
	return &st, nil
}

type MarketplaceFileType string

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

// Values returns all possible values for MarketplaceFileType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MarketplaceFileType) Values() []MarketplaceFileType {
	return []MarketplaceFileType{
		MarketplaceFileTypeApp,
		MarketplaceFileTypeEmbeddedNotebook,
		MarketplaceFileTypeProviderIcon,
	}
}

// Type always returns MarketplaceFileType to satisfy [pflag.Value] interface
func (f *MarketplaceFileType) Type() string {
	return "MarketplaceFileType"
}

func MarketplaceFileTypeToPb(st *MarketplaceFileType) (*marketplacepb.MarketplaceFileTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.MarketplaceFileTypePb(*st)
	return &pb, nil
}

func MarketplaceFileTypeFromPb(pb *marketplacepb.MarketplaceFileTypePb) (*MarketplaceFileType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MarketplaceFileType(*pb)
	return &st, nil
}

type PersonalizationRequest struct {

	// Wire name: 'comment'
	Comment string ``

	// Wire name: 'consumer_region'
	ConsumerRegion RegionInfo ``

	// Wire name: 'contact_info'
	ContactInfo *ContactInfo ``

	// Wire name: 'created_at'
	CreatedAt int64 ``

	// Wire name: 'id'
	Id string ``

	// Wire name: 'intended_use'
	IntendedUse string ``

	// Wire name: 'is_from_lighthouse'
	IsFromLighthouse bool ``

	// Wire name: 'listing_id'
	ListingId string ``

	// Wire name: 'listing_name'
	ListingName string ``

	// Wire name: 'metastore_id'
	MetastoreId string ``

	// Wire name: 'provider_id'
	ProviderId string ``

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType ``

	// Wire name: 'share'
	Share *ShareInfo ``

	// Wire name: 'status'
	Status PersonalizationRequestStatus ``

	// Wire name: 'status_message'
	StatusMessage string ``

	// Wire name: 'updated_at'
	UpdatedAt       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *PersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PersonalizationRequestToPb(st *PersonalizationRequest) (*marketplacepb.PersonalizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.PersonalizationRequestPb{}
	pb.Comment = st.Comment
	consumerRegionPb, err := RegionInfoToPb(&st.ConsumerRegion)
	if err != nil {
		return nil, err
	}
	if consumerRegionPb != nil {
		pb.ConsumerRegion = *consumerRegionPb
	}
	contactInfoPb, err := ContactInfoToPb(st.ContactInfo)
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
	recipientTypePb, err := DeltaSharingRecipientTypeToPb(&st.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypePb != nil {
		pb.RecipientType = *recipientTypePb
	}
	sharePb, err := ShareInfoToPb(st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = sharePb
	}
	statusPb, err := PersonalizationRequestStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StatusMessage = st.StatusMessage
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PersonalizationRequestFromPb(pb *marketplacepb.PersonalizationRequestPb) (*PersonalizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalizationRequest{}
	st.Comment = pb.Comment
	consumerRegionField, err := RegionInfoFromPb(&pb.ConsumerRegion)
	if err != nil {
		return nil, err
	}
	if consumerRegionField != nil {
		st.ConsumerRegion = *consumerRegionField
	}
	contactInfoField, err := ContactInfoFromPb(pb.ContactInfo)
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
	recipientTypeField, err := DeltaSharingRecipientTypeFromPb(&pb.RecipientType)
	if err != nil {
		return nil, err
	}
	if recipientTypeField != nil {
		st.RecipientType = *recipientTypeField
	}
	shareField, err := ShareInfoFromPb(pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = shareField
	}
	statusField, err := PersonalizationRequestStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StatusMessage = pb.StatusMessage
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PersonalizationRequestStatus string

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

// Values returns all possible values for PersonalizationRequestStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *PersonalizationRequestStatus) Values() []PersonalizationRequestStatus {
	return []PersonalizationRequestStatus{
		PersonalizationRequestStatusDenied,
		PersonalizationRequestStatusFulfilled,
		PersonalizationRequestStatusNew,
		PersonalizationRequestStatusRequestPending,
	}
}

// Type always returns PersonalizationRequestStatus to satisfy [pflag.Value] interface
func (f *PersonalizationRequestStatus) Type() string {
	return "PersonalizationRequestStatus"
}

func PersonalizationRequestStatusToPb(st *PersonalizationRequestStatus) (*marketplacepb.PersonalizationRequestStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.PersonalizationRequestStatusPb(*st)
	return &pb, nil
}

func PersonalizationRequestStatusFromPb(pb *marketplacepb.PersonalizationRequestStatusPb) (*PersonalizationRequestStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := PersonalizationRequestStatus(*pb)
	return &st, nil
}

type ProviderAnalyticsDashboard struct {

	// Wire name: 'id'
	Id string ``
}

func ProviderAnalyticsDashboardToPb(st *ProviderAnalyticsDashboard) (*marketplacepb.ProviderAnalyticsDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ProviderAnalyticsDashboardPb{}
	pb.Id = st.Id

	return pb, nil
}

func ProviderAnalyticsDashboardFromPb(pb *marketplacepb.ProviderAnalyticsDashboardPb) (*ProviderAnalyticsDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderAnalyticsDashboard{}
	st.Id = pb.Id

	return st, nil
}

type ProviderInfo struct {

	// Wire name: 'business_contact_email'
	BusinessContactEmail string ``

	// Wire name: 'company_website_link'
	CompanyWebsiteLink string ``

	// Wire name: 'dark_mode_icon_file_id'
	DarkModeIconFileId string ``

	// Wire name: 'dark_mode_icon_file_path'
	DarkModeIconFilePath string ``

	// Wire name: 'description'
	Description string ``

	// Wire name: 'icon_file_id'
	IconFileId string ``

	// Wire name: 'icon_file_path'
	IconFilePath string ``

	// Wire name: 'id'
	Id string ``
	// is_featured is accessible by consumers only
	// Wire name: 'is_featured'
	IsFeatured bool ``

	// Wire name: 'name'
	Name string ``

	// Wire name: 'privacy_policy_link'
	PrivacyPolicyLink string ``
	// published_by is only applicable to data aggregators (e.g. Crux)
	// Wire name: 'published_by'
	PublishedBy string ``

	// Wire name: 'support_contact_email'
	SupportContactEmail string ``

	// Wire name: 'term_of_service_link'
	TermOfServiceLink string   ``
	ForceSendFields   []string `tf:"-"`
}

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ProviderInfoToPb(st *ProviderInfo) (*marketplacepb.ProviderInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ProviderInfoPb{}
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

func ProviderInfoFromPb(pb *marketplacepb.ProviderInfoPb) (*ProviderInfo, error) {
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

type RegionInfo struct {

	// Wire name: 'cloud'
	Cloud string ``

	// Wire name: 'region'
	Region          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RegionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RegionInfoToPb(st *RegionInfo) (*marketplacepb.RegionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.RegionInfoPb{}
	pb.Cloud = st.Cloud
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RegionInfoFromPb(pb *marketplacepb.RegionInfoPb) (*RegionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegionInfo{}
	st.Cloud = pb.Cloud
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RemoveExchangeForListingRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func RemoveExchangeForListingRequestToPb(st *RemoveExchangeForListingRequest) (*marketplacepb.RemoveExchangeForListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.RemoveExchangeForListingRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func RemoveExchangeForListingRequestFromPb(pb *marketplacepb.RemoveExchangeForListingRequestPb) (*RemoveExchangeForListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveExchangeForListingRequest{}
	st.Id = pb.Id

	return st, nil
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	// Wire name: 'git_repo_url'
	GitRepoUrl string ``
}

func RepoInfoToPb(st *RepoInfo) (*marketplacepb.RepoInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.RepoInfoPb{}
	pb.GitRepoUrl = st.GitRepoUrl

	return pb, nil
}

func RepoInfoFromPb(pb *marketplacepb.RepoInfoPb) (*RepoInfo, error) {
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
	RepoName string ``
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	// Wire name: 'repo_path'
	RepoPath string ``
}

func RepoInstallationToPb(st *RepoInstallation) (*marketplacepb.RepoInstallationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.RepoInstallationPb{}
	pb.RepoName = st.RepoName
	pb.RepoPath = st.RepoPath

	return pb, nil
}

func RepoInstallationFromPb(pb *marketplacepb.RepoInstallationPb) (*RepoInstallation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoInstallation{}
	st.RepoName = pb.RepoName
	st.RepoPath = pb.RepoPath

	return st, nil
}

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
	Query           string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *SearchListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SearchListingsRequestToPb(st *SearchListingsRequest) (*marketplacepb.SearchListingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.SearchListingsRequestPb{}

	var assetsPb []marketplacepb.AssetTypePb
	for _, item := range st.Assets {
		itemPb, err := AssetTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			assetsPb = append(assetsPb, *itemPb)
		}
	}
	pb.Assets = assetsPb

	var categoriesPb []marketplacepb.CategoryPb
	for _, item := range st.Categories {
		itemPb, err := CategoryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			categoriesPb = append(categoriesPb, *itemPb)
		}
	}
	pb.Categories = categoriesPb
	pb.IsFree = st.IsFree
	pb.IsPrivateExchange = st.IsPrivateExchange
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ProviderIds = st.ProviderIds
	pb.Query = st.Query

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SearchListingsRequestFromPb(pb *marketplacepb.SearchListingsRequestPb) (*SearchListingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchListingsRequest{}

	var assetsField []AssetType
	for _, itemPb := range pb.Assets {
		item, err := AssetTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			assetsField = append(assetsField, *item)
		}
	}
	st.Assets = assetsField

	var categoriesField []Category
	for _, itemPb := range pb.Categories {
		item, err := CategoryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			categoriesField = append(categoriesField, *item)
		}
	}
	st.Categories = categoriesField
	st.IsFree = pb.IsFree
	st.IsPrivateExchange = pb.IsPrivateExchange
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ProviderIds = pb.ProviderIds
	st.Query = pb.Query

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SearchListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SearchListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SearchListingsResponseToPb(st *SearchListingsResponse) (*marketplacepb.SearchListingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.SearchListingsResponsePb{}

	var listingsPb []marketplacepb.ListingPb
	for _, item := range st.Listings {
		itemPb, err := ListingToPb(&item)
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

func SearchListingsResponseFromPb(pb *marketplacepb.SearchListingsResponsePb) (*SearchListingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchListingsResponse{}

	var listingsField []Listing
	for _, itemPb := range pb.Listings {
		item, err := ListingFromPb(&itemPb)
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

type ShareInfo struct {

	// Wire name: 'name'
	Name string ``

	// Wire name: 'type'
	Type ListingShareType ``
}

func ShareInfoToPb(st *ShareInfo) (*marketplacepb.ShareInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.ShareInfoPb{}
	pb.Name = st.Name
	typePb, err := ListingShareTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	return pb, nil
}

func ShareInfoFromPb(pb *marketplacepb.ShareInfoPb) (*ShareInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareInfo{}
	st.Name = pb.Name
	typeField, err := ListingShareTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	return st, nil
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	// Wire name: 'data_object_type'
	DataObjectType string ``
	// Name of the shared object
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SharedDataObjectToPb(st *SharedDataObject) (*marketplacepb.SharedDataObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.SharedDataObjectPb{}
	pb.DataObjectType = st.DataObjectType
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SharedDataObjectFromPb(pb *marketplacepb.SharedDataObjectPb) (*SharedDataObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObject{}
	st.DataObjectType = pb.DataObjectType
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TokenDetail struct {

	// Wire name: 'bearerToken'
	BearerToken string ``

	// Wire name: 'endpoint'
	Endpoint string ``

	// Wire name: 'expirationTime'
	ExpirationTime string ``
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	// Wire name: 'shareCredentialsVersion'
	ShareCredentialsVersion int      ``
	ForceSendFields         []string `tf:"-"`
}

func (s *TokenDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TokenDetailToPb(st *TokenDetail) (*marketplacepb.TokenDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.TokenDetailPb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ExpirationTime = st.ExpirationTime
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TokenDetailFromPb(pb *marketplacepb.TokenDetailPb) (*TokenDetail, error) {
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

type TokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string ``
	// Time at which this Recipient Token was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of Recipient Token creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``
	// Unique id of the Recipient Token.
	// Wire name: 'id'
	Id string ``
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of Recipient Token updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TokenInfoToPb(st *TokenInfo) (*marketplacepb.TokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.TokenInfoPb{}
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

func TokenInfoFromPb(pb *marketplacepb.TokenInfoPb) (*TokenInfo, error) {
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

type UpdateExchangeFilterRequest struct {

	// Wire name: 'filter'
	Filter ExchangeFilter ``

	// Wire name: 'id'
	Id string `tf:"-"`
}

func UpdateExchangeFilterRequestToPb(st *UpdateExchangeFilterRequest) (*marketplacepb.UpdateExchangeFilterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateExchangeFilterRequestPb{}
	filterPb, err := ExchangeFilterToPb(&st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = *filterPb
	}
	pb.Id = st.Id

	return pb, nil
}

func UpdateExchangeFilterRequestFromPb(pb *marketplacepb.UpdateExchangeFilterRequestPb) (*UpdateExchangeFilterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeFilterRequest{}
	filterField, err := ExchangeFilterFromPb(&pb.Filter)
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
	Filter *ExchangeFilter ``
}

func UpdateExchangeFilterResponseToPb(st *UpdateExchangeFilterResponse) (*marketplacepb.UpdateExchangeFilterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateExchangeFilterResponsePb{}
	filterPb, err := ExchangeFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	return pb, nil
}

func UpdateExchangeFilterResponseFromPb(pb *marketplacepb.UpdateExchangeFilterResponsePb) (*UpdateExchangeFilterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeFilterResponse{}
	filterField, err := ExchangeFilterFromPb(pb.Filter)
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
	Exchange Exchange ``

	// Wire name: 'id'
	Id string `tf:"-"`
}

func UpdateExchangeRequestToPb(st *UpdateExchangeRequest) (*marketplacepb.UpdateExchangeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateExchangeRequestPb{}
	exchangePb, err := ExchangeToPb(&st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = *exchangePb
	}
	pb.Id = st.Id

	return pb, nil
}

func UpdateExchangeRequestFromPb(pb *marketplacepb.UpdateExchangeRequestPb) (*UpdateExchangeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeRequest{}
	exchangeField, err := ExchangeFromPb(&pb.Exchange)
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
	Exchange *Exchange ``
}

func UpdateExchangeResponseToPb(st *UpdateExchangeResponse) (*marketplacepb.UpdateExchangeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateExchangeResponsePb{}
	exchangePb, err := ExchangeToPb(st.Exchange)
	if err != nil {
		return nil, err
	}
	if exchangePb != nil {
		pb.Exchange = exchangePb
	}

	return pb, nil
}

func UpdateExchangeResponseFromPb(pb *marketplacepb.UpdateExchangeResponsePb) (*UpdateExchangeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExchangeResponse{}
	exchangeField, err := ExchangeFromPb(pb.Exchange)
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
	Installation InstallationDetail ``

	// Wire name: 'installation_id'
	InstallationId string `tf:"-"`

	// Wire name: 'listing_id'
	ListingId string `tf:"-"`

	// Wire name: 'rotate_token'
	RotateToken     bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateInstallationRequestToPb(st *UpdateInstallationRequest) (*marketplacepb.UpdateInstallationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateInstallationRequestPb{}
	installationPb, err := InstallationDetailToPb(&st.Installation)
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

func UpdateInstallationRequestFromPb(pb *marketplacepb.UpdateInstallationRequestPb) (*UpdateInstallationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInstallationRequest{}
	installationField, err := InstallationDetailFromPb(&pb.Installation)
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

type UpdateInstallationResponse struct {

	// Wire name: 'installation'
	Installation *InstallationDetail ``
}

func UpdateInstallationResponseToPb(st *UpdateInstallationResponse) (*marketplacepb.UpdateInstallationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateInstallationResponsePb{}
	installationPb, err := InstallationDetailToPb(st.Installation)
	if err != nil {
		return nil, err
	}
	if installationPb != nil {
		pb.Installation = installationPb
	}

	return pb, nil
}

func UpdateInstallationResponseFromPb(pb *marketplacepb.UpdateInstallationResponsePb) (*UpdateInstallationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInstallationResponse{}
	installationField, err := InstallationDetailFromPb(pb.Installation)
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
	Listing Listing ``
}

func UpdateListingRequestToPb(st *UpdateListingRequest) (*marketplacepb.UpdateListingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateListingRequestPb{}
	pb.Id = st.Id
	listingPb, err := ListingToPb(&st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = *listingPb
	}

	return pb, nil
}

func UpdateListingRequestFromPb(pb *marketplacepb.UpdateListingRequestPb) (*UpdateListingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateListingRequest{}
	st.Id = pb.Id
	listingField, err := ListingFromPb(&pb.Listing)
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
	Listing *Listing ``
}

func UpdateListingResponseToPb(st *UpdateListingResponse) (*marketplacepb.UpdateListingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateListingResponsePb{}
	listingPb, err := ListingToPb(st.Listing)
	if err != nil {
		return nil, err
	}
	if listingPb != nil {
		pb.Listing = listingPb
	}

	return pb, nil
}

func UpdateListingResponseFromPb(pb *marketplacepb.UpdateListingResponsePb) (*UpdateListingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateListingResponse{}
	listingField, err := ListingFromPb(pb.Listing)
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
	Reason string ``

	// Wire name: 'request_id'
	RequestId string `tf:"-"`

	// Wire name: 'share'
	Share *ShareInfo ``

	// Wire name: 'status'
	Status          PersonalizationRequestStatus ``
	ForceSendFields []string                     `tf:"-"`
}

func (s *UpdatePersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePersonalizationRequestRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdatePersonalizationRequestRequestToPb(st *UpdatePersonalizationRequestRequest) (*marketplacepb.UpdatePersonalizationRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdatePersonalizationRequestRequestPb{}
	pb.ListingId = st.ListingId
	pb.Reason = st.Reason
	pb.RequestId = st.RequestId
	sharePb, err := ShareInfoToPb(st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = sharePb
	}
	statusPb, err := PersonalizationRequestStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdatePersonalizationRequestRequestFromPb(pb *marketplacepb.UpdatePersonalizationRequestRequestPb) (*UpdatePersonalizationRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalizationRequestRequest{}
	st.ListingId = pb.ListingId
	st.Reason = pb.Reason
	st.RequestId = pb.RequestId
	shareField, err := ShareInfoFromPb(pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = shareField
	}
	statusField, err := PersonalizationRequestStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdatePersonalizationRequestResponse struct {

	// Wire name: 'request'
	Request *PersonalizationRequest ``
}

func UpdatePersonalizationRequestResponseToPb(st *UpdatePersonalizationRequestResponse) (*marketplacepb.UpdatePersonalizationRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdatePersonalizationRequestResponsePb{}
	requestPb, err := PersonalizationRequestToPb(st.Request)
	if err != nil {
		return nil, err
	}
	if requestPb != nil {
		pb.Request = requestPb
	}

	return pb, nil
}

func UpdatePersonalizationRequestResponseFromPb(pb *marketplacepb.UpdatePersonalizationRequestResponsePb) (*UpdatePersonalizationRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalizationRequestResponse{}
	requestField, err := PersonalizationRequestFromPb(pb.Request)
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
	Version         int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateProviderAnalyticsDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateProviderAnalyticsDashboardRequestToPb(st *UpdateProviderAnalyticsDashboardRequest) (*marketplacepb.UpdateProviderAnalyticsDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateProviderAnalyticsDashboardRequestPb{}
	pb.Id = st.Id
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateProviderAnalyticsDashboardRequestFromPb(pb *marketplacepb.UpdateProviderAnalyticsDashboardRequestPb) (*UpdateProviderAnalyticsDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderAnalyticsDashboardRequest{}
	st.Id = pb.Id
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// id & version should be the same as the request
	// Wire name: 'id'
	Id string ``

	// Wire name: 'version'
	Version         int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateProviderAnalyticsDashboardResponseToPb(st *UpdateProviderAnalyticsDashboardResponse) (*marketplacepb.UpdateProviderAnalyticsDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateProviderAnalyticsDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.Id = st.Id
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateProviderAnalyticsDashboardResponseFromPb(pb *marketplacepb.UpdateProviderAnalyticsDashboardResponsePb) (*UpdateProviderAnalyticsDashboardResponse, error) {
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

type UpdateProviderRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'provider'
	Provider ProviderInfo ``
}

func UpdateProviderRequestToPb(st *UpdateProviderRequest) (*marketplacepb.UpdateProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateProviderRequestPb{}
	pb.Id = st.Id
	providerPb, err := ProviderInfoToPb(&st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = *providerPb
	}

	return pb, nil
}

func UpdateProviderRequestFromPb(pb *marketplacepb.UpdateProviderRequestPb) (*UpdateProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderRequest{}
	st.Id = pb.Id
	providerField, err := ProviderInfoFromPb(&pb.Provider)
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
	Provider *ProviderInfo ``
}

func UpdateProviderResponseToPb(st *UpdateProviderResponse) (*marketplacepb.UpdateProviderResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &marketplacepb.UpdateProviderResponsePb{}
	providerPb, err := ProviderInfoToPb(st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = providerPb
	}

	return pb, nil
}

func UpdateProviderResponseFromPb(pb *marketplacepb.UpdateProviderResponsePb) (*UpdateProviderResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProviderResponse{}
	providerField, err := ProviderInfoFromPb(pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = providerField
	}

	return st, nil
}

type Visibility string

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

// Values returns all possible values for Visibility.
//
// There is no guarantee on the order of the values in the slice.
func (f *Visibility) Values() []Visibility {
	return []Visibility{
		VisibilityPrivate,
		VisibilityPublic,
	}
}

// Type always returns Visibility to satisfy [pflag.Value] interface
func (f *Visibility) Type() string {
	return "Visibility"
}

func VisibilityToPb(st *Visibility) (*marketplacepb.VisibilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := marketplacepb.VisibilityPb(*st)
	return &pb, nil
}

func VisibilityFromPb(pb *marketplacepb.VisibilityPb) (*Visibility, error) {
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
	s := fmt.Sprintf("%.9fs", d.Seconds())
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
