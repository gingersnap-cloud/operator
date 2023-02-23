package v1alpha1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	LabelCache          = Group + "/cache"
	LabelCacheNamespace = LabelCache + "-namespace"
)

// CacheService defines the location of the Cache resource that this LazyCacheRule should be applied to
type CacheService struct {
	// Name is the name of the Cache resource that the LazyCacheRule will be applied to
	Name string `json:"name"`
	// Namespace is the namespace in which the Cache CR belongs
	Namespace string `json:"namespace"`
}

func (s CacheService) ApplyLabels(meta *metav1.ObjectMeta) {
	if meta.Labels == nil {
		meta.Labels = make(map[string]string, 2)
	}
	s.ApplyLabelsToMap(meta.Labels)
}

func (s CacheService) ApplyLabelsToMap(m map[string]string) {
	m[LabelCache] = s.Name
	m[LabelCacheNamespace] = s.Namespace
}

func (s CacheService) LabelSelector() map[string]string {
	return map[string]string{
		LabelCache:          s.Name,
		LabelCacheNamespace: s.Namespace,
	}
}

func (s CacheService) EagerCacheConfigMap() string {
	return fmt.Sprintf("%s-eager-cm", s.Name)
}

func (s CacheService) LazyCacheConfigMap() string {
	return fmt.Sprintf("%s-lazy-cm", s.Name)
}

func (s CacheService) String() string {
	return s.Namespace + "/" + s.Name
}

func (s CacheService) InternalSvcName() string {
	return fmt.Sprintf("%s-internal", s.Name)
}

func (s CacheService) InternalSvcHost() string {
	return s.svcHost(s.InternalSvcName())
}

func (s CacheService) UserServiceBindingSecret() string {
	return s.Name
}

func (s CacheService) UserSvcName() string {
	return s.Name
}

func (s CacheService) UserSvcHost() string {
	return s.svcHost(s.UserSvcName())
}

func (s CacheService) DBSyncerName() string {
	return fmt.Sprintf("%s-db-syncer", s.Name)
}

func (s CacheService) DBSyncerDataServiceBinding() string {
	return fmt.Sprintf("%s-data", s.DBSyncerName())
}

func (s CacheService) DBSyncerCacheServiceBinding() string {
	return fmt.Sprintf("%s-cache", s.DBSyncerName())
}

func (s CacheService) DBSyncerCacheServiceBindingSecret() string {
	return s.DBSyncerDataServiceBinding()
}

func (s CacheService) DataSourceServiceBinding() string {
	return fmt.Sprintf("%s-cache", s.Name)
}

func (s CacheService) SearchIndexName() string {
	return fmt.Sprintf("%s-index", s.Name)
}

func (s CacheService) SearchIndexSvcName() string {
	return s.SearchIndexName()
}

func (s CacheService) SearchIndexSvcHost() string {
	return s.svcHost(s.SearchIndexSvcName())
}

func (s CacheService) SearchIndexSvcPort() int {
	return 9200
}

func (s CacheService) svcHost(svc string) string {
	return fmt.Sprintf("%s.%s.svc.cluster.local", svc, s.Namespace)
}
