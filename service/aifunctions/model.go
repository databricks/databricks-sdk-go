// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"encoding/json"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// ai_classify
type AiClassifyOptions struct {
	// When true, includes a per-label confidence score in the response.
	EnableConfidenceScores bool `json:"enable_confidence_scores,omitempty"`
	// When true, includes a rationale explaining each classification in the
	// response.
	EnableRationales bool `json:"enable_rationales,omitempty"`
	// Natural-language guidance that steers how the text is classified (up to
	// 20,000 characters).
	Instructions string `json:"instructions,omitempty"`
	// When true, allows more than one label to be returned per input.
	Multilabel bool `json:"multilabel,omitempty"`
	// The function version to invoke. Defaults to the latest version. Supported
	// versions: ["2.1"].
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiClassifyOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiClassifyOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiClassifyRequest struct {
	// The content to classify. It accepts a plain string or the response object
	// of [ai_parse_document](:method:AiFunctions/AiParseDocument).
	Content json.RawMessage `json:"content"`
	// The label set to classify as. Either a JSON array of label strings (e.g.
	// ["spam", "not_spam"]), or a JSON object mapping each label to a
	// description (e.g. {"spam": "unsolicited bulk message", "not_spam": "a
	// legitimate message"}). Accepts 2 to 500 labels, each 1 to 100 characters.
	Labels json.RawMessage `json:"labels"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiClassifyOptions `json:"options,omitempty"`
}

func (s *AiClassifyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiClassifyResponse struct {
	// Additional metadata returned by AI Classify.
	Metadata *AiClassifyResponseMetadata `json:"metadata,omitempty"`
	// The function result as a JSON value. An array of per-label objects: one
	// element in single-label mode (the default), or multiple elements when
	// `multilabel` is true. When `enable_confidence_scores` and
	// `enable_rationales` are true, `confidence_score` and `rationale` are
	// included in each response value, respectively.
	Response *json.RawMessage `json:"response,omitempty"`
}

func (s *AiClassifyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiClassifyResponseMetadata struct {
	// The resolved function version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiClassifyResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiClassifyResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A bounding box on a source page; used by bbox-input citations.
type AiExtractBbox struct {
	// Pixel coordinates on the page image as [x0, y0, x1, y1].
	Coord []int64 `json:"coord,omitempty"`
	// 0-based page index the box is on.
	PageId int64 `json:"page_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractBbox) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractBbox) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A citation locating an extracted value in the source. start/stop are set for
// span (STRING input) citations, bbox for bbox (parsed-document input)
// citations.
type AiExtractCitation struct {
	// Bounding boxes locating the citation on the source pages; set for bbox
	// citations.
	Bbox []AiExtractBbox `json:"bbox,omitempty"`
	// Integer matching a citation_ids entry on an extracted field.
	Id int64 `json:"id,omitempty"`
	// Inclusive 0-based character offset into the input string; set for span
	// citations.
	Start int64 `json:"start,omitempty"`
	// Exclusive 0-based character offset into the input string; set for span
	// citations.
	Stop int64 `json:"stop,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractCitation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractCitation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ai_extract
type AiExtractOptions struct {
	// When true, includes citation metadata locating each extracted value in
	// the source. Depending on the type of input, citations can be one of two
	// types:
	//
	// For raw text (STRING) inputs, a citation is a span of text in the
	// original input. Each object in `metadata.citations` has an `id` (integer
	// matching a `citation_ids` entry on a field), a `start` (inclusive 0-based
	// character offset into the input string), and a `stop` (exclusive 0-based
	// character offset into the input string).
	//
	// For PDF documents and images (when using ai_extract downstream of
	// ai_parse_document), a citation is a bounding box in the original input.
	// Each object in `metadata.citations` has an `id` (integer matching a
	// `citation_ids` entry on a field) and a `bbox` (array of {coord, page_id}
	// objects, identical in shape to element.bbox in ai_parse_document output;
	// coord is pixel coordinates on the page image as [x0, y0, x1, y1], and
	// page_id is a 0-based page index).
	EnableCitations bool `json:"enable_citations,omitempty"`
	// When true, includes a per-field confidence score in the response.
	EnableConfidenceScores bool `json:"enable_confidence_scores,omitempty"`
	// Natural-language guidance that steers how data is extracted (up to 20,000
	// characters).
	Instructions string `json:"instructions,omitempty"`
	// Extraction mode. Supported modes: "precision" — more powerful
	// extraction for complex schemas, long documents, and reasoning-heavy
	// extractions. Defaults to none (standard extraction).
	Mode string `json:"mode,omitempty"`
	// The function version to invoke. Defaults to the latest version. Supported
	// versions: ["2.1"].
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiExtractRequest struct {
	// The text to extract from. It accepts a plain string or the response
	// object of [ai_parse_document](:method:AiFunctions/AiParseDocument).
	Content json.RawMessage `json:"content"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiExtractOptions `json:"options,omitempty"`
	// The extraction schema defining the fields to extract. Either a JSON array
	// of field names, assumed to be strings (e.g. ["company", "valuation"]), or
	// a JSON object mapping each field to its type/description/nullability
	// (e.g. {"company": {"type": "string", "description": "the company
	// name"}}). Accepts up to 256 fields, 12 levels of nesting, and 500 enum
	// values. Supported field types are string, integer, number, boolean, and
	// enum.
	Schema json.RawMessage `json:"schema"`
}

func (s *AiExtractRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiExtractResponse struct {
	// Additional metadata returned by AI Extract.
	Metadata *AiExtractResponseMetadata `json:"metadata,omitempty"`
	// The function result as a JSON value. When `enable_confidence_scores` and
	// `enable_citations` are true, `confidence` and `citation_ids` are included
	// in each response field, respectively.
	Response *json.RawMessage `json:"response,omitempty"`
}

func (s *AiExtractResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiExtractResponseMetadata struct {
	// How the source was chunked for citation offsets (span for text input,
	// bbox for parsed-document input); present when citations are enabled.
	ChunkType string `json:"chunk_type,omitempty"`
	// Citation objects locating each result in the source; present when
	// citations are enabled.
	Citations []AiExtractCitation `json:"citations,omitempty"`
	// The resolved extraction mode; present when a non-default mode was used.
	Mode string `json:"mode,omitempty"`
	// The resolved function version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata about the source file; present only for file-path input.
type AiParseDocumentFileMetadata struct {
	// Last-modified timestamp of the source file, as an HTTP date string.
	FileModificationTime string `json:"file_modification_time,omitempty"`
	// Base name of the source file.
	FileName string `json:"file_name,omitempty"`
	// Unity Catalog volume path of the source file.
	FilePath string `json:"file_path,omitempty"`
	// Size of the source file in bytes.
	FileSize int64 `json:"file_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentFileMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentFileMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ai_parse_document
type AiParseDocumentOptions struct {
	// Element types for which an AI-generated description is produced. Use "*"
	// (default) to generate descriptions for all supported element types,
	// "figure" to generate them for figures only, or "" (empty string) to
	// generate none. Only figure descriptions are supported for version "2.0",
	// so "*" and "figure" produce the same behavior.
	DescriptionElementTypes string `json:"description_element_types,omitempty"`
	// Unity Catalog volume path where rendered page and element images are
	// written.
	ImageOutputPath string `json:"image_output_path,omitempty"`
	// Pages to parse (1-indexed), as a comma-separated list of page numbers or
	// ranges (e.g. "1,3,5-10").
	PageRange string `json:"page_range,omitempty"`
	// The ai_parse_document output schema version. Supported value: "2.0".
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A single page that failed to parse while the overall request succeeded.
type AiParseDocumentPageError struct {
	// Message describing why the page failed.
	ErrorMessage string `json:"error_message,omitempty"`
	// 0-based index of the page that failed.
	PageId int64 `json:"page_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentPageError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentPageError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiParseDocumentRequest struct {
	// The document to parse, given as a Unity Catalog volume path to the source
	// file (the REST API accepts only a UC volume path, not inline binary
	// data). Supported formats: PDF, DOCX, DOC, PPTX, PPT, JPG, JPEG, PNG,
	// TIFF. Accepts up to 100 pages and 100 MB per document.
	Content string `json:"content"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiParseDocumentOptions `json:"options,omitempty"`
}

func (s *AiParseDocumentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiParseDocumentResponse struct {
	// The parsed document as a JSON value, containing the extracted pages and
	// elements.
	Document *json.RawMessage `json:"document,omitempty"`
	// Per-page partial-failure details; present when the request succeeds (2xx)
	// but individual pages fail.
	ErrorStatus []AiParseDocumentPageError `json:"error_status,omitempty"`
	// Additional metadata returned by AI Parse Document.
	Metadata *AiParseDocumentResponseMetadata `json:"metadata,omitempty"`
}

func (s *AiParseDocumentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiParseDocumentResponseMetadata struct {
	// Describes the source file; present only for file-path input.
	FileMetadata *AiParseDocumentFileMetadata `json:"file_metadata,omitempty"`
	// Unique identifier for the parse request.
	Id string `json:"id,omitempty"`
	// The resolved function version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
