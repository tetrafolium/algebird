package sarif

// Result is array element of Results in SARIF.
type Result struct {
	Level     string    `json:"level"`
	Message   Message   `json:"message"`
	Locations Locations `json:"locations"`
	Rank      float64   `json:"rank"`
}

// Message ...
type Message struct {
	Text string `json:"text"`
}

// Locations ...
type Locations []*Location

// Location ...
type Location struct {
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}

// PhysicalLocation ...
type PhysicalLocation struct {
	ArtifactLocation ArtifactLocation `json:"atifactLocation"`
	Region           Region           `json:"region"`
}

// ArtifactLocation ...
type ArtifactLocation struct {
	URI       string `json:"uri"`
	URIBaseID string `json:"uriBaseId"`
}

// Region ...
type Region struct {
	StartLine      int    `json:"startLine,omitempty"`
	StartColumn    int    `json:"startColumn,omitempty"`
	EndLine        int    `json:"endLine,omitempty"`
	EndColumn      int    `json:"endColumn,omitempty"`
	SourceLanguage string `json:"sourceLanguage,omitempty"`
}
