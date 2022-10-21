package main

import (
	sim "github.com/PlebusSupremus1234/FluidSim/src/simulation"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var h float32 = 16 // Particle width

	var simW, simH float32 = 1696, 800   // Simulation size
	var viewW, viewH float32 = 1696, 800 // Viewport size

	simulation := sim.New(
		h,

		simW, simH,
		viewW, viewH,
	)

	rl.InitWindow(int32(viewW), int32(viewH), "Fluid Simulation")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		simulation.Run()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
