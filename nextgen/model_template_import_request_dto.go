package nextgen

// Information of Templates import request DTO
type TemplatesImportRequestDto struct {
	TemplateName        string `json:"template_name,omitempty"`
	TemplateVersion        string `json:"template_version,omitempty"`
	TemplateDescription string `json:"template_description,omitempty"`
}