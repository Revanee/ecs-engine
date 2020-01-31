package system

import (
	"fmt"
	"reflect"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/event"
	"trashy-ecs/pkg/world"
)

type Jump struct {
	velType    component.Type
	eventTypes []event.Type
}

var _ System = (*Jump)(nil)
var _ event.Handler = (*Jump)(nil)
var _ EventHandler = (*Jump)(nil)

func NewJump() *Jump {
	return &Jump{
		eventTypes: []event.Type{reflect.TypeOf(event.JumpEvent{})},
		velType:    reflect.TypeOf(&component.Velocity{}),
	}
}

func (j *Jump) Handle(e event.Event, w world.World) error {
	fmt.Println("Received jump event")
	je, ok := e.(event.JumpEvent)
	if !ok {
		return fmt.Errorf("Received a wrong event type: %v", reflect.TypeOf(e))
	}
	ent := w.EntityHandle(je.Entity)
	velI, err := ent.ComponentOfType(j.velType)
	if err != nil {
		fmt.Println(w.EntitiesWithComponentTypes(j.velType))
		return err
	}

	vel, ok := velI.(*component.Velocity)
	if !ok {
		panic("Oof")
	}

	vel.Y = 300

	return nil
}

func (j *Jump) Types() []event.Type {
	return j.eventTypes
}
