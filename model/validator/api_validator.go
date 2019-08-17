package validator

import (
	"fmt"
	"errors"

	"fapi/iface"
	conf "fapi/model/configuration"
)

type APIValidator struct {
	iface.Validator
}

func (v APIValidator) Validate(c conf.Configuration) error {
	if !v.isValidStringField(c.Name) {
		return errors.New("Api must have a name.")
	}

	if !(c.Routes != nil && len(c.Routes) > 0) {
		return errors.New("Api must specify `routes` attribute.")
	}

	for i, r := range c.Routes {
		if !v.isValidStringField(r.Path) {
			message := fmt.Sprintf("Route %d must specify `path` attribute", i)
			return errors.New(message)
		}

		if !v.isValidStringField(r.Model) {
			message := fmt.Sprintf("Route %d must specify `model` attribute", i)
			return errors.New(message)
		}
	}

	return nil
}

func (v APIValidator) isValidStringField(value string) bool {
	return value != "" && len(value) > 0
}
