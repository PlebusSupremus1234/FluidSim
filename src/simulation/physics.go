package simulation

import (
	"github.com/PlebusSupremus1234/FluidSim/src/particle"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computeDensityPressure() {
	for _, i := range s.particles {
		var fluid float32 = 0
		var bound float32 = 0

		neighbours := s.neighbours[i.Index]

		for _, j := range neighbours.Fluid {
			rij := rl.Vector2Subtract(j.X, i.X)
			magSq := rl.Vector2LenSqr(rij)

			Wij := s.Poly6(magSq)
			fluid += Wij
		}

		fluid += s.Poly6(0) // Self

		for _, k := range neighbours.Bound {
			rij := rl.Vector2Subtract(k.X, i.X)
			magSq := rl.Vector2LenSqr(rij)

			Wij := s.Poly6(magSq)
			bound += Wij
		}

		density := i.M*fluid + i.M*bound

		i.Rho = density
		i.P = s.Stiffness * (density/s.RestDens - 1)
	}
}

func (s *Simulation) computeForces() {
	for _, i := range s.particles {
		Fpress := rl.Vector2Zero()
		Fvisc := rl.Vector2Zero()
		Fsurf := rl.Vector2Zero()

		neighbours := s.neighbours[i.Index]

		for _, j := range neighbours.Fluid {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)
			vij := rl.Vector2Subtract(j.V, i.V)

			Wpoly := s.Poly6(mag)
			Wspiky := s.SpikyGrad(mag)
			Wvisc := s.ViscLap(mag)

			// Compute pressure force
			multiplierP := (i.P/(i.Rho*i.Rho) + j.P/(j.Rho*j.Rho)) * Wspiky

			Fpress = rl.Vector2Add(Fpress, rl.Vector2Scale(normalized, multiplierP))

			// Compute viscosity force
			multiplierV := j.M / j.Rho * Wvisc
			Fvisc = rl.Vector2Add(Fvisc, rl.Vector2Scale(vij, multiplierV))

			// Compute surface tension force
			K := 2 * s.RestDens / (i.Rho + j.Rho) // Symmetric factor that amplifies the forces at the surface
			multiplierC := K * -s.Cohe * i.M * j.M * Wpoly
			Fsurf = rl.Vector2Add(Fsurf, rl.Vector2Scale(normalized, multiplierC))
		}

		for _, k := range neighbours.Bound {
			rij := rl.Vector2Subtract(i.X, k.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)

			Wspiky := s.SpikyGrad(mag)

			// Compute pressure force
			multiplierP := (i.P/(i.Rho*i.Rho) + k.P/(k.Rho*k.Rho)) * Wspiky
			Fpress = rl.Vector2Add(Fpress, rl.Vector2Scale(normalized, multiplierP))
		}

		Fpress = rl.Vector2Scale(Fpress, -i.M*i.M)
		Fvisc = rl.Vector2Scale(Fvisc, s.Visc)
		Fgravity := rl.Vector2Scale(s.Gravity, i.M/i.Rho)

		sum := rl.Vector2Add(Fpress, rl.Vector2Add(Fvisc, rl.Vector2Add(Fsurf, Fgravity)))
		i.A = sum
	}
}

func (s *Simulation) integrate() {
	for _, i := range s.particles {
		if i.T == particle.Fluid {
			i.V = rl.Vector2Add(i.V, rl.Vector2Scale(i.A, s.DT/i.Rho))
			i.X = rl.Vector2Add(i.X, rl.Vector2Scale(i.V, s.DT))
		}

		//if i.X.X-s.Eps < 0 {
		//	i.V.X *= s.BoundDamping
		//	i.X.X = s.Eps
		//}
		//if i.X.X+s.Eps > s.ViewWidth {
		//	i.V.X *= s.BoundDamping
		//	i.X.X = s.ViewWidth - s.Eps
		//}
		//if i.X.Y-s.Eps < 0 {
		//	i.V.Y *= s.BoundDamping
		//	i.X.Y = s.Eps
		//}
		//if i.X.Y+s.Eps > s.ViewHeight {
		//	i.V.Y *= s.BoundDamping
		//	i.X.Y = s.ViewHeight - s.Eps
		//}
	}
}
