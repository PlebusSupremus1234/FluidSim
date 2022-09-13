package simulation

func (s *Simulation) Run() {
	// Update grid for neighbour search
	s.updateGrid()

	// Compute density and pressure
	s.computeDensityPressure()

	// Compute non-pressure forces
	s.computeNonPressForces()

	// Compute pressure forces
	s.computePressForces()

	// Integration
	s.integrate()

	// Render scene
	s.Draw(s.particles, s.boundaries)
}
