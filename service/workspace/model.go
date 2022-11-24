// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

// all definitions in this file are in alphabetical order

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// ``false`` by default. Please note this deleting directory is not atomic.
	// If it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`
}

// This specifies the format of the file to be imported. By default, this is
// “SOURCE“. However it may be one of: “SOURCE“, “HTML“, “JUPYTER“,
// “DBC“. The value is case sensitive.
type ExportFormat string

const ExportFormatDbc ExportFormat = `DBC`

const ExportFormatHtml ExportFormat = `HTML`

const ExportFormatJupyter ExportFormat = `JUPYTER`

const ExportFormatRMarkdown ExportFormat = `R_MARKDOWN`

const ExportFormatSource ExportFormat = `SOURCE`

type ExportRequest struct {
	// Flag to enable direct download. If it is ``true``, the response will be
	// the exported file itself. Otherwise, the response contains content as
	// base64 encoded string.
	DirectDownload bool `json:"-" url:"direct_download,omitempty"`
	// This specifies the format of the exported file. By default, this is
	// ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
	// ``DBC``.
	//
	// The value is case sensitive.
	Format ExportFormat `json:"-" url:"format,omitempty"`
	// The absolute path of the notebook or directory. Exporting directory is
	// only support for ``DBC`` format.
	Path string `json:"-" url:"path,omitempty"`
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown.
	Content string `json:"content,omitempty"`
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path,omitempty"`
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown. This parameter might be
	// absent, and instead a posted file will be used.
	Content string `json:"content,omitempty"`
	// This specifies the format of the file to be imported. By default, this is
	// ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
	// ``DBC``. The value is case sensitive.
	Format ExportFormat `json:"format,omitempty"`
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	Language Language `json:"language,omitempty"`
	// The flag that specifies whether to overwrite existing object. It is
	// ``false`` by default. For ``DBC`` format, ``overwrite`` is not supported
	// since it may contain a directory.
	Overwrite bool `json:"overwrite,omitempty"`
	// The absolute path of the notebook or directory. Importing directory is
	// only support for ``DBC`` format.
	Path string `json:"path"`
}

// The language of the object. This value is set only if the object type is
// “NOTEBOOK“.
type Language string

const LanguagePython Language = `PYTHON`

const LanguageR Language = `R`

const LanguageScala Language = `SCALA`

const LanguageSql Language = `SQL`

type ListRequest struct {
	// <content needed>
	NotebooksModifiedAfter int `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path,omitempty"`
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `json:"objects,omitempty"`
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`
}

type ObjectInfo struct {
	// <content needed>
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	Language Language `json:"language,omitempty"`
	// <content needed>
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// <content needed>
	ObjectId int64 `json:"object_id,omitempty"`
	// The type of the object in workspace.
	ObjectType ObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	Path string `json:"path,omitempty"`
	// <content needed>
	Size int64 `json:"size,omitempty"`
}

// The type of the object in workspace.
type ObjectType string

const ObjectTypeDirectory ObjectType = `DIRECTORY`

const ObjectTypeFile ObjectType = `FILE`

const ObjectTypeLibrary ObjectType = `LIBRARY`

const ObjectTypeNotebook ObjectType = `NOTEBOOK`

const ObjectTypeRepo ObjectType = `REPO`
