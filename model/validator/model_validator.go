package validator

import (
	"fmt"
	"errors"

	"fapi/iface"
	conf "fapi/model/configuration"
)

var Types = map[string]bool{
	"string": true,
	"int": true,
	"float": true,
	"date": true,
	"datetime": true,
	"bool": true,
}

type ModelValidator struct {
	iface.Validator
}

func (v ModelValidator) Validate(c conf.Configuration) error {
	for _, r := range c.Routes {
		for field, schema := range *r.Schema {
			s := schema.(map[string]interface{})
			if err := v.validateField(field, s); err != nil {
				return err
			}
		}
	}

	return nil
}

func (v ModelValidator) validateField(name string, schema map[string]interface{}) error {
	value, ok := schema["type"]
	templateMessage := "Field `%s` has invalid %s value, `%s`."

	if !ok || !v.existType(value.(string)) {
		message := fmt.Sprintf(templateMessage, name, "type", value)
		return errors.New(message)
	}

	template, ok := schema["template"]
	if (value == "date" || value == "datetime") && !ok {
		message := fmt.Sprintf(templateMessage, name, "template", template)
		return errors.New(message)
	}

	return nil
}

func (v ModelValidator) existType(t string) bool {
	return Types[t] == true
}
