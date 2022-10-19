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

// Filter by whether the Resource Group has a particular Resource.
type ResourceSelectorFilter struct {
	// Filter by Resource type
	ResourceType string `json:"resource_type"`
	// Filter by Resource identifier
	ResourceSlug string `json:"resource_slug,omitempty"`
}
