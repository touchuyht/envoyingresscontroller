package config

import (
	"sync"

	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
)

var Config Configure
var once sync.Once

type Configure struct {
	v1alpha1.EnvoyControlPlaneConfig
}

func InitConfigure(m *v1alpha1.EnvoyControlPlaneConfig) {
	once.Do(func() {
		Config = Configure{
			*m,
		}
	})
}
