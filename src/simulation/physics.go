package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Simulation) computeDensityPressure() {
	// Compute the particle's density and pressure
	for _, i := range s.particles {
		var densityF float32 = 0
		var densityB float32 = 0

		// Fluid portion
		for _, j := range s.findNeighbours(i) {
			rij := rl.Vector2Subtract(i.X, j.X)
			magSq := rl.Vector2LenSqr(rij)
			W := s.poly6(magSq)

			densityF += j.M * W
		}

		densityF += i.M * s.poly6(0) // Self

		// Boundary portion
		boundVol := s.computeBoundaryVolume(i)

		r := rl.Vector2Subtract(i.X, boundVol.Closest)
		W := s.poly6(rl.Vector2LenSqr(r))
		densityB = boundVol.Vol * s.rho0 * W

		i.Rho = densityF + densityB            // Density
		i.P = s.stiffness * (i.Rho/s.rho0 - 1) // Pressure
		i.A = rl.Vector2Zero()                 // Reset acceleration

		// Store volume for later use
		s.volumeMap[i.Index] = boundVol
	}
}

func (s *Simulation) computeNonPressForces() {
	// Compute non-pressure forces
	for _, i := range s.particles {
		Fvisc := rl.Vector2Zero()

		for _, j := range s.findNeighbours(i) {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			vji := rl.Vector2Subtract(j.V, i.V)

			Wlap := s.viscLap(mag)

			// Compute viscosity force
			multiplierV := j.M / j.Rho * Wlap
			Fvisc = rl.Vector2Add(Fvisc, rl.Vector2Scale(vji, multiplierV))
		}

		Fvisc = rl.Vector2Scale(Fvisc, s.nu)
		Fgravity := rl.Vector2Scale(s.gravity, i.M/i.Rho)

		sum := rl.Vector2Add(Fvisc, Fgravity)
		i.A = rl.Vector2Add(i.A, sum)
	}
}

func (s *Simulation) computePressForces() {
	// Compute pressure forces
	for _, i := range s.particles {
		FpressF := rl.Vector2Zero()

		// Compute fluid pressure force potion
		for _, j := range s.findNeighbours(i) {
			rij := rl.Vector2Subtract(i.X, j.X)
			mag := rl.Vector2Length(rij)

			normalized := rl.Vector2Normalize(rij)

			Wgrad := s.spikyGrad(mag)

			// Compute pressure force
			multiplier := (i.P/(i.Rho*i.Rho) + j.P/(j.Rho*j.Rho)) * Wgrad
			FpressF = rl.Vector2Add(FpressF, rl.Vector2Scale(normalized, multiplier))
		}

		// Compute boundary pressure force portion
		volume := s.volumeMap[i.Index]

		diff := rl.Vector2Subtract(i.X, volume.Closest)

		Wgrad := s.spikyGrad(rl.Vector2Length(diff))
		multiplier := -i.M * volume.Vol * i.M * (i.P/(i.Rho*i.Rho) + i.P/(s.rho0*s.rho0)) * Wgrad
		FpressB := rl.Vector2Scale(rl.Vector2Normalize(diff), multiplier)

		FpressF = rl.Vector2Scale(FpressF, -i.M*i.M)
		Fpress := rl.Vector2Add(FpressF, FpressB)

		i.A = rl.Vector2Add(i.A, Fpress)
	}
}

func (s *Simulation) integrate() {
	// Integrate the particles
	for _, i := range s.particles {
		i.V = rl.Vector2Add(i.V, rl.Vector2Scale(i.A, s.dt/i.Rho))

		i.X = rl.Vector2Add(i.X, rl.Vector2Scale(i.V, s.dt))

		i.Draw()
	}
}
