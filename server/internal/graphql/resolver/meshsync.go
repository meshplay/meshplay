package resolver

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/khulnasoft/meshplay/server/internal/graphql/model"
	"github.com/khulnasoft/meshplay/server/models"
	"github.com/layer5io/meshkit/broker"
	"github.com/layer5io/meshkit/models/meshmodel/registry"
	"github.com/layer5io/meshkit/utils"
	meshsyncmodel "github.com/layer5io/meshsync/pkg/model"
	"github.com/spf13/viper"
)

// Global singleton instance of k8s connection tracker to map Each K8sContext to a unique Broker URL
var connectionTrackerSingleton = model.NewK8sConnctionTracker()
var (
	MeshSyncSubscriptionError = model.Error{
		Description: "Failed to get MeshSync data",
		Code:        ErrResolverMeshsyncSubscriptionCode,
	}
	MeshSyncMeshplayClientMissingError = model.Error{
		Code:        ErrResolverMeshsyncSubscriptionCode,
		Description: "Cannot find Meshplay Client",
	}
)

func (r *Resolver) resyncCluster(ctx context.Context, provider models.Provider, actions *model.ReSyncActions, k8scontextID string) (model.Status, error) {
	if actions.ClearDb == "true" {
		// copies the contents .meshplay/config/meshplaydb.sql to .meshplay/config/.archive/meshplaydb.sql
		// then drops all the DB table and then migrate/create tables, missing foreign keys, constraints, columns and indexes.
		if actions.HardReset == "true" {
			meshplaydbPath := path.Join(utils.GetHome(), ".meshplay/config")
			err := os.Mkdir(path.Join(meshplaydbPath, ".archive"), os.ModePerm)
			if err != nil && os.IsNotExist(err) {
				return "", err
			}

			src := path.Join(meshplaydbPath, "meshplaydb.sql")
			dst := path.Join(meshplaydbPath, ".archive/meshplaydb.sql")

			fin, err := os.Open(src)
			if err != nil {
				return "", err
			}
			defer fin.Close()

			fout, err := os.Create(dst)
			if err != nil {
				return "", err
			}
			defer fout.Close()

			_, err = io.Copy(fout, fin)
			if err != nil {
				return "", err
			}

			dbHandler := provider.GetGenericPersister()
			if dbHandler == nil {
				return "", model.ErrEmptyHandler
			}

			dbHandler.Lock()
			defer dbHandler.Unlock()

			r.Log.Info("Dropping Meshplay Database")
			tables, err := dbHandler.Migrator().GetTables()
			if err != nil {
				r.Log.Error(ErrGormDatabase(err))
				return "", err
			}

			for _, table := range tables {
				if table == "events" {
					continue
				}
				if err := dbHandler.Migrator().DropTable(table); err != nil {
					r.Log.Error(ErrGormDatabase(err))
					return "", err
				}
			}

			r.Log.Info("Migrating Meshplay Database")
			err = dbHandler.AutoMigrate(
				&meshsyncmodel.KubernetesKeyValue{},
				&meshsyncmodel.KubernetesResource{},
				&meshsyncmodel.KubernetesResourceSpec{},
				&meshsyncmodel.KubernetesResourceStatus{},
				&meshsyncmodel.KubernetesResourceObjectMeta{},
				&models.PerformanceProfile{},
				&models.MeshplayResult{},
				&models.MeshplayPattern{},
				&models.MeshplayFilter{},
				&models.PatternResource{},
				&models.MeshplayApplication{},
				&models.UserPreference{},
				&models.PerformanceTestConfig{},
				&models.SmiResultWithID{},
				&models.K8sContext{},
			)
			if err != nil {
				r.Log.Error(err)
				return "", err
			}

			krh, err := models.NewKeysRegistrationHelper(dbHandler, r.Log)
			if err != nil {
				return "", err
			}

			rm, err := registry.NewRegistryManager(dbHandler)
			if err != nil {
				return "", err
			}

			go func() {
				models.SeedComponents(r.Log, r.Config, rm)
				krh.SeedKeys(viper.GetString("KEYS_PATH"))
			}()
			r.Log.Info("Hard reset complete.")
		} else { //Delete meshsync objects coming from a particular cluster
			k8sctxs, ok := ctx.Value(models.AllKubeClusterKey).([]models.K8sContext)
			if !ok || len(k8sctxs) == 0 {
				r.Log.Error(ErrEmptyCurrentK8sContext)
				return "", ErrEmptyCurrentK8sContext
			}
			var sid string
			for _, k8ctx := range k8sctxs {
				if k8ctx.ID == k8scontextID && k8ctx.KubernetesServerID != nil {
					sid = k8ctx.KubernetesServerID.String()
					break
				}
			}
			if provider.GetGenericPersister() == nil {
				return "", model.ErrEmptyHandler
			}

			err := provider.GetGenericPersister().Where("id IN (?)", provider.GetGenericPersister().Table("kubernetes_resources").Select("id").Where("cluster_id=?", sid)).Delete(&meshsyncmodel.KubernetesKeyValue{}).Error
			if err != nil {
				return "", model.ErrEmptyHandler
			}

			err = provider.GetGenericPersister().Where("id IN (?)", provider.GetGenericPersister().Table("kubernetes_resources").Select("id").Where("cluster_id=?", sid)).Delete(&meshsyncmodel.KubernetesResourceSpec{}).Error
			if err != nil {
				return "", model.ErrEmptyHandler
			}

			err = provider.GetGenericPersister().Where("id IN (?)", provider.GetGenericPersister().Table("kubernetes_resources").Select("id").Where("cluster_id=?", sid)).Delete(&meshsyncmodel.KubernetesResourceStatus{}).Error
			if err != nil {
				return "", model.ErrEmptyHandler
			}

			err = provider.GetGenericPersister().Where("id IN (?)", provider.GetGenericPersister().Table("kubernetes_resources").Select("id").Where("cluster_id=?", sid)).Delete(&meshsyncmodel.KubernetesResourceObjectMeta{}).Error
			if err != nil {
				return "", model.ErrEmptyHandler
			}

			err = provider.GetGenericPersister().Where("cluster_id = ?", sid).Delete(&meshsyncmodel.KubernetesResource{}).Error
			if err != nil {
				return "", model.ErrEmptyHandler
			}
		}
	}

	if actions.ReSync == "true" {
		if r.BrokerConn.Info() != broker.NotConnected {
			err := r.BrokerConn.Publish(model.RequestSubject, &broker.Message{
				Request: &broker.RequestObject{
					Entity: broker.ReSyncDiscoveryEntity,
				},
			})
			if err != nil {
				return "", ErrPublishBroker(err)
			}
		}
	}
	return model.StatusProcessing, nil
}