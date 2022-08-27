package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computeDensityPressure() {
	for i, Pi := range s.particles {
		var density float32 = 0

		neighbours := s.neighbours[i]

		for _, Pj := range neighbours.Fluid {
			rij := rl.Vector2Subtract(Pj.X, Pi.X)
			magSq := rl.Vector2LenSqr(rij)

			Wij := s.Poly6(magSq)
			density += Wij
		}

		for _, Pj := range neighbours.Bound {
			rij := rl.Vector2Subtract(Pj.X, Pi.X)
			magSq := rl.Vector2LenSqr(rij)

			Wij := s.Poly6(magSq)
			density += Wij
		}

		density += s.Poly6(0) // Self

		Pi.Rho = Pi.M * density
		Pi.P = s.Stiffness * (density/s.RestDens - 1)
	}
}

func (s *Simulation) computeForces() {
	for i, Pi := range s.particles {
		Fpress := rl.Vector2Zero()
		Fvisc := rl.Vector2Zero()
		Fsurf := rl.Vector2Zero()

		neighbours := s.neighbours[i]

		for _, Pj := range neighbours.Fluid {
			rij := rl.Vector2Subtract(Pi.X, Pj.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)
			vij := rl.Vector2Subtract(Pj.V, Pi.V)

			Wpoly := s.Poly6(mag)
			Wspiky := s.SpikyGrad(mag)
			Wvisc := s.ViscLap(mag)

			// Compute pressure force
			multiplierP := (Pi.P/(Pi.Rho*Pi.Rho) + Pj.P/(Pj.Rho*Pj.Rho)) * Wspiky

			Fpress = rl.Vector2Add(Fpress, rl.Vector2Scale(normalized, multiplierP))

			// Compute viscosity force
			multiplierV := Pj.M / Pj.Rho * Wvisc
			Fvisc = rl.Vector2Add(Fvisc, rl.Vector2Scale(vij, multiplierV))

			// Compute surface tension force
			K := 2 * s.RestDens / (Pi.Rho + Pj.Rho) // Symmetric factor that amplifies the forces at the surface
			multiplierC := K * -s.Cohe * Pi.M * Pj.M * Wpoly
			Fsurf = rl.Vector2Add(Fsurf, rl.Vector2Scale(normalized, multiplierC))
		}

		for _, Pj := range neighbours.Bound {
			rij := rl.Vector2Subtract(Pi.X, Pj.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)

			Wspiky := s.SpikyGrad(mag)

			// Compute pressure force
			multiplierP := (Pi.P/(Pi.Rho*Pi.Rho) + Pj.P/(Pj.Rho*Pj.Rho)) * Wspiky
			Fpress = rl.Vector2Add(Fpress, rl.Vector2Scale(normalized, multiplierP))
		}

		Fpress = rl.Vector2Scale(Fpress, -Pi.M*Pi.M)
		Fvisc = rl.Vector2Scale(Fvisc, s.Visc)
		Fgravity := rl.Vector2Scale(s.Gravity, Pi.M/Pi.Rho)

		sum := rl.Vector2Add(Fpress, rl.Vector2Add(Fvisc, rl.Vector2Add(Fsurf, Fgravity)))
		Pi.A = sum
	}
}

func (s *Simulation) integrate() {
	for _, p := range s.particles {
		if p.T == particle.Fluid {
			p.V = rl.Vector2Add(p.V, rl.Vector2Scale(p.A, s.DT/p.Rho))
			p.X = rl.Vector2Add(p.X, rl.Vector2Scale(p.V, s.DT))
		}

		//if p.X.X-s.Eps < 0 {
		//	p.V.X *= s.BoundDamping
		//	p.X.X = s.Eps
		//}
		//if p.X.X+s.Eps > s.ViewWidth {
		//	p.V.X *= s.BoundDamping
		//	p.X.X = s.ViewWidth - s.Eps
		//}
		//if p.X.Y-s.Eps < 0 {
		//	p.V.Y *= s.BoundDamping
		//	p.X.Y = s.Eps
		//}
		//if p.X.Y+s.Eps > s.ViewHeight {
		//	p.V.Y *= s.BoundDamping
		//	p.X.Y = s.ViewHeight - s.Eps
		//}
	}
}
