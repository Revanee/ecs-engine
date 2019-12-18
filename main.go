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
	engine.AddSystem(system.PosRend{})
	engine.AddSystem(system.NewMotion())
	engine.AddSystem(system.NewGravity())
	w := engine.World()
	e1 := w.NewEntity()
	e2 := w.NewEntity()
	e1.AddComponent(&component.Position{
		X: 0,
		Y: 0,
	})
	e2.AddComponent(&component.Position{
		X: 50,
		Y: 50,
	})
	e2.AddComponent(&component.Velocity{
		X: 10,
		Y: 1,
	})
	e2.AddComponent(&component.Mass{
		Grams: 200,
	})

	for !rl.WindowShouldClose() {
		engine.Update()
	}

	rl.CloseWindow()
}
