package manager

import (
	"fmt"
	"reflect"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/entity"
)

// ComponentManager handles associations between components and entities
type ComponentManager interface {
	AddComponentToEntity(component.Component, entity.Entity) error
	RemoveComponentsOfEntity(entity.Entity) error
	RemoveComponentWithTypeFromEntity(component.Type, entity.Entity) error
	ComponentsOfEntity(entity.Entity) ([]component.Component, error)
	ComponentWithTypeFromEntity(component.Type, entity.Entity) (component.Component, error)
	EntitiesWithComponentType(component.Type) ([]entity.Entity, error)
}

type componentMap map[component.Type]component.Component

// IComponentManager is the implementation of a ComponentManager
type IComponentManager struct {
	entitiesWithComponents map[entity.Entity]componentMap
}

var _ ComponentManager = (*IComponentManager)(nil)

// NewComponentManager instantiates an IComponentManager
func NewComponentManager() *IComponentManager {
	return &IComponentManager{
		entitiesWithComponents: make(map[entity.Entity]componentMap, 0),
	}
}

// AddComponentToEntity adds a component to an entity
func (cm *IComponentManager) AddComponentToEntity(c component.Component, e entity.Entity) error {
	cType := reflect.TypeOf(c)
	if cMap, exists := cm.entitiesWithComponents[e]; exists {
		if _, exists := cMap[cType]; exists {
			return fmt.Errorf("Component of type %v already registered for entity %v",
				reflect.TypeOf(c), e)
		}
	} else {
		cm.entitiesWithComponents[e] = componentMap{}
	}
	cm.entitiesWithComponents[e][cType] = c
	return nil
}

// RemoveComponentWithTypeFromEntity the component of a type of an entity
func (cm *IComponentManager) RemoveComponentWithTypeFromEntity(cType component.Type,
	entity entity.Entity) error {
	cMap, exists := cm.entitiesWithComponents[entity]
	if !exists {
		return fmt.Errorf("Entity %v not registered", entity)
	}
	if _, exists := cMap[cType]; !exists {
		return fmt.Errorf("Entity %v does not have a component of type %v",
			entity, cType)
	}
	cMap[cType] = nil
	cm.entitiesWithComponents[entity] = cMap
	return nil
}

// RemoveComponentsOfEntity removes all components of an entity
func (cm *IComponentManager) RemoveComponentsOfEntity(entity entity.Entity) error {
	if _, exists := cm.entitiesWithComponents[entity]; !exists {
		return fmt.Errorf("Entity %v not registered", entity)
	}
	cm.entitiesWithComponents[entity] = nil
	return nil
}

// ComponentWithTypeFromEntity returns a component of type of entity
func (cm IComponentManager) ComponentWithTypeFromEntity(cType component.Type,
	entity entity.Entity) (component.Component, error) {
	cMap, exists := cm.entitiesWithComponents[entity]
	if !exists {
		return nil, fmt.Errorf("No components for entity: %v", entity)
	}
	component, exists := cMap[cType]
	if !exists {
		return nil, fmt.Errorf("No component of type %v for entity %v", cType, entity)
	}
	return component, nil
}

// ComponentsOfEntity returns all components of an entity
func (cm IComponentManager) ComponentsOfEntity(entity entity.Entity) ([]component.Component, error) {
	compMap, exists := cm.entitiesWithComponents[entity]
	if !exists {
		return nil, fmt.Errorf("No entity %v", entity)
	}
	var components []component.Component
	for _, comp := range compMap {
		components = append(components, comp)
	}
	return components, nil
}

// EntitiesWithComponentType returns all entities with a component of type
func (cm *IComponentManager) EntitiesWithComponentType(cType component.Type) ([]entity.Entity, error) {
	var entities []entity.Entity
	for entity, cMap := range cm.entitiesWithComponents {
		if _, exists := cMap[cType]; exists {
			entities = append(entities, entity)
		}
	}
	return entities, nil
}
