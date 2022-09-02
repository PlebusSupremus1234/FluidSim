package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/renderer"
)

func (s *Simulation) Run() {
	// Update grid for neighbour search
	s.UpdateGrid()

	// Find neighbours
	s.UpdateNeighbours()

	// Reconstruct density and pressure
	s.computeDensityPressure()

	// Compute non-pressure forces
	s.computeNonPressForces()

	// Compute pressure forces
	s.computePressForces()

	// Integration
	s.integrate()

	// Render scene
	renderer.Draw(s.particles, s.boundaries)
}
