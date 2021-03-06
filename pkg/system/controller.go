package system

import (
	"reflect"
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/event"
	"trashy-ecs/pkg/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Controller struct {
	jumpType component.Type
}

var _ System = (*Controller)(nil)
var _ Updater = (*Controller)(nil)

func NewController() *Controller {
	return &Controller{
		jumpType: reflect.TypeOf(&component.Jump{}),
	}
}

func (c *Controller) Update(w world.World, eb event.Bus) error {
	if !rl.IsKeyPressed(rl.KeySpace) {
		return nil
	}
	entities, err := w.EntitiesWithComponentTypes(c.jumpType)
	if err != nil {
		return err
	}
	for _, e := range entities {
		eb.Publish(event.JumpEvent{
			Entity: e.ID(),
		}, w)
	}
	return nil
}
