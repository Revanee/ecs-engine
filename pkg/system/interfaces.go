package system

import "trashy-ecs/pkg/manager"

type System interface{}

type Updater interface {
	Update(manager.EntityManager, manager.ComponentManager) error
}

type Renderer interface {
	Render(manager.EntityManager, manager.ComponentManager) error
}
