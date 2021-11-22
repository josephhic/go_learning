package main

import "math"

type Vector struct {
	x, y, z float64
}

func (a *Vector) add(b *Vector) Vector {
	return Vector{x: a.x + b.x, y: a.y + b.y, z: a.z + b.z}
}

func (a *Vector) subtract(b *Vector) Vector {
	return Vector{x: a.x - b.x, y: a.y - b.y, z: a.z - b.z}
}

func (a *Vector) dot(b *Vector) float64 {
	return (a.x * b.x) + (a.y * b.y) + (a.z * b.z)
}

func (a *Vector) magnitude() float64 {
	return math.Pow(
		math.Pow(a.x, 2)+
			math.Pow(a.y, 2)+
			math.Pow(a.z, 2),
		0.5)
}
