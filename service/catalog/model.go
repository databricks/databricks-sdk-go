// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import "fmt"

// all definitions in this file are in alphabetical order

type AwsIamRole struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `json:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`
}

type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId string `json:"application_id"`
	// The client secret generated for the above app ID in AAD.
	ClientSecret string `json:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId string `json:"directory_id"`
}

type CatalogInfo struct {
	// The type of the catalog.
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of catalog creator.
	CreatedBy string `json:"created_by,omitempty"`

	EffectiveAutoMaintenanceFlag *EffectiveAutoMaintenanceFlag `json:"effective_auto_maintenance_flag,omitempty"`
	// Whether auto maintenance should be enabled for this object and objects
	// under it.
	EnableAutoMaintenance EnableAutoMaintenance `json:"enable_auto_maintenance,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified catalog.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// The type of the catalog.
type CatalogType string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (ct *CatalogType) String() string {
	return string(*ct)
}

// Set raw string value and validate it against allowed values
func (ct *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `MANAGED_CATALOG`, `SYSTEM_CATALOG`:
		*ct = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "MANAGED_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (ct *CatalogType) Type() string {
	return "CatalogType"
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`

	Mask *ColumnMask `json:"mask,omitempty"`
	// Name of Column.
	Name string `json:"name,omitempty"`
	// Whether field may be Null (default: true).
	Nullable bool `json:"nullable,omitempty"`
	// Partition index for column.
	PartitionIndex int `json:"partition_index,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position,omitempty"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type specification, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type specification as SQL/catalogString text.
	TypeText string `json:"type_text,omitempty"`
}

type ColumnMask struct {
	// The full name of the column maks SQL UDF.
	FunctionName string `json:"function_name,omitempty"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []string `json:"using_column_names,omitempty"`
}

// Name of type (INT, STRUCT, MAP, etc.).
type ColumnTypeName string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTableType ColumnTypeName = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeName = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

// String representation for [fmt.Print]
func (ctn *ColumnTypeName) String() string {
	return string(*ctn)
}

// Set raw string value and validate it against allowed values
func (ctn *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`:
		*ctn = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE"`, v)
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (ctn *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `json:"storage_root,omitempty"`
}

type CreateExternalLocation struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name"`
	// Name of the external location.
	Name string `json:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	Url string `json:"url"`
}

type CreateFunction struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Scalar function return data type.
	DataType ColumnTypeName `json:"data_type"`
	// External function language.
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	FullDataType string `json:"full_data_type"`
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	InputParams []FunctionParameterInfo `json:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic bool `json:"is_deterministic"`
	// Function null call.
	IsNullCall bool `json:"is_null_call"`
	// Name of function, relative to parent schema.
	Name string `json:"name"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle CreateFunctionParameterStyle `json:"parameter_style"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Table function return parameters.
	ReturnParams []FunctionParameterInfo `json:"return_params"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody CreateFunctionRoutineBody `json:"routine_body"`
	// Function body.
	RoutineDefinition string `json:"routine_definition"`
	// Function dependencies.
	RoutineDependencies []Dependency `json:"routine_dependencies"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name"`
	// Function security type.
	SecurityType CreateFunctionSecurityType `json:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `json:"specific_name"`
	// Function SQL data access.
	SqlDataAccess CreateFunctionSqlDataAccess `json:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
}

// Function parameter style. **S** is the value for SQL.
type CreateFunctionParameterStyle string

const CreateFunctionParameterStyleS CreateFunctionParameterStyle = `S`

// String representation for [fmt.Print]
func (cfps *CreateFunctionParameterStyle) String() string {
	return string(*cfps)
}

// Set raw string value and validate it against allowed values
func (cfps *CreateFunctionParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*cfps = CreateFunctionParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns CreateFunctionParameterStyle to satisfy [pflag.Value] interface
func (cfps *CreateFunctionParameterStyle) Type() string {
	return "CreateFunctionParameterStyle"
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type CreateFunctionRoutineBody string

const CreateFunctionRoutineBodyExternal CreateFunctionRoutineBody = `EXTERNAL`

const CreateFunctionRoutineBodySql CreateFunctionRoutineBody = `SQL`

// String representation for [fmt.Print]
func (cfrb *CreateFunctionRoutineBody) String() string {
	return string(*cfrb)
}

// Set raw string value and validate it against allowed values
func (cfrb *CreateFunctionRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*cfrb = CreateFunctionRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns CreateFunctionRoutineBody to satisfy [pflag.Value] interface
func (cfrb *CreateFunctionRoutineBody) Type() string {
	return "CreateFunctionRoutineBody"
}

// Function security type.
type CreateFunctionSecurityType string

const CreateFunctionSecurityTypeDefiner CreateFunctionSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (cfst *CreateFunctionSecurityType) String() string {
	return string(*cfst)
}

// Set raw string value and validate it against allowed values
func (cfst *CreateFunctionSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*cfst = CreateFunctionSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns CreateFunctionSecurityType to satisfy [pflag.Value] interface
func (cfst *CreateFunctionSecurityType) Type() string {
	return "CreateFunctionSecurityType"
}

// Function SQL data access.
type CreateFunctionSqlDataAccess string

const CreateFunctionSqlDataAccessContainsSql CreateFunctionSqlDataAccess = `CONTAINS_SQL`

const CreateFunctionSqlDataAccessNoSql CreateFunctionSqlDataAccess = `NO_SQL`

const CreateFunctionSqlDataAccessReadsSqlData CreateFunctionSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (cfsda *CreateFunctionSqlDataAccess) String() string {
	return string(*cfsda)
}

// Set raw string value and validate it against allowed values
func (cfsda *CreateFunctionSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*cfsda = CreateFunctionSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns CreateFunctionSqlDataAccess to satisfy [pflag.Value] interface
func (cfsda *CreateFunctionSqlDataAccess) Type() string {
	return "CreateFunctionSqlDataAccess"
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	Name string `json:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`). If
	// this field is omitted, the region of the workspace receiving the request
	// will be used.
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root"`
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName string `json:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type CreateSchema struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint TableConstraint `json:"constraint"`
	// The full name of the table referenced by the constraint.
	FullNameArg string `json:"full_name_arg"`
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The name of the volume
	Name string `json:"name"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud
	StorageLocation string `json:"storage_location,omitempty"`

	VolumeType VolumeType `json:"volume_type"`
}

// Data source format
type DataSourceFormat string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatText DataSourceFormat = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormat = `UNITY_CATALOG`

// String representation for [fmt.Print]
func (dsf *DataSourceFormat) String() string {
	return string(*dsf)
}

// Set raw string value and validate it against allowed values
func (dsf *DataSourceFormat) Set(v string) error {
	switch v {
	case `AVRO`, `CSV`, `DELTA`, `DELTASHARING`, `JSON`, `ORC`, `PARQUET`, `TEXT`, `UNITY_CATALOG`:
		*dsf = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "CSV", "DELTA", "DELTASHARING", "JSON", "ORC", "PARQUET", "TEXT", "UNITY_CATALOG"`, v)
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (dsf *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest struct {
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`
}

// Delete a function
type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	Force bool `json:"-" url:"force,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`
}

// Delete a metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
}

// Delete a schema
type DeleteSchemaRequest struct {
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Delete a table constraint
type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint.
	//
	// If false, reject this operation if the current constraint has any child
	// constraints.
	Cascade bool `json:"-" url:"cascade"`
	// The name of the constraint to delete.
	ConstraintName string `json:"-" url:"constraint_name"`
	// Full name of the table referenced by the constraint.
	FullName string `json:"-" url:"-"`
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	FullNameArg string `json:"-" url:"-"`
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function *FunctionDependency `json:"function,omitempty"`
	// A table that is dependent on a SQL object.
	Table *TableDependency `json:"table,omitempty"`
}

type EffectiveAutoMaintenanceFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType EffectiveAutoMaintenanceFlagInheritedFromType `json:"inherited_from_type,omitempty"`
	// Whether auto maintenance should be enabled for this object and objects
	// under it.
	Value EnableAutoMaintenance `json:"value"`
}

// The type of the object from which the flag was inherited. If there was no
// inheritance, this field is left blank.
type EffectiveAutoMaintenanceFlagInheritedFromType string

const EffectiveAutoMaintenanceFlagInheritedFromTypeCatalog EffectiveAutoMaintenanceFlagInheritedFromType = `CATALOG`

const EffectiveAutoMaintenanceFlagInheritedFromTypeSchema EffectiveAutoMaintenanceFlagInheritedFromType = `SCHEMA`

// String representation for [fmt.Print]
func (eamfift *EffectiveAutoMaintenanceFlagInheritedFromType) String() string {
	return string(*eamfift)
}

// Set raw string value and validate it against allowed values
func (eamfift *EffectiveAutoMaintenanceFlagInheritedFromType) Set(v string) error {
	switch v {
	case `CATALOG`, `SCHEMA`:
		*eamfift = EffectiveAutoMaintenanceFlagInheritedFromType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "SCHEMA"`, v)
	}
}

// Type always returns EffectiveAutoMaintenanceFlagInheritedFromType to satisfy [pflag.Value] interface
func (eamfift *EffectiveAutoMaintenanceFlagInheritedFromType) Type() string {
	return "EffectiveAutoMaintenanceFlagInheritedFromType"
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments []EffectivePrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	//
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object that conveys this privilege via inheritance.
	//
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromType SecurableType `json:"inherited_from_type,omitempty"`
	// The privilege assigned to the principal.
	Privilege Privilege `json:"privilege,omitempty"`
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []EffectivePrivilege `json:"privileges,omitempty"`
}

// Whether auto maintenance should be enabled for this object and objects under
// it.
type EnableAutoMaintenance string

const EnableAutoMaintenanceDisable EnableAutoMaintenance = `DISABLE`

const EnableAutoMaintenanceEnable EnableAutoMaintenance = `ENABLE`

const EnableAutoMaintenanceInherit EnableAutoMaintenance = `INHERIT`

// String representation for [fmt.Print]
func (eam *EnableAutoMaintenance) String() string {
	return string(*eam)
}

// Set raw string value and validate it against allowed values
func (eam *EnableAutoMaintenance) Set(v string) error {
	switch v {
	case `DISABLE`, `ENABLE`, `INHERIT`:
		*eam = EnableAutoMaintenance(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLE", "ENABLE", "INHERIT"`, v)
	}
}

// Type always returns EnableAutoMaintenance to satisfy [pflag.Value] interface
func (eam *EnableAutoMaintenance) Type() string {
	return "EnableAutoMaintenance"
}

type ExternalLocationInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of external location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the location's storage credential.
	CredentialId string `json:"credential_id,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external location.
	Name string `json:"name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the external location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Path URL of the external location.
	Url string `json:"url,omitempty"`
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
	// Column names for this constraint.
	ParentColumns []string `json:"parent_columns"`
	// The full name of the parent constraint.
	ParentTable string `json:"parent_table"`
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName string `json:"function_full_name"`
}

type FunctionInfo struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of function creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Scalar function return data type.
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// External function language.
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	FullDataType string `json:"full_data_type,omitempty"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName string `json:"full_name,omitempty"`
	// Id of Function, relative to parent schema.
	FunctionId string `json:"function_id,omitempty"`
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	InputParams []FunctionParameterInfo `json:"input_params,omitempty"`
	// Whether the function is deterministic.
	IsDeterministic bool `json:"is_deterministic,omitempty"`
	// Function null call.
	IsNullCall bool `json:"is_null_call,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of function, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of function.
	Owner string `json:"owner,omitempty"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle FunctionInfoParameterStyle `json:"parameter_style,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Table function return parameters.
	ReturnParams []FunctionParameterInfo `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody FunctionInfoRoutineBody `json:"routine_body,omitempty"`
	// Function body.
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// Function dependencies.
	RoutineDependencies []Dependency `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// Function security type.
	SecurityType FunctionInfoSecurityType `json:"security_type,omitempty"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `json:"specific_name,omitempty"`
	// Function SQL data access.
	SqlDataAccess FunctionInfoSqlDataAccess `json:"sql_data_access,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified function.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// Function parameter style. **S** is the value for SQL.
type FunctionInfoParameterStyle string

const FunctionInfoParameterStyleS FunctionInfoParameterStyle = `S`

// String representation for [fmt.Print]
func (fips *FunctionInfoParameterStyle) String() string {
	return string(*fips)
}

// Set raw string value and validate it against allowed values
func (fips *FunctionInfoParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*fips = FunctionInfoParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns FunctionInfoParameterStyle to satisfy [pflag.Value] interface
func (fips *FunctionInfoParameterStyle) Type() string {
	return "FunctionInfoParameterStyle"
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type FunctionInfoRoutineBody string

const FunctionInfoRoutineBodyExternal FunctionInfoRoutineBody = `EXTERNAL`

const FunctionInfoRoutineBodySql FunctionInfoRoutineBody = `SQL`

// String representation for [fmt.Print]
func (firb *FunctionInfoRoutineBody) String() string {
	return string(*firb)
}

// Set raw string value and validate it against allowed values
func (firb *FunctionInfoRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*firb = FunctionInfoRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns FunctionInfoRoutineBody to satisfy [pflag.Value] interface
func (firb *FunctionInfoRoutineBody) Type() string {
	return "FunctionInfoRoutineBody"
}

// Function security type.
type FunctionInfoSecurityType string

const FunctionInfoSecurityTypeDefiner FunctionInfoSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (fist *FunctionInfoSecurityType) String() string {
	return string(*fist)
}

// Set raw string value and validate it against allowed values
func (fist *FunctionInfoSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*fist = FunctionInfoSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns FunctionInfoSecurityType to satisfy [pflag.Value] interface
func (fist *FunctionInfoSecurityType) Type() string {
	return "FunctionInfoSecurityType"
}

// Function SQL data access.
type FunctionInfoSqlDataAccess string

const FunctionInfoSqlDataAccessContainsSql FunctionInfoSqlDataAccess = `CONTAINS_SQL`

const FunctionInfoSqlDataAccessNoSql FunctionInfoSqlDataAccess = `NO_SQL`

const FunctionInfoSqlDataAccessReadsSqlData FunctionInfoSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (fisda *FunctionInfoSqlDataAccess) String() string {
	return string(*fisda)
}

// Set raw string value and validate it against allowed values
func (fisda *FunctionInfoSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*fisda = FunctionInfoSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns FunctionInfoSqlDataAccess to satisfy [pflag.Value] interface
func (fisda *FunctionInfoSqlDataAccess) Type() string {
	return "FunctionInfoSqlDataAccess"
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of parameter.
	Name string `json:"name"`
	// Default value of the parameter.
	ParameterDefault string `json:"parameter_default,omitempty"`
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// The type of function parameter.
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type spec, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName ColumnTypeName `json:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type spec, SQL/catalogString text.
	TypeText string `json:"type_text"`
}

// The mode of the function parameter.
type FunctionParameterMode string

const FunctionParameterModeIn FunctionParameterMode = `IN`

// String representation for [fmt.Print]
func (fpm *FunctionParameterMode) String() string {
	return string(*fpm)
}

// Set raw string value and validate it against allowed values
func (fpm *FunctionParameterMode) Set(v string) error {
	switch v {
	case `IN`:
		*fpm = FunctionParameterMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN"`, v)
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (fpm *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

// The type of function parameter.
type FunctionParameterType string

const FunctionParameterTypeColumn FunctionParameterType = `COLUMN`

const FunctionParameterTypeParam FunctionParameterType = `PARAM`

// String representation for [fmt.Print]
func (fpt *FunctionParameterType) String() string {
	return string(*fpt)
}

// Set raw string value and validate it against allowed values
func (fpt *FunctionParameterType) Set(v string) error {
	switch v {
	case `COLUMN`, `PARAM`:
		*fpt = FunctionParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN", "PARAM"`, v)
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (fpt *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

type GcpServiceAccountKey struct {
	// The email of the service account.
	Email string `json:"email"`
	// The service account's RSA private key.
	PrivateKey string `json:"private_key"`
	// The ID of the service account's private key.
	PrivateKeyId string `json:"private_key_id"`
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Get a catalog
type GetCatalogRequest struct {
	// The name of the catalog.
	Name string `json:"-" url:"-"`
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`
}

// Get an external location
type GetExternalLocationRequest struct {
	// Name of the external location.
	Name string `json:"-" url:"-"`
}

// Get a function
type GetFunctionRequest struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`
}

// Get permissions
type GetGrantRequest struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`
}

// Get a metastore
type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// The scope of Delta Sharing enabled for the metastore.
type GetMetastoreSummaryResponseDeltaSharingScope string

const GetMetastoreSummaryResponseDeltaSharingScopeInternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL`

const GetMetastoreSummaryResponseDeltaSharingScopeInternalAndExternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (gmsrdss *GetMetastoreSummaryResponseDeltaSharingScope) String() string {
	return string(*gmsrdss)
}

// Set raw string value and validate it against allowed values
func (gmsrdss *GetMetastoreSummaryResponseDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*gmsrdss = GetMetastoreSummaryResponseDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns GetMetastoreSummaryResponseDeltaSharingScope to satisfy [pflag.Value] interface
func (gmsrdss *GetMetastoreSummaryResponseDeltaSharingScope) Type() string {
	return "GetMetastoreSummaryResponseDeltaSharingScope"
}

// Get a schema
type GetSchemaRequest struct {
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
}

// List functions
type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Parent schema of functions.
	SchemaName string `json:"-" url:"schema_name"`
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	Schemas []FunctionInfo `json:"schemas,omitempty"`
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
}

// List schemas
type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName string `json:"-" url:"catalog_name"`
}

type ListSchemasResponse struct {
	// An array of schema information objects.
	Schemas []SchemaInfo `json:"schemas,omitempty"`
}

// List table summaries
type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Maximum number of tables to return (page length). Defaults to 10000.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string `json:"-" url:"schema_name_pattern,omitempty"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string `json:"-" url:"table_name_pattern,omitempty"`
}

type ListTableSummariesResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of table summaries.
	Tables []TableSummary `json:"tables,omitempty"`
}

// List tables
type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
	// Maximum number of tables to return (page length). If not set, all
	// accessible tables in the schema are returned. If set to:
	//
	// * greater than 0, page length is the minimum of this value and a server
	// configured value. * equal to 0, page length is set to a server configured
	// value. * lesser than 0, invalid parameter error.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of tables.
	SchemaName string `json:"-" url:"schema_name"`
}

type ListTablesResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of table information objects.
	Tables []TableInfo `json:"tables,omitempty"`
}

// List Volumes
type ListVolumesRequest struct {
	// The identifier of the catalog
	CatalogName string `json:"-" url:"catalog_name"`
	// The identifier of the schema
	SchemaName string `json:"-" url:"schema_name"`
}

type ListVolumesResponseContent struct {
	Volumes []VolumeInfo `json:"volumes,omitempty"`
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// The unique ID of the Databricks workspace.
	WorkspaceId string `json:"workspace_id"`
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope MetastoreInfoDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// The scope of Delta Sharing enabled for the metastore.
type MetastoreInfoDeltaSharingScope string

const MetastoreInfoDeltaSharingScopeInternal MetastoreInfoDeltaSharingScope = `INTERNAL`

const MetastoreInfoDeltaSharingScopeInternalAndExternal MetastoreInfoDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (midss *MetastoreInfoDeltaSharingScope) String() string {
	return string(*midss)
}

// Set raw string value and validate it against allowed values
func (midss *MetastoreInfoDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*midss = MetastoreInfoDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns MetastoreInfoDeltaSharingScope to satisfy [pflag.Value] interface
func (midss *MetastoreInfoDeltaSharingScope) Type() string {
	return "MetastoreInfoDeltaSharingScope"
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name string `json:"name"`
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add []Privilege `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []Privilege `json:"remove,omitempty"`
}

type PermissionsList struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
}

type Privilege string

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

// String representation for [fmt.Print]
func (p *Privilege) String() string {
	return string(*p)
}

// Set raw string value and validate it against allowed values
func (p *Privilege) Set(v string) error {
	switch v {
	case `ALL_PRIVILEGES`, `CREATE`, `CREATE_CATALOG`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `EXECUTE`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`:
		*p = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_PRIVILEGES", "CREATE", "CREATE_CATALOG", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "EXECUTE", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES"`, v)
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (p *Privilege) Type() string {
	return "Privilege"
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	Privileges []Privilege `json:"privileges,omitempty"`
}

// Get a Volume
type ReadVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	FullNameArg string `json:"-" url:"-"`
}

type SchemaInfo struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The type of the parent catalog.
	CatalogType string `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of schema creator.
	CreatedBy string `json:"created_by,omitempty"`

	EffectiveAutoMaintenanceFlag *EffectiveAutoMaintenanceFlag `json:"effective_auto_maintenance_flag,omitempty"`
	// Whether auto maintenance should be enabled for this object and objects
	// under it.
	EnableAutoMaintenance EnableAutoMaintenance `json:"enable_auto_maintenance,omitempty"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of schema.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Storage location for managed tables within schema.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified schema.
	UpdatedBy string `json:"updated_by,omitempty"`
}

// A map of key-value properties attached to the securable.
type SecurablePropertiesMap map[string]string

// The type of Unity Catalog securable
type SecurableType string

const SecurableTypeCatalog SecurableType = `CATALOG`

const SecurableTypeExternalLocation SecurableType = `EXTERNAL_LOCATION`

const SecurableTypeFunction SecurableType = `FUNCTION`

const SecurableTypeMetastore SecurableType = `METASTORE`

const SecurableTypeProvider SecurableType = `PROVIDER`

const SecurableTypeRecipient SecurableType = `RECIPIENT`

const SecurableTypeSchema SecurableType = `SCHEMA`

const SecurableTypeShare SecurableType = `SHARE`

const SecurableTypeStorageCredential SecurableType = `STORAGE_CREDENTIAL`

const SecurableTypeTable SecurableType = `TABLE`

// String representation for [fmt.Print]
func (st *SecurableType) String() string {
	return string(*st)
}

// Set raw string value and validate it against allowed values
func (st *SecurableType) Set(v string) error {
	switch v {
	case `CATALOG`, `EXTERNAL_LOCATION`, `FUNCTION`, `METASTORE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STORAGE_CREDENTIAL`, `TABLE`:
		*st = SecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "EXTERNAL_LOCATION", "FUNCTION", "METASTORE", "PROVIDER", "RECIPIENT", "SCHEMA", "SHARE", "STORAGE_CREDENTIAL", "TABLE"`, v)
	}
}

// Type always returns SecurableType to satisfy [pflag.Value] interface
func (st *SecurableType) Type() string {
	return "SecurableType"
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {
	ForeignKeyConstraint *ForeignKeyConstraint `json:"foreign_key_constraint,omitempty"`

	NamedTableConstraint *NamedTableConstraint `json:"named_table_constraint,omitempty"`

	PrimaryKeyConstraint *PrimaryKeyConstraint `json:"primary_key_constraint,omitempty"`
}

type TableConstraintList struct {
	// List of table constraints.
	TableConstraints []TableConstraint `json:"table_constraints,omitempty"`
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName string `json:"table_full_name"`
}

type TableInfo struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo `json:"columns,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// Data source format
	DataSourceFormat DataSourceFormat `json:"data_source_format,omitempty"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs any `json:"delta_runtime_properties_kvpairs,omitempty"`

	EffectiveAutoMaintenanceFlag *EffectiveAutoMaintenanceFlag `json:"effective_auto_maintenance_flag,omitempty"`
	// Whether auto maintenance should be enabled for this object and objects
	// under it.
	EnableAutoMaintenance EnableAutoMaintenance `json:"enable_auto_maintenance,omitempty"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of table, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of table.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	RowFilter *TableRowFilter `json:"row_filter,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation string `json:"storage_location,omitempty"`

	TableConstraints *TableConstraintList `json:"table_constraints,omitempty"`
	// Name of table, relative to parent schema.
	TableId string `json:"table_id,omitempty"`

	TableType TableType `json:"table_type,omitempty"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition string `json:"view_definition,omitempty"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies []Dependency `json:"view_dependencies,omitempty"`
}

type TableRowFilter struct {
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []string `json:"input_column_names"`
	// The full name of the row filter SQL UDF.
	Name string `json:"name"`
}

type TableSummary struct {
	// The full name of the table.
	FullName string `json:"full_name,omitempty"`

	TableType TableType `json:"table_type,omitempty"`
}

type TableType string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeManaged TableType = `MANAGED`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (tt *TableType) String() string {
	return string(*tt)
}

// Set raw string value and validate it against allowed values
func (tt *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*tt = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (tt *TableType) Type() string {
	return "TableType"
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId string `json:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type UpdateAutoMaintenance struct {
	// Whether to enable auto maintenance on the metastore.
	Enable bool `json:"enable"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id"`
}

type UpdateAutoMaintenanceResponse struct {
	// Whether auto maintenance is enabled on the metastore.
	State bool `json:"state,omitempty"`
	// Id of the auto maintenance service principal. This will be the user used
	// to run maintenance tasks.
	UserId int64 `json:"user_id,omitempty"`
	// Name of the auto maintenance service principal.
	Username string `json:"username,omitempty"`
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of catalog.
	Name string `json:"name,omitempty" url:"-"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
}

type UpdateExternalLocation struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool `json:"force,omitempty"`
	// Name of the external location.
	Name string `json:"name,omitempty" url:"-"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Path URL of the external location.
	Url string `json:"url,omitempty"`
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`
	// Username of current owner of function.
	Owner string `json:"owner,omitempty"`
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope UpdateMetastoreDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
	// Databricks Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog for the metastore.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id,omitempty" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// The scope of Delta Sharing enabled for the metastore.
type UpdateMetastoreDeltaSharingScope string

const UpdateMetastoreDeltaSharingScopeInternal UpdateMetastoreDeltaSharingScope = `INTERNAL`

const UpdateMetastoreDeltaSharingScopeInternalAndExternal UpdateMetastoreDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (umdss *UpdateMetastoreDeltaSharingScope) String() string {
	return string(*umdss)
}

// Set raw string value and validate it against allowed values
func (umdss *UpdateMetastoreDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*umdss = UpdateMetastoreDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns UpdateMetastoreDeltaSharingScope to satisfy [pflag.Value] interface
func (umdss *UpdateMetastoreDeltaSharingScope) Type() string {
	return "UpdateMetastoreDeltaSharingScope"
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`
}

type UpdateSchema struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of schema.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force bool `json:"force,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name,omitempty" url:"-"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the volume
	FullNameArg string `json:"-" url:"-"`
	// The name of the volume
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty"`
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The name of an existing external location to validate.
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// The GCP service account key configuration.
	GcpServiceAccountKey *GcpServiceAccountKey `json:"gcp_service_account_key,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// The name of the storage credential to validate.
	StorageCredentialName any `json:"storage_credential_name,omitempty"`
	// The external location url to validate.
	Url string `json:"url,omitempty"`
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []ValidationResult `json:"results,omitempty"`
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The operation tested.
	Operation ValidationResultOperation `json:"operation,omitempty"`
	// The results of the tested operation.
	Result ValidationResultResult `json:"result,omitempty"`
}

// The operation tested.
type ValidationResultOperation string

const ValidationResultOperationDelete ValidationResultOperation = `DELETE`

const ValidationResultOperationList ValidationResultOperation = `LIST`

const ValidationResultOperationRead ValidationResultOperation = `READ`

const ValidationResultOperationWrite ValidationResultOperation = `WRITE`

// String representation for [fmt.Print]
func (vro *ValidationResultOperation) String() string {
	return string(*vro)
}

// Set raw string value and validate it against allowed values
func (vro *ValidationResultOperation) Set(v string) error {
	switch v {
	case `DELETE`, `LIST`, `READ`, `WRITE`:
		*vro = ValidationResultOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE", "LIST", "READ", "WRITE"`, v)
	}
}

// Type always returns ValidationResultOperation to satisfy [pflag.Value] interface
func (vro *ValidationResultOperation) Type() string {
	return "ValidationResultOperation"
}

// The results of the tested operation.
type ValidationResultResult string

const ValidationResultResultFail ValidationResultResult = `FAIL`

const ValidationResultResultPass ValidationResultResult = `PASS`

const ValidationResultResultSkip ValidationResultResult = `SKIP`

// String representation for [fmt.Print]
func (vrr *ValidationResultResult) String() string {
	return string(*vrr)
}

// Set raw string value and validate it against allowed values
func (vrr *ValidationResultResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*vrr = ValidationResultResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Type always returns ValidationResultResult to satisfy [pflag.Value] interface
func (vrr *ValidationResultResult) Type() string {
	return "ValidationResultResult"
}

type VolumeInfo struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the volume
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the volume
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the volume
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud
	StorageLocation string `json:"storage_location,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the volume last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// The unique identifier of the volume
	VolumeId string `json:"volume_id,omitempty"`

	VolumeType VolumeType `json:"volume_type,omitempty"`
}

type VolumeType string

const VolumeTypeExternal VolumeType = `EXTERNAL`

const VolumeTypeManaged VolumeType = `MANAGED`

// String representation for [fmt.Print]
func (vt *VolumeType) String() string {
	return string(*vt)
}

// Set raw string value and validate it against allowed values
func (vt *VolumeType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`:
		*vt = VolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED"`, v)
	}
}

// Type always returns VolumeType to satisfy [pflag.Value] interface
func (vt *VolumeType) Type() string {
	return "VolumeType"
}
