package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshplay/server/internal/sql"
	"gopkg.in/yaml.v2"
)

// MeshplayFilter represents the filters that needs to be saved
type MeshplayFilter struct {
	ID *uuid.UUID `json:"id,omitempty"`

	Name       string `json:"name,omitempty"`
	FilterFile []byte `json:"filter_file"`
	// Meshplay doesn't have the user id fields
	// but the remote provider is allowed to provide one
	UserID *string `json:"user_id"`

	Location       sql.Map    `json:"location"`
	Visibility     string     `json:"visibility"`
	CatalogData    sql.Map    `json:"catalog_data"`
	FilterResource string     `json:"filter_resource"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
}

type MeshplayFilterPayload struct {
	ID *uuid.UUID `json:"id,omitempty"`

	Name       string `json:"name,omitempty"`
	FilterFile []byte `json:"filter_file"`
	// Meshplay doesn't have the user id fields
	// but the remote provider is allowed to provide one
	UserID *string `json:"user_id"`

	Location       sql.Map    `json:"location"`
	Visibility     string     `json:"visibility"`
	CatalogData    sql.Map    `json:"catalog_data"`
	FilterResource string     `json:"filter_resource"`
	Config         string     `json:"config"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
}

// MeshplayCatalogFilterRequestBody refers to the type of request body that PublishCatalogFilter would receive
type MeshplayCatalogFilterRequestBody struct {
	ID          uuid.UUID `json:"id,omitempty"`
	CatalogData sql.Map   `json:"catalog_data,omitempty"`
}

// MeshplayCatalogFilterRequestBody refers to the type of request body
// that CloneMeshplayFilterHandler would receive
type MeshplayCloneFilterRequestBody struct {
	Name string `json:"name,omitempty"`
}

// MeshplayFilterRequestBody refers to the type of request body that
// SaveMeshplayFilter would receive
type MeshplayFilterRequestBody struct {
	URL        string                `json:"url,omitempty"`
	Path       string                `json:"path,omitempty"`
	Save       bool                  `json:"save,omitempty"`
	Config     string                `json:"config,omitempty"`
	FilterData *MeshplayFilterPayload `json:"filter_data,omitempty"`
}

// GetFilterName takes in a stringified filterfile and extracts the name from it
func GetFilterName(stringifiedFile string) (string, error) {
	out := map[string]interface{}{}

	if err := yaml.Unmarshal([]byte(stringifiedFile), &out); err != nil {
		return "", err
	}

	// Get Name from the file
	name, ok := out["name"].(string)
	if !ok {
		return "", ErrFilterFileName
	}

	return name, nil
}
