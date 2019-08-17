package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	conf "fapi/model/configuration"
)

type APIValidatorSuite struct {
	suite.Suite
}

func (s APIValidatorSuite) TestValidate() {
	v := new(APIValidator)
	c := conf.Configuration{
		Name: "api-name",
		Routes: []*conf.RouteConfiguration{
			&conf.RouteConfiguration{
				Path: "v1",
				Model: "test/model-v1.yaml",
			},
		},
	}

	s.Nil(v.Validate(c))
}

func (s APIValidatorSuite) TestValidateName() {
	v := new(APIValidator)
	c := conf.Configuration{
		Name: "",
		Routes: []*conf.RouteConfiguration{
			&conf.RouteConfiguration{
				Path: "v1",
				Model: "test/model-v1.yaml",
			},
		},
	}

	expected := errors.New("Api must have a name.")
	s.Equal(expected, v.Validate(c))
}

func (s APIValidatorSuite) TestValidateRoutesLength() {
	v := new(APIValidator)
	c := conf.Configuration{
		Name: "api-name",
		Routes: []*conf.RouteConfiguration{},
	}

	expected := errors.New("Api must specify `routes` attribute.")
	s.Equal(expected, v.Validate(c))
}

func (s APIValidatorSuite) TestValidateRoutesPath() {
	v := new(APIValidator)
	c := conf.Configuration{
		Name: "api-name",
		Routes: []*conf.RouteConfiguration{
			&conf.RouteConfiguration{
				Path: "",
				Model: "test/model-v1.yaml",
			},
		},
	}

	expected := errors.New("Route 0 must specify `path` attribute")
	s.Equal(expected, v.Validate(c))
}

func (s APIValidatorSuite) TestValidateRoutesModel() {
	v := new(APIValidator)
	c := conf.Configuration{
		Name: "api-name",
		Routes: []*conf.RouteConfiguration{
			&conf.RouteConfiguration{
				Path: "v1",
				Model: "",
			},
		},
	}

	expected := errors.New("Route 0 must specify `model` attribute")
	s.Equal(expected, v.Validate(c))
}

func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(APIValidatorSuite))
}
