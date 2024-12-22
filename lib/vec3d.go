package lib

import "math"

type Vec3d struct {
	X, Y, Z int
}

func (v Vec3d) Add(v2 Vec3d) Vec3d {
	return Vec3d{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3d) Sub(v2 Vec3d) Vec3d {
	return Vec3d{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec3d) MulScalar(s int) Vec3d {
	return Vec3d{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3d) DivScalar(s int) Vec3d {
	return Vec3d{v.X / s, v.Y / s, v.Z / s}
}

func (v Vec3d) Dot(v2 Vec3d) int {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3d) Cross(v2 Vec3d) Vec3d {
	return Vec3d{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vec3d) Length() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

func (v Vec3d) Normalize() Vec3d {
	return v.DivScalar(int(v.Length()))
}

func (v Vec3d) Angle() int {
	return int(math.Round(math.Atan2(float64(v.Y), float64(v.X)) * 180 / math.Pi))
}

func (v Vec3d) AngleTo(v2 Vec3d) int {
	return v2.Sub(v).Angle()
}

func (v Vec3d) Distance(v2 Vec3d) float64 {
	return v.Sub(v2).Length()
}

// Rotate90 rotates the vector 90 degrees around the given axis. The axis is 0 for X, 1 for Y, and 2 for Z.
func (v Vec3d) Rotate90(axis int) Vec3d {
	switch axis {
	case 0:
		return Vec3d{-v.Y, v.X, v.Z}
	case 1:
		return Vec3d{v.X, -v.Z, v.Y}
	case 2:
		return Vec3d{v.Z, v.Y, -v.X}
	}
	return Vec3d{}
}

// Rotate270 rotates the vector 270 degrees around the given axis. The axis is 0 for X, 1 for Y, and 2 for Z.
func (v Vec3d) Rotate270(axis int) Vec3d {
	switch axis {
	case 0:
		return Vec3d{v.Y, -v.X, v.Z}
	case 1:
		return Vec3d{-v.X, v.Z, v.Y}
	case 2:
		return Vec3d{-v.Z, v.Y, v.X}
	}
	return Vec3d{}
}
