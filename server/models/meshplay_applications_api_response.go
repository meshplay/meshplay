package models

// ApplicationsAPIResponse response retruned by patternfile endpoint on meshplay server
type ApplicationsAPIResponse struct {
	Page         uint                 `json:"page"`
	PageSize     uint                 `json:"page_size"`
	TotalCount   uint                 `json:"total_count"`
	Applications []MeshplayApplication `json:"applications"`
}

type ApplicationSourceTypesAPIResponse struct {
	ApplicationType     string   `json:"application_type"`
	SupportedExtensions []string `json:"supported_extensions"`
}
