package simulation

import rl "github.com/gen2brain/raylib-go/raylib"

func (s *Simulation) computePressForces() {
	for _, i := range s.particles {
		FpressF := rl.Vector2Zero()

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)

			WGrad := s.SpikyGrad(mag)

			// Compute pressure force
			multiplier := (i.P/(i.Rho*i.Rho) + j.P/(j.Rho*j.Rho)) * WGrad
			FpressF = rl.Vector2Add(FpressF, rl.Vector2Scale(normalized, multiplier))
		}

		volume, closest := s.computeBoundaryVolume(i)

		WGrad := s.SpikyGrad(rl.Vector2Length(rl.Vector2Subtract(i.X, closest)))
		multiplier := (i.P/(i.Rho*i.Rho) + i.P/(s.RestDens*s.RestDens)) * WGrad
		FpressB := rl.Vector2Scale(rl.Vector2Normalize(rl.Vector2Subtract(i.X, closest)), multiplier)

		FpressF = rl.Vector2Scale(FpressF, -i.M*i.M)
		FpressB = rl.Vector2Scale(FpressB, -i.M*i.M*volume)

		Fpress := rl.Vector2Add(FpressF, FpressB)
		i.A = rl.Vector2Add(i.A, Fpress)
	}
}
