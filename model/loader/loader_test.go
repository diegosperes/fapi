package loader

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"fapi/test"
	conf "fapi/model/configuration"
)

type LoaderSuite struct {
	test.Suite
}

func (s LoaderSuite) TestLoadInvalidPath() {
	loader := NewLoader("invalid/path/file.yaml")

	s.NotNil(loader.Load())
}

func (s LoaderSuite) TestLoadInvalidYAMLContent() {
	loader := NewLoader(s.GetFixture("invalid-api.yaml"))

	s.NotNil(loader.Load())
}

func (s LoaderSuite) TestLoadValidPath() {
	loader := NewLoader(s.GetFixture("api.yaml"))

	s.Nil(loader.Load())
}

func (s LoaderSuite) TestConfiguration() {
	loader := NewLoader(s.GetFixture("api.yaml"))
	expected := conf.Configuration{
		Name: "api-name",
		Routes: []*conf.RouteConfiguration{
			&conf.RouteConfiguration{
				Path: "v1",
				Model: "model-v1.yaml",
			},
			&conf.RouteConfiguration{
				Path: "v2",
				Model: "model-v1.yaml",
			},
		},
	}

	s.Nil(loader.Load())
	s.Equal(expected, loader.Configuration())
}

func TestLoaderSuite(t *testing.T) {
	suite.Run(t, new(LoaderSuite))
}
