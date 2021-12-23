package utils

import "math"

func Rotate(x, y, z float64, yaw, pitch, roll float64) (float64, float64, float64) {
	cosa := math.Cos(yaw)
	sina := math.Sin(yaw)
	cosb := math.Cos(pitch)
	sinb := math.Sin(pitch)
	cosc := math.Cos(roll)
	sinc := math.Sin(roll)

	Axx := cosa*cosb
	Axy := cosa*sinb*sinc - sina*cosc
	Axz := cosa*sinb*cosc + sina*sinc
	Ayx := sina*cosb
	Ayy := sina*sinb*sinc + cosa*cosc
	Ayz := sina*sinb*cosc - cosa*sinc
	Azx := -sinb
	Azy := cosb*sinc
	Azz := cosb*cosc

	px := x
	py := y
	pz := z

	outX := Axx*px + Axy*py + Axz*pz
	outY := Ayx*px + Ayy*py + Ayz*pz
	outZ := Azx*px + Azy*py + Azz*pz

	return outX, outY, outZ
}

func DegreeToRad(degree float64) float64 {
	return degree * (math.Pi/180)
}