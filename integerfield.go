package gforms

import (
	"errors"
	"reflect"
	"strconv"
)

type IntegerField struct {
	BaseField
}

func (self *IntegerField) Html(rd RawData) string {
	return fieldToHtml(self, rd)
}

func (self *IntegerField) html(vs ...string) string {
	return renderTemplate("TextTypeField", newTemplateContext(self, vs...))
}

// Create a new field for integer value.
func NewIntegerField(name string, vs Validators, ws ...Widget) *IntegerField {
	self := new(IntegerField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}

func (self *IntegerField) Clean(data Data) (*V, error) {
	m, hasField := data[self.GetName()]
	if hasField {
		v := m.rawValueAsString()
		m.Kind = reflect.Int
		if v != nil && (*v) != "" {
			iv, err := strconv.Atoi(*v)
			if err == nil {
				m.Value = iv
				m.IsNil = false
				return m, nil
			}
			return nil, errors.New("This field should be specified as int.")
		}
	}
	return nilV(), nil
}
