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

// Default response when a service is returned
type ServiceResponse struct {
	Service *Service `json:"service,omitempty"`
	// Creation timestamp for Service.
	Created int64 `json:"created,omitempty"`
	// Last modification timestamp for Service.
	Updated int64 `json:"updated,omitempty"`
}
