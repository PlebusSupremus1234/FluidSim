package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/renderer"
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

	// Render particles
	renderer.Draw(s.particles)
}
