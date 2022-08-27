package simulation

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) Run() {
	// Update grid for neighbour search
	s.UpdateGrid()

	// Find neighbours
	s.UpdateNeighbours()

	// Compute density and pressure
	s.computeDensityPressure()

	// Compute forces
	s.computeForces()

	// Integration
	s.integrate()

	// Draw particles
	s.Draw()

	// Stats
	rl.DrawText(fmt.Sprintf("FPS: %.2f", rl.GetFPS()), 10, 10, 10, rl.White)
}

func (s *Simulation) Draw() {
	for _, p := range s.particles {
		p.Draw()
	}
}
