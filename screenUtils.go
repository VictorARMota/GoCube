package main

import (
	"fmt"
	"math"
)

type ScreenUtils struct {
	distanceFromCam int
	consoleWidth    int
	consoleHeight   int
	buffer          [][]string
}

func NewScreenUtils(distanceFromCam int, consoleWidth int, consoleHeight int) *ScreenUtils {
	buffer := make([][]string, consoleHeight)
	for i := range buffer {
		buffer[i] = make([]string, consoleWidth)
	}

	return &ScreenUtils{distanceFromCam: distanceFromCam,
		consoleWidth:  consoleWidth,
		consoleHeight: consoleHeight,
		buffer:        buffer}
}

func (s *ScreenUtils) ScaleObject(object Mesh) {
	scaleFactor := float64(s.distanceFromCam) / 100.0
	for i := range object.tris {
		for j := range object.tris[i].p {
			object.tris[i].p[j].x *= scaleFactor
			object.tris[i].p[j].y *= scaleFactor
			object.tris[i].p[j].z *= scaleFactor
		}
	}
}

func (s *ScreenUtils) OffsetFromCenter(object Mesh) {
	for i := range object.tris {
		for j := range object.tris[i].p {
			object.tris[i].p[j].x = (float64(s.consoleWidth) - (object.tris[i].p[j].x * float64(s.consoleWidth))) / 2
			object.tris[i].p[j].y = (float64(s.consoleWidth) - (object.tris[i].p[j].y * float64(s.consoleWidth))) / 2
			object.tris[i].p[j].z = (float64(s.consoleWidth) - (object.tris[i].p[j].z * float64(s.consoleWidth))) / 2
		}
	}
}

func (s *ScreenUtils) ScaleToView(vector *Vec3d) {
	vector.x += 1.0
	vector.y += 1.0
	vector.x *= 0.5 * float64(s.consoleWidth)
	vector.y *= 0.5 * float64(s.consoleHeight)
}

func (s *ScreenUtils) DrawTriangle(triangle Triangle, c string, color string) {
	s.DrawLine(triangle.p[0].x, triangle.p[0].y, triangle.p[1].x, triangle.p[1].y, c, color)
	s.DrawLine(triangle.p[1].x, triangle.p[1].y, triangle.p[2].x, triangle.p[2].y, c, color)
	s.DrawLine(triangle.p[2].x, triangle.p[2].y, triangle.p[0].x, triangle.p[0].y, c, color)
}

func (s *ScreenUtils) DrawLine(x1, y1, x2, y2 float64, c string, color string) {
	var x, y, xe, ye int

	dx := x2 - x1
	dy := y2 - y1
	dx1 := math.Abs(dx)
	dy1 := math.Abs(dy)
	px := 2*dy1 - dx1
	py := 2*dx1 - dy1

	if dy1 <= dx1 {
		if dx >= 0 {
			x = int(x1)
			y = int(y1)
			xe = int(x2)
		} else {
			x = int(x2)
			y = int(y2)
			xe = int(x1)
		}

		s.Draw(x, y, c, color)

		for i := 0; x < xe; i++ {
			x += 1
			if px < 0 {
				px = px + 2*dy1
			} else {
				if (dx < 0 && dy < 0) || (dx > 0 && dy > 0) {
					y += 1
				} else {
					y -= 1
				}
				px = px + 2*(dy1-dx1)
			}

			s.Draw(x, y, c, color)
		}
	} else {
		if dy >= 0 {
			x = int(x1)
			y = int(y1)
			ye = int(y2)
		} else {
			x = int(x2)
			y = int(y2)
			ye = int(y1)
		}

		s.Draw(x, y, c, color)

		for i := 0; y < ye; i++ {
			y++
			if py <= 0 {
				py = py + 2*dx1
			} else {
				if (dx < 0 && dy < 0) || (dx > 0 && dy > 0) {
					x++
				} else {
					x--
				}
				py = py + 2*(dx1-dy1)
			}

			s.Draw(x, y, c, color)
		}
	}
}

func (s *ScreenUtils) Draw(x, y int, c string, color string) {
	s.Clip(&x, &y)
	if x >= 0 && x < s.consoleWidth && y >= 0 && y < s.consoleHeight {
		s.buffer[y][x] = color + c
	}
}

func (s *ScreenUtils) Clip(x, y *int) {
	if *x < 0 {
		*x = 0
	}
	if *x >= s.consoleWidth {
		*x = s.consoleWidth
	}
	if *y < 0 {
		*y = 0
	}
	if *y >= s.consoleHeight {
		*y = s.consoleHeight
	}
}

func (s *ScreenUtils) ClearBuffer() {
	for i := 0; i < len(s.buffer); i++ {
		for j := 0; j < len(s.buffer[i]); j++ {
			s.buffer[i][j] = " "
		}
	}
}

func (s *ScreenUtils) PrintBuffer() {
	for i := 0; i < len(s.buffer); i++ {
		for j := 0; j < len(s.buffer[i]); j++ {
			// Prints from buffer and reset color
			fmt.Printf(s.buffer[i][j] + "\033[0m")
		}
		fmt.Println() // Move to the next row
	}
}
