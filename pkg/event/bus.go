package event

import (
	"fmt"
	"reflect"
	"trashy-ecs/pkg/world"
)

type Bus interface {
	Publish(Event, world.World)
	Subscribe(Type, Handler)
}

type IBus struct {
	eventHandlers map[Type][]Handler
}

var _ Bus = (*IBus)(nil)

// NewBus instantiates an event handler for a World
func NewBus() *IBus {
	ehs := make(map[Type][]Handler)
	return &IBus{
		eventHandlers: ehs,
	}
}

// Publish a new Event to the Bus
func (b IBus) Publish(e Event, w world.World) {
	t := reflect.TypeOf(e)
	for _, eh := range b.eventHandlers[t] {
		err := eh.Handle(e, w)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Subscribe to an Event type
func (b *IBus) Subscribe(t Type, eh Handler) {
	b.eventHandlers[t] = append(b.eventHandlers[t], eh)
}
