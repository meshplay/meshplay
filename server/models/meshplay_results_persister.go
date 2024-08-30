package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshkit/database"
	"github.com/khulnasoft/meshkit/logger"
)

type MeshplayResultsPersister struct {
	DB *database.Handler
}

// MeshplayResultPage - represents a page of meshplay results
type MeshplayResultPage struct {
	Page       uint64           `json:"page"`
	PageSize   uint64           `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Results    []*MeshplayResult `json:"results"`
}

type localMeshplayResultDBRepresentation struct {
	ID                 uuid.UUID  `json:"meshplay_id,omitempty"`
	Name               string     `json:"name,omitempty"`
	Mesh               string     `json:"mesh,omitempty"`
	PerformanceProfile *uuid.UUID `json:"performance_profile,omitempty"`
	Result             []byte     `json:"runner_results,omitempty" gorm:"type:JSONB"`

	ServerMetrics     interface{} `json:"server_metrics,omitempty" gorm:"type:JSONB"`
	ServerBoardConfig interface{} `json:"server_board_config,omitempty" gorm:"type:JSONB"`

	TestStartTime          *time.Time         `json:"test_start_time,omitempty"`
	PerformanceProfileInfo PerformanceProfile `json:"-" gorm:"constraint:OnDelete:SET NULL;foreignKey:PerformanceProfile"`
}

func (mrp *MeshplayResultsPersister) GetResults(page, pageSize uint64, profileID string, log logger.Handler) ([]byte, error) {
	var res []*localMeshplayResultDBRepresentation
	var count int64
	query := mrp.DB.Where("performance_profile = ?", profileID)

	err := query.Table("meshplay_results").Count(&count).Error
	if err != nil {
		return nil, err
	}
	err = Paginate(uint(page), uint(pageSize))(query).Find(&res).Error

	resultPage := &MeshplayResultPage{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: int(count),
		Results:    convertLocalRepresentationSliceToMeshplayResultSlice(res, log),
	}

	return marshalMeshplayResultsPage(resultPage), err
}

func (mrp *MeshplayResultsPersister) GetAllResults(page, pageSize uint64, log logger.Handler) ([]byte, error) {
	var res []*localMeshplayResultDBRepresentation
	var count int64
	query := mrp.DB.Table("meshplay_results")

	err := query.Count(&count).Error
	if err != nil {
		return nil, err
	}
	err = Paginate(uint(page), uint(pageSize))(query).Find(&res).Error

	resultPage := &MeshplayResultPage{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: int(count),
		Results:    convertLocalRepresentationSliceToMeshplayResultSlice(res, log),
	}

	return marshalMeshplayResultsPage(resultPage), err
}

func (mrp *MeshplayResultsPersister) GetResult(key uuid.UUID, log logger.Handler) (*MeshplayResult, error) {
	var lres localMeshplayResultDBRepresentation

	err := mrp.DB.Table("meshplay_results").Find(&lres).Where("id = ?", key).Error
	res := convertLocalRepresentationToMeshplayResult(&lres, log)
	return res, err
}

func (mrp *MeshplayResultsPersister) WriteResult(key uuid.UUID, result []byte) error {
	var data MeshplayResult
	if err := json.Unmarshal(result, &data); err != nil {
		return err
	}

	data.ID = key

	t := time.Now()
	data.TestStartTime = &t
	return mrp.DB.Table("meshplay_results").Save(convertMeshplayResultToLocalRepresentation(&data)).Error
}

func marshalMeshplayResultsPage(mrp *MeshplayResultPage) []byte {
	res, _ := json.Marshal(mrp)

	return res
}

func convertLocalRepresentationSliceToMeshplayResultSlice(local []*localMeshplayResultDBRepresentation, log logger.Handler) (res []*MeshplayResult) {
	for _, val := range local {
		res = append(res, convertLocalRepresentationToMeshplayResult(val, log))
	}

	return
}

func convertLocalRepresentationToMeshplayResult(local *localMeshplayResultDBRepresentation, log logger.Handler) *MeshplayResult {
	var jsonmap map[string]interface{}
	if err := json.Unmarshal(local.Result, &jsonmap); err != nil {
		err = ErrUnmarshal(err, "MeshplayResult")
		log.Error(err)
		return nil
	}

	res := &MeshplayResult{
		ID:                 local.ID,
		Name:               local.Name,
		Mesh:               local.Mesh,
		PerformanceProfile: local.PerformanceProfile,
		Result:             jsonmap,
		ServerMetrics:      local.ServerMetrics,
		ServerBoardConfig:  local.ServerMetrics,
		TestStartTime:      local.TestStartTime,
	}

	return res
}

func convertMeshplayResultToLocalRepresentation(mr *MeshplayResult) *localMeshplayResultDBRepresentation {
	byt, _ := json.Marshal(mr.Result)

	res := &localMeshplayResultDBRepresentation{
		ID:                 mr.ID,
		Name:               mr.Name,
		Mesh:               mr.Mesh,
		PerformanceProfile: mr.PerformanceProfile,
		Result:             byt,
		ServerMetrics:      mr.ServerMetrics,
		ServerBoardConfig:  mr.ServerMetrics,
		TestStartTime:      mr.TestStartTime,
	}

	return res
}
