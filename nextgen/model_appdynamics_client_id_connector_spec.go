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

// This contains details of the appdynamics connector with client secrets
type AppdynamicsClientIdConnectorSpec struct {
	Type_ string `json:"type"`
	// List of unique delegate selectors
	DelegateSelectors []string `json:"delegate_selectors,omitempty"`
	// appdymanics account name
	AccountName string `json:"account_name"`
	// appdynamics controller url
	ControllerUrl string `json:"controller_url"`
	// appdynamics client id
	ClientId string `json:"client_id,omitempty"`
	// Reference to encrypted Harness secret for appdynamics client secret
	ClientSecretRef string `json:"client_secret_ref,omitempty"`
}