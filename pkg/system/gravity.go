package system

import "trashy-ecs/pkg/manager"

import "trashy-ecs/pkg/component"

import "reflect"

type Gravity struct {
	massType      component.Type
	velType       component.Type
	requiredTypes []component.Type
}

func NewGravity() Gravity {
	massType := reflect.TypeOf(&component.Mass{})
	velType := reflect.TypeOf(&component.Velocity{})
	requiredComponents := []component.Type{massType, velType}
	return Gravity{
		massType,
		velType,
		requiredComponents,
	}
}

func (g Gravity) Update(_ manager.EntityManager, cm manager.ComponentManager) error {
	entites, err := cm.EntitiesWithComponentTypes(g.requiredTypes)
	if err != nil {
		panic(err)
	}
	for _, e := range entites {
		massI, err := cm.ComponentWithTypeFromEntity(g.massType, e)
		if err != nil {
			panic(err)
		}
		velI, err := cm.ComponentWithTypeFromEntity(g.velType, e)
		if err != nil {
			panic(err)
		}
		mass := massI.(*component.Mass)
		vel := velI.(*component.Velocity)
		vel.Y -= mass.Grams * 0.001 * 9.8
	}
	return nil
}
