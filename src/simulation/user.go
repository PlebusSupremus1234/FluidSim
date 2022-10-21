package simulation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Spawn particles if the space is pressed
func (s *Simulation) spawnIfSpacePressed() {
	if rl.IsKeyDown(rl.KeySpace) {
		if !s.spaceDown {
			s.spawnParticles()

			s.spaceDown = true
		}
	} else {
		s.spaceDown = false
	}
}

// User interaction forces
func (s *Simulation) userForces() {
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		if !s.mouseDown { // First frame the mouse is down
			s.firstMouseDown = true
			s.mouseDown = true
		} else { // Not first frame
			s.firstMouseDown = false
		}
	} else {
		s.firstMouseDown = false
		s.mouseDown = false
	}
}
