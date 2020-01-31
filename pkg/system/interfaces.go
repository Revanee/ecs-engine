package system

import (
	"trashy-ecs/pkg/event"
	"trashy-ecs/pkg/world"
)

type System interface{}

type Updater interface {
	Update(world.World, event.Bus) error
}

type Renderer interface {
	Render(world.World) error
}

type EventHandler interface {
	event.Handler
	Types() []event.Type
}
