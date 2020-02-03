package engine

import (
	"reflect"
	"trashy-ecs/pkg/event"
	"trashy-ecs/pkg/world"
	"trashy-ecs/pkg/world/manager"

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
	systems  []system.System
	world    world.World
	eventBus event.Bus
}

func NewEngine() IEngine {
	cm := manager.NewComponentManager()
	em := manager.NewEntityManager()
	w := world.NewWorld(em, cm)
	eb := event.NewBus()
	return IEngine{
		ComponentManager: cm,
		EntityManager:    em,
		world:            w,
		eventBus:         eb,
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
			err := upd.Update(e.world, e.eventBus)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (e *IEngine) AddSystem(s system.System) {
	e.systems = append(e.systems, s)
	if eh, ok := s.(system.EventHandler); ok {
		for _, et := range eh.Types() {
			e.eventBus.Subscribe(et, eh)
		}
		fmt.Println("Registered event handler:", reflect.Indirect(reflect.ValueOf(eh)).Type().Name())
	} else {
		fmt.Println("Registered system:", reflect.Indirect(reflect.ValueOf(s)).Type().Name())
	}
}

func (e *IEngine) World() world.World {
	return e.world
}
