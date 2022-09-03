// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

// all definitions in this file are in alphabetical order

type DeleteRequest struct {
    // The path of the file or directory to delete. The path should be the
    // absolute DBFS path (e.g. &#34;/mnt/foo/&#34;).
    Path string `json:"path"`
    // Whether or not to recursively delete the directory&#39;s contents. Deleting
    // empty directories can be done without providing the recursive flag.
    Recursive bool `json:"recursive,omitempty"`
}


type ExportRequest struct {
    // Flag to enable direct download. If it is ``true``, the response will be
    // the exported file itself. Otherwise, the response contains content as
    // base64 encoded string. See :ref:`workspace-api-export-example` for more
    // information about how to use it.
    DirectDownload bool ` url:"direct_download,omitempty"`
    // This specifies the format of the exported file. By default, this is
    // ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
    // ``DBC``. The value is case sensitive.
    Format string ` url:"format,omitempty"`
    // The absolute path of the notebook or directory. Exporting directory is
    // only support for ``DBC`` format.
    Path string ` url:"path,omitempty"`
}


type ExportResponse struct {
    // The base64-encoded content. If the limit (10MB) is exceeded, exception
    // with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown.
    Content string `json:"content,omitempty"`
}


type GetStatusRequest struct {
    // The absolute path of the notebook or directory.
    Path string ` url:"path,omitempty"`
}


type GetStatusResponse struct {
    // The length of the file in bytes or zero if the path is a directory.
    FileSize int64 `json:"file_size,omitempty"`
    // True if the path is a directory.
    IsDir bool `json:"is_dir,omitempty"`
    // Last modification time of given file/dir in milliseconds since Epoch.
    ModificationTime int64 `json:"modification_time,omitempty"`
    // The path of the file or directory.
    Path string `json:"path,omitempty"`
}


type ImportRequest struct {
    // The base64-encoded content. This has a limit of 10 MB. If the limit
    // (10MB) is exceeded, exception with error code
    // **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown. This parameter might be
    // absent, and instead a posted file will be used. See
    // :ref:`workspace-api-import-example` for more information about how to use
    // it.
    Content string `json:"content,omitempty"`
    // This specifies the format of the file to be imported. By default, this is
    // ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
    // ``DBC``. The value is case sensitive.
    Format ImportRequestFormat `json:"format,omitempty"`
    // The language. If format is set to ``SOURCE``, this field is required;
    // otherwise, it will be ignored.
    Language ImportRequestLanguage `json:"language,omitempty"`
    // The flag that specifies whether to overwrite existing object. It is
    // ``false`` by default. For ``DBC`` format, ``overwrite`` is not supported
    // since it may contain a directory.
    Overwrite bool `json:"overwrite,omitempty"`
    // The absolute path of the notebook or directory. Importing directory is
    // only support for ``DBC`` format.
    Path string `json:"path"`
}

// This specifies the format of the file to be imported. By default, this is
// ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
// ``DBC``. The value is case sensitive.
type ImportRequestFormat string


const ImportRequestFormatDbc ImportRequestFormat = `DBC`

const ImportRequestFormatHtml ImportRequestFormat = `HTML`

const ImportRequestFormatJupyter ImportRequestFormat = `JUPYTER`

const ImportRequestFormatRMarkdown ImportRequestFormat = `R_MARKDOWN`

const ImportRequestFormatSource ImportRequestFormat = `SOURCE`
// The language. If format is set to ``SOURCE``, this field is required;
// otherwise, it will be ignored.
type ImportRequestLanguage string


const ImportRequestLanguagePython ImportRequestLanguage = `PYTHON`

const ImportRequestLanguageR ImportRequestLanguage = `R`

const ImportRequestLanguageScala ImportRequestLanguage = `SCALA`

const ImportRequestLanguageSql ImportRequestLanguage = `SQL`

type ListRequest struct {
    
    NotebooksModifiedAfter int ` url:"notebooks_modified_after,omitempty"`
    // The absolute path of the notebook or directory.
    Path string ` url:"path,omitempty"`
}


type ListResponse struct {
    // List of objects.
    Objects []ObjectInfo `json:"objects,omitempty"`
}


type MkdirsRequest struct {
    // The absolute path of the directory. If the parent directories do not
    // exist, it will also create them. If the directory already exists, this
    // command will do nothing and succeed.
    Path string `json:"path"`
}


type ObjectInfo struct {
    // The location (bucket and prefix) enum value of the content blob. This
    // field is used in conjunction with the blob_path field to determine where
    // the blob is located.
    BlobLocation ObjectInfoBlobLocation `json:"blob_location,omitempty"`
    // ========= File metadata. These values are set only if the object type is
    // ``FILE``. ===========//
    BlobPath string `json:"blob_path,omitempty"`
    
    ContentSha256Hex string `json:"content_sha256_hex,omitempty"`
    
    CreatedAt int64 `json:"created_at,omitempty"`
    // The language of the object. This value is set only if the object type is
    // ``NOTEBOOK``.
    Language ObjectInfoLanguage `json:"language,omitempty"`
    
    MetadataVersion int `json:"metadata_version,omitempty"`
    
    ModifiedAt int64 `json:"modified_at,omitempty"`
    
    ObjectId int64 `json:"object_id,omitempty"`
    
    ObjectType ObjectInfoObjectType `json:"object_type,omitempty"`
    // The absolute path of the object.
    Path string `json:"path,omitempty"`
    
    Size int64 `json:"size,omitempty"`
}

// The location (bucket and prefix) enum value of the content blob. This field
// is used in conjunction with the blob_path field to determine where the blob
// is located.
type ObjectInfoBlobLocation string


const ObjectInfoBlobLocationDbfsRoot ObjectInfoBlobLocation = `DBFS_ROOT`

const ObjectInfoBlobLocationInternalDbfsJobs ObjectInfoBlobLocation = `INTERNAL_DBFS_JOBS`
// The language of the object. This value is set only if the object type is
// ``NOTEBOOK``.
type ObjectInfoLanguage string


const ObjectInfoLanguagePython ObjectInfoLanguage = `PYTHON`

const ObjectInfoLanguageR ObjectInfoLanguage = `R`

const ObjectInfoLanguageScala ObjectInfoLanguage = `SCALA`

const ObjectInfoLanguageSql ObjectInfoLanguage = `SQL`

type ObjectInfoObjectType string


const ObjectInfoObjectTypeDirectory ObjectInfoObjectType = `DIRECTORY`

const ObjectInfoObjectTypeFile ObjectInfoObjectType = `FILE`

const ObjectInfoObjectTypeLibrary ObjectInfoObjectType = `LIBRARY`

const ObjectInfoObjectTypeMlflowExperiment ObjectInfoObjectType = `MLFLOW_EXPERIMENT`

const ObjectInfoObjectTypeNotebook ObjectInfoObjectType = `NOTEBOOK`

const ObjectInfoObjectTypeProject ObjectInfoObjectType = `PROJECT`

const ObjectInfoObjectTypeRepo ObjectInfoObjectType = `REPO`
