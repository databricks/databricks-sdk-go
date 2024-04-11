package openapi

// DataPlane is the Databricks OpenAPI Extension for direct access to DataPlane APIs
type DataPlane struct {
	InfoOperationId string   `json:"configMethod"`
	InfoField       []string `json:"field"`
}

// Pagination is the Databricks OpenAPI Extension for retrieving
// lists of entities through multiple API calls
type Pagination struct {
	Offset    string   `json:"offset,omitempty"`
	Limit     string   `json:"limit,omitempty"`
	Results   string   `json:"results,omitempty"`
	Increment int      `json:"increment,omitempty"`
	Inline    bool     `json:"inline,omitempty"`
	Token     *Binding `json:"token,omitempty"`
}

// Wait is the Databricks OpenAPI Extension for long-running result polling
type Wait struct {
	Poll         string             `json:"poll"`
	Bind         string             `json:"bind"`
	BindResponse string             `json:"bindResponse,omitempty"`
	Binding      map[string]Binding `json:"binding,omitempty"`
	Field        []string           `json:"field"`
	Message      []string           `json:"message"`
	Success      []string           `json:"success"`
	Failure      []string           `json:"failure"`
	Timeout      int                `json:"timeout,omitempty"`
}

// Binding is a relationship between request and/or response
type Binding struct {
	Request  string `json:"request,omitempty"`
	Response string `json:"response,omitempty"`
}
