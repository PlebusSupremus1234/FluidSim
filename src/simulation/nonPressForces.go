package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computeNonPressForces() {
	for _, i := range s.particles {
		Fvisc := rl.Vector2Zero()
		Fsurf := rl.Vector2Zero()

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)
			vji := rl.Vector2Subtract(j.V, i.V)

			W := s.Poly6(mag)
			WLap := s.ViscLap(mag)

			// Compute viscosity force
			multiplierV := j.M / j.Rho * WLap
			Fvisc = rl.Vector2Add(Fvisc, rl.Vector2Scale(vji, multiplierV))

			// Compute surface tension force
			K := 2 * s.RestDens / (i.Rho + j.Rho) // Symmetric factor that amplifies the forces at the surface
			multiplierC := K * -s.Cohe * i.M * j.M * W
			Fsurf = rl.Vector2Add(Fsurf, rl.Vector2Scale(normalized, multiplierC))
		}

		Fvisc = rl.Vector2Scale(Fvisc, s.Visc)
		Fgravity := rl.Vector2Scale(s.Gravity, i.M/i.Rho)

		sum := rl.Vector2Add(Fvisc, rl.Vector2Add(Fsurf, Fgravity))
		i.A = rl.Vector2Add(i.A, sum)
	}
}
