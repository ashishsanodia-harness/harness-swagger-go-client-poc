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

// Specifies resources in Resource Group.
type ResourceFilter struct {
	// Specifies the type of Resource.
	ResourceType string `json:"resource_type"`
	// Identifiers for the Resource Type.
	Identifiers []string `json:"identifiers,omitempty"`
	// Attribute name on the basis of which filtering will be done.
	AttributeName string `json:"attribute_name,omitempty"`
	// Attribute values selected.
	AttributeValues []string `json:"attribute_values,omitempty"`
}
