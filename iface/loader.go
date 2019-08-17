package iface

import (
	conf "fapi/model/configuration"
)

type Loader interface {
	Load() error
	Configuration() *conf.Configuration
}