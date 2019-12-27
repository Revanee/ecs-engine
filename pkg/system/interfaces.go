package system

import (
	"trashy-ecs/pkg/engine/world"
	"trashy-ecs/pkg/event"
)

type System interface{}

type Updater interface {
	Update(world.World) error
}

type Renderer interface {
	Render(world.World) error
}

type EventHandler interface {
	Handle(event.Event) error
}
