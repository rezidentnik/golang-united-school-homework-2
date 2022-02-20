package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type NumberOfSides = int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: NumberOfSides(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const SidesCircle NumberOfSides = 0
const SidesTriangle NumberOfSides = 3
const SidesSquare NumberOfSides = 4

func CalcSquare(sideLen float64, sidesNum NumberOfSides) float64 {
	return math.Pi
}
