package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sideType uint64

// const (
// 	SidesCircle sideType = iota
// 	_
// 	_
// 	SidesTriangle
// 	SidesSquare
// )
const SidesCircle sideType = 0
const SidesTriangle sideType = 3
const SidesSquare sideType = 4

func CalcSquare(sideLen float64, sidesNum sideType) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case 3:
		return math.Sqrt(sideLen) * math.Pow(sideLen, 2) / 4
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
