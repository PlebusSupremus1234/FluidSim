package simulation

import (
	"github.com/PlebusSupremus1234/Fluid-Simulation/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computeDensityPressure() {
	for i, Pi := range s.particles {
		var density float32 = 0

		for _, Pj := range s.neighbours[i] {
			rij := rl.Vector2Subtract(Pj.X, Pi.X)
			magSq := rl.Vector2LenSqr(rij)

			if magSq <= s.H*s.H {
				Wij := s.Poly6(magSq)
				density += Pj.M * Wij
			}
		}

		density += Pi.M * s.Poly6(0) // Self

		Pi.Rho = density
		Pi.P = s.Stiffness * (density - s.RestDens)
	}
}

func (s *Simulation) computeForces() {
	for i, Pi := range s.particles {
		Fpress := rl.Vector2Zero()
		Fvisc := rl.Vector2Zero()
		Fsurf := rl.Vector2Zero()

		for _, Pj := range s.neighbours[i] {
			rij := rl.Vector2Subtract(Pi.X, Pj.X)
			mag := rl.Vector2Length(rij)

			if mag != 0 && mag <= s.H {
				normalized := rl.Vector2Normalize(rij)
				vij := rl.Vector2Subtract(Pj.V, Pi.V)

				Wpoly := s.Poly6(mag)
				Wspiky := s.SpikyGrad(mag)
				Wvisc := s.ViscLap(mag)

				// Compute pressure force
				multiplierP := Pj.M * (Pi.P + Pj.P) / (2 * Pj.Rho) * Wspiky
				Fpress = rl.Vector2Add(Fpress, rl.Vector2Scale(normalized, multiplierP))

				// Compute viscosity force
				// multiplierV := Pj.M / Pj.Rho * Wvisc
				// Fvisc = rl.Vector2Add(Fvisc, rl.Vector2Scale(velDifference, multiplierV))

				multiplierV := Pj.M / Pj.Rho * Wvisc
				Fvisc = rl.Vector2Add(Fvisc, rl.Vector2Scale(vij, multiplierV))

				// Compute surface tension force
				K := 2 * s.RestDens / (Pi.Rho + Pj.Rho) // Symmetric factor that amplifies the forces at the surface
				multiplierC := K * -s.Cohe * Pi.M * Pj.M * Wpoly
				Fsurf = rl.Vector2Add(Fsurf, rl.Vector2Scale(normalized, multiplierC))
			}
		}

		Fvisc = rl.Vector2Scale(Fvisc, s.Visc)
		Fgravity := rl.Vector2Scale(s.Gravity, Pi.M/Pi.Rho)

		sum := rl.Vector2Add(Fpress, rl.Vector2Add(Fvisc, rl.Vector2Add(Fsurf, Fgravity)))
		Pi.A = sum
	}
}

func (s *Simulation) integrate() {
	for _, p := range s.particles {
		if p.T == particle.PARTICLE {
			p.V = rl.Vector2Add(p.V, rl.Vector2Scale(p.A, s.DT/p.Rho))
			p.X = rl.Vector2Add(p.X, rl.Vector2Scale(p.V, s.DT))
		}

		// if p.X.X-s.EPS < 0 {
		// 	p.V.X *= s.BoundDamping
		// 	p.X.X = s.EPS
		// }
		// if p.X.X+s.EPS > s.VIEW_WIDTH {
		// 	p.V.X *= s.BoundDamping
		// 	p.X.X = s.VIEW_WIDTH - s.EPS
		// }
		// if p.X.Y-s.EPS < 0 {
		// 	p.V.Y *= s.BoundDamping
		// 	p.X.Y = s.EPS
		// }
		// if p.X.Y+s.EPS > s.VIEW_HEIGHT {
		// 	p.V.Y *= s.BoundDamping
		// 	p.X.Y = s.VIEW_HEIGHT - s.EPS
		// }
	}
}
