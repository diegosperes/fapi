package loader

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"fapi/iface"
	"fapi/model/validator"
	conf "fapi/model/configuration"
)

type Loader struct {
	iface.Loader
	Path string
	apiValidator iface.Validator
	configuration *conf.Configuration
}

func NewLoader(path string) *Loader {
	return &Loader {
		Path: path,
		apiValidator: new(validator.APIValidator),
		configuration: new(conf.Configuration),
	}
}

func (l Loader) Load() error {
	data, err := ioutil.ReadFile(l.Path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &l.configuration)

	if err != nil {
		return err
	}

	return l.apiValidator.Validate(*l.configuration)
}

func (l Loader) Configuration() conf.Configuration {
	return *l.configuration
}
