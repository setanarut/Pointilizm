// CC BY-NC 3.0
// setanarut

package utils

import "fmt"

// Vec is a Vector with x and y
type Vec struct {
	X float64
	Y float64
}

// Prt print vector
func (v *Vec) Prt() {
	fmt.Println(v.X, v.Y)
}

// Randomize position with min and max range
func (v *Vec) Randomize(min float64, max float64) {
	v.X += RandRange(min, max)
	v.Y += RandRange(min, max)
}
