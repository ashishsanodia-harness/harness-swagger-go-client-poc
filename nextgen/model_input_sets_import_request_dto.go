package nextgen

// Information of InputSets import request DTO
type InputSetsImportRequestDto struct {
	InputSetName        string `json:"input_set_name,omitempty"`
	InputSetDescription string `json:"input_set_description,omitempty"`
}