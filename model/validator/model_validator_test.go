package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"fapi/test"
	conf "fapi/model/configuration"
)

type ModelValidatorSuite struct {
	test.Suite
}

func (s ModelValidatorSuite) Configuration(field map[string]interface{}) conf.Configuration {
	return conf.Configuration{
		Name: "api-name",
		Routes: []*conf.RouteConfiguration{
			&conf.RouteConfiguration{
				Path: "v1",
				Model: "invalid/path/file.yaml",
				Schema: &map[string]interface{}{
					"field": field,
				},
			},
		},
	}
}

func (s ModelValidatorSuite) TestValidateStringType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "string",
		},
	)
	v := new(ModelValidator)
	s.Nil(v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateIntType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "int",
		},
	)
	v := new(ModelValidator)
	s.Nil(v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateFloatType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "float",
		},
	)
	v := new(ModelValidator)
	s.Nil(v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateDateType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "date",
			"template": "",
		},
	)
	v := new(ModelValidator)
	s.Nil(v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateDateTimeType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "datetime",
			"template": "",
		},
	)
	v := new(ModelValidator)
	s.Nil(v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateBoolType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "bool",
		},
	)
	v := new(ModelValidator)
	s.Nil(v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateDateTypeWithoutTemplate() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "date",
		},
	)
	v := new(ModelValidator)
	s.Equal(errors.New("Field `field` has invalid template value, `%!s(<nil>)`."), v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateDateTimeTypeWithoutTemplate() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "datetime",
		},
	)
	v := new(ModelValidator)
	s.Equal(errors.New("Field `field` has invalid template value, `%!s(<nil>)`."), v.Validate(c))
}

func (s ModelValidatorSuite) TestValidateInvalidType() {
	c := s.Configuration(
		map[string]interface{}{
			"type": "something",
		},
	)
	v := new(ModelValidator)
	s.Equal(errors.New("Field `field` has invalid type value, `something`."), v.Validate(c))
}


func (s ModelValidatorSuite) TestValidateNilType() {
	field := make(map[string]interface{})
	c := s.Configuration(field)
	v := new(ModelValidator)
	s.Equal(errors.New("Field `field` has invalid type value, `%!s(<nil>)`."), v.Validate(c))
}

func TestModelValidatorSuite(t *testing.T) {
	suite.Run(t, new(ModelValidatorSuite))
}
