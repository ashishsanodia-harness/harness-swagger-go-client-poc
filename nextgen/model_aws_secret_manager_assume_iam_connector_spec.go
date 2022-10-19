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

// This contains details of the AWS connector and Harness will authenticate using the IAM role assigned to the AWS host running the Delegate, you select using a Delegate Selector.
type AwsSecretManagerAssumeIamConnectorSpec struct {
	// This specifies the type of connector
	Type_ string `json:"type"`
	// AWS Region for kms
	Region string `json:"region"`
	// Boolean value to indicate if the Secret Manager is your default Secret Manager
	Default_ bool `json:"default,omitempty"`
	// Text that is prepended to the Secret name as a prefix
	SecretNamePrefix string `json:"secret_name_prefix,omitempty"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager
	DelegateSelectors []string `json:"delegate_selectors,omitempty"`
}
