package simulation

func (s *Simulation) Run() {
	// Compute density and pressure
	s.computeDensityPressure()

	// Compute non-pressure forces
	s.computeNonPressForces()

	// Compute pressure forces
	s.computePressForces()

	// Integrate and draw the particles
	particleCount := 0

	node := s.particles.Head
	for node != nil {
		p := node.Value
		prev := p.X

		p.Integrate(s.dt) // Integrate

		s.enforceBoundaries(p) // Enforce boundaries

		s.updateGridParticle(prev, p) // Update the grid

		particleCount++ // Increment particle count

		p.Draw() // Draw

		node = node.Next
	}

	// Draw scene
	s.drawScene(particleCount)
}
