package vec3

import "math"

type Vec3 struct {
	X, Y, Z float32
}

func Add(a, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Multiplies(a Vec3, b float32) Vec3 {
	return Vec3{a.X * b, a.Y * b, a.Z * b}
}

func (a Vec3) length() float32 {
	return float32(math.Sqrt(float64(a.X*a.X+a.Y*a.Y+a.Z*a.Z)))
}

func Distance(a, b Vec3) float32 {
	xDiff := a.X-b.X
	yDiff := a.Y-b.Y
	zDiff := a.Z-b.Z
	return float32(math.Sqrt(float64(xDiff*xDiff+yDiff*yDiff+zDiff*zDiff)))
}

func DistanceSquared(a, b Vec3) float32 {
	xDiff := a.X-b.X
	yDiff := a.Y-b.Y
	zDiff := a.Z-b.Z
	return xDiff*xDiff+yDiff*yDiff+zDiff*zDiff
}

func Normalize(a Vec3) Vec3 {
	length := a.length()
	return Vec3{a.X/length,a.Y/length,a.Z/length}
}