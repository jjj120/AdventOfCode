package lib

import "math"

type Vec2d struct {
	X, Y int
}

func (v Vec2d) Add(v2 Vec2d) Vec2d {
	return Vec2d{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2d) Sub(v2 Vec2d) Vec2d {
	return Vec2d{v.X - v2.X, v.Y - v2.Y}
}

func (v Vec2d) MulScalar(s int) Vec2d {
	return Vec2d{v.X * s, v.Y * s}
}

func (v Vec2d) DivScalar(s int) Vec2d {
	return Vec2d{v.X / s, v.Y / s}
}

func (v Vec2d) Dot(v2 Vec2d) int {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2d) Cross(v2 Vec2d) int {
	return v.X*v2.Y - v.Y*v2.X
}

func (v Vec2d) Length() float64 {
	return float64(v.X*v.X + v.Y*v.Y)
}

func (v Vec2d) Normalize() Vec2d {
	return v.DivScalar(int(v.Length()))
}

func (v Vec2d) Rotate90() Vec2d {
	return Vec2d{-v.Y, v.X}
}

func (v Vec2d) Rotate270() Vec2d {
	return Vec2d{v.Y, -v.X}
}

func (v Vec2d) Angle() int {
	return int(math.Round(math.Atan2(float64(v.Y), float64(v.X)) * 180 / math.Pi))
}

func (v Vec2d) AngleTo(v2 Vec2d) int {
	return v2.Sub(v).Angle()
}

func (v Vec2d) Distance(v2 Vec2d) float64 {
	return v.Sub(v2).Length()
}
