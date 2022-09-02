package test

import (
	"fmt"
	"github.com/PlebusSupremus1234/FluidSim/src/boundary"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Test() {
	l := boundary.New(rl.NewVector2(100, 100), rl.NewVector2(100, 300), 16)

	x := rl.NewVector2(100, 150)
	fmt.Println(l.Contains(x))
	fmt.Println(boundary.Sdf(x, l))

	rl.DrawRectangleV(x, rl.NewVector2(5, 5), rl.Red)
	l.Draw()

	//a := rl.NewVector2(100, 500)
	//b := rl.NewVector2(400, 300)
	//
	//c := math.Atan2(float64(b.Y-a.Y), float64(b.X-a.X))
	//fmt.Println(c)
}
