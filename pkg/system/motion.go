package system

import (
	"reflect"
	"time"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/manager"
)

const nanoseconsInOneSecond float64 = 1000000000

type Motion struct {
	posType             component.Type
	velType             component.Type
	requiredComponents  []component.Type
	lastTimeStep        int64
	nanosecondsPerFrame float64
}

func NewMotion() *Motion {
	posType := reflect.TypeOf(&component.Position{})
	velType := reflect.TypeOf(&component.Velocity{})
	requiredComponents := []component.Type{posType, velType}
	lastTimeStep := time.Now().UnixNano()

	var targetFPS float64 = 60
	nanosecondsPerFrame := nanoseconsInOneSecond / targetFPS

	return &Motion{
		posType,
		velType,
		requiredComponents,
		lastTimeStep,
		nanosecondsPerFrame,
	}
}

func (m *Motion) Update(em manager.EntityManager, cm manager.ComponentManager) error {
	currentTime := time.Now().UnixNano()
	elapsedNanoseconds := float64(currentTime - m.lastTimeStep)
	var deltaT float64 = 0.0
	deltaT = elapsedNanoseconds / nanoseconsInOneSecond
	m.lastTimeStep = currentTime

	entities, err := cm.EntitiesWithComponentTypes(m.requiredComponents)
	if err != nil {
		panic(err)
	}
	for _, e := range entities {
		posI, err := cm.ComponentWithTypeFromEntity(m.posType, e)
		if err != nil {
			panic(err)
		}
		velI, err := cm.ComponentWithTypeFromEntity(m.velType, e)
		if err != nil {
			panic(err)
		}
		pos, ok := posI.(*component.Position)
		if !ok {
			panic("Could not get Position pointer")
		}
		vel, ok := velI.(*component.Velocity)
		if !ok {
			panic("Could not get Velocity pointer")
		}
		pos.X += vel.X * deltaT
		pos.Y += -vel.Y * deltaT
		pos.Z += vel.Z * deltaT
	}
	return nil
}
