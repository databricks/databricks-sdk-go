// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplacepb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddExchangeForListingRequestPb struct {
	ExchangeId string `json:"exchange_id"`
	ListingId  string `json:"listing_id"`
}

type AddExchangeForListingResponsePb struct {
	ExchangeForListing *ExchangeListingPb `json:"exchange_for_listing,omitempty"`
}

type AssetTypePb string

const AssetTypeAssetTypeApp AssetTypePb = `ASSET_TYPE_APP`

const AssetTypeAssetTypeDataTable AssetTypePb = `ASSET_TYPE_DATA_TABLE`

const AssetTypeAssetTypeGitRepo AssetTypePb = `ASSET_TYPE_GIT_REPO`

const AssetTypeAssetTypeMedia AssetTypePb = `ASSET_TYPE_MEDIA`

const AssetTypeAssetTypeModel AssetTypePb = `ASSET_TYPE_MODEL`

const AssetTypeAssetTypeNotebook AssetTypePb = `ASSET_TYPE_NOTEBOOK`

const AssetTypeAssetTypePartnerIntegration AssetTypePb = `ASSET_TYPE_PARTNER_INTEGRATION`

type BatchGetListingsRequestPb struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

type BatchGetListingsResponsePb struct {
	Listings []ListingPb `json:"listings,omitempty"`
}

type BatchGetProvidersRequestPb struct {
	Ids []string `json:"-" url:"ids,omitempty"`
}

type BatchGetProvidersResponsePb struct {
	Providers []ProviderInfoPb `json:"providers,omitempty"`
}

type CategoryPb string

const CategoryAdvertisingAndMarketing CategoryPb = `ADVERTISING_AND_MARKETING`

const CategoryClimateAndEnvironment CategoryPb = `CLIMATE_AND_ENVIRONMENT`

const CategoryCommerce CategoryPb = `COMMERCE`

const CategoryDemographics CategoryPb = `DEMOGRAPHICS`

const CategoryEconomics CategoryPb = `ECONOMICS`

const CategoryEducation CategoryPb = `EDUCATION`

const CategoryEnergy CategoryPb = `ENERGY`

const CategoryFinancial CategoryPb = `FINANCIAL`

const CategoryGaming CategoryPb = `GAMING`

const CategoryGeospatial CategoryPb = `GEOSPATIAL`

const CategoryHealth CategoryPb = `HEALTH`

const CategoryLookupTables CategoryPb = `LOOKUP_TABLES`

const CategoryManufacturing CategoryPb = `MANUFACTURING`

const CategoryMedia CategoryPb = `MEDIA`

const CategoryOther CategoryPb = `OTHER`

const CategoryPublicSector CategoryPb = `PUBLIC_SECTOR`

const CategoryRetail CategoryPb = `RETAIL`

const CategoryScienceAndResearch CategoryPb = `SCIENCE_AND_RESEARCH`

const CategorySecurity CategoryPb = `SECURITY`

const CategorySports CategoryPb = `SPORTS`

const CategoryTransportationAndLogistics CategoryPb = `TRANSPORTATION_AND_LOGISTICS`

const CategoryTravelAndTourism CategoryPb = `TRAVEL_AND_TOURISM`

type ConsumerTermsPb struct {
	Version string `json:"version"`
}

type ContactInfoPb struct {
	Company   string `json:"company,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ContactInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ContactInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CostPb string

const CostFree CostPb = `FREE`

const CostPaid CostPb = `PAID`

type CreateExchangeFilterRequestPb struct {
	Filter ExchangeFilterPb `json:"filter"`
}

type CreateExchangeFilterResponsePb struct {
	FilterId string `json:"filter_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateExchangeFilterResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateExchangeFilterResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateExchangeRequestPb struct {
	Exchange ExchangePb `json:"exchange"`
}

type CreateExchangeResponsePb struct {
	ExchangeId string `json:"exchange_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateExchangeResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateExchangeResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateFileRequestPb struct {
	DisplayName         string                `json:"display_name,omitempty"`
	FileParent          FileParentPb          `json:"file_parent"`
	MarketplaceFileType MarketplaceFileTypePb `json:"marketplace_file_type"`
	MimeType            string                `json:"mime_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateFileRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateFileRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateFileResponsePb struct {
	FileInfo  *FileInfoPb `json:"file_info,omitempty"`
	UploadUrl string      `json:"upload_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateFileResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateFileResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateInstallationRequestPb struct {
	AcceptedConsumerTerms *ConsumerTermsPb            `json:"accepted_consumer_terms,omitempty"`
	CatalogName           string                      `json:"catalog_name,omitempty"`
	ListingId             string                      `json:"-" url:"-"`
	RecipientType         DeltaSharingRecipientTypePb `json:"recipient_type,omitempty"`
	RepoDetail            *RepoInstallationPb         `json:"repo_detail,omitempty"`
	ShareName             string                      `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateInstallationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateInstallationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateListingRequestPb struct {
	Listing ListingPb `json:"listing"`
}

type CreateListingResponsePb struct {
	ListingId string `json:"listing_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateListingResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateListingResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePersonalizationRequestPb struct {
	AcceptedConsumerTerms ConsumerTermsPb             `json:"accepted_consumer_terms"`
	Comment               string                      `json:"comment,omitempty"`
	Company               string                      `json:"company,omitempty"`
	FirstName             string                      `json:"first_name,omitempty"`
	IntendedUse           string                      `json:"intended_use"`
	IsFromLighthouse      bool                        `json:"is_from_lighthouse,omitempty"`
	LastName              string                      `json:"last_name,omitempty"`
	ListingId             string                      `json:"-" url:"-"`
	RecipientType         DeltaSharingRecipientTypePb `json:"recipient_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePersonalizationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePersonalizationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePersonalizationRequestResponsePb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePersonalizationRequestResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePersonalizationRequestResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateProviderAnalyticsDashboardRequestPb struct {
}

type CreateProviderRequestPb struct {
	Provider ProviderInfoPb `json:"provider"`
}

type CreateProviderResponsePb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateProviderResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateProviderResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataRefreshPb string

const DataRefreshDaily DataRefreshPb = `DAILY`

const DataRefreshHourly DataRefreshPb = `HOURLY`

const DataRefreshMinute DataRefreshPb = `MINUTE`

const DataRefreshMonthly DataRefreshPb = `MONTHLY`

const DataRefreshNone DataRefreshPb = `NONE`

const DataRefreshQuarterly DataRefreshPb = `QUARTERLY`

const DataRefreshSecond DataRefreshPb = `SECOND`

const DataRefreshWeekly DataRefreshPb = `WEEKLY`

const DataRefreshYearly DataRefreshPb = `YEARLY`

type DataRefreshInfoPb struct {
	Interval int64         `json:"interval"`
	Unit     DataRefreshPb `json:"unit"`
}

type DeleteExchangeFilterRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteExchangeFilterResponsePb struct {
}

type DeleteExchangeRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteExchangeResponsePb struct {
}

type DeleteFileRequestPb struct {
	FileId string `json:"-" url:"-"`
}

type DeleteFileResponsePb struct {
}

type DeleteInstallationRequestPb struct {
	InstallationId string `json:"-" url:"-"`
	ListingId      string `json:"-" url:"-"`
}

type DeleteInstallationResponsePb struct {
}

type DeleteListingRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteListingResponsePb struct {
}

type DeleteProviderRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteProviderResponsePb struct {
}

type DeltaSharingRecipientTypePb string

const DeltaSharingRecipientTypeDeltaSharingRecipientTypeDatabricks DeltaSharingRecipientTypePb = `DELTA_SHARING_RECIPIENT_TYPE_DATABRICKS`

const DeltaSharingRecipientTypeDeltaSharingRecipientTypeOpen DeltaSharingRecipientTypePb = `DELTA_SHARING_RECIPIENT_TYPE_OPEN`

type ExchangePb struct {
	Comment        string              `json:"comment,omitempty"`
	CreatedAt      int64               `json:"created_at,omitempty"`
	CreatedBy      string              `json:"created_by,omitempty"`
	Filters        []ExchangeFilterPb  `json:"filters,omitempty"`
	Id             string              `json:"id,omitempty"`
	LinkedListings []ExchangeListingPb `json:"linked_listings,omitempty"`
	Name           string              `json:"name"`
	UpdatedAt      int64               `json:"updated_at,omitempty"`
	UpdatedBy      string              `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExchangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExchangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExchangeFilterPb struct {
	CreatedAt   int64                `json:"created_at,omitempty"`
	CreatedBy   string               `json:"created_by,omitempty"`
	ExchangeId  string               `json:"exchange_id"`
	FilterType  ExchangeFilterTypePb `json:"filter_type"`
	FilterValue string               `json:"filter_value"`
	Id          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	UpdatedAt   int64                `json:"updated_at,omitempty"`
	UpdatedBy   string               `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExchangeFilterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExchangeFilterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExchangeFilterTypePb string

const ExchangeFilterTypeGlobalMetastoreId ExchangeFilterTypePb = `GLOBAL_METASTORE_ID`

type ExchangeListingPb struct {
	CreatedAt    int64  `json:"created_at,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	ExchangeId   string `json:"exchange_id,omitempty"`
	ExchangeName string `json:"exchange_name,omitempty"`
	Id           string `json:"id,omitempty"`
	ListingId    string `json:"listing_id,omitempty"`
	ListingName  string `json:"listing_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExchangeListingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExchangeListingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileInfoPb struct {
	CreatedAt           int64                 `json:"created_at,omitempty"`
	DisplayName         string                `json:"display_name,omitempty"`
	DownloadLink        string                `json:"download_link,omitempty"`
	FileParent          *FileParentPb         `json:"file_parent,omitempty"`
	Id                  string                `json:"id,omitempty"`
	MarketplaceFileType MarketplaceFileTypePb `json:"marketplace_file_type,omitempty"`
	MimeType            string                `json:"mime_type,omitempty"`
	Status              FileStatusPb          `json:"status,omitempty"`
	StatusMessage       string                `json:"status_message,omitempty"`
	UpdatedAt           int64                 `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileParentPb struct {
	FileParentType FileParentTypePb `json:"file_parent_type,omitempty" url:"file_parent_type,omitempty"`
	ParentId       string           `json:"parent_id,omitempty" url:"parent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileParentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileParentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileParentTypePb string

const FileParentTypeListing FileParentTypePb = `LISTING`

const FileParentTypeListingResource FileParentTypePb = `LISTING_RESOURCE`

const FileParentTypeProvider FileParentTypePb = `PROVIDER`

type FileStatusPb string

const FileStatusFileStatusPublished FileStatusPb = `FILE_STATUS_PUBLISHED`

const FileStatusFileStatusSanitizationFailed FileStatusPb = `FILE_STATUS_SANITIZATION_FAILED`

const FileStatusFileStatusSanitizing FileStatusPb = `FILE_STATUS_SANITIZING`

const FileStatusFileStatusStaging FileStatusPb = `FILE_STATUS_STAGING`

type FulfillmentTypePb string

const FulfillmentTypeInstall FulfillmentTypePb = `INSTALL`

const FulfillmentTypeRequestAccess FulfillmentTypePb = `REQUEST_ACCESS`

type GetExchangeRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetExchangeResponsePb struct {
	Exchange *ExchangePb `json:"exchange,omitempty"`
}

type GetFileRequestPb struct {
	FileId string `json:"-" url:"-"`
}

type GetFileResponsePb struct {
	FileInfo *FileInfoPb `json:"file_info,omitempty"`
}

type GetLatestVersionProviderAnalyticsDashboardRequestPb struct {
}

type GetLatestVersionProviderAnalyticsDashboardResponsePb struct {
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetLatestVersionProviderAnalyticsDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetLatestVersionProviderAnalyticsDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetListingContentMetadataRequestPb struct {
	ListingId string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetListingContentMetadataRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetListingContentMetadataRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetListingContentMetadataResponsePb struct {
	NextPageToken     string               `json:"next_page_token,omitempty"`
	SharedDataObjects []SharedDataObjectPb `json:"shared_data_objects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetListingContentMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetListingContentMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetListingRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetListingResponsePb struct {
	Listing *ListingPb `json:"listing,omitempty"`
}

type GetListingsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetListingsResponsePb struct {
	Listings      []ListingPb `json:"listings,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetListingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetListingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPersonalizationRequestRequestPb struct {
	ListingId string `json:"-" url:"-"`
}

type GetPersonalizationRequestResponsePb struct {
	PersonalizationRequests []PersonalizationRequestPb `json:"personalization_requests,omitempty"`
}

type GetProviderRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetProviderResponsePb struct {
	Provider *ProviderInfoPb `json:"provider,omitempty"`
}

type InstallationPb struct {
	Installation *InstallationDetailPb `json:"installation,omitempty"`
}

type InstallationDetailPb struct {
	CatalogName   string                      `json:"catalog_name,omitempty"`
	ErrorMessage  string                      `json:"error_message,omitempty"`
	Id            string                      `json:"id,omitempty"`
	InstalledOn   int64                       `json:"installed_on,omitempty"`
	ListingId     string                      `json:"listing_id,omitempty"`
	ListingName   string                      `json:"listing_name,omitempty"`
	RecipientType DeltaSharingRecipientTypePb `json:"recipient_type,omitempty"`
	RepoName      string                      `json:"repo_name,omitempty"`
	RepoPath      string                      `json:"repo_path,omitempty"`
	ShareName     string                      `json:"share_name,omitempty"`
	Status        InstallationStatusPb        `json:"status,omitempty"`
	TokenDetail   *TokenDetailPb              `json:"token_detail,omitempty"`
	Tokens        []TokenInfoPb               `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstallationDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstallationDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstallationStatusPb string

const InstallationStatusFailed InstallationStatusPb = `FAILED`

const InstallationStatusInstalled InstallationStatusPb = `INSTALLED`

type ListAllInstallationsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAllInstallationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAllInstallationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAllInstallationsResponsePb struct {
	Installations []InstallationDetailPb `json:"installations,omitempty"`
	NextPageToken string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAllInstallationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAllInstallationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAllPersonalizationRequestsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAllPersonalizationRequestsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAllPersonalizationRequestsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAllPersonalizationRequestsResponsePb struct {
	NextPageToken           string                     `json:"next_page_token,omitempty"`
	PersonalizationRequests []PersonalizationRequestPb `json:"personalization_requests,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAllPersonalizationRequestsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAllPersonalizationRequestsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExchangeFiltersRequestPb struct {
	ExchangeId string `json:"-" url:"exchange_id"`
	PageSize   int    `json:"-" url:"page_size,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExchangeFiltersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExchangeFiltersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExchangeFiltersResponsePb struct {
	Filters       []ExchangeFilterPb `json:"filters,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExchangeFiltersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExchangeFiltersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExchangesForListingRequestPb struct {
	ListingId string `json:"-" url:"listing_id"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExchangesForListingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExchangesForListingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExchangesForListingResponsePb struct {
	ExchangeListing []ExchangeListingPb `json:"exchange_listing,omitempty"`
	NextPageToken   string              `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExchangesForListingResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExchangesForListingResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExchangesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExchangesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExchangesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExchangesResponsePb struct {
	Exchanges     []ExchangePb `json:"exchanges,omitempty"`
	NextPageToken string       `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListExchangesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListExchangesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFilesRequestPb struct {
	FileParent FileParentPb `json:"-" url:"file_parent"`
	PageSize   int          `json:"-" url:"page_size,omitempty"`
	PageToken  string       `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFilesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFilesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFilesResponsePb struct {
	FileInfos     []FileInfoPb `json:"file_infos,omitempty"`
	NextPageToken string       `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFilesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFilesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFulfillmentsRequestPb struct {
	ListingId string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFulfillmentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFulfillmentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFulfillmentsResponsePb struct {
	Fulfillments  []ListingFulfillmentPb `json:"fulfillments,omitempty"`
	NextPageToken string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFulfillmentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFulfillmentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListInstallationsRequestPb struct {
	ListingId string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListInstallationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListInstallationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListInstallationsResponsePb struct {
	Installations []InstallationDetailPb `json:"installations,omitempty"`
	NextPageToken string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListInstallationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListInstallationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListListingsForExchangeRequestPb struct {
	ExchangeId string `json:"-" url:"exchange_id"`
	PageSize   int    `json:"-" url:"page_size,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListListingsForExchangeRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListListingsForExchangeRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListListingsForExchangeResponsePb struct {
	ExchangeListings []ExchangeListingPb `json:"exchange_listings,omitempty"`
	NextPageToken    string              `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListListingsForExchangeResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListListingsForExchangeResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListListingsRequestPb struct {
	Assets            []AssetTypePb  `json:"-" url:"assets,omitempty"`
	Categories        []CategoryPb   `json:"-" url:"categories,omitempty"`
	IsFree            bool           `json:"-" url:"is_free,omitempty"`
	IsPrivateExchange bool           `json:"-" url:"is_private_exchange,omitempty"`
	IsStaffPick       bool           `json:"-" url:"is_staff_pick,omitempty"`
	PageSize          int            `json:"-" url:"page_size,omitempty"`
	PageToken         string         `json:"-" url:"page_token,omitempty"`
	ProviderIds       []string       `json:"-" url:"provider_ids,omitempty"`
	Tags              []ListingTagPb `json:"-" url:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListListingsResponsePb struct {
	Listings      []ListingPb `json:"listings,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListListingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListListingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProviderAnalyticsDashboardRequestPb struct {
}

type ListProviderAnalyticsDashboardResponsePb struct {
	DashboardId string `json:"dashboard_id"`
	Id          string `json:"id"`
	Version     int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProviderAnalyticsDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProviderAnalyticsDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProvidersRequestPb struct {
	IsFeatured bool   `json:"-" url:"is_featured,omitempty"`
	PageSize   int    `json:"-" url:"page_size,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProvidersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProvidersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProvidersResponsePb struct {
	NextPageToken string           `json:"next_page_token,omitempty"`
	Providers     []ProviderInfoPb `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProvidersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProvidersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListingPb struct {
	Detail  *ListingDetailPb `json:"detail,omitempty"`
	Id      string           `json:"id,omitempty"`
	Summary ListingSummaryPb `json:"summary"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListingDetailPb struct {
	Assets                    []AssetTypePb      `json:"assets,omitempty"`
	CollectionDateEnd         int64              `json:"collection_date_end,omitempty"`
	CollectionDateStart       int64              `json:"collection_date_start,omitempty"`
	CollectionGranularity     *DataRefreshInfoPb `json:"collection_granularity,omitempty"`
	Cost                      CostPb             `json:"cost,omitempty"`
	DataSource                string             `json:"data_source,omitempty"`
	Description               string             `json:"description,omitempty"`
	DocumentationLink         string             `json:"documentation_link,omitempty"`
	EmbeddedNotebookFileInfos []FileInfoPb       `json:"embedded_notebook_file_infos,omitempty"`
	FileIds                   []string           `json:"file_ids,omitempty"`
	GeographicalCoverage      string             `json:"geographical_coverage,omitempty"`
	License                   string             `json:"license,omitempty"`
	PricingModel              string             `json:"pricing_model,omitempty"`
	PrivacyPolicyLink         string             `json:"privacy_policy_link,omitempty"`
	Size                      float64            `json:"size,omitempty"`
	SupportLink               string             `json:"support_link,omitempty"`
	Tags                      []ListingTagPb     `json:"tags,omitempty"`
	TermsOfService            string             `json:"terms_of_service,omitempty"`
	UpdateFrequency           *DataRefreshInfoPb `json:"update_frequency,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListingDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListingDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListingFulfillmentPb struct {
	FulfillmentType FulfillmentTypePb           `json:"fulfillment_type,omitempty"`
	ListingId       string                      `json:"listing_id"`
	RecipientType   DeltaSharingRecipientTypePb `json:"recipient_type,omitempty"`
	RepoInfo        *RepoInfoPb                 `json:"repo_info,omitempty"`
	ShareInfo       *ShareInfoPb                `json:"share_info,omitempty"`
}

type ListingSettingPb struct {
	Visibility VisibilityPb `json:"visibility,omitempty"`
}

type ListingShareTypePb string

const ListingShareTypeFull ListingShareTypePb = `FULL`

const ListingShareTypeSample ListingShareTypePb = `SAMPLE`

type ListingStatusPb string

const ListingStatusDraft ListingStatusPb = `DRAFT`

const ListingStatusPending ListingStatusPb = `PENDING`

const ListingStatusPublished ListingStatusPb = `PUBLISHED`

const ListingStatusSuspended ListingStatusPb = `SUSPENDED`

type ListingSummaryPb struct {
	Categories     []CategoryPb      `json:"categories,omitempty"`
	CreatedAt      int64             `json:"created_at,omitempty"`
	CreatedBy      string            `json:"created_by,omitempty"`
	CreatedById    int64             `json:"created_by_id,omitempty"`
	ExchangeIds    []string          `json:"exchange_ids,omitempty"`
	GitRepo        *RepoInfoPb       `json:"git_repo,omitempty"`
	ListingType    ListingTypePb     `json:"listingType"`
	Name           string            `json:"name"`
	ProviderId     string            `json:"provider_id,omitempty"`
	ProviderRegion *RegionInfoPb     `json:"provider_region,omitempty"`
	PublishedAt    int64             `json:"published_at,omitempty"`
	PublishedBy    string            `json:"published_by,omitempty"`
	Setting        *ListingSettingPb `json:"setting,omitempty"`
	Share          *ShareInfoPb      `json:"share,omitempty"`
	Status         ListingStatusPb   `json:"status,omitempty"`
	Subtitle       string            `json:"subtitle,omitempty"`
	UpdatedAt      int64             `json:"updated_at,omitempty"`
	UpdatedBy      string            `json:"updated_by,omitempty"`
	UpdatedById    int64             `json:"updated_by_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListingSummaryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListingSummaryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListingTagPb struct {
	TagName   ListingTagTypePb `json:"tag_name,omitempty" url:"tag_name,omitempty"`
	TagValues []string         `json:"tag_values,omitempty" url:"tag_values,omitempty"`
}

type ListingTagTypePb string

const ListingTagTypeListingTagTypeLanguage ListingTagTypePb = `LISTING_TAG_TYPE_LANGUAGE`

const ListingTagTypeListingTagTypeTask ListingTagTypePb = `LISTING_TAG_TYPE_TASK`

type ListingTypePb string

const ListingTypePersonalized ListingTypePb = `PERSONALIZED`

const ListingTypeStandard ListingTypePb = `STANDARD`

type MarketplaceFileTypePb string

const MarketplaceFileTypeApp MarketplaceFileTypePb = `APP`

const MarketplaceFileTypeEmbeddedNotebook MarketplaceFileTypePb = `EMBEDDED_NOTEBOOK`

const MarketplaceFileTypeProviderIcon MarketplaceFileTypePb = `PROVIDER_ICON`

type PersonalizationRequestPb struct {
	Comment          string                         `json:"comment,omitempty"`
	ConsumerRegion   RegionInfoPb                   `json:"consumer_region"`
	ContactInfo      *ContactInfoPb                 `json:"contact_info,omitempty"`
	CreatedAt        int64                          `json:"created_at,omitempty"`
	Id               string                         `json:"id,omitempty"`
	IntendedUse      string                         `json:"intended_use,omitempty"`
	IsFromLighthouse bool                           `json:"is_from_lighthouse,omitempty"`
	ListingId        string                         `json:"listing_id,omitempty"`
	ListingName      string                         `json:"listing_name,omitempty"`
	MetastoreId      string                         `json:"metastore_id,omitempty"`
	ProviderId       string                         `json:"provider_id,omitempty"`
	RecipientType    DeltaSharingRecipientTypePb    `json:"recipient_type,omitempty"`
	Share            *ShareInfoPb                   `json:"share,omitempty"`
	Status           PersonalizationRequestStatusPb `json:"status,omitempty"`
	StatusMessage    string                         `json:"status_message,omitempty"`
	UpdatedAt        int64                          `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PersonalizationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PersonalizationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PersonalizationRequestStatusPb string

const PersonalizationRequestStatusDenied PersonalizationRequestStatusPb = `DENIED`

const PersonalizationRequestStatusFulfilled PersonalizationRequestStatusPb = `FULFILLED`

const PersonalizationRequestStatusNew PersonalizationRequestStatusPb = `NEW`

const PersonalizationRequestStatusRequestPending PersonalizationRequestStatusPb = `REQUEST_PENDING`

type ProviderAnalyticsDashboardPb struct {
	Id string `json:"id"`
}

type ProviderInfoPb struct {
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

func (st *ProviderInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ProviderInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegionInfoPb struct {
	Cloud  string `json:"cloud,omitempty"`
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RemoveExchangeForListingRequestPb struct {
	Id string `json:"-" url:"-"`
}

type RemoveExchangeForListingResponsePb struct {
}

type RepoInfoPb struct {
	GitRepoUrl string `json:"git_repo_url"`
}

type RepoInstallationPb struct {
	RepoName string `json:"repo_name"`
	RepoPath string `json:"repo_path"`
}

type SearchListingsRequestPb struct {
	Assets            []AssetTypePb `json:"-" url:"assets,omitempty"`
	Categories        []CategoryPb  `json:"-" url:"categories,omitempty"`
	IsFree            bool          `json:"-" url:"is_free,omitempty"`
	IsPrivateExchange bool          `json:"-" url:"is_private_exchange,omitempty"`
	PageSize          int           `json:"-" url:"page_size,omitempty"`
	PageToken         string        `json:"-" url:"page_token,omitempty"`
	ProviderIds       []string      `json:"-" url:"provider_ids,omitempty"`
	Query             string        `json:"-" url:"query"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchListingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchListingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchListingsResponsePb struct {
	Listings      []ListingPb `json:"listings,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SearchListingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SearchListingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ShareInfoPb struct {
	Name string             `json:"name"`
	Type ListingShareTypePb `json:"type"`
}

type SharedDataObjectPb struct {
	DataObjectType string `json:"data_object_type,omitempty"`
	Name           string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SharedDataObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SharedDataObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenDetailPb struct {
	BearerToken             string `json:"bearerToken,omitempty"`
	Endpoint                string `json:"endpoint,omitempty"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	ShareCredentialsVersion int    `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenInfoPb struct {
	ActivationUrl  string `json:"activation_url,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
	ExpirationTime int64  `json:"expiration_time,omitempty"`
	Id             string `json:"id,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	UpdatedBy      string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateExchangeFilterRequestPb struct {
	Filter ExchangeFilterPb `json:"filter"`
	Id     string           `json:"-" url:"-"`
}

type UpdateExchangeFilterResponsePb struct {
	Filter *ExchangeFilterPb `json:"filter,omitempty"`
}

type UpdateExchangeRequestPb struct {
	Exchange ExchangePb `json:"exchange"`
	Id       string     `json:"-" url:"-"`
}

type UpdateExchangeResponsePb struct {
	Exchange *ExchangePb `json:"exchange,omitempty"`
}

type UpdateInstallationRequestPb struct {
	Installation   InstallationDetailPb `json:"installation"`
	InstallationId string               `json:"-" url:"-"`
	ListingId      string               `json:"-" url:"-"`
	RotateToken    bool                 `json:"rotate_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateInstallationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateInstallationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateInstallationResponsePb struct {
	Installation *InstallationDetailPb `json:"installation,omitempty"`
}

type UpdateListingRequestPb struct {
	Id      string    `json:"-" url:"-"`
	Listing ListingPb `json:"listing"`
}

type UpdateListingResponsePb struct {
	Listing *ListingPb `json:"listing,omitempty"`
}

type UpdatePersonalizationRequestRequestPb struct {
	ListingId string                         `json:"-" url:"-"`
	Reason    string                         `json:"reason,omitempty"`
	RequestId string                         `json:"-" url:"-"`
	Share     *ShareInfoPb                   `json:"share,omitempty"`
	Status    PersonalizationRequestStatusPb `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdatePersonalizationRequestRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdatePersonalizationRequestRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdatePersonalizationRequestResponsePb struct {
	Request *PersonalizationRequestPb `json:"request,omitempty"`
}

type UpdateProviderAnalyticsDashboardRequestPb struct {
	Id      string `json:"-" url:"-"`
	Version int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateProviderAnalyticsDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateProviderAnalyticsDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateProviderAnalyticsDashboardResponsePb struct {
	DashboardId string `json:"dashboard_id"`
	Id          string `json:"id"`
	Version     int64  `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateProviderAnalyticsDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateProviderAnalyticsDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateProviderRequestPb struct {
	Id       string         `json:"-" url:"-"`
	Provider ProviderInfoPb `json:"provider"`
}

type UpdateProviderResponsePb struct {
	Provider *ProviderInfoPb `json:"provider,omitempty"`
}

type VisibilityPb string

const VisibilityPrivate VisibilityPb = `PRIVATE`

const VisibilityPublic VisibilityPb = `PUBLIC`
