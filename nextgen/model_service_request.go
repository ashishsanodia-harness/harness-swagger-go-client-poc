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

// Service Request Body 
type ServiceRequest struct {
	// Identifier of the Service
	Slug string `json:"slug"`
	// Name of the Service
	Name string `json:"name"`
	// Description of the entity
	Description string `json:"description,omitempty"`
	// Service tags
	Tags map[string]string `json:"tags,omitempty"`
	// YAML for the Service Request
	Yaml string `json:"yaml,omitempty"`
}
