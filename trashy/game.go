package trashy

import (
	"trashy-ecs/pkg/component"
	"trashy-ecs/pkg/engine"
	"trashy-ecs/pkg/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = int32(800)
const screenHeight = int32(450)

func Play() {

	// RayLib init
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	// Engine init
	engine := engine.NewEngine()
	engine.AddSystem(system.PosRend{})
	engine.AddSystem(system.NewMotion())
	engine.AddSystem(system.NewGravity())
	engine.AddSystem(system.NewJump())
	engine.AddSystem(system.NewController())

	// Scene init
	w := engine.World()
	e1 := w.NewEntity()
	e1.AddComponent(&component.Position{
		X: 0,
		Y: 0,
	})

	NewBird(w)

	// Start game loop
	for !rl.WindowShouldClose() {
		engine.Update()
	}

	rl.CloseWindow()
}
