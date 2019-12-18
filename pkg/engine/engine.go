package engine

import (
	"trashy-ecs/pkg/manager"

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
}

func NewEngine() IEngine {
	return IEngine{
		ComponentManager: manager.NewComponentManager(),
		EntityManager:    manager.NewEntityManager(),
	}
}

func (e *IEngine) Update() {
	rl.BeginDrawing()
	for _, sys := range e.systems {
		rend, ok := sys.(system.Renderer)
		if ok {
			err := rend.Render(e.EntityManager, e.ComponentManager)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	rl.EndDrawing()

	for _, sys := range e.systems {
		upd, ok := sys.(system.Updater)
		if ok {
			err := upd.Update(e.EntityManager, e.ComponentManager)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (e *IEngine) Add(system system.System) {
	e.systems = append(e.systems, system)
}
