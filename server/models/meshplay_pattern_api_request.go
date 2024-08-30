package models

// MeshplayPatternDeleteRequestBody refers to the type of request body
// that DeleteMultiMeshplayPatternsHandler would receive
type MeshplayPatternDeleteRequestBody struct {
	Patterns []deletePatternModel `json:"patterns,omitempty"`
}

// DeletePatternModel is the model for individual patterns to be deleted
type deletePatternModel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
