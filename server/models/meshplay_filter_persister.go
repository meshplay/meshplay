package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/layer5io/meshkit/database"
)

// MeshplayFilterPersister is the persister for persisting
// wasm filters on the database
type MeshplayFilterPersister struct {
	DB *database.Handler
}

// MeshplayFilterPage represents a page of filters
type MeshplayFilterPage struct {
	Page       uint64           `json:"page"`
	PageSize   uint64           `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Filters    []*MeshplayFilter `json:"filters"`
}

// GetMeshplayFilters returns all of the 'private' filters. Though private has no meaning here since there is only
// one local user. We make this distinction to be consistent with the remote provider
func (mfp *MeshplayFilterPersister) GetMeshplayFilters(search, order string, page, pageSize uint64, visibility []string) ([]byte, error) {
	order = SanitizeOrderInput(order, []string{"created_at", "updated_at", "name"})

	if order == "" {
		order = "updated_at desc"
	}

	count := int64(0)
	filters := []*MeshplayFilter{}

	query := mfp.DB.Table("meshplay_filters")

	if len(visibility) > 0 {
		query = query.Where("visibility in (?)", visibility)
	}

	query = query.Order(order)

	if search != "" {
		like := "%" + strings.ToLower(search) + "%"
		query = query.Where("(lower(meshplay_filters.name) like ?)", like)
	}

	query.Count(&count)
	Paginate(uint(page), uint(pageSize))(query).Find(&filters)

	meshplayFilterPage := &MeshplayFilterPage{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: int(count),
		Filters:    filters,
	}

	return marshalMeshplayFilterPage(meshplayFilterPage), nil
}

// GetMeshplayCatalogFilters returns all of the published filters
func (mfp *MeshplayFilterPersister) GetMeshplayCatalogFilters(page, pageSize, search, order string) ([]byte, error) {
	var err error
	order = SanitizeOrderInput(order, []string{"created_at", "updated_at", "name"})

	if order == "" {
		order = "updated_at desc"
	}

	var pg int
	if page != "" {
		pg, err = strconv.Atoi(page)

		if err != nil || pg < 0 {
			pg = 0
		}
	} else {
		pg = 0
	}

	// 0 page size is for all records
	var pgSize int
	if pageSize != "" {
		pgSize, err = strconv.Atoi(pageSize)

		if err != nil || pgSize < 0 {
			pgSize = 0
		}
	} else {
		pgSize = 0
	}

	filters := []MeshplayFilter{}

	query := mfp.DB.Where("visibility = ?", Published).Order(order)

	if search != "" {
		like := "%" + strings.ToLower(search) + "%"
		query = query.Where("(lower(meshplay_filters.name) like ?)", like)
	}

	var count int64
	err = query.Model(&MeshplayFilter{}).Count(&count).Error

	if err != nil {
		return nil, err
	}

	if pgSize != 0 {
		Paginate(uint(pg), uint(pgSize))(query).Find(&filters)
	} else {
		query.Find(&filters)
	}

	response := FiltersAPIResponse{
		Page:       uint(pg),
		PageSize:   uint(pgSize),
		TotalCount: uint(count),
		Filters:    filters,
	}

	marshalledResponse, _ := json.Marshal(response)
	return marshalledResponse, nil
}

// CloneMeshplayFilter clones meshplay filter to private
func (mfp *MeshplayFilterPersister) CloneMeshplayFilter(filterID string, cloneFilterRequest *MeshplayCloneFilterRequestBody) ([]byte, error) {
	var meshplayFilter MeshplayFilter
	filterUUID, _ := uuid.FromString(filterID)
	err := mfp.DB.First(&meshplayFilter, filterUUID).Error
	if err != nil || *meshplayFilter.ID == uuid.Nil {
		return nil, fmt.Errorf("unable to get filter: %w", err)
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	meshplayFilter.Visibility = Private
	meshplayFilter.ID = &id
	meshplayFilter.Name = cloneFilterRequest.Name

	return mfp.SaveMeshplayFilter(&meshplayFilter)
}

// DeleteMeshplayFilter takes in a profile id and delete it if it already exists
func (mfp *MeshplayFilterPersister) DeleteMeshplayFilter(id uuid.UUID) ([]byte, error) {
	filter := MeshplayFilter{ID: &id}
	mfp.DB.Delete(&filter)

	return marshalMeshplayFilter(&filter), nil
}

func (mfp *MeshplayFilterPersister) SaveMeshplayFilter(filter *MeshplayFilter) ([]byte, error) {
	if filter.Visibility == "" {
		filter.Visibility = Private
	}
	if filter.ID == nil {
		id, err := uuid.NewV4()
		if err != nil {
			return nil, ErrGenerateUUID(err)
		}

		filter.ID = &id
	}

	return marshalMeshplayFilters([]MeshplayFilter{*filter}), mfp.DB.Save(filter).Error
}

// SaveMeshplayFilters batch inserts the given filters
func (mfp *MeshplayFilterPersister) SaveMeshplayFilters(filters []MeshplayFilter) ([]byte, error) {
	finalFilters := []MeshplayFilter{}
	nilUserID := ""
	for _, filter := range filters {
		if filter.Visibility == "" {
			filter.Visibility = Private
		}
		filter.UserID = &nilUserID
		if filter.ID == nil {
			id, err := uuid.NewV4()
			if err != nil {
				return nil, ErrGenerateUUID(err)
			}

			filter.ID = &id
		}

		finalFilters = append(finalFilters, filter)
	}

	return marshalMeshplayFilters(finalFilters), mfp.DB.Create(finalFilters).Error
}

func (mfp *MeshplayFilterPersister) GetMeshplayFilter(id uuid.UUID) ([]byte, error) {
	var meshplayFilter MeshplayFilter

	err := mfp.DB.First(&meshplayFilter, id).Error
	return marshalMeshplayFilter(&meshplayFilter), err
}

func (mfp *MeshplayFilterPersister) GetMeshplayFilterFile(id uuid.UUID) ([]byte, error) {
	var meshplayFilter MeshplayFilter

	err := mfp.DB.First(&meshplayFilter, id).Error
	return []byte(meshplayFilter.FilterFile), err
}

func marshalMeshplayFilterPage(mfp *MeshplayFilterPage) []byte {
	res, _ := json.Marshal(mfp)

	return res
}

func marshalMeshplayFilter(mf *MeshplayFilter) []byte {
	res, _ := json.Marshal(mf)

	return res
}

func marshalMeshplayFilters(mps []MeshplayFilter) []byte {
	res, _ := json.Marshal(mps)

	return res
}
