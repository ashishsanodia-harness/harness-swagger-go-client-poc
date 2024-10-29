package nextgen

// Templates import request body
type TemplatesImportRequestBody struct {
	GitImportDetails         *GitImportDetails           `json:"git_import_details,omitempty"`
	TemplatesImportRequest *TemplatesImportRequestDto `json:"template_import_request,omitempty"`
}