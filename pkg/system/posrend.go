package system

import (
	"fmt"
	"reflect"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PosRend struct{}

var _ System = (*PosRend)(nil)
var _ Renderer = (*PosRend)(nil)

func (pr PosRend) Render(w world.World) error {
	var components []*component.Position
	cType := reflect.TypeOf(&component.Position{})
	entities, err := w.EntitiesWithComponentTypes(cType)
	if err != nil {
		return fmt.Errorf("Could not get entities with components of type %v", cType)
	}
	for _, entity := range entities {
		c, _ := entity.ComponentOfType(cType)
		components = append(components, c.(*component.Position))
	}

	rl.ClearBackground(rl.RayWhite)

	for _, c := range components {
		rl.DrawCircleV(rl.Vector2{
			X: float32(c.X),
			Y: float32(c.Y),
		}, 20, rl.Blue)
	}

	return nil
}
