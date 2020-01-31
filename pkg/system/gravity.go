package system

import (
	"reflect"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/event"
	"trashy-ecs/pkg/world"
)

type Gravity struct {
	massType      component.Type
	velType       component.Type
	requiredTypes []component.Type
}

var _ System = (*Gravity)(nil)
var _ Updater = (*Gravity)(nil)

func NewGravity() Gravity {
	massType := reflect.TypeOf(&component.Mass{})
	velType := reflect.TypeOf(&component.Velocity{})
	requiredComponents := []component.Type{massType, velType}
	return Gravity{
		massType,
		velType,
		requiredComponents,
	}
}

func (g Gravity) Update(w world.World, eb event.Bus) error {
	entites, err := w.EntitiesWithComponentTypes(g.requiredTypes...)
	if err != nil {
		panic(err)
	}
	for _, e := range entites {
		massI, _ := e.ComponentOfType(g.massType)
		velI, _ := e.ComponentOfType(g.velType)
		mass, _ := massI.(*component.Mass)
		vel := velI.(*component.Velocity)
		vel.Y -= mass.Grams * 0.01 * 9.8
	}
	return nil
}
