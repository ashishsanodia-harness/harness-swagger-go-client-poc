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

// Templates Create Request Body
type TemplateCreateRequestBody struct {
	// Yaml for creating new Template
	TemplateYaml string `json:"template_yaml,omitempty"`
	GitDetails *GitCreateDetails1 `json:"git_details,omitempty"`
	// True if given version for template to be set as stable
	IsStable bool `json:"is_stable,omitempty"`
	// Specify comment with respect to changes  
	Comments string `json:"comments,omitempty"`
}
