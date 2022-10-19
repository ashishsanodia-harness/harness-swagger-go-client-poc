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

// Pipeline response body.
type PipelineGetResponseBody struct {
	// Pipeline YAML (returned as a String).
	PipelineYaml string `json:"pipeline_yaml,omitempty"`
	// Pipeline YAML after resolving Templates (returned as a String).
	TemplateAppliedPipelineYaml string `json:"template_applied_pipeline_yaml,omitempty"`
	// Pipeline identifier
	Slug string `json:"slug,omitempty"`
	// Pipeline name
	Name string `json:"name,omitempty"`
	// Pipeline description
	Description string `json:"description,omitempty"`
	// Pipeline tags
	Tags map[string]string `json:"tags,omitempty"`
	// Modules utilised in the Pipeline.
	Modules []string `json:"modules,omitempty"`
	GitDetails *GitDetails `json:"git_details,omitempty"`
	// Specifies whether Pipeline is a valid or not.
	Valid bool `json:"valid,omitempty"`
	// YAML schema errors.
	YamlErrorWrapper []YamlSchemaErrorWrapper `json:"yaml_error_wrapper,omitempty"`
	// Creation timestamp for Pipeline.
	Created int64 `json:"created,omitempty"`
	// Last modification timestamp for Pipeline.
	Updated int64 `json:"updated,omitempty"`
}
