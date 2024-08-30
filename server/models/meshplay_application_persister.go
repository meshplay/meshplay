package models

import (
	"encoding/json"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshkit/database"
)

// MeshplayApplicationPersister is the persister for persisting
// applications on the database
type MeshplayApplicationPersister struct {
	DB *database.Handler
}

// MeshplayApplicationPage represents a page of applications
type MeshplayApplicationPage struct {
	Page         uint64                `json:"page"`
	PageSize     uint64                `json:"page_size"`
	TotalCount   int                   `json:"total_count"`
	Applications []*MeshplayApplication `json:"applications"`
}

// GetMeshplayApplications returns all of the applications
func (maap *MeshplayApplicationPersister) GetMeshplayApplications(search, order string, page, pageSize uint64, updatedAfter string) ([]byte, error) {
	order = SanitizeOrderInput(order, []string{"created_at", "updated_at", "name"})

	if order == "" {
		order = "updated_at desc"
	}

	count := int64(0)
	applications := []*MeshplayApplication{}

	query := maap.DB.Where("updated_at > ?", updatedAfter).Order(order)

	if search != "" {
		like := "%" + strings.ToLower(search) + "%"
		query = query.Where("(lower(meshplay_applications.name) like ?)", like)
	}

	query.Table("meshplay_applications").Count(&count)

	Paginate(uint(page), uint(pageSize))(query).Find(&applications)

	meshplayApplicationPage := &MeshplayApplicationPage{
		Page:         page,
		PageSize:     pageSize,
		TotalCount:   int(count),
		Applications: applications,
	}

	return marshalMeshplayApplicationPage(meshplayApplicationPage), nil
}

// DeleteMeshplayApplication takes in an application id and delete it if it already exists
func (maap *MeshplayApplicationPersister) DeleteMeshplayApplication(id uuid.UUID) ([]byte, error) {
	application := MeshplayApplication{ID: &id}
	err := maap.DB.Delete(&application).Error

	return marshalMeshplayApplication(&application), err
}

func (maap *MeshplayApplicationPersister) SaveMeshplayApplication(application *MeshplayApplication) ([]byte, error) {
	if application.ID == nil {
		id, err := uuid.NewV4()
		if err != nil {
			return nil, ErrGenerateUUID(err)
		}

		application.ID = &id
	}

	return marshalMeshplayApplications([]MeshplayApplication{*application}), maap.DB.Save(application).Error
}

// SaveMeshplayApplications batch inserts the given applications
func (maap *MeshplayApplicationPersister) SaveMeshplayApplications(applications []MeshplayApplication) ([]byte, error) {
	finalApplications := []MeshplayApplication{}
	for _, application := range applications {
		if application.ID == nil {
			id, err := uuid.NewV4()
			if err != nil {
				return nil, ErrGenerateUUID(err)
			}

			application.ID = &id
		}

		finalApplications = append(finalApplications, application)
	}

	return marshalMeshplayApplications(finalApplications), maap.DB.Create(finalApplications).Error
}

func (maap *MeshplayApplicationPersister) GetMeshplayApplication(id uuid.UUID) ([]byte, error) {
	var meshplayApplication MeshplayApplication
	err := maap.DB.First(&meshplayApplication, id).Error
	return marshalMeshplayApplication(&meshplayApplication), err
}

func (maap *MeshplayApplicationPersister) GetMeshplayApplicationSource(id uuid.UUID) ([]byte, error) {
	var meshplayApplication MeshplayApplication
	err := maap.DB.First(&meshplayApplication, id).Error
	return meshplayApplication.SourceContent, err
}

func marshalMeshplayApplicationPage(maap *MeshplayApplicationPage) []byte {
	res, _ := json.Marshal(maap)

	return res
}

func marshalMeshplayApplication(ma *MeshplayApplication) []byte {
	res, _ := json.Marshal(ma)

	return res
}

func marshalMeshplayApplications(mas []MeshplayApplication) []byte {
	res, _ := json.Marshal(mas)

	return res
}
