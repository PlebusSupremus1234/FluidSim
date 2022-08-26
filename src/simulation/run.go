package simulation

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
}

func (s *Simulation) Draw() {
	for _, p := range s.particles {
		p.Draw()
	}
}
