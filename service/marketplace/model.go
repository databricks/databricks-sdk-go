// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddExchangeForListingRequest struct {
	ExchangeId string `json:"exchange_id"`

	ListingId string `json:"listing_id"`
}

type AddExchangeForListingResponse struct {
	ExchangeForListing *ExchangeListing `json:"exchange_for_listing,omitempty"`
}

type AssetType string

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
	case `ASSET_TYPE_DATA_TABLE`, `ASSET_TYPE_GIT_REPO`, `ASSET_TYPE_MEDIA`, `ASSET_TYPE_MODEL`, `ASSET_TYPE_NOTEBOOK`, `ASSET_TYPE_PARTNER_INTEGRATION`:
		*f = AssetType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASSET_TYPE_DATA_TABLE", "ASSET_TYPE_GIT_REPO", "ASSET_TYPE_MEDIA", "ASSET_TYPE_MODEL", "ASSET_TYPE_NOTEBOOK", "ASSET_TYPE_PARTNER_INTEGRATION"`, v)
	}
}

// Type always returns AssetType to satisfy [pflag.Value] interface
func (f *AssetType) Type() string {
	return "AssetType"
}

// Get one batch of listings. One may specify up to 50 IDs per request.
type BatchGetListingsRequest struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

type BatchGetListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`
}

// Get one batch of providers. One may specify up to 50 IDs per request.
type BatchGetProvidersRequest struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

type BatchGetProvidersResponse struct {
	Providers []ProviderInfo `json:"providers,omitempty"`
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

// Type always returns Category to satisfy [pflag.Value] interface
func (f *Category) Type() string {
	return "Category"
}

type ConsumerTerms struct {
	Version string `json:"version"`
}

// contact info for the consumer requesting data or performing a listing
// installation
type ContactInfo struct {
	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContactInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContactInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns Cost to satisfy [pflag.Value] interface
func (f *Cost) Type() string {
	return "Cost"
}

type CreateExchangeFilterRequest struct {
	Filter ExchangeFilter `json:"filter"`
}

type CreateExchangeFilterResponse struct {
	FilterId string `json:"filter_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExchangeFilterResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeFilterResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExchangeRequest struct {
	Exchange Exchange `json:"exchange"`
}

type CreateExchangeResponse struct {
	ExchangeId string `json:"exchange_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFileRequest struct {
	DisplayName string `json:"display_name,omitempty"`

	FileParent FileParent `json:"file_parent"`

	MarketplaceFileType MarketplaceFileType `json:"marketplace_file_type"`

	MimeType string `json:"mime_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFileResponse struct {
	FileInfo *FileInfo `json:"file_info,omitempty"`
	// Pre-signed POST URL to blob storage
	UploadUrl string `json:"upload_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFileResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFileResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateInstallationRequest struct {
	AcceptedConsumerTerms *ConsumerTerms `json:"accepted_consumer_terms,omitempty"`

	CatalogName string `json:"catalog_name,omitempty"`

	ListingId string `json:"-" url:"-"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`
	// for git repo installations
	RepoDetail *RepoInstallation `json:"repo_detail,omitempty"`

	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateListingRequest struct {
	Listing Listing `json:"listing"`
}

type CreateListingResponse struct {
	ListingId string `json:"listing_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Data request messages also creates a lead (maybe)
type CreatePersonalizationRequest struct {
	AcceptedConsumerTerms ConsumerTerms `json:"accepted_consumer_terms"`

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

func (s *CreatePersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePersonalizationRequestResponse struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePersonalizationRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePersonalizationRequestResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateProviderRequest struct {
	Provider ProviderInfo `json:"provider"`
}

type CreateProviderResponse struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateProviderResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProviderResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns DataRefresh to satisfy [pflag.Value] interface
func (f *DataRefresh) Type() string {
	return "DataRefresh"
}

type DataRefreshInfo struct {
	Interval int64 `json:"interval"`

	Unit DataRefresh `json:"unit"`
}

// Delete an exchange filter
type DeleteExchangeFilterRequest struct {
	Id string `json:"-" url:"-"`
}

type DeleteExchangeFilterResponse struct {
}

// Delete an exchange
type DeleteExchangeRequest struct {
	Id string `json:"-" url:"-"`
}

type DeleteExchangeResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	FileId string `json:"-" url:"-"`
}

type DeleteFileResponse struct {
}

// Uninstall from a listing
type DeleteInstallationRequest struct {
	InstallationId string `json:"-" url:"-"`

	ListingId string `json:"-" url:"-"`
}

type DeleteInstallationResponse struct {
}

// Delete a listing
type DeleteListingRequest struct {
	Id string `json:"-" url:"-"`
}

type DeleteListingResponse struct {
}

// Delete provider
type DeleteProviderRequest struct {
	Id string `json:"-" url:"-"`
}

type DeleteProviderResponse struct {
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

// Type always returns DeltaSharingRecipientType to satisfy [pflag.Value] interface
func (f *DeltaSharingRecipientType) Type() string {
	return "DeltaSharingRecipientType"
}

type Exchange struct {
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	Filters []ExchangeFilter `json:"filters,omitempty"`

	Id string `json:"id,omitempty"`

	LinkedListings []ExchangeListing `json:"linked_listings,omitempty"`

	Name string `json:"name"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Exchange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Exchange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExchangeFilter struct {
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

func (s *ExchangeFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns ExchangeFilterType to satisfy [pflag.Value] interface
func (f *ExchangeFilterType) Type() string {
	return "ExchangeFilterType"
}

type ExchangeListing struct {
	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	ExchangeId string `json:"exchange_id,omitempty"`

	ExchangeName string `json:"exchange_name,omitempty"`

	Id string `json:"id,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExchangeListing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeListing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileInfo struct {
	CreatedAt int64 `json:"created_at,omitempty"`
	// Name displayed to users for applicable files, e.g. embedded notebooks
	DisplayName string `json:"display_name,omitempty"`

	DownloadLink string `json:"download_link,omitempty"`

	FileParent *FileParent `json:"file_parent,omitempty"`

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

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileParent struct {
	FileParentType FileParentType `json:"file_parent_type,omitempty"`
	// TODO make the following fields required
	ParentId string `json:"parent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileParent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileParent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileParentType string

const FileParentTypeListing FileParentType = `LISTING`

const FileParentTypeProvider FileParentType = `PROVIDER`

// String representation for [fmt.Print]
func (f *FileParentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FileParentType) Set(v string) error {
	switch v {
	case `LISTING`, `PROVIDER`:
		*f = FileParentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LISTING", "PROVIDER"`, v)
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

// Type always returns FulfillmentType to satisfy [pflag.Value] interface
func (f *FulfillmentType) Type() string {
	return "FulfillmentType"
}

// Get an exchange
type GetExchangeRequest struct {
	Id string `json:"-" url:"-"`
}

type GetExchangeResponse struct {
	Exchange *Exchange `json:"exchange,omitempty"`
}

// Get a file
type GetFileRequest struct {
	FileId string `json:"-" url:"-"`
}

type GetFileResponse struct {
	FileInfo *FileInfo `json:"file_info,omitempty"`
}

type GetLatestVersionProviderAnalyticsDashboardResponse struct {
	// version here is latest logical version of the dashboard template
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLatestVersionProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLatestVersionProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get listing content metadata
type GetListingContentMetadataRequest struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingContentMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingContentMetadataResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	SharedDataObjects []SharedDataObject `json:"shared_data_objects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingContentMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingContentMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get listing
type GetListingRequest struct {
	Id string `json:"-" url:"-"`
}

type GetListingResponse struct {
	Listing *Listing `json:"listing,omitempty"`
}

// List listings
type GetListingsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the personalization request for a listing
type GetPersonalizationRequestRequest struct {
	ListingId string `json:"-" url:"-"`
}

type GetPersonalizationRequestResponse struct {
	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`
}

// Get a provider
type GetProviderRequest struct {
	Id string `json:"-" url:"-"`
}

type GetProviderResponse struct {
	Provider *ProviderInfo `json:"provider,omitempty"`
}

type Installation struct {
	Installation *InstallationDetail `json:"installation,omitempty"`
}

type InstallationDetail struct {
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

	TokenDetail *TokenDetail `json:"token_detail,omitempty"`

	Tokens []TokenInfo `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstallationDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstallationDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns InstallationStatus to satisfy [pflag.Value] interface
func (f *InstallationStatus) Type() string {
	return "InstallationStatus"
}

// List all installations
type ListAllInstallationsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllInstallationsResponse struct {
	Installations []InstallationDetail `json:"installations,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all personalization requests
type ListAllPersonalizationRequestsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllPersonalizationRequestsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAllPersonalizationRequestsResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	PersonalizationRequests []PersonalizationRequest `json:"personalization_requests,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAllPersonalizationRequestsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAllPersonalizationRequestsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List exchange filters
type ListExchangeFiltersRequest struct {
	ExchangeId string `json:"-" url:"exchange_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangeFiltersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangeFiltersResponse struct {
	Filters []ExchangeFilter `json:"filters,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangeFiltersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangeFiltersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List exchanges for listing
type ListExchangesForListingRequest struct {
	ListingId string `json:"-" url:"listing_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesForListingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesForListingResponse struct {
	ExchangeListing []ExchangeListing `json:"exchange_listing,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesForListingResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesForListingResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List exchanges
type ListExchangesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExchangesResponse struct {
	Exchanges []Exchange `json:"exchanges,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExchangesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExchangesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List files
type ListFilesRequest struct {
	FileParent FileParent `json:"-" url:"file_parent"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFilesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFilesResponse struct {
	FileInfos []FileInfo `json:"file_infos,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFilesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFilesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all listing fulfillments
type ListFulfillmentsRequest struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFulfillmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFulfillmentsResponse struct {
	Fulfillments []ListingFulfillment `json:"fulfillments,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFulfillmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFulfillmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List installations for a listing
type ListInstallationsRequest struct {
	ListingId string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListInstallationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListInstallationsResponse struct {
	Installations []InstallationDetail `json:"installations,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListInstallationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListInstallationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List listings for exchange
type ListListingsForExchangeRequest struct {
	ExchangeId string `json:"-" url:"exchange_id"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsForExchangeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsForExchangeResponse struct {
	ExchangeListings []ExchangeListing `json:"exchange_listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsForExchangeResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsForExchangeResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List listings
type ListListingsRequest struct {
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
	Tags []ListingTag `json:"-" url:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProviderAnalyticsDashboardResponse struct {
	// dashboard_id will be used to open Lakeview dashboard.
	DashboardId string `json:"dashboard_id"`

	Id string `json:"id"`

	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List providers
type ListProvidersRequest struct {
	IsFeatured bool `json:"-" url:"is_featured,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProvidersResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Providers []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Listing struct {
	Detail *ListingDetail `json:"detail,omitempty"`

	Id string `json:"id,omitempty"`
	// Next Number: 26
	Summary ListingSummary `json:"summary"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Listing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Listing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingDetail struct {
	// Type of assets included in the listing. eg. GIT_REPO, DATA_TABLE, MODEL,
	// NOTEBOOK
	Assets []AssetType `json:"assets,omitempty"`
	// The ending date timestamp for when the data spans
	CollectionDateEnd int64 `json:"collection_date_end,omitempty"`
	// The starting date timestamp for when the data spans
	CollectionDateStart int64 `json:"collection_date_start,omitempty"`
	// Smallest unit of time in the dataset
	CollectionGranularity *DataRefreshInfo `json:"collection_granularity,omitempty"`
	// Whether the dataset is free or paid
	Cost Cost `json:"cost,omitempty"`
	// Where/how the data is sourced
	DataSource string `json:"data_source,omitempty"`

	Description string `json:"description,omitempty"`

	DocumentationLink string `json:"documentation_link,omitempty"`

	EmbeddedNotebookFileInfos []FileInfo `json:"embedded_notebook_file_infos,omitempty"`

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
	Tags []ListingTag `json:"tags,omitempty"`

	TermsOfService string `json:"terms_of_service,omitempty"`
	// How often data is updated
	UpdateFrequency *DataRefreshInfo `json:"update_frequency,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListingDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingFulfillment struct {
	FulfillmentType FulfillmentType `json:"fulfillment_type,omitempty"`

	ListingId string `json:"listing_id"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	RepoInfo *RepoInfo `json:"repo_info,omitempty"`

	ShareInfo *ShareInfo `json:"share_info,omitempty"`
}

type ListingSetting struct {
	Visibility Visibility `json:"visibility,omitempty"`
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

// Type always returns ListingStatus to satisfy [pflag.Value] interface
func (f *ListingStatus) Type() string {
	return "ListingStatus"
}

// Next Number: 26
type ListingSummary struct {
	Categories []Category `json:"categories,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty"`

	CreatedById int64 `json:"created_by_id,omitempty"`

	ExchangeIds []string `json:"exchange_ids,omitempty"`
	// if a git repo is being created, a listing will be initialized with this
	// field as opposed to a share
	GitRepo *RepoInfo `json:"git_repo,omitempty"`

	ListingType ListingType `json:"listingType"`

	Name string `json:"name"`

	ProviderId string `json:"provider_id,omitempty"`

	ProviderRegion *RegionInfo `json:"provider_region,omitempty"`

	PublishedAt int64 `json:"published_at,omitempty"`

	PublishedBy string `json:"published_by,omitempty"`

	Setting *ListingSetting `json:"setting,omitempty"`

	Share *ShareInfo `json:"share,omitempty"`
	// Enums
	Status ListingStatus `json:"status,omitempty"`

	Subtitle string `json:"subtitle,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy string `json:"updated_by,omitempty"`

	UpdatedById int64 `json:"updated_by_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListingSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListingSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListingTag struct {
	// Tag name (enum)
	TagName ListingTagType `json:"tag_name,omitempty"`
	// String representation of the tag value. Values should be string literals
	// (no complex types)
	TagValues []string `json:"tag_values,omitempty"`
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

// Type always returns ListingType to satisfy [pflag.Value] interface
func (f *ListingType) Type() string {
	return "ListingType"
}

type MarketplaceFileType string

const MarketplaceFileTypeEmbeddedNotebook MarketplaceFileType = `EMBEDDED_NOTEBOOK`

const MarketplaceFileTypeProviderIcon MarketplaceFileType = `PROVIDER_ICON`

// String representation for [fmt.Print]
func (f *MarketplaceFileType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MarketplaceFileType) Set(v string) error {
	switch v {
	case `EMBEDDED_NOTEBOOK`, `PROVIDER_ICON`:
		*f = MarketplaceFileType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMBEDDED_NOTEBOOK", "PROVIDER_ICON"`, v)
	}
}

// Type always returns MarketplaceFileType to satisfy [pflag.Value] interface
func (f *MarketplaceFileType) Type() string {
	return "MarketplaceFileType"
}

type PersonalizationRequest struct {
	Comment string `json:"comment,omitempty"`

	ConsumerRegion RegionInfo `json:"consumer_region"`
	// contact info for the consumer requesting data or performing a listing
	// installation
	ContactInfo *ContactInfo `json:"contact_info,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	Id string `json:"id,omitempty"`

	IntendedUse string `json:"intended_use,omitempty"`

	IsFromLighthouse bool `json:"is_from_lighthouse,omitempty"`

	ListingId string `json:"listing_id,omitempty"`

	ListingName string `json:"listing_name,omitempty"`

	MetastoreId string `json:"metastore_id,omitempty"`

	ProviderId string `json:"provider_id,omitempty"`

	RecipientType DeltaSharingRecipientType `json:"recipient_type,omitempty"`

	Share *ShareInfo `json:"share,omitempty"`

	Status PersonalizationRequestStatus `json:"status,omitempty"`

	StatusMessage string `json:"status_message,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PersonalizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalizationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns PersonalizationRequestStatus to satisfy [pflag.Value] interface
func (f *PersonalizationRequestStatus) Type() string {
	return "PersonalizationRequestStatus"
}

type ProviderAnalyticsDashboard struct {
	Id string `json:"id"`
}

type ProviderInfo struct {
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

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegionInfo struct {
	Cloud string `json:"cloud,omitempty"`

	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Remove an exchange for listing
type RemoveExchangeForListingRequest struct {
	Id string `json:"-" url:"-"`
}

type RemoveExchangeForListingResponse struct {
}

type RepoInfo struct {
	// the git repo url e.g. https://github.com/databrickslabs/dolly.git
	GitRepoUrl string `json:"git_repo_url"`
}

type RepoInstallation struct {
	// the user-specified repo name for their installed git repo listing
	RepoName string `json:"repo_name"`
	// refers to the full url file path that navigates the user to the repo's
	// entrypoint (e.g. a README.md file, or the repo file view in the unified
	// UI) should just be a relative path
	RepoPath string `json:"repo_path"`
}

// Search listings
type SearchListingsRequest struct {
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

func (s *SearchListingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchListingsResponse struct {
	Listings []Listing `json:"listings,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchListingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchListingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ShareInfo struct {
	Name string `json:"name"`

	Type ListingShareType `json:"type"`
}

type SharedDataObject struct {
	// The type of the data object. Could be one of: TABLE, SCHEMA,
	// NOTEBOOK_FILE, MODEL, VOLUME
	DataObjectType string `json:"data_object_type,omitempty"`
	// Name of the shared object
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenDetail struct {
	BearerToken string `json:"bearerToken,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol. Original
	// message: RetrieveToken.Response in
	// managed-catalog/api/messages/recipient.proto
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
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

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateExchangeFilterRequest struct {
	Filter ExchangeFilter `json:"filter"`

	Id string `json:"-" url:"-"`
}

type UpdateExchangeFilterResponse struct {
	Filter *ExchangeFilter `json:"filter,omitempty"`
}

type UpdateExchangeRequest struct {
	Exchange Exchange `json:"exchange"`

	Id string `json:"-" url:"-"`
}

type UpdateExchangeResponse struct {
	Exchange *Exchange `json:"exchange,omitempty"`
}

type UpdateInstallationRequest struct {
	Installation InstallationDetail `json:"installation"`

	InstallationId string `json:"-" url:"-"`

	ListingId string `json:"-" url:"-"`

	RotateToken bool `json:"rotate_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateInstallationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInstallationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateInstallationResponse struct {
	Installation *InstallationDetail `json:"installation,omitempty"`
}

type UpdateListingRequest struct {
	Id string `json:"-" url:"-"`

	Listing Listing `json:"listing"`
}

type UpdateListingResponse struct {
	Listing *Listing `json:"listing,omitempty"`
}

type UpdatePersonalizationRequestRequest struct {
	ListingId string `json:"-" url:"-"`

	Reason string `json:"reason,omitempty"`

	RequestId string `json:"-" url:"-"`

	Share *ShareInfo `json:"share,omitempty"`

	Status PersonalizationRequestStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdatePersonalizationRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePersonalizationRequestRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePersonalizationRequestResponse struct {
	Request *PersonalizationRequest `json:"request,omitempty"`
}

type UpdateProviderAnalyticsDashboardRequest struct {
	// id is immutable property and can't be updated.
	Id string `json:"-" url:"-"`
	// this is the version of the dashboard template we want to update our user
	// to current expectation is that it should be equal to latest version of
	// the dashboard template
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateProviderAnalyticsDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProviderAnalyticsDashboardResponse struct {
	// this is newly created Lakeview dashboard for the user
	DashboardId string `json:"dashboard_id"`
	// id & version should be the same as the request
	Id string `json:"id"`

	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateProviderAnalyticsDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProviderAnalyticsDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProviderRequest struct {
	Id string `json:"-" url:"-"`

	Provider ProviderInfo `json:"provider"`
}

type UpdateProviderResponse struct {
	Provider *ProviderInfo `json:"provider,omitempty"`
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

// Type always returns Visibility to satisfy [pflag.Value] interface
func (f *Visibility) Type() string {
	return "Visibility"
}
