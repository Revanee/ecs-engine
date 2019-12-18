package component

import (
	"reflect"
)

// Type is a Component type
type Type interface {
	reflect.Type
}

// Component holds data for an Entity
type Component interface{}
