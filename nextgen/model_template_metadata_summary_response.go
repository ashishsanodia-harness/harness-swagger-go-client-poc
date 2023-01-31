/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Single Template Metadata Model
type TemplateMetadataSummaryResponse struct {
	// Account identifier
	Account string `json:"account,omitempty"`
	// Organization identifier
	Org string `json:"org,omitempty"`
	// Project identifier
	Project string `json:"project,omitempty"`
	Slug    string `json:"slug,omitempty"`
	// Template identifier
	Identifier string `json:"identifier,omitempty"`
	// Template Name
	Name string `json:"name,omitempty"`
	// Template description
	Description string `json:"description,omitempty"`
	// Template tags
	Tags map[string]string `json:"tags,omitempty"`
	// Version label of template
	VersionLabel string `json:"version_label,omitempty"`
	// Type of Template
	EntityType string `json:"entity_type,omitempty"`
	// Defines child template type
	ChildType string `json:"child_type,omitempty"`
	// Scope of template
	Scope string `json:"scope,omitempty"`
	// Version of template
	Version    int64             `json:"version,omitempty"`
	GitDetails *EntityGitDetails `json:"git_details,omitempty"`
	// Last modification timestamp for Service.
	Updated int64 `json:"updated,omitempty"`
	// Specifies whether the Entity is to be stored in Git or not (for Git Experience).
	StoreType string `json:"store_type,omitempty"`
	// Identifier of the Harness Connector used for CRUD operations on the Entity (for Git Experience).
	ConnectorRef string `json:"connector_ref,omitempty"`
	// True if this version is stable version of Template
	StableTemplate bool `json:"stable_template,omitempty"`
}
