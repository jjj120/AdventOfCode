package lib

import "math"

type Matrix2x2 struct {
	// M11, M12
	// M21, M22
	M11, M12, M21, M22 int
}

func (m Matrix2x2) Mul(m2 Matrix2x2) Matrix2x2 {
	return Matrix2x2{
		M11: m.M11*m2.M11 + m.M12*m2.M21,
		M12: m.M11*m2.M12 + m.M12*m2.M22,
		M21: m.M21*m2.M11 + m.M22*m2.M21,
		M22: m.M21*m2.M12 + m.M22*m2.M22,
	}
}

func (m Matrix2x2) MulVec2d(v Vec2d) Vec2d {
	return Vec2d{
		X: m.M11*v.X + m.M12*v.Y,
		Y: m.M21*v.X + m.M22*v.Y,
	}
}

func (m Matrix2x2) Det() int {
	return m.M11*m.M22 - m.M12*m.M21
}

func (m Matrix2x2) Inverse() Matrix2x2 {
	det := m.Det()
	if det == 0 {
		panic("Matrix is not invertible")
	}
	return Matrix2x2{
		M11: m.M22 / det,
		M12: -m.M12 / det,
		M21: -m.M21 / det,
		M22: m.M11 / det,
	}
}

func (m Matrix2x2) Transpose() Matrix2x2 {
	return Matrix2x2{
		M11: m.M11,
		M12: m.M21,
		M21: m.M12,
		M22: m.M22,
	}
}

type Matrix3x3 struct {
	// M11, M12, M13
	// M21, M22, M23
	// M31, M32, M33
	M11, M12, M13, M21, M22, M23, M31, M32, M33 int
}

func (m Matrix3x3) Mul(m2 Matrix3x3) Matrix3x3 {
	return Matrix3x3{
		M11: m.M11*m2.M11 + m.M12*m2.M21 + m.M13*m2.M31,
		M12: m.M11*m2.M12 + m.M12*m2.M22 + m.M13*m2.M32,
		M13: m.M11*m2.M13 + m.M12*m2.M23 + m.M13*m2.M33,
		M21: m.M21*m2.M11 + m.M22*m2.M21 + m.M23*m2.M31,
		M22: m.M21*m2.M12 + m.M22*m2.M22 + m.M23*m2.M32,
		M23: m.M21*m2.M13 + m.M22*m2.M23 + m.M23*m2.M33,
		M31: m.M31*m2.M11 + m.M32*m2.M21 + m.M33*m2.M31,
		M32: m.M31*m2.M12 + m.M32*m2.M22 + m.M33*m2.M32,
		M33: m.M31*m2.M13 + m.M32*m2.M23 + m.M33*m2.M33,
	}
}

func (m Matrix3x3) MulVec3d(v Vec3d) Vec3d {
	return Vec3d{
		X: m.M11*v.X + m.M12*v.Y + m.M13*v.Z,
		Y: m.M21*v.X + m.M22*v.Y + m.M23*v.Z,
		Z: m.M31*v.X + m.M32*v.Y + m.M33*v.Z,
	}
}

func (m Matrix3x3) Det() int {
	return m.M11*m.M22*m.M33 + m.M12*m.M23*m.M31 + m.M13*m.M21*m.M32 - m.M13*m.M22*m.M31 - m.M12*m.M21*m.M33 - m.M11*m.M23*m.M32
}

func (m Matrix3x3) Inverse() Matrix3x3 {
	det := m.Det()
	if det == 0 {
		panic("Matrix is not invertible")
	}
	return Matrix3x3{
		M11: (m.M22*m.M33 - m.M23*m.M32) / det,
		M12: (m.M13*m.M32 - m.M12*m.M33) / det,
		M13: (m.M12*m.M23 - m.M13*m.M22) / det,
		M21: (m.M23*m.M31 - m.M21*m.M33) / det,
		M22: (m.M11*m.M33 - m.M13*m.M31) / det,
		M23: (m.M13*m.M21 - m.M11*m.M23) / det,
		M31: (m.M21*m.M32 - m.M22*m.M31) / det,
		M32: (m.M12*m.M31 - m.M11*m.M32) / det,
		M33: (m.M11*m.M22 - m.M12*m.M21) / det,
	}
}

func (m Matrix3x3) Transpose() Matrix3x3 {
	return Matrix3x3{
		M11: m.M11,
		M12: m.M21,
		M13: m.M31,
		M21: m.M12,
		M22: m.M22,
		M23: m.M32,
		M31: m.M13,
		M32: m.M23,
		M33: m.M33,
	}
}

func Identity2x2() Matrix2x2 {
	return Matrix2x2{
		M11: 1,
		M22: 1,
	}
}

func Identity3x3() Matrix3x3 {
	return Matrix3x3{
		M11: 1,
		M22: 1,
		M33: 1,
	}
}

func Rotation2x2(angle int) Matrix2x2 {
	angle = angle % 360
	rad := float64(angle) * math.Pi / 180
	cos := int(math.Round(math.Cos(rad)))
	sin := int(math.Round(math.Sin(rad)))
	return Matrix2x2{
		M11: cos,
		M12: -sin,
		M21: sin,
		M22: cos,
	}
}

func Rotation3x3(angle int, axis int) Matrix3x3 {
	// axis is 0 for X, 1 for Y, and 2 for Z
	angle = angle % 360
	rad := float64(angle) * math.Pi / 180
	cos := int(math.Round(math.Cos(rad)))
	sin := int(math.Round(math.Sin(rad)))
	switch axis {
	case 0:
		return Matrix3x3{
			M11: 1,
			M12: 0,
			M13: 0,
			M21: 0,
			M22: cos,
			M23: -sin,
			M31: 0,
			M32: sin,
			M33: cos,
		}
	case 1:
		return Matrix3x3{
			M11: cos,
			M12: 0,
			M13: sin,
			M21: 0,
			M22: 1,
			M23: 0,
			M31: -sin,
			M32: 0,
			M33: cos,
		}
	case 2:
		return Matrix3x3{
			M11: cos,
			M12: -sin,
			M13: 0,
			M21: sin,
			M22: cos,
			M23: 0,
			M31: 0,
			M32: 0,
			M33: 1,
		}
	}
	return Matrix3x3{}
}
