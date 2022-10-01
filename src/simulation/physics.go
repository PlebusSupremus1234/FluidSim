package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Compute the particle's density and pressure
func (s *Simulation) computeDensityPressure() {
	node := s.particles.Head

	for node != nil {
		i := node.Value
		var density float32 = 0

		neighbours := s.findNeighbours(i)
		s.neighbours[i.Index] = neighbours

		for _, j := range neighbours {
			rij := rl.Vector2Subtract(i.X, j.X)
			W := s.poly6(rij)

			density += j.M * W
		}

		density += i.M * s.poly6(rl.Vector2Zero()) // Self

		i.Rho = density                        // Density
		i.P = s.stiffness * (i.Rho/s.rho0 - 1) // Pressure

		node = node.Next
	}
}

// Compute non-pressure forces
func (s *Simulation) computeNonPressForces() {
	node := s.particles.Head

	for node != nil {
		i := node.Value

		viscForce := rl.Vector2Zero() // Viscosity force

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			vij := rl.Vector2Subtract(i.V, j.V)

			Wlap := s.viscLap(rij)

			// Compute viscosity force
			multiplierV := j.M / j.Rho * Wlap
			viscForce = rl.Vector2Add(viscForce, rl.Vector2Scale(vij, multiplierV))
		}

		viscForce = rl.Vector2Scale(viscForce, -i.M*s.nu)
		Fgravity := rl.Vector2Scale(s.gravity, i.M/i.Rho)

		sum := rl.Vector2Add(viscForce, Fgravity)

		i.V = rl.Vector2Add(i.V, rl.Vector2Scale(sum, s.dt/i.Rho))

		node = node.Next
	}
}

// Compute pressure forces
func (s *Simulation) computePressForces() {
	node := s.particles.Head

	for node != nil {
		i := node.Value

		pressureForce := rl.Vector2Zero()
		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)

			Wgrad := s.spikyGrad(rij)

			// Compute pressure force
			multiplier := i.P/(i.Rho*i.Rho) + j.P/(j.Rho*j.Rho)
			pressureForce = rl.Vector2Add(pressureForce, rl.Vector2Scale(Wgrad, multiplier))
		}

		pressureForce = rl.Vector2Scale(pressureForce, -i.M*i.M)

		i.V = rl.Vector2Add(i.V, rl.Vector2Scale(pressureForce, s.dt/i.Rho))

		node = node.Next
	}
}

// Enforce boundaries
func (s *Simulation) enforceBoundaries(p *particle.Particle) {
	if p.X.X < s.h {
		p.V.X *= -0.5
		p.X.X = s.h
	}

	if p.X.X > s.simW-s.h {
		p.V.X *= -0.5
		p.X.X = s.simW - s.h
	}

	if p.X.Y < s.h {
		p.V.Y *= -0.5
		p.X.Y = s.h
	}

	if p.X.Y > s.simH-s.h {
		p.V.Y *= -0.5
		p.X.Y = s.simH - s.h
	}
}

// Integrate and draw particles
func (s *Simulation) integrateAndDraw() int {
	particleCount := 0

	node := s.particles.Head
	for node != nil {
		p := node.Value
		prev := p.X

		p.Integrate(s.dt) // Integrate

		s.enforceBoundaries(p) // Enforce boundaries

		s.updateGridParticle(prev, p) // Update the grid

		particleCount++ // Increment particle count

		p.Draw(s.scale) // Draw

		node = node.Next
	}

	return particleCount
}
