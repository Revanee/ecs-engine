package system

import (
	"fmt"
	"reflect"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/manager"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PosRend struct{}

func (pr PosRend) Update(em manager.EntityManager, cm manager.ComponentManager) error {
	var components []*component.Position
	cType := reflect.TypeOf(&component.Position{})
	entities, err := cm.EntitiesWithComponentType(cType)
	if err != nil {
		return fmt.Errorf("Could not get entities with components of type %v", cType)
	}
	for _, entity := range entities {
		c, err := cm.ComponentWithTypeFromEntity(cType, entity)
		if err != nil {
			panic(err)
		}
		components = append(components, c.(*component.Position))
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	for _, c := range components {
		rl.DrawCircleV(rl.Vector2{
			X: float32(c.X),
			Y: float32(c.Y),
		}, 20, rl.Blue)
	}

	rl.EndDrawing()

	return nil
}
