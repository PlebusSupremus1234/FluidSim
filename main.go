package main

import (
	sim "github.com/PlebusSupremus1234/FluidSim/src/simulation"
	"github.com/PlebusSupremus1234/FluidSim/src/test"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var h float32 = 16                     // Particle width
	var cols, rows float32 = 106, 50       // Grid size
	var width, height = cols * h, rows * h // Simulation size

	simulation := sim.New(
		h,
		cols, rows,
		width, height,
	)

	spaceDown := false

	rl.InitWindow(int32(width), int32(height), "Fluid Simulation")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if rl.IsKeyDown(rl.KeySpace) {
			if !spaceDown {
				simulation.SpawnParticles()

				spaceDown = true
			}
		} else {
			spaceDown = false
		}

		simulation.Run()

		rl.EndDrawing()
	}

	rl.CloseWindow()

	test.Test()
}
