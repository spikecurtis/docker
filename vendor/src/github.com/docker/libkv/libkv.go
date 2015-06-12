package libkv

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/mock"
	"github.com/docker/libkv/store/zookeeper"
)

// Initialize creates a new Store object, initializing the client
type Initialize func(addrs []string, options *store.Config) (store.Store, error)

var (
	// Backend initializers
	initializers = map[store.Backend]Initialize{
		store.MOCK:   mock.InitializeMock,
		store.CONSUL: consul.InitializeConsul,
		store.ETCD:   etcd.InitializeEtcd,
		store.ZK:     zookeeper.InitializeZookeeper,
	}
)

// NewStore creates a an instance of store
func NewStore(backend store.Backend, addrs []string, options *store.Config) (store.Store, error) {
	if init, exists := initializers[backend]; exists {
		log.WithFields(log.Fields{"backend": backend}).Debug("Initializing store service")
		return init(addrs, options)
	}

	return nil, store.ErrNotSupported
}
