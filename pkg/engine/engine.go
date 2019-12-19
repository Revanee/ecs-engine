package engine

import (
	"trashy-ecs/pkg/engine/manager"
	"trashy-ecs/pkg/engine/world"

	rl "github.com/gen2brain/raylib-go/raylib"

	"trashy-ecs/pkg/system"

	"fmt"
)

type Engine interface {
	manager.EntityManager
	manager.ComponentManager
}

type IEngine struct {
	manager.ComponentManager
	manager.EntityManager
	systems []system.System
	world   world.World
}

func NewEngine() IEngine {
	cm := manager.NewComponentManager()
	em := manager.NewEntityManager()
	w := world.NewWorld(em, cm)
	return IEngine{
		ComponentManager: cm,
		EntityManager:    em,
		world:            w,
	}
}

func (e *IEngine) Update() {
	rl.BeginDrawing()
	for _, sys := range e.systems {
		rend, ok := sys.(system.Renderer)
		if ok {
			err := rend.Render(e.world)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	rl.EndDrawing()

	for _, sys := range e.systems {
		upd, ok := sys.(system.Updater)
		if ok {
			err := upd.Update(e.world)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (e *IEngine) AddSystem(system system.System) {
	e.systems = append(e.systems, system)
}

func (e *IEngine) World() world.World {
	return e.world
}
