package iface

import (
	conf "fapi/model/configuration"
)

type Validator interface {
	Validate(conf.Configuration) error
}
