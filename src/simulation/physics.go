package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computeDensityPressure() {
	// Compute the particle's density and pressure

	node := s.particles.Head

	for node != nil {
		i := node.Value
		var density float32 = 0

		neighbours := s.findNeighbours(i)
		s.neighbours[i.Index] = neighbours

		for _, j := range neighbours {
			rij := rl.Vector2Subtract(i.X, j.X)
			magSq := rl.Vector2LenSqr(rij)
			W := s.poly6(magSq)

			density += j.M * W
		}

		density += i.M * s.poly6(0) // Self

		i.Rho = density                        // Density
		i.P = s.stiffness * (i.Rho/s.rho0 - 1) // Pressure
		i.A = rl.Vector2Zero()                 // Reset acceleration

		node = node.Next
	}
}

func (s *Simulation) computeNonPressForces() {
	// Compute non-pressure forces

	node := s.particles.Head

	for node != nil {
		i := node.Value

		viscForce := rl.Vector2Zero()

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			vji := rl.Vector2Subtract(j.V, i.V)

			Wlap := s.viscLap(mag)

			// Compute viscosity force
			multiplierV := j.M / j.Rho * Wlap
			viscForce = rl.Vector2Add(viscForce, rl.Vector2Scale(vji, multiplierV))
		}

		viscForce = rl.Vector2Scale(viscForce, s.nu)
		Fgravity := rl.Vector2Scale(s.gravity, i.M/i.Rho)

		sum := rl.Vector2Add(viscForce, Fgravity)
		i.A = rl.Vector2Add(i.A, sum)

		node = node.Next
	}
}

func (s *Simulation) computePressForces() {
	// Compute pressure forces

	node := s.particles.Head

	for node != nil {
		i := node.Value

		pressureForce := rl.Vector2Zero()

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)

			Wgrad := s.spikyGrad(mag)

			// Compute pressure force
			multiplier := (i.P/(i.Rho*i.Rho) + j.P/(j.Rho*j.Rho)) * Wgrad
			pressureForce = rl.Vector2Add(pressureForce, rl.Vector2Scale(normalized, multiplier))
		}

		pressureForce = rl.Vector2Scale(pressureForce, -i.M*i.M)

		i.A = rl.Vector2Add(i.A, pressureForce)

		node = node.Next
	}
}

func (s *Simulation) enforceBoundaries(p *particle.Particle) {
	// Enforce boundaries

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
}
