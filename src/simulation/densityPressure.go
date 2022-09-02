package simulation

import rl "github.com/gen2brain/raylib-go/raylib"

func (s *Simulation) computeDensityPressure() {
	for _, i := range s.particles {
		var density float32 = 0

		for _, j := range s.neighbours[i.Index] {
			rij := rl.Vector2Subtract(i.X, j.X)
			magSq := rl.Vector2LenSqr(rij)

			W := s.Poly6(magSq)
			density += W
		}

		density += s.Poly6(0) // Self
		density *= i.M

		volume := s.computeBoundaryVolume(i)

		distSq := rl.Vector2LenSqr(rl.Vector2Subtract(i.X, volume.Closest))
		density += i.M * volume.Vol * s.Poly6(distSq)

		i.Rho = density
		i.P = s.Stiffness * (density/s.RestDens - 1)
		i.A = rl.Vector2Zero()

		// Store volume for later use
		s.volumeMap[i.Index] = volume
	}
}
