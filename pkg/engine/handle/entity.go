package handle

import "trashy-ecs/pkg/component"

import "trashy-ecs/pkg/entity"

import "trashy-ecs/pkg/world/manager"

import "reflect"

type Entity interface {
	AddComponent(component.Component)
	RemoveComponent(component.Component)
	RemoveComponentOfType(component.Type)
	ComponentOfType(component.Type) (component.Component, error)
	ID() entity.Entity
}

type IEntity struct {
	componentManager manager.ComponentManager
	entity           entity.Entity
}

var _ Entity = (*IEntity)(nil)

func NewEntity(cm manager.ComponentManager, e entity.Entity) IEntity {
	return IEntity{
		componentManager: cm,
		entity:           e,
	}
}
func (eh IEntity) AddComponent(c component.Component) {
	eh.componentManager.AddComponentToEntity(c, eh.entity)
}

func (eh IEntity) RemoveComponent(c component.Component) {
	cType := reflect.TypeOf(c)
	eh.componentManager.RemoveComponentWithTypeFromEntity(cType, eh.entity)
}

func (eh IEntity) RemoveComponentOfType(cType component.Type) {
	eh.componentManager.RemoveComponentWithTypeFromEntity(cType, eh.entity)
}

func (eh IEntity) ID() entity.Entity {
	return entity.Entity(eh.entity)
}

func (eh IEntity) ComponentOfType(cType component.Type) (component.Component, error) {
	c, err := eh.componentManager.ComponentWithTypeFromEntity(cType, eh.entity)
	return c, err
}
