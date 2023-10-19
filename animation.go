package main

import (
	"fmt"
	"math"
	"time"
)

var A, B, C float64

var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
var colors []string = []string{Red, Green, Yellow, Blue, Purple, Cyan, Gray, White}

func main() {
	cube := MeshCube()
	screen := NewScreenUtils(10, 83, 22)
	for {
		colorIterator := 0
		fmt.Print("\033[H\033[2J") // Clear the console screen
		screen.ClearBuffer()

		for i, tri := range cube.tris {
			rotatedTriangleVertices := make([]Vec3d, 3)

			for j, vec := range tri.p {
				x := calculateRotationX(vec.x, vec.y, vec.z)
				y := calculateRotationY(vec.x, vec.y, vec.z)
				z := calculateRotationZ(vec.x, vec.y, vec.z)

				rotatedVec := Vec3d{x, y, z}
				screen.ScaleToView(&rotatedVec)

				rotatedTriangleVertices[j] = rotatedVec
			}

			rotatedTriangle := Triangle{p: [3]Vec3d{rotatedTriangleVertices[0], rotatedTriangleVertices[1], rotatedTriangleVertices[2]}}

			if i > 0 && i%2 == 0 {
				colorIterator++
			}

			if colorIterator > len(colors) {
				colorIterator = 0
			}

			screen.DrawTriangle(rotatedTriangle, "+", colors[colorIterator])
		}

		screen.PrintBuffer()

		A += 0.05
		B += 0.05
		C += 0.01

		time.Sleep(100 * time.Millisecond)
	}
}

func calculateRotationX(i, j, k float64) float64 {
	return j*math.Sin(A)*math.Sin(B)*math.Cos(C) - k*math.Cos(A)*math.Sin(B)*math.Cos(C) +
		j*math.Cos(A)*math.Sin(C) + k*math.Sin(A)*math.Sin(C) + i*math.Cos(B)*math.Cos(C)
}

func calculateRotationY(i, j, k float64) float64 {
	return j*math.Cos(A)*math.Cos(C) + k*math.Sin(A)*math.Cos(C) -
		j*math.Sin(A)*math.Sin(B)*math.Sin(C) + k*math.Cos(A)*math.Sin(B)*math.Sin(C) -
		i*math.Cos(B)*math.Sin(C)
}

func calculateRotationZ(i, j, k float64) float64 {
	return k*math.Cos(A)*math.Cos(B) - j*math.Sin(A)*math.Cos(B) + i*math.Sin(B)
}
