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

	// Draw scene
	s.drawScene(s.particles)

	// Integrate and draw the particles
	for _, p := range s.particles {
		p.Integrate(s.dt)

		if p.X.X < s.h {
			p.V.X *= -0.5
			p.X.X = s.h
		}

		if p.X.X > s.viewW-s.h {
			p.V.X *= -0.5
			p.X.X = s.viewW - s.h
		}

		if p.X.Y < s.h {
			p.V.Y *= -0.5
			p.X.Y = s.h
		}

		if p.X.Y > s.viewH-s.h {
			p.V.Y *= -0.5
			p.X.Y = s.viewH - s.h
		}

		p.Draw()
	}
}
