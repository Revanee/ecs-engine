package trashy

import (
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/engine/handle"
	"trashy-ecs/pkg/world"
)

func NewBird(w world.World) handle.Entity {
	bird := w.NewEntity()
	bird.AddComponent(&component.Position{
		X: 50,
		Y: 50,
	})
	bird.AddComponent(&component.Velocity{
		X: 10,
		Y: 1,
	})
	bird.AddComponent(&component.Mass{
		Grams: 200,
	})
	bird.AddComponent(&component.Jump{})
	return bird
}
