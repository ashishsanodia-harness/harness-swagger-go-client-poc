package nextgen

// InputSet Validation Response Body
type InputSetValidationResponseBody struct {
	// Status of the InputSet Validation Event
	Status string `json:"status,omitempty"`
	// Result of Policy Evaluations on the InputSet
	PolicyEval *interface{} `json:"policy_eval,omitempty"`
	// Start time of the Evaluation
	StartTs int64 `json:"start_ts,omitempty"`
	// End time of the Evaluation
	EndTs                      int64                           `json:"end_ts,omitempty"`
	TemplateValidationResponse *TemplateValidationResponseBody `json:"template_validation_response,omitempty"`
}