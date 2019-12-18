package system

import "trashy-ecs/pkg/manager"

type System interface {
	Update(manager.EntityManager, manager.ComponentManager) error
}
