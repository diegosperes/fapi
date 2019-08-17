package loader

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"fapi/test"
)

type ConfiguationSuite struct {
	test.Suite
}

func (s ConfiguationSuite) TestExpand() {
	c := Configuration{
		Name: "api-name",
		Routes: []*RouteConfiguration{
			&RouteConfiguration{
				Path: "v1",
				Model: s.GetFixture("model-v1.yaml"),
			},
			&RouteConfiguration{
				Path: "v2",
				Model: s.GetFixture("model-v1.yaml"),
			},
		},
	}

	s.Nil(c.Routes[0].Schema)
	s.Nil(c.Routes[1].Schema)

	err := c.Expand()

	s.Nil(err)
	s.NotNil(c.Routes[0].Schema)
	s.NotNil(c.Routes[1].Schema)
}

func TestConfiguationSuite(t *testing.T) {
	suite.Run(t, new(ConfiguationSuite))
}
