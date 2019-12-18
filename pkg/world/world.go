package world

import (
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/handle"
	"trashy-ecs/pkg/manager"
)

type World interface {
	EntitiesWithComponentTypes(...component.Type) ([]handle.Entity, error)
	NewEntity() handle.Entity
}

type IWorld struct {
	componentManager manager.ComponentManager
	entityManager    manager.EntityManager
}

var _ World = (*IWorld)(nil)

func NewWorld(em manager.EntityManager, cm manager.ComponentManager) World {
	return &IWorld{
		componentManager: cm,
		entityManager:    em,
	}
}

func (w IWorld) EntitiesWithComponentTypes(cTypes ...component.Type) ([]handle.Entity, error) {
	entities, err := w.componentManager.EntitiesWithComponentTypes(cTypes)
	if err != nil {
		panic(err)
	}
	entityHandles := make([]handle.Entity, 0)
	for _, e := range entities {
		entityHandles = append(entityHandles, handle.NewEntity(w.componentManager, e))
	}
	return entityHandles, nil
}

func (w IWorld) NewEntity() handle.Entity {
	e := w.entityManager.NewEntity()
	eh := handle.NewEntity(w.componentManager, e)
	return eh
}
