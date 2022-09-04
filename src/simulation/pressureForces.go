package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computePressForces() {
	for _, i := range s.particles {
		FpressF := rl.Vector2Zero()

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)

			Wgrad := s.SpikyGrad(mag)

			// Compute pressure force
			multiplier := (i.P/(i.Rho*i.Rho) + j.P/(j.Rho*j.Rho)) * Wgrad
			FpressF = rl.Vector2Add(FpressF, rl.Vector2Scale(normalized, multiplier))
		}

		volume := s.volumeMap[i.Index]

		diff := rl.Vector2Subtract(i.X, volume.Closest)

		Wgrad := s.SpikyGrad(rl.Vector2Length(diff))
		multiplier := -i.M * i.M * volume.Vol * (i.P/(i.Rho*i.Rho) + i.P/(s.RestDens*s.RestDens)) * Wgrad
		FpressB := rl.Vector2Scale(rl.Vector2Normalize(diff), multiplier)

		//fmt.Printf("%.2f  {%.2f %.2f}  {%.2f %.2f}\n", volume.Vol, FpressB.X, FpressB.Y, i.V.X, i.V.Y)

		FpressF = rl.Vector2Scale(FpressF, -i.M*i.M)
		Fpress := rl.Vector2Add(FpressF, FpressB)

		i.A = rl.Vector2Add(i.A, Fpress)
	}
}
