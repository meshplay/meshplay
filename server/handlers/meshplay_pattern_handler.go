package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	guid "github.com/google/uuid"
	"github.com/gorilla/mux"
	helpers "github.com/khulnasoft/meshplay/server/helpers/utils"
	isql "github.com/khulnasoft/meshplay/server/internal/sql"
	"github.com/khulnasoft/meshplay/server/meshes"
	"github.com/khulnasoft/meshplay/server/models"
	pCore "github.com/khulnasoft/meshplay/server/models/pattern/core"
	"github.com/khulnasoft/meshplay/server/models/pattern/resource/selector"
	patternutils "github.com/khulnasoft/meshplay/server/models/pattern/utils"
	"github.com/khulnasoft/meshkit/encoding"
	"github.com/khulnasoft/meshkit/errors"
	_errors "github.com/pkg/errors"

	"github.com/khulnasoft/meshkit/logger"
	"github.com/khulnasoft/meshkit/models/catalog/v1alpha1"
	"github.com/khulnasoft/meshkit/models/events"
	meshmodel "github.com/khulnasoft/meshkit/models/meshmodel/registry"
	"github.com/khulnasoft/meshkit/models/oci"
	"github.com/khulnasoft/meshkit/utils"
	"github.com/khulnasoft/meshkit/utils/catalog"
	"github.com/khulnasoft/meshkit/utils/kubernetes"
	"github.com/khulnasoft/meshkit/utils/kubernetes/kompose"
	"github.com/khulnasoft/meshkit/utils/walker"

	regv1beta1 "github.com/khulnasoft/meshkit/models/meshmodel/registry/v1beta1"
	"github.com/meshplay/schemas/models/v1alpha2"
	"github.com/meshplay/schemas/models/v1beta1/component"
	"github.com/meshplay/schemas/models/v1beta1/connection"
	"github.com/meshplay/schemas/models/v1beta1/pattern"
	"gopkg.in/yaml.v2"
)

// MeshplayPatternRequestBody refers to the type of request body that
// SaveMeshplayPattern would receive
type MeshplayPatternPOSTRequestBody struct {
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url,omitempty"`
	Path          string                 `json:"path,omitempty"`
	Save          bool                   `json:"save,omitempty"`
	PatternData   *meshplayPatternPayload `json:"pattern_data,omitempty"`
	CytoscapeJSON string                 `json:"cytoscape_json,omitempty"`
}

type MeshplayPatternUPDATERequestBody struct {
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url,omitempty"`
	Path          string                 `json:"path,omitempty"`
	Save          bool                   `json:"save,omitempty"`
	PatternData   *models.MeshplayPattern `json:"pattern_data,omitempty"`
	CytoscapeJSON string                 `json:"cytoscape_json,omitempty"`
}

type meshplayPatternPayload struct {
	ID *uuid.UUID `json:"id,omitempty"`

	Name        string `json:"name,omitempty"`
	PatternFile []byte `json:"pattern_file"`
	FileName    string `json:"file_name"`
	// Meshplay doesn't have the user id fields
	// but the remote provider is allowed to provide one
	UserID *string `json:"user_id"`

	Location      isql.Map             `json:"location"`
	Visibility    string               `json:"visibility"`
	CatalogData   v1alpha1.CatalogData `json:"catalog_data,omitempty"`
	Type          sql.NullString       `json:"type"`
	SourceContent []byte               `json:"source_content"`
}

// PatternFileRequestHandler will handle requests of both type GET and POST
// on the route /api/pattern
func (h *Handler) PatternFileRequestHandler(
	rw http.ResponseWriter,
	r *http.Request,
	prefObj *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	if r.Method == http.MethodGet {
		h.GetMeshplayPatternsHandler(rw, r, prefObj, user, provider)
		return
	}

	if r.Method == http.MethodPost {
		h.handlePatternPOST(rw, r, prefObj, user, provider)
		return
	}
}

// swagger:route POST /api/pattern PatternsAPI idPostPatternFile
// Handle POST requests for patterns
//
// Edit/update a meshplay pattern
// responses:
// 	200: meshplayPatternResponseWrapper

func (h *Handler) handlePatternPOST(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	var err error
	action := models.Create
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("create").ActedUpon(userID).WithSeverity(events.Informational)

	sourcetype := mux.Vars(r)["sourcetype"]
	parsedBody := &MeshplayPatternPOSTRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrRequestBody(err),
		}).WithDescription("Unable to parse uploaded design.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	actedUpon := &userID
	if parsedBody.PatternData != nil && parsedBody.PatternData.ID != nil {
		actedUpon = parsedBody.PatternData.ID
		action = models.Update
	}

	eventBuilder.ActedUpon(*actedUpon)

	token, err := provider.GetProviderToken(r)
	if err != nil {
		h.log.Error(ErrRetrieveUserToken(err))
		http.Error(rw, ErrRetrieveUserToken(err).Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Critical).WithMetadata(map[string]interface{}{
			"error": ErrRetrieveUserToken(err),
		}).WithDescription("No auth token provided in the request.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}

	format := r.URL.Query().Get("output")
	meshplayPattern := &models.MeshplayPattern{} // pattern to be saved in the database

	// If Content is not empty then assume it's a local upload
	if parsedBody.PatternData != nil {
		meshplayPattern.Name = parsedBody.PatternData.Name
		// Assign a location if no location is specified
		if parsedBody.PatternData.Location == nil {
			parsedBody.PatternData.Location = map[string]interface{}{
				"host":   "",
				"path":   "",
				"type":   "local",
				"branch": "",
			}
		}

		bytPattern := parsedBody.PatternData.PatternFile
		fileName := parsedBody.PatternData.FileName
		meshplayPattern.SourceContent = bytPattern
		if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) {
			var k8sres string
			if sourcetype == string(models.DockerCompose) {
				k8sres, err = kompose.Convert(bytPattern) // convert the docker compose file into kubernetes manifest
				if err != nil {
					h.log.Error(ErrConvertingDockerComposeToDesign(err))
					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": ErrConvertingDockerComposeToDesign(err),
					}).WithDescription(fmt.Sprintf("Failed to convert Docker Compose application \"%s\"", meshplayPattern.Name)).Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)

					http.Error(rw, ErrConvertingDockerComposeToDesign(err).Error(), http.StatusInternalServerError)

					return
				}
				meshplayPattern.Type = sql.NullString{
					String: string(models.DockerCompose),
					Valid:  true,
				}
			} else if sourcetype == string(models.K8sManifest) {
				k8sres = string(bytPattern)
				meshplayPattern.Type = sql.NullString{
					String: string(models.K8sManifest),
					Valid:  true,
				}
			}
			patternFile, err := pCore.NewPatternFileFromK8sManifest(k8sres, fileName, false, h.registryManager)
			if err != nil {
				h.log.Error(ErrConvertingK8sManifestToDesign(err))
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrConvertingK8sManifestToDesign(err),
				}).WithDescription(fmt.Sprintf("Failed converting %s \"%s\" to design file format.", sourcetype, meshplayPattern.Name)).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				http.Error(rw, ErrConvertingK8sManifestToDesign(err).Error(), http.StatusInternalServerError)

				return
			}

			pfByt, _ := encoding.Marshal(patternFile)
			meshplayPattern.PatternFile = string(pfByt)
		} else {
			patternFile := &pattern.PatternFile{}
			var err error
			patternFileStr := string(parsedBody.PatternData.PatternFile)
			isOldFormat, err := patternutils.IsDesignInAlpha2Format(string(parsedBody.PatternData.PatternFile))
			if err != nil {
				err = ErrPatternFile(err)
				event := events.NewEvent().ActedUpon(*actedUpon).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("view").WithDescription(fmt.Sprintf("Failed to parse design \"%s\".", parsedBody.PatternData.ID)).WithMetadata(map[string]interface{}{"error": err, "id": *actedUpon}).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				h.log.Error(err)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			if isOldFormat {
				eventBuilder := events.NewEvent().ActedUpon(*actedUpon).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("convert")
				patternFile, patternFileStr, err = h.convertV1alpha2ToV1beta1(&models.MeshplayPattern{
					ID:          actedUpon,
					Name:        parsedBody.PatternData.Name,
					PatternFile: string(parsedBody.PatternData.PatternFile),
				}, eventBuilder)
				event := eventBuilder.Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				if err != nil {
					h.log.Error(err)
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			meshplayPattern.PatternFile = patternFileStr
			meshplayPattern.CatalogData = parsedBody.PatternData.CatalogData

			if parsedBody.PatternData.ID != nil {
				meshplayPattern.ID = parsedBody.PatternData.ID
			}

			// assume the design is in OCI Artifact format
			uncompressedDesign, err := unCompressOCIArtifactIntoDesign(parsedBody.PatternData.PatternFile)
			// if errors occurs in decompressing OCI Artifact into design file
			// then fall back to importing design as yaml file
			if err != nil {
				h.log.Info("Falling back to importing design as yaml file")
			} else {
				h.log.Info("OCI Artifact decompressed.")
				event := eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("OCI Artifact decompressed into %s design file", meshplayPattern.Name)).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				meshplayPattern = uncompressedDesign
			}
			meshplayPattern.Type = sql.NullString{
				String: string(models.Design),
				Valid:  true,
			}
			// Check if the pattern is valid
			// Replace with/add validation

			// Assign a name if no name is provided
			if parsedBody.PatternData.Name == "" {
				meshplayPattern.Name = patternFile.Name
			}

			if parsedBody.PatternData.Visibility != "" {
				meshplayPattern.Visibility = parsedBody.PatternData.Visibility
			}

			if parsedBody.Save {
				resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
				if err != nil {
					h.log.Error(ErrSavePattern(err))
					http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": ErrSavePattern(err),
					}).WithDescription(ErrSavePattern(err).Error()).Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)
					return
				}

				h.formatPatternOutput(rw, resp, format, sourcetype, eventBuilder, parsedBody.URL, action)
				event := eventBuilder.Build()
				_ = provider.PersistEvent(event)
				// Create the event but do not notify the client immediately, as the evaluations are frequent and takes up the view area.
				// go h.config.EventBroadcaster.Publish(userID, event)
				go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
				return
			}

			byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
			if err != nil {
				h.log.Error(ErrEncodePattern(err))
				http.Error(rw, ErrEncodePattern(err).Error(), http.StatusInternalServerError)

				return
			}

			h.formatPatternOutput(rw, byt, format, sourcetype, eventBuilder, parsedBody.URL, action)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}
	}

	if parsedBody.URL != "" {
		latestKuberVersion := getLatestKubeVersionFromRegistry(h.registryManager)
		if sourcetype == string(models.HelmChart) {
			helmSourceResp, err := http.Get(parsedBody.URL)
			defer func() {
				if helmSourceResp == nil {
					return
				}
				_ = helmSourceResp.Body.Close()
			}()
			if err != nil {
				obj := "import"
				importErr := ErrApplicationFailure(err, obj)
				h.log.Error(importErr)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": importErr,
				}).WithDescription(fmt.Sprintf("Failed importing design from URL %s.", parsedBody.URL)).Build()
				_ = provider.PersistEvent(event)

				go h.config.EventBroadcaster.Publish(userID, event)
				http.Error(rw, importErr.Error(), http.StatusInternalServerError)

				return
			}
			sourceContent, err := io.ReadAll(helmSourceResp.Body)
			if err != nil {
				http.Error(rw, "error read body", http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSaveApplication(err),
				}).WithDescription(fmt.Sprintf("error reading design from the remote URL %s, URL is malformed or not reachable.", parsedBody.URL)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				return
			}

			resp, err := kubernetes.ConvertHelmChartToK8sManifest(kubernetes.ApplyHelmChartConfig{
				URL:               parsedBody.URL,
				KubernetesVersion: latestKuberVersion,
			})
			if err != nil {
				h.log.Error(ErrConvertingHelmChartToDesign(err))
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrConvertingHelmChartToDesign(err),
				}).WithDescription(fmt.Sprintf("Failed converting Helm Chart %s to K8s Manifest.", parsedBody.URL)).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				http.Error(rw, ErrConvertingHelmChartToDesign(err).Error(), http.StatusInternalServerError)

				return
			}

			fileName := strings.TrimSuffix(path.Base(parsedBody.URL), filepath.Ext(path.Base(parsedBody.URL)))

			result := string(resp)
			pattern, err := pCore.NewPatternFileFromK8sManifest(result, fileName, false, h.registryManager)
			if err != nil {
				h.log.Error(ErrConvertingHelmChartToDesign(err))
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrConvertingHelmChartToDesign(err),
				}).WithDescription(fmt.Sprintf("Failed converting Helm Chart %s to K8s Manifest.", parsedBody.URL)).Build()
				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				http.Error(rw, ErrConvertingHelmChartToDesign(err).Error(), http.StatusInternalServerError)

				return
			}
			bytPattern, _ := encoding.Marshal(pattern)

			meshplayPattern = &models.MeshplayPattern{
				Name:        parsedBody.Name,
				PatternFile: string(bytPattern),
				Type: sql.NullString{
					String: string(models.HelmChart),
					Valid:  true,
				},
				Location: map[string]interface{}{
					"type":   "http",
					"host":   parsedBody.URL,
					"path":   "",
					"branch": "",
				},
				SourceContent: sourceContent,
			}
		} else if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) {
			parsedURL, err := url.Parse(parsedBody.URL)
			if err != nil {
				err := ErrSaveApplication(fmt.Errorf("error parsing URL"))
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": err,
				}).WithDescription(fmt.Sprintf("Invalid URL provided %s", parsedBody.URL)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			var pfs []models.MeshplayPattern
			// Check if hostname is github
			if parsedURL.Host == "github.com" {
				parsedPath := strings.Split(parsedURL.Path, "/")
				if parsedPath[3] == "tree" {
					parsedPath = append(parsedPath[0:3], parsedPath[4:]...)
				}
				if len(parsedPath) < 3 {
					http.Error(rw, "malformed URL: url should be of type github.com/<owner>/<repo>/[branch]", http.StatusNotAcceptable)
				}

				owner := parsedPath[1]
				repo := parsedPath[2]
				branch := "master"
				path := parsedBody.Path
				if len(parsedPath) == 4 {
					branch = parsedPath[3]
				}
				if path == "" && len(parsedPath) > 4 {
					path = strings.Join(parsedPath[4:], "/")
				}

				pfs, err = githubRepoDesignScan(owner, repo, path, branch, sourcetype, h.registryManager)
				if err != nil {
					remoteApplicationErr := ErrRemoteApplication(err)
					http.Error(rw, remoteApplicationErr.Error(), http.StatusInternalServerError)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": err,
					}).WithDescription(fmt.Sprintf("Failed to retrieve remote design at %s", parsedBody.URL)).Build()

					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)

					return
				}
				meshplayPattern = &pfs[0]
			} else {
				// Fallback to generic HTTP import
				pfs, err = genericHTTPDesignFile(parsedBody.URL, parsedBody.Name, sourcetype, h.registryManager, h.log)
				if err != nil {
					remoteApplicationErr := ErrRemoteApplication(err)
					http.Error(rw, remoteApplicationErr.Error(), http.StatusInternalServerError)

					event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
						"error": err,
					}).WithDescription(fmt.Sprintf("Failed to retrieve remote design at %s", parsedBody.URL)).Build()
					_ = provider.PersistEvent(event)
					go h.config.EventBroadcaster.Publish(userID, event)

					return
				}
				meshplayPattern = &pfs[0]
			}
			if parsedBody.Name != "" {
				meshplayPattern.Name = parsedBody.Name
			}
		} else {

			if parsedBody.PatternData == nil {
				parsedBody.PatternData = &meshplayPatternPayload{}
			}
			parsedBody.PatternData.Type = sql.NullString{
				String: string(models.Design),
				Valid:  true,
			}
			result, err := genericHTTPDesignFile(parsedBody.URL, parsedBody.Name, sourcetype, h.registryManager, h.log)

			if err != nil {
				h.log.Error(ErrImportPattern(err))
				http.Error(rw, ErrImportPattern(err).Error(), http.StatusInternalServerError)
				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrImportPattern(err),
				}).WithDescription(ErrImportPattern(err).Error()).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}

			resp, err := provider.SaveMeshplayPattern(token, &result[0])
			if err != nil {
				h.log.Error(ErrSavePattern(err))
				http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": ErrSavePattern(err),
				}).WithDescription(ErrSavePattern(err).Error()).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)
				return
			}
			h.formatPatternOutput(rw, resp, format, sourcetype, eventBuilder, parsedBody.URL, action)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}
	}

	if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) || sourcetype == string(models.HelmChart) {
		var savedPatternID *uuid.UUID

		if parsedBody.Save {
			resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
			if err != nil {
				obj := "save"

				saveErr := ErrApplicationFailure(err, obj)
				h.log.Error(saveErr)
				http.Error(rw, saveErr.Error(), http.StatusInternalServerError)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": saveErr,
				}).WithDescription(fmt.Sprintf("Failed persisting design %s", parsedBody.Name)).Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				return
			}

			h.formatPatternOutput(rw, resp, format, sourcetype, eventBuilder, parsedBody.URL, action)

			eventBuilder.WithSeverity(events.Informational)
			event := eventBuilder.Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			var meshplayPatternContent []models.MeshplayPattern
			err = json.Unmarshal(resp, &meshplayPatternContent)
			if err != nil {
				obj := "pattern"
				h.log.Error(models.ErrEncoding(err, obj))
				http.Error(rw, models.ErrEncoding(err, obj).Error(), http.StatusInternalServerError)
				return
			}
			savedPatternID = meshplayPatternContent[0].ID
			err = provider.SaveMeshplayPatternSourceContent(token, (savedPatternID).String(), meshplayPattern.SourceContent)

			if err != nil {
				obj := "upload"
				uploadSourceContentErr := ErrApplicationSourceContent(err, obj)

				h.log.Error(uploadSourceContentErr)
				http.Error(rw, uploadSourceContentErr.Error(), http.StatusInternalServerError)

				event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
					"error": uploadSourceContentErr,
				}).WithDescription("Failed uploading original design content to remote provider.").Build()

				_ = provider.PersistEvent(event)
				go h.config.EventBroadcaster.Publish(userID, event)

				return
			}
			go h.config.PatternChannel.Publish(userID, struct{}{})
			event = eventBuilder.WithDescription(fmt.Sprintf("Design %s source content uploaded", meshplayPatternContent[0].Name)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}

		meshplayPattern.ID = savedPatternID
		byt, err := json.Marshal([]models.MeshplayPattern{*meshplayPattern})
		if err != nil {
			obj := "design"
			h.log.Error(models.ErrEncoding(err, obj))
			http.Error(rw, models.ErrEncoding(err, obj).Error(), http.StatusInternalServerError)
			return
		}

		h.formatPatternOutput(rw, byt, format, sourcetype, eventBuilder, parsedBody.URL, action)
		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
	}

}

// Verifies and converts a pattern to design format if required.
// A pattern is required to be converted to design format iff,
// 1. pattern_file attribute is empty, and
// 2. The "type" (sourcetype/original content) is not Design. [is one of compose/helmchart/manifests]

func (h *Handler) VerifyAndConvertToDesign(
	ctx context.Context,
	meshplayPattern *models.MeshplayPattern,
	provider models.Provider,
) error {

	if meshplayPattern.Type.Valid && meshplayPattern.Type.String != string(models.Design) && meshplayPattern.PatternFile == "" {
		token, _ := ctx.Value(models.TokenCtxKey).(string)

		sourceContent, err := provider.GetDesignSourceContent(token, meshplayPattern.ID.String())
		if err != nil {
			return err
		}

		meshplayPattern.SourceContent = sourceContent
		sourcetype := meshplayPattern.Type.String

		if sourcetype == string(models.DockerCompose) || sourcetype == string(models.K8sManifest) {
			var k8sres string
			if sourcetype == string(models.DockerCompose) {
				k8sres, err = kompose.Convert(sourceContent) // convert the docker compose file into kubernetes manifest
				if err != nil {
					err = ErrConvertingDockerComposeToDesign(err)
					return err
				}

			} else if sourcetype == string(models.K8sManifest) {
				k8sres = string(sourceContent)
			}
			pattern, err := pCore.NewPatternFileFromK8sManifest(k8sres, "", false, h.registryManager)
			if err != nil {
				err = ErrConvertingK8sManifestToDesign(err)
				return err
			}
			bytPattern, _ := yaml.Marshal(pattern)
			meshplayPattern.PatternFile = string(bytPattern)
		}

		resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
		if err != nil {
			obj := "save"
			saveErr := ErrApplicationFailure(err, obj)
			return saveErr
		}

		contentMeshplayPatternSlice := make([]models.MeshplayPattern, 0)

		if err := json.Unmarshal(resp, &contentMeshplayPatternSlice); err != nil {
			return models.ErrUnmarshal(err, "pattern")
		}
	}

	return nil
}

func unCompressOCIArtifactIntoDesign(artifact []byte) (*models.MeshplayPattern, error) {

	// Assume design is in OCI Tarball Format
	tmpDir, err := oci.CreateTempOCIContentDir()
	if err != nil {
		return nil, ErrCreateDir(err, "OCI")
	}
	defer os.RemoveAll(tmpDir)

	tmpInputDesignFile := filepath.Join(tmpDir, "design.tar")
	file, err := os.Create(tmpInputDesignFile)
	if err != nil {
		return nil, ErrCreateFile(err, tmpInputDesignFile)
	}
	defer file.Close()

	reader := bytes.NewReader(artifact)
	if _, err := io.Copy(file, reader); err != nil {
		return nil, ErrWritingIntoFile(err, tmpInputDesignFile)
	}

	tmpOutputDesignFile := filepath.Join(tmpDir, "output")
	// Extract the tarball
	if err := oci.UnCompressOCIArtifact(tmpInputDesignFile, tmpOutputDesignFile); err != nil {
		return nil, ErrUnCompressOCIArtifact(err)
	}

	files, err := walker.WalkLocalDirectory(tmpOutputDesignFile)
	if err != nil {
		return nil, ErrWaklingLocalDirectory(err)
	}

	// TODO: Add support to merge multiple designs into one
	// Currently, assumes to save only the first design
	if len(files) == 0 {
		return nil, ErrEmptyOCIImage(fmt.Errorf("No design file detected in the imported OCI image"))
	}
	design := files[0]

	var patternFile pattern.PatternFile

	err = encoding.Unmarshal([]byte(design.Content), &patternFile)
	if err != nil {
		return nil, ErrDecodePattern(err)
	}
	meshplayPattern := &models.MeshplayPattern{
		PatternFile: design.Content,
		Name:        design.Name,
	}

	return meshplayPattern, nil
}

func githubRepoDesignScan(
	owner,
	repo,
	path,
	branch,
	sourceType string,
	reg *meshmodel.RegistryManager,
) ([]models.MeshplayPattern, error) {
	var mu sync.Mutex
	ghWalker := walker.NewGit()
	result := make([]models.MeshplayPattern, 0)
	err := ghWalker.
		Owner(owner).
		Repo(repo).
		Branch(branch).
		Root(path).
		RegisterFileInterceptor(func(f walker.File) error {
			ext := filepath.Ext(f.Name)
			var k8sres string
			var err error
			k8sres = f.Content
			if ext == ".yml" || ext == ".yaml" {
				if sourceType == string(models.DockerCompose) {
					k8sres, err = kompose.Convert([]byte(f.Content))
					if err != nil {
						return ErrRemoteApplication(err)
					}
				}
				pattern, err := pCore.NewPatternFileFromK8sManifest(k8sres, "", false, reg)
				if err != nil {
					return err //always a meshkit error
				}

				patternByt, _ := encoding.Marshal(pattern)

				af := models.MeshplayPattern{
					Name:        strings.TrimSuffix(f.Name, ext),
					PatternFile: string(patternByt),
					Location: map[string]interface{}{
						"type":   "github",
						"host":   fmt.Sprintf("github.com/%s/%s", owner, repo),
						"path":   f.Path,
						"branch": branch,
					},
					Type: sql.NullString{
						String: string(sourceType),
						Valid:  true,
					},
					SourceContent: []byte(f.Content),
				}

				mu.Lock()
				result = append(result, af)
				mu.Unlock()
			}

			return nil
		}).
		Walk()

	return result, ErrRemoteApplication(err)
}

// Always returns a meshplay pattern slice of length 1 otherwise an error is returned
func genericHTTPDesignFile(fileURL, patternName, sourceType string, reg *meshmodel.RegistryManager, log logger.Handler) ([]models.MeshplayPattern, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, ErrRemoteApplication(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrRemoteApplication(fmt.Errorf("file not found"))
	}

	defer models.SafeClose(resp.Body, log)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrRemoteApplication(err)
	}

	res := string(body)

	if sourceType == string(models.DockerCompose) {
		res, err = kompose.Convert(body)
		if err != nil {
			return nil, ErrRemoteApplication(err)
		}
	}

	var pattern pattern.PatternFile
	if sourceType == string(models.DockerCompose) || sourceType == string(models.K8sManifest) {
		var err error
		pattern, err = pCore.NewPatternFileFromK8sManifest(res, "", false, reg)
		if err != nil {
			return nil, err //This error is already a meshkit error
		}
	} else {
		err := encoding.Unmarshal([]byte(res), &pattern)
		if err != nil {
			return nil, ErrDecodePattern(err)
		}
	}

	if patternName != "" {
		pattern.Name = patternName
	}

	patternByt, _ := encoding.Marshal(pattern)

	url := strings.Split(fileURL, "/")
	af := models.MeshplayPattern{
		Name:        url[len(url)-1],
		PatternFile: string(patternByt),
		Location: map[string]interface{}{
			"type":   "http",
			"host":   fileURL,
			"path":   "",
			"branch": "",
		},
		Type: sql.NullString{
			String: string(sourceType),
			Valid:  true,
		},
		SourceContent: body,
	}
	return []models.MeshplayPattern{af}, nil
}

// swagger:route GET /api/pattern PatternsAPI idGetPatternFiles
// Handle GET request for patterns
//
// Returns the list of all the patterns saved by the current user
// This will return all the patterns with their details
//
// ```?order={field}``` orders on the passed field
//
// ```?search=<design name>``` A string matching is done on the specified design name
//
// ```?page={page-number}``` Default page number is 1
//
// ```?pagesize={pagesize}``` Default pagesize is 10
//
// ```?visibility={[visibility]}``` Default visibility is public + private; Mulitple visibility filters can be passed as an array
// Eg: ```?visibility=["public", "published"]``` will return public and published designs
//
// ```?metrics``` Returns metrics like deployment/share/clone/view/download count for desings, default is false,
//
// responses:
//
//	200: meshplayPatternsResponseWrapper
func (h *Handler) GetMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	q := r.URL.Query()
	tokenString := r.Context().Value(models.TokenCtxKey).(string)
	updateAfter := q.Get("updated_after")
	includeMetrics := q.Get("metrics")
	err := r.ParseForm() // necessary to get r.Form["visibility"], i.e, ?visibility=public&visbility=private
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}
	filter := struct {
		Visibility []string `json:"visibility"`
	}{}

	visibility := q.Get("visibility")
	if visibility != "" {
		err := json.Unmarshal([]byte(visibility), &filter.Visibility)
		if err != nil {
			h.log.Error(ErrFetchPattern(err))
			http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	resp, err := provider.GetMeshplayPatterns(tokenString, q.Get("page"), q.Get("pagesize"), q.Get("search"), q.Get("order"), updateAfter, filter.Visibility, includeMetrics)
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	// mc := NewContentModifier(token, provider, prefObj, user.UserID)
	// //acts like a middleware, modifying the bytes lazily just before sending them back
	// err = mc.AddMetadataForPatterns(r.Context(), &resp)
	// if err != nil {
	// 	fmt.Println("Could not add metadata about pattern's current support ", err.Error())
	// }
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/catalog PatternsAPI idGetCatalogMeshplayPatternsHandler
// Handle GET request for catalog patterns
//
// # Patterns can be further filtered through query parameter
//
// ```?order={field}``` orders on the passed field
//
// ```?page={page-number}``` Default page number is 0
//
// ```?pagesize={pagesize}``` Default pagesize is 10.
//
// ```?search={patternname}``` If search is non empty then a greedy search is performed
//
// ```?metrics``` Returns metrics like deployment/share/clone/view/download count for desings, default false,
// responses:
//
//	200: meshplayPatternsResponseWrapper
func (h *Handler) GetCatalogMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	q := r.URL.Query()
	tokenString := r.Context().Value(models.TokenCtxKey).(string)

	resp, err := provider.GetCatalogMeshplayPatterns(tokenString, q.Get("page"), q.Get("pagesize"), q.Get("search"), q.Get("order"), q.Get("metrics"))
	if err != nil {
		h.log.Error(ErrFetchPattern(err))
		http.Error(rw, ErrFetchPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/pattern/{id} PatternsAPI idDeleteMeshplayPattern
// Handle Delete for a Meshplay Pattern
//
// Deletes a meshplay pattern with ID: id
// responses:
//
//	200: noContentWrapper
//
// DeleteMeshplayPatternHandler deletes a pattern with the given id
func (h *Handler) DeleteMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("delete").ActedUpon(uuid.FromStringOrNil(patternID))

	meshplayPattern := models.MeshplayPattern{}

	resp, err := provider.DeleteMeshplayPattern(r, patternID)
	if err != nil {
		errPatternDelete := ErrDeletePattern(err)

		h.log.Error(errPatternDelete)
		http.Error(rw, errPatternDelete.Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": errPatternDelete,
		}).WithDescription("Error deleting pattern.").Build()
		http.Error(rw, errPatternDelete.Error(), http.StatusInternalServerError)
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		return
	}

	_ = json.Unmarshal(resp, &meshplayPattern)
	event := eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("Pattern %s deleted.", meshplayPattern.Name)).Build()
	_ = provider.PersistEvent(event)
	go h.config.EventBroadcaster.Publish(userID, event)
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/download/{id} PatternsAPI idGetMeshplayPattern
// Handle GET request for Meshplay Pattern with the given id
//
// ?oci={true|false} - If true, returns the pattern in OCI Artifact format
//
// Get the pattern with the given id
// responses:
//  200:

// GetMeshplayPatternHandler returns the pattern file with the given id

func (h *Handler) DownloadMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {

	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("download").ActedUpon(userID).WithSeverity(events.Informational)

	patternID := mux.Vars(r)["id"]
	ociFormat, _ := strconv.ParseBool(r.URL.Query().Get("oci"))
	ahpkg, _ := strconv.ParseBool(r.URL.Query().Get("pkg"))
	var unmarshalledPatternFile pattern.PatternFile

	resp, err := provider.GetMeshplayPattern(r, patternID, "false")
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrGetPattern(err),
		}).WithDescription(fmt.Sprintf("Failed to fetch design file for ID: %s.", patternID)).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}
	pattern := &models.MeshplayPattern{}
	err = json.Unmarshal(resp, &pattern)
	if err != nil {
		obj := "download pattern"
		h.log.Error(models.ErrUnmarshal(err, obj))
		http.Error(rw, models.ErrUnmarshal(err, obj).Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": models.ErrUnmarshal(err, obj),
		}).WithDescription(fmt.Sprintf("Failed to unmarshal design file for ID: %s.", patternID)).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}

	err = h.VerifyAndConvertToDesign(r.Context(), pattern, provider)
	if err != nil {
		event := events.NewEvent().ActedUpon(*pattern.ID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("convert").WithDescription(fmt.Sprintf("The \"%s\" is not in the design format, failed to convert and persist the original source content from \"%s\" to design file format", pattern.Name, pattern.Type.String)).WithMetadata(map[string]interface{}{"error": err}).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		h.log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// v1beta1
	isOldFormat, err := patternutils.IsDesignInAlpha2Format(pattern.PatternFile)
	if err != nil {
		err = ErrPatternFile(err)
		event := events.NewEvent().ActedUpon(*pattern.ID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("download").WithDescription(fmt.Sprintf("Failed to parse design \"%s\".", pattern.Name)).WithMetadata(map[string]interface{}{"error": err, "id": pattern.ID}).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		h.log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if isOldFormat {

		eventBuilder := events.NewEvent().ActedUpon(*pattern.ID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("convert")
		_, patternFileStr, err := h.convertV1alpha2ToV1beta1(pattern, eventBuilder)
		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		if err != nil {
			h.log.Error(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		pattern.PatternFile = patternFileStr
	}

	if ociFormat {
		tmpDir, err := oci.CreateTempOCIContentDir()
		if err != nil {
			h.log.Error(ErrCreateDir(err, "OCI"))
			http.Error(rw, ErrCreateDir(err, "OCI").Error(), http.StatusInternalServerError)

			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrCreateDir(err, "OCI"),
			}).WithDescription("Error creating tmp directory under ~/.meshplay/content/").Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}
		defer os.RemoveAll(tmpDir)

		tmpDesignFile := filepath.Join(tmpDir, "design.yml")
		file, err := os.Create(tmpDesignFile)
		if err != nil {
			h.log.Error(ErrCreateFile(err, tmpDesignFile))
			http.Error(rw, ErrCreateFile(err, tmpDesignFile).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrCreateFile(err, tmpDesignFile),
			}).WithDescription(fmt.Sprintf("Error creating tmp file %s", tmpDesignFile)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}
		defer file.Close()

		ymlDesign, err := yaml.Marshal(pattern.PatternFile)
		if err != nil {
			err = ErrEncodePattern(err)
			h.log.Error(err)
			http.Error(rw, fmt.Sprintf("Failed to export design \"%s\" as OCI image.", pattern.Name), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": err,
			}).WithDescription(fmt.Sprintf("Failed to export design \"%s\" as OCI image.", pattern.Name)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		if _, err := file.Write(ymlDesign); err != nil {
			h.log.Error(ErrWritingIntoFile(err, tmpDesignFile))
			http.Error(rw, ErrWritingIntoFile(err, tmpDesignFile).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrWritingIntoFile(err, tmpDesignFile),
			}).WithDescription(fmt.Sprintf("Error writing into tmp design file %s at %s", pattern.Name, tmpDesignFile)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		artifactHubPkgFilePath := filepath.Join(tmpDir, "artifacthub-pkg.yml")
		artifactHubPkgFile, err := os.Create(artifactHubPkgFilePath)
		if err != nil {
			h.log.Error(ErrCreateFile(err, "artifacthub-pkg.yml"))
			eb := *eventBuilder
			event := eb.WithSeverity(events.Error).WithDescription(fmt.Sprintf("Unable to create artifacthub pkg for the design \"%s\"", pattern.Name)).WithMetadata(map[string]interface{}{"error": err}).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
		}

		data, err := createArtifactHubPkg(pattern, strings.Trim(fmt.Sprintf("%s %s", user.FirstName, user.LastName), " "))
		if err != nil {
			h.log.Error(err)
			eb := *eventBuilder
			event := eb.WithSeverity(events.Error).WithDescription(fmt.Sprintf("Unable to create artifacthub pkg for the design \"%s\"", pattern.Name)).WithMetadata(map[string]interface{}{"error": err}).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
		}

		_, err = artifactHubPkgFile.Write(data)
		if err != nil {
			err = ErrWritingIntoFile(err, "artifacthub-pkg.yml")
			h.log.Error(err)
			eb := *eventBuilder
			event := eb.WithSeverity(events.Error).WithDescription(fmt.Sprintf("Unable to create artifacthub pkg for the design \"%s\"", pattern.Name)).WithMetadata(map[string]interface{}{"error": err}).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
		}

		ociImg, err := oci.BuildImage(tmpDir)
		if err != nil {
			h.log.Error(ErrBuildOCIImg(err))
			http.Error(rw, ErrBuildOCIImg(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrBuildOCIImg(err),
			}).WithDescription(fmt.Sprintf("Error building OCI Image from contents at %s", tmpDesignFile)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		digest, err := ociImg.Digest()
		if err != nil {
			h.log.Error(ErrBuildOCIImg(err))
			http.Error(rw, ErrBuildOCIImg(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrBuildOCIImg(err),
			}).WithDescription(fmt.Sprintf("Error getting image digest for %s", tmpDesignFile)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		size, err := ociImg.Size()
		if err != nil {
			h.log.Error(ErrBuildOCIImg(err))
			http.Error(rw, ErrBuildOCIImg(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrBuildOCIImg(err),
			}).WithDescription(fmt.Sprintf("Error getting calculating image size for %s", tmpDesignFile)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		h.log.Info(fmt.Sprintf("OCI Image built. Digest: %v, Size: %v", digest, size))

		eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("OCI Image built. Digest: %v, Size: %v", digest, size))
		event := eventBuilder.Build()
		go h.config.EventBroadcaster.Publish(userID, event)
		_ = provider.PersistEvent(event)

		pretifiedName := strings.ToLower(strings.Replace(pattern.Name, " ", "", -1)) // ensures that tag validation passes
		tmpOCITarFilePath := filepath.Join(tmpDir, pretifiedName+".tar")
		err = oci.SaveOCIArtifact(ociImg, tmpOCITarFilePath, pretifiedName)
		if err != nil {
			h.log.Error(ErrSaveOCIArtifact(err))
			http.Error(rw, ErrSaveOCIArtifact(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSaveOCIArtifact(err),
			}).WithDescription(fmt.Sprintf("Failed to save OCI Artifact %s temporarily", tmpOCITarFilePath)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		file, err = os.OpenFile(tmpOCITarFilePath, os.O_RDONLY, 0444)
		if err != nil {
			h.log.Error(ErrOpenFile(tmpOCITarFilePath))
			http.Error(rw, ErrOpenFile(tmpOCITarFilePath).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrOpenFile(tmpOCITarFilePath),
			}).WithDescription(fmt.Sprintf("Failed to read contents of OCI Artifact %s", tmpOCITarFilePath)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}
		content, err := io.ReadAll(file)
		if err != nil {
			h.log.Error(ErrIOReader(err))
			http.Error(rw, ErrIOReader(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrIOReader(err),
			}).WithDescription(fmt.Sprintf("Failed to read contents of OCI artifact %s", tmpOCITarFilePath)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}

		h.log.Info("OCI Artifact saved at: ", tmpOCITarFilePath)

		eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("OCI Artifact temporarily saved at: %s", tmpOCITarFilePath))
		event = eventBuilder.Build()
		go h.config.EventBroadcaster.Publish(userID, event)
		_ = provider.PersistEvent(event)

		rw.Header().Set("Content-Type", "application/tar")
		rw.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s.tar", pattern.Name))

		reader := bytes.NewReader(content)
		if _, err := io.Copy(rw, reader); err != nil {
			http.Error(rw, ErrIOReader(err).Error(), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrIOReader(err),
			}).WithDescription("Failed to save contents of OCI Artifact at requested path").Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}
		return
	}

	if ahpkg {
		rw.Header().Set("Content-Type", "application/zip")
		rw.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s.zip", pattern.Name))

		tarWriter := utils.NewTarWriter()
		data, _ := createArtifactHubPkg(pattern, strings.Trim(fmt.Sprintf("%s %s", user.FirstName, user.LastName), " "))
		err = tarWriter.Compress("artifacthub-pkg.yml", data)
		if err != nil {
			h.log.Error(err)
			eb := *eventBuilder
			event := eb.WithSeverity(events.Error).WithDescription(fmt.Sprintf("Unable to create artifacthub pkg for the design \"%s\"", pattern.Name)).WithMetadata(map[string]interface{}{"error": err}).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
		}
		ymlDesign, err := yaml.Marshal(pattern.PatternFile)
		if err != nil {
			err = ErrEncodePattern(err)
			h.log.Error(err)
			http.Error(rw, fmt.Sprintf("Failed to export design \"%s\" as OCI image.", pattern.Name), http.StatusInternalServerError)
			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": err,
			}).WithDescription(fmt.Sprintf("Failed to export design \"%s\" as OCI image.", pattern.Name)).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)

			return
		}
		err = tarWriter.Compress(pattern.Name+".yml", ymlDesign)
		if err != nil {
			h.log.Error(err)
			eb := *eventBuilder
			event := eb.WithSeverity(events.Error).WithDescription(fmt.Sprintf("Unable to compress design \"%s\" and artifacthub pkg.", pattern.Name)).WithMetadata(map[string]interface{}{"error": err}).Build()
			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			tarWriter.Close()
			return
		}

		tarWriter.Close()

		_, err = io.Copy(rw, tarWriter.Buffer)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = encoding.Unmarshal([]byte(pattern.PatternFile), &unmarshalledPatternFile)
	if err != nil {
		err = ErrParsePattern(err)
		h.log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/yaml")
	err = yaml.NewEncoder(rw).Encode(unmarshalledPatternFile)
	if err != nil {
		err = ErrEncodePattern(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/pattern/clone/{id} PatternsAPI idCloneMeshplayPattern
// Handle Clone for a Meshplay Pattern
//
// Creates a local copy of a published pattern with id: id
// responses:
//
//	200 : noContentWrapper
//
// CloneMeshplayPatternHandler clones a pattern with the given id
func (h *Handler) CloneMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	patternUUID := uuid.FromStringOrNil(patternID)

	userID := uuid.FromStringOrNil(user.ID)
	token, _ := r.Context().Value(models.TokenCtxKey).(string)

	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("clone").ActedUpon(patternUUID).WithSeverity(events.Informational)

	var parsedBody *models.MeshplayClonePatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil || patternID == "" {
		h.log.Error(ErrRequestBody(err))
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}

	defer func() {
		_ = r.Body.Close()
	}()

	meshplayPatternByt, err := provider.GetMeshplayPattern(r, patternID, "false")
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrGetPattern(err),
		}).WithDescription(fmt.Sprintf("Failed to fetch meshplay pattern \"%s\" with id: %s.", parsedBody.Name, patternID)).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}

	pattern := &models.MeshplayPattern{}
	err = json.Unmarshal(meshplayPatternByt, &pattern)
	if err != nil {
		obj := "pattern: " + patternID
		h.log.Error(models.ErrUnmarshal(err, obj))
		http.Error(rw, models.ErrUnmarshal(err, obj).Error(), http.StatusInternalServerError)
		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": models.ErrUnmarshal(err, obj),
		}).WithDescription(fmt.Sprintf("Failed to fetch meshplay pattern \"%s\" with ID: %s.", parsedBody.Name, patternID)).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}

	isOldFormat, err := patternutils.IsDesignInAlpha2Format(pattern.PatternFile)
	if err != nil {
		err = ErrPatternFile(err)
		event := events.NewEvent().ActedUpon(*pattern.ID).FromSystem(*h.SystemID).FromUser(userID).WithDescription(fmt.Sprintf("Failed to parse design \"%s\".", pattern.Name)).WithMetadata(map[string]interface{}{"error": err, "id": patternID}).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		h.log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if isOldFormat {
		eventBuilder := events.NewEvent().ActedUpon(*pattern.ID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("convert")
		_, patternFileStr, err := h.convertV1alpha2ToV1beta1(pattern, eventBuilder)
		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		if err != nil {
			h.log.Error(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		pattern.PatternFile = patternFileStr
		_, err = provider.SaveMeshplayPattern(token, pattern)
		if err != nil {
			h.log.Error(ErrSavePattern(err))
			http.Error(rw, ErrSavePattern(err).Error(), http.StatusInternalServerError)

			event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
				"error": ErrSavePattern(_errors.Wrapf(err, "failed to persist converted v1beta1 design file \"%s\" with id: %s", parsedBody.Name, patternID)),
			}).WithDescription(ErrSavePattern(err).Error()).Build()

			_ = provider.PersistEvent(event)
			go h.config.EventBroadcaster.Publish(userID, event)
			return
		}
	}

	resp, err := provider.CloneMeshplayPattern(r, patternID, parsedBody)
	if err != nil {
		h.log.Error(ErrClonePattern(err))
		http.Error(rw, ErrClonePattern(err).Error(), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route POST /api/pattern/catalog/publish PatternsAPI idPublishCatalogPatternHandler
// Handle Publish for a Meshplay Pattern
//
// Publishes pattern to Meshplay Catalog by setting visibility to published and setting catalog data
// responses:
//
//	202: noContentWrapper
//
// PublishCatalogPatternHandler sets visibility of pattern with given id as published
func (h *Handler) PublishCatalogPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().
		FromUser(userID).
		FromSystem(*h.SystemID).
		WithCategory("pattern").
		WithAction("publish").
		ActedUpon(userID)

	var parsedBody *models.MeshplayCatalogPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrRequestBody(err),
			}).
			WithDescription("Error parsing design payload.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}
	resp, err := provider.PublishCatalogPattern(r, parsedBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrPublishCatalogPattern(err),
			}).
			WithDescription("Error publishing design.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	var respBody *models.CatalogRequest
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrPublishCatalogPattern(err),
		}).WithDescription("Error parsing response.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
	}

	e := eventBuilder.WithSeverity(events.Informational).ActedUpon(parsedBody.ID).WithDescription(fmt.Sprintf("Request to publish '%s' design submitted with status: %s", respBody.ContentName, respBody.Status)).Build()
	_ = provider.PersistEvent(e)
	go h.config.EventBroadcaster.Publish(userID, e)

	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/pattern/catalog/unpublish PatternsAPI idUnPublishCatalogPatternHandler
// Handle Publish for a Meshplay Pattern
//
// Unpublishes pattern from Meshplay Catalog by setting visibility to private and removing catalog data from website
// responses:
//
//	200: noContentWrapper
//
// UnPublishCatalogPatternHandler sets visibility of pattern with given id as private
func (h *Handler) UnPublishCatalogPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()

	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().
		FromUser(userID).
		FromSystem(*h.SystemID).
		WithCategory("pattern").
		WithAction("unpublish_request").
		ActedUpon(userID)

	var parsedBody *models.MeshplayCatalogPatternRequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		h.log.Error(ErrRequestBody(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrRequestBody(err),
			}).
			WithDescription("Error parsing design payload.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrRequestBody(err).Error(), http.StatusBadRequest)
		return
	}
	resp, err := provider.UnPublishCatalogPattern(r, parsedBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).
			WithMetadata(map[string]interface{}{
				"error": ErrPublishCatalogPattern(err),
			}).
			WithDescription("Error publishing design.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	var respBody *models.CatalogRequest
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		h.log.Error(ErrPublishCatalogPattern(err))
		e := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrPublishCatalogPattern(err),
		}).WithDescription("Error parsing response.").Build()
		_ = provider.PersistEvent(e)
		go h.config.EventBroadcaster.Publish(userID, e)
		http.Error(rw, ErrPublishCatalogPattern(err).Error(), http.StatusInternalServerError)
	}

	e := eventBuilder.WithSeverity(events.Informational).ActedUpon(parsedBody.ID).WithDescription(fmt.Sprintf("'%s' design unpublished", respBody.ContentName)).Build()
	_ = provider.PersistEvent(e)
	go h.config.EventBroadcaster.Publish(userID, e)

	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route DELETE /api/patterns PatternsAPI idDeleteMeshplayPattern
// Handle Delete for multiple Meshplay Patterns
//
// DeleteMultiMeshplayPatternsHandler deletes patterns with the given ids
func (h *Handler) DeleteMultiMeshplayPatternsHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error(models.ErrDataRead(err, "Request Body"))
	}
	var patterns models.MeshplayPatternDeleteRequestBody
	err = json.Unmarshal([]byte(body), &patterns)
	if err != nil {
		h.log.Error(models.ErrMarshal(err, "pattern"))
	}

	h.log.Debug("patterns to be deleted: ", patterns)

	resp, err := provider.DeleteMeshplayPatterns(r, patterns)

	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to delete the pattern: %s", err), http.StatusInternalServerError)
		return
	}
	go h.config.PatternChannel.Publish(uuid.FromStringOrNil(user.ID), struct{}{})
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(resp))
}

// swagger:route GET /api/pattern/{id} PatternsAPI idGetMeshplayPattern
// Handle GET for a Meshplay Pattern
//
// ```?metrics``` Returns metrics like deployment/share/clone/view/download count for desings, default false,
//
// Fetches the pattern with the given id
// responses:
// 	200: meshplayPatternResponseWrapper

// GetMeshplayPatternHandler fetched the pattern with the given id
func (h *Handler) GetMeshplayPatternHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	patternID := mux.Vars(r)["id"]
	patternUUID := uuid.FromStringOrNil(patternID)
	userID := uuid.FromStringOrNil(user.ID)

	resp, err := provider.GetMeshplayPattern(r, patternID, r.URL.Query().Get("metrics"))
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	pattern := &models.MeshplayPattern{}
	err = json.Unmarshal(resp, &pattern)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusInternalServerError)
		return
	}

	err = h.VerifyAndConvertToDesign(r.Context(), pattern, provider)
	if err != nil {
		event := events.NewEvent().ActedUpon(patternUUID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("convert").WithDescription(fmt.Sprintf("The \"%s\" is not in the design format, failed to convert and persist the original source content from \"%s\" to design file format", pattern.Name, pattern.Type.String)).WithMetadata(map[string]interface{}{"error": err}).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		h.log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	isOldFormat, err := patternutils.IsDesignInAlpha2Format(pattern.PatternFile)
	if err != nil {
		err = ErrPatternFile(err)
		event := events.NewEvent().ActedUpon(patternUUID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("view").WithDescription(fmt.Sprintf("Failed to parse design \"%s\".", pattern.Name)).WithMetadata(map[string]interface{}{"error": err, "id": pattern.ID}).Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		h.log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if isOldFormat {
		eventBuilder := events.NewEvent().ActedUpon(*pattern.ID).FromSystem(*h.SystemID).FromUser(userID).WithCategory("pattern").WithAction("convert")
		_, patternFileStr, err := h.convertV1alpha2ToV1beta1(pattern, eventBuilder)
		event := eventBuilder.Build()
		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		if err != nil {
			h.log.Error(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		pattern.PatternFile = patternFileStr
	}

	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(pattern); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) formatPatternOutput(rw http.ResponseWriter, content []byte, format, sourcetype string, eventBuilder *events.EventBuilder, URL, action string) {
	contentMeshplayPatternSlice := make([]models.MeshplayPattern, 0)

	if err := json.Unmarshal(content, &contentMeshplayPatternSlice); err != nil {
		http.Error(rw, ErrDecodePattern(err).Error(), http.StatusInternalServerError)
		return
	}

	result := []models.MeshplayPattern{}
	names := []string{}
	for _, content := range contentMeshplayPatternSlice {
		if content.ID != nil {
			eventBuilder.ActedUpon(*content.ID)
		}

		result = append(result, content)
		names = append(names, content.Name)
	}

	data, err := json.Marshal(&result)
	if err != nil {
		obj := "pattern file"
		http.Error(rw, models.ErrMarshal(err, obj).Error(), http.StatusInternalServerError)
		return
	}
	var response string
	if URL == "" {
		actionDesc := "updated"
		if action == models.Create {
			actionDesc = "created"
		}
		response = fmt.Sprintf("%s \"%s\" %s.", sourcetype, strings.Join(names, ","), actionDesc)
	} else {
		response = fmt.Sprintf("%s \"%s\" imported from URL %s", sourcetype, strings.Join(names, ","), URL)
	}
	eventBuilder.WithDescription(response)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(data))
}

// Since the client currently does not support pattern imports and externalized variables, the first(import) stage of pattern engine
// is evaluated here to simplify the pattern file such that it is valid when a deploy takes place

//unsued currently

// func evalImportAndReferenceStage(p *pattern.PatternFile) (newp pattern.PatternFile) {
// 	chain := stages.CreateChain()
// 	chain.
// 		// Add(stages.Import(sip, sap)). enable this
// 		Add(stages.Filler(false)).
// 		Add(func(data *stages.Data, err error, next stages.ChainStageNextFunction) {
// 			data.Lock.Lock()
// 			newp = *data.Pattern
// 			data.Lock.Unlock()
// 		}).
// 		Process(&stages.Data{
// 			Pattern: p,
// 		})
// 	return newp
// }

// Only pass Meshkit err here or there will be a panic
func addMeshkitErr(res *meshes.EventsResponse, err error) {
	if err != nil {
		res.EventType = meshes.EventType_ERROR
		res.ProbableCause = errors.GetCause(err)
		res.SuggestedRemediation = errors.GetRemedy(err)
		res.Details = err.Error()
		res.Summary = errors.GetSDescription(err)
		res.ErrorCode = errors.GetCode(err)
	}
}

// swagger:route PUT /api/pattern/{sourcetype} PatternsAPI idUpdateMeshplayPattern
// Handle PUT request for Meshplay Pattern with the given payload
//
// Updates the pattern with the given payload
// responses:
//
//	200: meshplayPatternResponseWrapper
func (h *Handler) handlePatternUpdate(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	defer func() {
		_ = r.Body.Close()
	}()
	userID := uuid.FromStringOrNil(user.ID)
	eventBuilder := events.NewEvent().FromUser(userID).FromSystem(*h.SystemID).WithCategory("pattern").WithAction("update").ActedUpon(userID)

	res := meshes.EventsResponse{
		Component:     "core",
		ComponentName: "Design",
		OperationId:   guid.NewString(),
		EventType:     meshes.EventType_INFO,
	}

	sourcetype := mux.Vars(r)["sourcetype"]
	if sourcetype == "" {
		http.Error(rw, "missing route variable \"source-type\"", http.StatusBadRequest)

		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": ErrSaveApplication(fmt.Errorf("missing route variable \"source-type\" (one of %s, %s, %s)", models.K8sManifest, models.DockerCompose, models.HelmChart)),
		}).WithDescription("Please provide design source-type").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		go h.EventsBuffer.Publish(&res)
		return
	}

	var parsedBody *MeshplayPatternUPDATERequestBody
	if err := json.NewDecoder(r.Body).Decode(&parsedBody); err != nil {
		http.Error(rw, ErrRetrieveData(err).Error(), http.StatusBadRequest)
		return
	}

	actedUpon := &userID
	if parsedBody.PatternData != nil && parsedBody.PatternData.ID != nil {
		actedUpon = parsedBody.PatternData.ID
	}

	eventBuilder.ActedUpon(*actedUpon)

	token, err := provider.GetProviderToken(r)
	if err != nil {
		event := eventBuilder.WithSeverity(events.Critical).WithMetadata(map[string]interface{}{
			"error": ErrRetrieveUserToken(err),
		}).WithDescription("No auth token provided in the request.").Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)
		http.Error(rw, ErrRetrieveUserToken(err).Error(), http.StatusInternalServerError)
		return
	}
	format := r.URL.Query().Get("output")

	meshplayPattern := parsedBody.PatternData
	meshplayPattern.Type = sql.NullString{
		String: sourcetype,
		Valid:  true,
	}
	resp, err := provider.SaveMeshplayPattern(token, meshplayPattern)
	if err != nil {
		errAppSave := ErrSaveApplication(err)
		h.log.Error(errAppSave)

		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "%s", err)

		event := eventBuilder.WithSeverity(events.Error).WithMetadata(map[string]interface{}{
			"error": errAppSave,
		}).WithDescription(fmt.Sprintf("Error saving design %s", parsedBody.PatternData.Name)).Build()

		_ = provider.PersistEvent(event)
		go h.config.EventBroadcaster.Publish(userID, event)

		return
	}
	go h.config.PatternChannel.Publish(userID, struct{}{})

	eventBuilder.WithSeverity(events.Informational)
	h.formatPatternOutput(rw, resp, format, sourcetype, eventBuilder, parsedBody.URL, models.Update)
	event := eventBuilder.Build()
	_ = provider.PersistEvent(event)
	go h.config.EventBroadcaster.Publish(userID, event)

}

// swagger:route POST /api/pattern/{sourcetype} PatternsAPI idPostPatternFileRequest
// Handle POST request for Pattern Files
//
// Creates a new Pattern with source-content
// responses:
//  200: meshplayPatternResponseWrapper

// PatternFileRequestHandler will handle requests of both type GET and POST
// on the route /api/pattern
func (h *Handler) DesignFileRequestHandlerWithSourceType(
	rw http.ResponseWriter,
	r *http.Request,
	prefObj *models.Preference,
	user *models.User,
	provider models.Provider,
) {
	if r.Method == http.MethodPost {
		h.handlePatternPOST(rw, r, prefObj, user, provider)
		return
	}

	if r.Method == http.MethodPut {
		h.handlePatternUpdate(rw, r, prefObj, user, provider)
		return
	}
}

// swagger:route GET /api/pattern/types PatternsAPI typeGetMeshplayPatternTypesHandler
// Handle GET request for Meshplay Pattern types
//
// Get pattern file types
// responses:
//
//	200: meshplayApplicationTypesResponseWrapper
func (h *Handler) GetMeshplayDesignTypesHandler(
	rw http.ResponseWriter,
	_ *http.Request,
	_ *models.Preference,
	_ *models.User,
	_ models.Provider,
) {
	response := models.GetDesignsTypes()
	b, err := json.Marshal(response)
	if err != nil {
		obj := "available types"
		h.log.Error(models.ErrMarshal(err, obj))
		http.Error(rw, models.ErrMarshal(err, obj).Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(b))
}

// swagger:route GET /api/pattern/download/{id}/{sourcetype} PatternsAPI typeGetPatternSourceContent
// Handle GET request for Meshplay Patterns with of source content
//
// Get the pattern source-content
// responses:
//  200: meshplayPatternSourceContentResponseWrapper

// GetMeshplayPatternHandler fetched the design using the given id and sourcetype
func (h *Handler) GetMeshplayPatternSourceHandler(
	rw http.ResponseWriter,
	r *http.Request,
	_ *models.Preference,
	_ *models.User,
	provider models.Provider,
) {
	designID := mux.Vars(r)["id"]
	token, _ := r.Context().Value(models.TokenCtxKey).(string)

	resp, err := provider.GetDesignSourceContent(token, designID)
	if err != nil {
		h.log.Error(ErrGetPattern(err))
		http.Error(rw, ErrGetPattern(err).Error(), http.StatusNotFound)
		return
	}

	var mimeType string
	sourcetype := mux.Vars(r)["sourcetype"]

	switch models.DesignType(sourcetype) {
	case models.HelmChart:
		mimeType = "application/x-tar"
	default: // docker-compose, k8smanifest
		mimeType = "application/x-yaml"
	}

	reader := bytes.NewReader(resp)
	rw.Header().Set("Content-Type", mimeType)
	_, err = io.Copy(rw, reader)
	if err != nil {
		h.log.Error(ErrApplicationSourceContent(err, "download"))
		http.Error(rw, ErrApplicationSourceContent(err, "download").Error(), http.StatusInternalServerError)
	}
}

func createArtifactHubPkg(pattern *models.MeshplayPattern, user string) ([]byte, error) {
	isCatalogItem := pattern.Visibility == models.Published
	var version string
	if isCatalogItem {
		version = pattern.CatalogData.PublishedVersion
	}
	artifactHubPkg := catalog.BuildArtifactHubPkg(pattern.Name, "", user, version, pattern.CreatedAt.Format(time.RFC3339), &pattern.CatalogData)

	data, err := yaml.Marshal(artifactHubPkg)
	if err != nil {
		return nil, models.ErrMarshalYAML(err, "artifacthub-pkg")
	}

	return data, nil
}

func (h *Handler) convertV1alpha2ToV1beta1(meshplayPattern *models.MeshplayPattern, eventBuilder *events.EventBuilder) (*pattern.PatternFile, string, error) {

	v1alpha1PatternFile := v1alpha2.PatternFile{}

	v1beta1PatternFile := pattern.PatternFile{}

	err := encoding.Unmarshal([]byte(meshplayPattern.PatternFile), &v1alpha1PatternFile)
	if err != nil {
		return nil, "", ErrParsePattern(err)
	}

	for _, svc := range v1alpha1PatternFile.Services {
		svc.Settings = helpers.RecursiveCastMapStringInterfaceToMapStringInterface(svc.Settings)
		svc.Traits = helpers.RecursiveCastMapStringInterfaceToMapStringInterface(svc.Traits)
	}

	err = v1beta1PatternFile.ConvertFrom(&v1alpha1PatternFile)
	if err != nil {
		return nil, "", err
	}

	v1beta1PatternFile.Id = *meshplayPattern.ID
	v1beta1PatternFile.Version = v1alpha1PatternFile.Version

	h.log.Infof("Converted design file with id \"%s\" to v1beta1 format", *meshplayPattern.ID)

	err = mapModelRelatedData(h.registryManager, &v1beta1PatternFile)
	if err != nil {
		eventBuilder.WithDescription("Design converted to v1beta1 format but failed to assign styles and metadata").
			WithMetadata(map[string]interface{}{"error": ErrGetComponentDefinition(err), "id": *meshplayPattern.ID}).WithSeverity(events.Warning)
		return nil, "", err
	}

	v1beta1PatternByt, err := encoding.Marshal(v1beta1PatternFile)
	if err != nil {
		return nil, "", utils.ErrMarshal(err)
	}
	eventBuilder.WithSeverity(events.Informational).WithDescription(fmt.Sprintf("Converted design file \"%s\" with id \"%s\" to v1beta1 format", meshplayPattern.Name, *meshplayPattern.ID))
	return &v1beta1PatternFile, string(v1beta1PatternByt), nil
}

func mapModelRelatedData(reg *meshmodel.RegistryManager, patternFile *pattern.PatternFile) error {
	s := selector.New(reg)
	for _, comp := range patternFile.Components {
		if comp == nil {
			continue
		}

		wc, err := s.GetDefinition(comp.Component.Kind, comp.Model.Model.Version, comp.Model.Name, comp.Component.Version, true)
		if err != nil {
			m := []string{"meshplay", "meshplay-core", "meshplay-shapes", "meshplay-flowchart"}
			// if model is one of those defined in the slice above as meshplay, and no matching defs were found,
			// try to find the component just by name, this ensures the component is upgraded to newer model.
			// Eg: Some old designs contains "Comment" component under "meshplay" model instead of "meshplay-core"
			
			
			// Update the component kind to reflect the current registry.
			// Eg: The Connection component for k8s, had "kind" updated to "KuberntesConnection",hence any designs which has model k8s and kind "Connection" will fail, to ensure it gets converted, update the kind
			if comp.Model.Name == "kubernetes" && comp.Component.Kind == "Connection" {
				comp.Component.Kind = "KubernetesConnection"
			} else if comp.Model.Name == "aws" || comp.Model.Name == "gcp" {
				comp.Component.Kind = fmt.Sprintf("%s %s", strings.ToUpper(comp.Model.Name), comp.Component.Kind)
			} else if !slices.Contains(m, comp.Model.Name) {
				return err
			}

			entities, _, _, _ := reg.GetEntities(&regv1beta1.ComponentFilter{
				Name:       comp.Component.Kind,
				APIVersion: comp.Component.Version,
			})
			comp, found := selector.FindCompDefinitionWithVersion(entities, comp.Model.Model.Version)

			if found {
				wc = *comp
			}

		}

		comp.Model = wc.Model
		if wc.Model.Registrant.Status == "" {
			comp.Model.Registrant.Status = connection.Registered
		}
		comp.Format = wc.Format
		comp.Version = wc.Version
		status := component.ComponentDefinitionStatus(wc.Model.Status)
		comp.Status = &status

		// Replace the SVG value with the  svg path. UI uses the path to fetch the SVG from the server.

		// helpers.WriteSVGsOnFileSystem()
		if comp.Model.Metadata.SvgComplete != nil && *comp.Model.Metadata.SvgComplete == "" {
			comp.Model.Metadata.SvgComplete = nil
		}
		comp.Capabilities = wc.Capabilities
		if comp.Capabilities == nil {
			comp.Capabilities = models.K8sMeshModelMetadata.Capabilities
		}
		comp.Metadata.Genealogy = wc.Metadata.Genealogy
		comp.Metadata.IsAnnotation = wc.Metadata.IsAnnotation
		comp.Metadata.Published = wc.Metadata.Published

		var styles component.Styles

		if comp.Styles != nil {
			styles = *comp.Styles
		} else {
			comp.Styles = &component.Styles{}
		}

		// Assign the other styles and reassign the position.
		if wc.Styles != nil {
			comp.Styles = wc.Styles
		}
		if styles.Position != nil {
			comp.Styles.Position = styles.Position
		}

	}

	return nil
}
