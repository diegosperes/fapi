package loader

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RouteConfiguration struct {
	Path string `json: "path" yaml: "path"`
	Model string `json: "model" yaml: "model"`
	Schema *map[string]interface{} `json: "schema" yaml: "schema"`
}

type Configuration struct {
	Name string `json: "name" yaml:"name"`
	Routes []*RouteConfiguration `json: "routes" yaml: "routes"`
}

func (c Configuration) Expand() error {
	for _, r := range c.Routes {
		data, err := ioutil.ReadFile(r.Model)

		if err != nil {
			return err
		}

		err = yaml.Unmarshal(data, &r.Schema)

		if err != nil {
			return err
		}
	}

	return nil
}
