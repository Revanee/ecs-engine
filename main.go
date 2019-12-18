package main

import (
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/engine"
	"trashy-ecs/pkg/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	engine := engine.NewEngine()
	engine.Add(system.PosRend{})
	e1 := engine.EntityManager.NewEntity()
	e2 := engine.EntityManager.NewEntity()
	engine.ComponentManager.AddComponentToEntity(component.Position{
		X: 0,
		Y: 0,
	}, e1)
	engine.ComponentManager.AddComponentToEntity(component.Position{
		X: 50,
		Y: 50,
	}, e2)

	for !rl.WindowShouldClose() {
		engine.Update()
	}

	rl.CloseWindow()
}
