package event

import (
	"reflect"
	"trashy-ecs/pkg/world"
)

type Event interface{}

type Type reflect.Type

type Handler interface {
	Handle(Event, world.World) error
}
