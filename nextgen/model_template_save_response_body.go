package nextgen

// Response body for Templates save.
type TemplatesSaveResponseBody struct {
	Identifier         string               `json:"identifier,omitempty"`
	GovernanceMetadata []GovernanceMetadata `json:"governance_metadata,omitempty"`
}