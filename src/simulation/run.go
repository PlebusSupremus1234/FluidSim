package simulation

func (s *Simulation) Run() {
	// User interactivity
	s.spawnIfSpacePressed()
	s.userForces()

	// Compute density and pressure
	s.computeDensityPressure()

	// Compute forces
	s.computeNonPressForces()

	// Compute pressure forces
	s.computePressForces()

	// Integrate and draw the particles
	particleCount := s.integrateAndDraw()

	// Draw scene
	s.drawScene(particleCount)
}
