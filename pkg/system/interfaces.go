package system

import (
	"trashy-ecs/pkg/world"
)

type System interface{}

type Updater interface {
	Update(world.World) error
}

type Renderer interface {
	Render(world.World) error
}
