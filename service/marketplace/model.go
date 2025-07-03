// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AddExchangeForListingRequest struct {

	// Wire name: 'exchange_id'
	ExchangeId string `json:"exchange_id"`

	// Wire name: 'listing_id'
	ListingId string `json:"listing_id"`
}

func (st *AddExchangeForListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := addExchangeForListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AddExchangeForListingResponse struct {

	// Wire name: 'exchange_for_listing'
	ExchangeForListing *ExchangeListing `json:"exchange_for_listing,omitempty"`
}

func (st *AddExchangeForListingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := addExchangeForListingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type BatchGetListingsRequest struct {
	Ids []string `json:"-" tf:"-"`
}

func (st *BatchGetListingsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := batchGetListingsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type BatchGetListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing `json:"listings,omitempty"`
}

func (st *BatchGetListingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := batchGetListingsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type BatchGetProvidersRequest struct {
	Ids []string `json:"-" tf:"-"`
}

func (st *BatchGetProvidersRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := batchGetProvidersRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type BatchGetProvidersResponse struct {

	// Wire name: 'providers'
	Providers []ProviderInfo `json:"providers,omitempty"`
}

func (st *BatchGetProvidersResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := batchGetProvidersResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ConsumerTerms struct {

	// Wire name: 'version'
	Version string `json:"version"`
}

func (st *ConsumerTerms) EncodeValues(key string, v *url.Values) error {
	pb, err := consumerTermsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {

	// Wire name: 'company'
	Company string `json:"company,omitempty"`

	// Wire name: 'email'
	Email string `json:"email,omitempty"`

	// Wire name: 'first_name'
	FirstName string `json:"first_name,omitempty"`

	// Wire name: 'last_name'
	LastName string `json:"last_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ContactInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := contactInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateExchangeFilterRequest struct {

	// Wire name: 'filter'
	Filter ExchangeFilter `json:"filter"`
}

func (st *CreateExchangeFilterRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createExchangeFilterRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateExchangeFilterResponse struct {

	// Wire name: 'filter_id'
	FilterId string `json:"filter_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateExchangeFilterResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createExchangeFilterResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateExchangeRequest struct {

	// Wire name: 'exchange'
	Exchange Exchange `json:"exchange"`
}

func (st *CreateExchangeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createExchangeRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateExchangeResponse struct {

	// Wire name: 'exchange_id'
	ExchangeId string `json:"exchange_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateExchangeResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createExchangeResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateFileRequest struct {

	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`

	// Wire name: 'file_parent'
	FileParent FileParent `json:"file_parent"`

	// Wire name: 'marketplace_file_type'
	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type"`

	// Wire name: 'mime_type'
	MimeType string `json:"mime_type"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateFileRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createFileRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateFileResponse struct {

	// Wire name: 'file_info'
	FileInfo *FileInfo `json:"file_info,omitempty"`
	// Pre-signed POST URL to blob storage
	// Wire name: 'upload_url'
	UploadUrl string `json:"upload_url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateFileResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createFileResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateInstallationRequest struct {

	// Wire name: 'accepted_consumer_terms'
	AcceptedConsumerTerms *ConsumerTerms `json:"accepted_consumer_terms,omitempty"`

	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`

	ListingId string `json:"-" tf:"-"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	// for git repo installations
	// Wire name: 'repo_detail'
	RepoDetail *RepoInstallation `json:"repo_detail,omitempty"`

	// Wire name: 'share_name'
	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateInstallationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createInstallationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateListingRequest struct {

	// Wire name: 'listing'
	Listing Listing `json:"listing"`
}

func (st *CreateListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateListingResponse struct {

	// Wire name: 'listing_id'
	ListingId string `json:"listing_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateListingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createListingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {

	// Wire name: 'accepted_consumer_terms'
	AcceptedConsumerTerms ConsumerTerms `json:"accepted_consumer_terms"`

	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`

	// Wire name: 'company'
	Company string `json:"company,omitempty"`

	// Wire name: 'first_name'
	FirstName string `json:"first_name,omitempty"`

	// Wire name: 'intended_use'
	IntendedUse string `json:"intended_use"`

	// Wire name: 'is_from_lighthouse'
	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`

	// Wire name: 'last_name'
	LastName string `json:"last_name,omitempty"`

	ListingId string `json:"-" tf:"-"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePersonalizationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createPersonalizationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreatePersonalizationRequestResponse struct {

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePersonalizationRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createPersonalizationRequestResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateProviderRequest struct {

	// Wire name: 'provider'
	Provider ProviderInfo `json:"provider"`
}

func (st *CreateProviderRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createProviderRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateProviderResponse struct {

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateProviderResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createProviderResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DataRefreshInfo struct {

	// Wire name: 'interval'
	Interval int64 `json:"interval"`

	// Wire name: 'unit'
	Unit DataRefresh `json:"unit"`
}

func (st *DataRefreshInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := dataRefreshInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteExchangeFilterRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *DeleteExchangeFilterRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExchangeFilterRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteExchangeRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *DeleteExchangeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExchangeRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteFileRequest struct {
	FileId string `json:"-" tf:"-"`
}

func (st *DeleteFileRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteFileRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteInstallationRequest struct {
	InstallationId string `json:"-" tf:"-"`

	ListingId string `json:"-" tf:"-"`
}

func (st *DeleteInstallationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteInstallationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteListingRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *DeleteListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteProviderRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *DeleteProviderRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteProviderRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Exchange struct {

	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`

	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'filters'
	Filters []ExchangeFilter `json:"filters,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'linked_listings'
	LinkedListings []ExchangeListing `json:"linked_listings,omitempty"`

	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Exchange) EncodeValues(key string, v *url.Values) error {
	pb, err := exchangeToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ExchangeFilter struct {

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`

	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'exchange_id'
	ExchangeId string `json:"exchange_id"`

	// Wire name: 'filter_type'
	FilterType ExchangeFilterType `json:"filter_type"`

	// Wire name: 'filter_value'
	FilterValue string `json:"filter_value"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ExchangeFilter) EncodeValues(key string, v *url.Values) error {
	pb, err := exchangeFilterToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ExchangeListing struct {

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`

	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'exchange_id'
	ExchangeId string `json:"exchange_id,omitempty"`

	// Wire name: 'exchange_name'
	ExchangeName string `json:"exchange_name,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'listing_id'
	ListingId string `json:"listing_id,omitempty"`

	// Wire name: 'listing_name'
	ListingName string `json:"listing_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ExchangeListing) EncodeValues(key string, v *url.Values) error {
	pb, err := exchangeListingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FileInfo struct {

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`

	// Wire name: 'download_link'
	DownloadLink string `json:"download_link,omitempty"`

	// Wire name: 'file_parent'
	FileParent *FileParent `json:"file_parent,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'marketplace_file_type'
	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type,omitempty"`

	// Wire name: 'mime_type'
	MimeType string `json:"mime_type,omitempty"`

	// Wire name: 'status'
	Status FileStatus `json:"status,omitempty"`
	// Populated if status is in a failed state with more information on reason
	// for the failure.
	// Wire name: 'status_message'
	StatusMessage string `json:"status_message,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *FileInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := fileInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FileParent struct {

	// Wire name: 'file_parent_type'
	FileParentType FileParentType `json:"file_parent_type,omitempty"`
	// TODO make the following fields required
	// Wire name: 'parent_id'
	ParentId string `json:"parent_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *FileParent) EncodeValues(key string, v *url.Values) error {
	pb, err := fileParentToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetExchangeRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *GetExchangeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getExchangeRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetExchangeResponse struct {

	// Wire name: 'exchange'
	Exchange *Exchange `json:"exchange,omitempty"`
}

func (st *GetExchangeResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getExchangeResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetFileRequest struct {
	FileId string `json:"-" tf:"-"`
}

func (st *GetFileRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getFileRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetFileResponse struct {

	// Wire name: 'file_info'
	FileInfo *FileInfo `json:"file_info,omitempty"`
}

func (st *GetFileResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getFileResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	// Wire name: 'version'
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetLatestVersionProviderAnalyticsDashboardResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getLatestVersionProviderAnalyticsDashboardResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetListingContentMetadataRequest struct {
	ListingId string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetListingContentMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getListingContentMetadataRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetListingContentMetadataResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'shared_data_objects'
	SharedDataObjects []SharedDataObject `json:"shared_data_objects,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetListingContentMetadataResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getListingContentMetadataResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetListingRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *GetListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetListingResponse struct {

	// Wire name: 'listing'
	Listing *Listing `json:"listing,omitempty"`
}

func (st *GetListingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getListingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetListingsRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetListingsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getListingsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing `json:"listings,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetListingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getListingsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetPersonalizationRequestRequest struct {
	ListingId string `json:"-" tf:"-"`
}

func (st *GetPersonalizationRequestRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPersonalizationRequestRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetPersonalizationRequestResponse struct {

	// Wire name: 'personalization_requests'
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`
}

func (st *GetPersonalizationRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getPersonalizationRequestResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetProviderRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *GetProviderRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getProviderRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetProviderResponse struct {

	// Wire name: 'provider'
	Provider *ProviderInfo `json:"provider,omitempty"`
}

func (st *GetProviderResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getProviderResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Installation struct {

	// Wire name: 'installation'
	Installation *InstallationDetail `json:"installation,omitempty"`
}

func (st *Installation) EncodeValues(key string, v *url.Values) error {
	pb, err := installationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type InstallationDetail struct {

	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`

	// Wire name: 'error_message'
	ErrorMessage string `json:"error_message,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'installed_on'
	InstalledOn int64 `json:"installed_on,omitempty"`

	// Wire name: 'listing_id'
	ListingId string `json:"listing_id,omitempty"`

	// Wire name: 'listing_name'
	ListingName string `json:"listing_name,omitempty"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	// Wire name: 'repo_name'
	RepoName string `json:"repo_name,omitempty"`

	// Wire name: 'repo_path'
	RepoPath string `json:"repo_path,omitempty"`

	// Wire name: 'share_name'
	ShareName string `json:"share_name,omitempty"`

	// Wire name: 'status'
	Status InstallationStatus `json:"status,omitempty"`

	// Wire name: 'token_detail'
	TokenDetail *TokenDetail `json:"token_detail,omitempty"`

	// Wire name: 'tokens'
	Tokens []TokenInfo `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstallationDetail) EncodeValues(key string, v *url.Values) error {
	pb, err := installationDetailToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListAllInstallationsRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAllInstallationsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listAllInstallationsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListAllInstallationsResponse struct {

	// Wire name: 'installations'
	Installations []InstallationDetail `json:"installations,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAllInstallationsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listAllInstallationsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListAllPersonalizationRequestsRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAllPersonalizationRequestsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listAllPersonalizationRequestsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListAllPersonalizationRequestsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'personalization_requests'
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAllPersonalizationRequestsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listAllPersonalizationRequestsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExchangeFiltersRequest struct {
	ExchangeId string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListExchangeFiltersRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExchangeFiltersRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExchangeFiltersResponse struct {

	// Wire name: 'filters'
	Filters []ExchangeFilter `json:"filters,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListExchangeFiltersResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExchangeFiltersResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExchangesForListingRequest struct {
	ListingId string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListExchangesForListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExchangesForListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExchangesForListingResponse struct {

	// Wire name: 'exchange_listing'
	ExchangeListing []ExchangeListing `json:"exchange_listing,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListExchangesForListingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExchangesForListingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExchangesRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListExchangesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExchangesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExchangesResponse struct {

	// Wire name: 'exchanges'
	Exchanges []Exchange `json:"exchanges,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListExchangesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExchangesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListFilesRequest struct {
	FileParent FileParent `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListFilesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listFilesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListFilesResponse struct {

	// Wire name: 'file_infos'
	FileInfos []FileInfo `json:"file_infos,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListFilesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listFilesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListFulfillmentsRequest struct {
	ListingId string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListFulfillmentsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listFulfillmentsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListFulfillmentsResponse struct {

	// Wire name: 'fulfillments'
	Fulfillments []ListingFulfillment `json:"fulfillments,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListFulfillmentsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listFulfillmentsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListInstallationsRequest struct {
	ListingId string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListInstallationsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listInstallationsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListInstallationsResponse struct {

	// Wire name: 'installations'
	Installations []InstallationDetail `json:"installations,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListInstallationsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listInstallationsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListListingsForExchangeRequest struct {
	ExchangeId string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListListingsForExchangeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listListingsForExchangeRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListListingsForExchangeResponse struct {

	// Wire name: 'exchange_listings'
	ExchangeListings []ExchangeListing `json:"exchange_listings,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListListingsForExchangeResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listListingsForExchangeResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `json:"-" tf:"-"`
	// Matches any of the following categories
	Categories []Category `json:"-" tf:"-"`
	// Filters each listing based on if it is free.
	IsFree bool `json:"-" tf:"-"`
	// Filters each listing based on if it is a private exchange.
	IsPrivateExchange bool `json:"-" tf:"-"`
	// Filters each listing based on whether it is a staff pick.
	IsStaffPick bool `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`
	// Matches any of the following provider ids
	ProviderIds []string `json:"-" tf:"-"`
	// Matches any of the following tags
	Tags []ListingTag `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListListingsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listListingsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing `json:"listings,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListListingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listListingsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id"`

	// Wire name: 'id'
	Id string `json:"id"`

	// Wire name: 'version'
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProviderAnalyticsDashboardResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listProviderAnalyticsDashboardResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListProvidersRequest struct {
	IsFeatured bool `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProvidersRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listProvidersRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListProvidersResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'providers'
	Providers []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProvidersResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listProvidersResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Listing struct {

	// Wire name: 'detail'
	Detail *ListingDetail `json:"detail,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'summary'
	Summary ListingSummary `json:"summary"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Listing) EncodeValues(key string, v *url.Values) error {
	pb, err := listingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	// Wire name: 'assets'
	Assets []AssetType `json:"assets,omitempty"`
	// The ending date timestamp for when the data spans
	// Wire name: 'collection_date_end'
	CollectionDateEnd int64 `json:"collection_date_end,omitempty"`
	// The starting date timestamp for when the data spans
	// Wire name: 'collection_date_start'
	CollectionDateStart int64 `json:"collection_date_start,omitempty"`
	// Smallest unit of time in the dataset
	// Wire name: 'collection_granularity'
	CollectionGranularity *DataRefreshInfo `json:"collection_granularity,omitempty"`
	// Whether the dataset is free or paid
	// Wire name: 'cost'
	Cost Cost `json:"cost,omitempty"`
	// Where/how the data is sourced
	// Wire name: 'data_source'
	DataSource string `json:"data_source,omitempty"`

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'documentation_link'
	DocumentationLink string `json:"documentation_link,omitempty"`

	// Wire name: 'embedded_notebook_file_infos'
	EmbeddedNotebookFileInfos []FileInfo `json:"embedded_notebook_file_infos,omitempty"`

	// Wire name: 'file_ids'
	FileIds []string `json:"file_ids,omitempty"`
	// Which geo region the listing data is collected from
	// Wire name: 'geographical_coverage'
	GeographicalCoverage string `json:"geographical_coverage,omitempty"`
	// ID 20, 21 removed don't use License of the data asset - Required for
	// listings with model based assets
	// Wire name: 'license'
	License string `json:"license,omitempty"`
	// What the pricing model is (e.g. paid, subscription, paid upfront); should
	// only be present if cost is paid TODO: Not used yet, should deprecate if
	// we will never use it
	// Wire name: 'pricing_model'
	PricingModel string `json:"pricing_model,omitempty"`

	// Wire name: 'privacy_policy_link'
	PrivacyPolicyLink string `json:"privacy_policy_link,omitempty"`
	// size of the dataset in GB
	// Wire name: 'size'
	Size float64 `json:"size,omitempty"`

	// Wire name: 'support_link'
	SupportLink string `json:"support_link,omitempty"`
	// Listing tags - Simple key value pair to annotate listings. When should I
	// use tags vs dedicated fields? Using tags avoids the need to add new
	// columns in the database for new annotations. However, this should be used
	// sparingly since tags are stored as key value pair. Use tags only: 1. If
	// the field is optional and won't need to have NOT NULL integrity check 2.
	// The value is fairly fixed, static and low cardinality (eg. enums). 3. The
	// value won't be used in filters or joins with other tables.
	// Wire name: 'tags'
	Tags []ListingTag `json:"tags,omitempty"`

	// Wire name: 'terms_of_service'
	TermsOfService string `json:"terms_of_service,omitempty"`
	// How often data is updated
	// Wire name: 'update_frequency'
	UpdateFrequency *DataRefreshInfo `json:"update_frequency,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListingDetail) EncodeValues(key string, v *url.Values) error {
	pb, err := listingDetailToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListingFulfillment struct {

	// Wire name: 'fulfillment_type'
	FulfillmentType FulfillmentType `json:"fulfillment_type,omitempty"`

	// Wire name: 'listing_id'
	ListingId string `json:"listing_id"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	// Wire name: 'repo_info'
	RepoInfo *RepoInfo `json:"repo_info,omitempty"`

	// Wire name: 'share_info'
	ShareInfo *ShareInfo `json:"share_info,omitempty"`
}

func (st *ListingFulfillment) EncodeValues(key string, v *url.Values) error {
	pb, err := listingFulfillmentToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListingSetting struct {

	// Wire name: 'visibility'
	Visibility Visibility `json:"visibility,omitempty"`
}

func (st *ListingSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := listingSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListingSummary struct {

	// Wire name: 'categories'
	Categories []Category `json:"categories,omitempty"`

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`

	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'created_by_id'
	CreatedById int64 `json:"created_by_id,omitempty"`

	// Wire name: 'exchange_ids'
	ExchangeIds []string `json:"exchange_ids,omitempty"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	// Wire name: 'git_repo'
	GitRepo *RepoInfo `json:"git_repo,omitempty"`

	// Wire name: 'listingType'
	ListingType ListingType `json:"listingType"`

	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'provider_id'
	ProviderId string `json:"provider_id,omitempty"`

	// Wire name: 'provider_region'
	ProviderRegion *RegionInfo `json:"provider_region,omitempty"`

	// Wire name: 'published_at'
	PublishedAt int64 `json:"published_at,omitempty"`

	// Wire name: 'published_by'
	PublishedBy string `json:"published_by,omitempty"`

	// Wire name: 'setting'
	Setting *ListingSetting `json:"setting,omitempty"`

	// Wire name: 'share'
	Share *ShareInfo `json:"share,omitempty"`

	// Wire name: 'status'
	Status ListingStatus `json:"status,omitempty"`

	// Wire name: 'subtitle'
	Subtitle string `json:"subtitle,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	// Wire name: 'updated_by_id'
	UpdatedById int64 `json:"updated_by_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListingSummary) EncodeValues(key string, v *url.Values) error {
	pb, err := listingSummaryToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListingTag struct {
	// Tag name (enum)
	// Wire name: 'tag_name'
	TagName ListingTagType `json:"tag_name,omitempty"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	// Wire name: 'tag_values'
	TagValues []string `json:"tag_values,omitempty"`
}

func (st *ListingTag) EncodeValues(key string, v *url.Values) error {
	pb, err := listingTagToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type PersonalizationRequest struct {

	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`

	// Wire name: 'consumer_region'
	ConsumerRegion RegionInfo `json:"consumer_region"`

	// Wire name: 'contact_info'
	ContactInfo *ContactInfo `json:"contact_info,omitempty"`

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'intended_use'
	IntendedUse string `json:"intended_use,omitempty"`

	// Wire name: 'is_from_lighthouse'
	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`

	// Wire name: 'listing_id'
	ListingId string `json:"listing_id,omitempty"`

	// Wire name: 'listing_name'
	ListingName string `json:"listing_name,omitempty"`

	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`

	// Wire name: 'provider_id'
	ProviderId string `json:"provider_id,omitempty"`

	// Wire name: 'recipient_type'
	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	// Wire name: 'share'
	Share *ShareInfo `json:"share,omitempty"`

	// Wire name: 'status'
	Status PersonalizationRequestStatus `json:"status,omitempty"`

	// Wire name: 'status_message'
	StatusMessage string `json:"status_message,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PersonalizationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := personalizationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ProviderAnalyticsDashboard struct {

	// Wire name: 'id'
	Id string `json:"id"`
}

func (st *ProviderAnalyticsDashboard) EncodeValues(key string, v *url.Values) error {
	pb, err := providerAnalyticsDashboardToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ProviderInfo struct {

	// Wire name: 'business_contact_email'
	BusinessContactEmail string `json:"business_contact_email"`

	// Wire name: 'company_website_link'
	CompanyWebsiteLink string `json:"company_website_link,omitempty"`

	// Wire name: 'dark_mode_icon_file_id'
	DarkModeIconFileId string `json:"dark_mode_icon_file_id,omitempty"`

	// Wire name: 'dark_mode_icon_file_path'
	DarkModeIconFilePath string `json:"dark_mode_icon_file_path,omitempty"`

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'icon_file_id'
	IconFileId string `json:"icon_file_id,omitempty"`

	// Wire name: 'icon_file_path'
	IconFilePath string `json:"icon_file_path,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// is_featured is accessible by consumers only
	// Wire name: 'is_featured'
	IsFeatured bool `json:"is_featured,omitempty"`

	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'privacy_policy_link'
	PrivacyPolicyLink string `json:"privacy_policy_link"`
	// published_by is only applicable to data aggregators (e.g. Crux)
	// Wire name: 'published_by'
	PublishedBy string `json:"published_by,omitempty"`

	// Wire name: 'support_contact_email'
	SupportContactEmail string `json:"support_contact_email,omitempty"`

	// Wire name: 'term_of_service_link'
	TermOfServiceLink string `json:"term_of_service_link"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ProviderInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := providerInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RegionInfo struct {

	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`

	// Wire name: 'region'
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RegionInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := regionInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RemoveExchangeForListingRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *RemoveExchangeForListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := removeExchangeForListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	// Wire name: 'git_repo_url'
	GitRepoUrl string `json:"git_repo_url"`
}

func (st *RepoInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := repoInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	// Wire name: 'repo_name'
	RepoName string `json:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	// Wire name: 'repo_path'
	RepoPath string `json:"repo_path"`
}

func (st *RepoInstallation) EncodeValues(key string, v *url.Values) error {
	pb, err := repoInstallationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SearchListingsRequest struct {
	// Matches any of the following asset types
	Assets []AssetType `json:"-" tf:"-"`
	// Matches any of the following categories
	Categories []Category `json:"-" tf:"-"`

	IsFree bool `json:"-" tf:"-"`

	IsPrivateExchange bool `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`
	// Matches any of the following provider ids
	ProviderIds []string `json:"-" tf:"-"`
	// Fuzzy matches query
	Query string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SearchListingsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := searchListingsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SearchListingsResponse struct {

	// Wire name: 'listings'
	Listings []Listing `json:"listings,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SearchListingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := searchListingsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ShareInfo struct {

	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'type'
	Type ListingShareType `json:"type"`
}

func (st *ShareInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := shareInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	// Wire name: 'data_object_type'
	DataObjectType string `json:"data_object_type,omitempty"`
	// Name of the shared object
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SharedDataObject) EncodeValues(key string, v *url.Values) error {
	pb, err := sharedDataObjectToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TokenDetail struct {

	// Wire name: 'bearerToken'
	BearerToken string `json:"bearerToken,omitempty"`

	// Wire name: 'endpoint'
	Endpoint string `json:"endpoint,omitempty"`

	// Wire name: 'expirationTime'
	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	// Wire name: 'shareCredentialsVersion'
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenDetail) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenDetailToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TokenInfo struct {
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this Recipient Token was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Recipient Token creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique id of the Recipient Token.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Time at which this Recipient Token was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of Recipient Token updater.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateExchangeFilterRequest struct {

	// Wire name: 'filter'
	Filter ExchangeFilter `json:"filter"`

	Id string `json:"-" tf:"-"`
}

func (st *UpdateExchangeFilterRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExchangeFilterRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateExchangeFilterResponse struct {

	// Wire name: 'filter'
	Filter *ExchangeFilter `json:"filter,omitempty"`
}

func (st *UpdateExchangeFilterResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExchangeFilterResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateExchangeRequest struct {

	// Wire name: 'exchange'
	Exchange Exchange `json:"exchange"`

	Id string `json:"-" tf:"-"`
}

func (st *UpdateExchangeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExchangeRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateExchangeResponse struct {

	// Wire name: 'exchange'
	Exchange *Exchange `json:"exchange,omitempty"`
}

func (st *UpdateExchangeResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExchangeResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateInstallationRequest struct {

	// Wire name: 'installation'
	Installation InstallationDetail `json:"installation"`

	InstallationId string `json:"-" tf:"-"`

	ListingId string `json:"-" tf:"-"`

	// Wire name: 'rotate_token'
	RotateToken bool `json:"rotate_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateInstallationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateInstallationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateInstallationResponse struct {

	// Wire name: 'installation'
	Installation *InstallationDetail `json:"installation,omitempty"`
}

func (st *UpdateInstallationResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateInstallationResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateListingRequest struct {
	Id string `json:"-" tf:"-"`

	// Wire name: 'listing'
	Listing Listing `json:"listing"`
}

func (st *UpdateListingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateListingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateListingResponse struct {

	// Wire name: 'listing'
	Listing *Listing `json:"listing,omitempty"`
}

func (st *UpdateListingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateListingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdatePersonalizationRequestRequest struct {
	ListingId string `json:"-" tf:"-"`

	// Wire name: 'reason'
	Reason string `json:"reason,omitempty"`

	RequestId string `json:"-" tf:"-"`

	// Wire name: 'share'
	Share *ShareInfo `json:"share,omitempty"`

	// Wire name: 'status'
	Status PersonalizationRequestStatus `json:"status"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdatePersonalizationRequestRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updatePersonalizationRequestRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdatePersonalizationRequestResponse struct {

	// Wire name: 'request'
	Request *PersonalizationRequest `json:"request,omitempty"`
}

func (st *UpdatePersonalizationRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updatePersonalizationRequestResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id string `json:"-" tf:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	// Wire name: 'version'
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateProviderAnalyticsDashboardRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateProviderAnalyticsDashboardRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id"`
	// id & version should be the same as the request
	// Wire name: 'id'
	Id string `json:"id"`

	// Wire name: 'version'
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateProviderAnalyticsDashboardResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateProviderAnalyticsDashboardResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateProviderRequest struct {
	Id string `json:"-" tf:"-"`

	// Wire name: 'provider'
	Provider ProviderInfo `json:"provider"`
}

func (st *UpdateProviderRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateProviderRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateProviderResponse struct {

	// Wire name: 'provider'
	Provider *ProviderInfo `json:"provider,omitempty"`
}

func (st *UpdateProviderResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateProviderResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
