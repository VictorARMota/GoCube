package main

type Vec3d struct {
	x, y, z float64
}

type Triangle struct {
	p [3]Vec3d
}

type Mesh struct {
	tris []Triangle
}

func MeshCube() Mesh {
	tempCube := Mesh{
		tris: []Triangle{
			// South
			{[3]Vec3d{{0.0, 0.0, 0.0}, {0.0, 1.0, 0.0}, {1.0, 1.0, 0.0}}},
			{[3]Vec3d{{0.0, 0.0, 0.0}, {1.0, 1.0, 0.0}, {1.0, 0.0, 0.0}}},
			// East
			{[3]Vec3d{{1.0, 0.0, 0.0}, {1.0, 1.0, 0.0}, {1.0, 1.0, 1.0}}},
			{[3]Vec3d{{1.0, 0.0, 0.0}, {1.0, 1.0, 1.0}, {1.0, 0.0, 1.0}}},
			// North
			{[3]Vec3d{{1.0, 0.0, 1.0}, {1.0, 1.0, 1.0}, {0.0, 1.0, 1.0}}},
			{[3]Vec3d{{1.0, 0.0, 1.0}, {0.0, 1.0, 1.0}, {0.0, 0.0, 1.0}}},
			// West
			{[3]Vec3d{{0.0, 0.0, 1.0}, {0.0, 1.0, 1.0}, {0.0, 1.0, 0.0}}},
			{[3]Vec3d{{0.0, 0.0, 1.0}, {0.0, 1.0, 0.0}, {0.0, 0.0, 0.0}}},
			// Top
			{[3]Vec3d{{0.0, 1.0, 0.0}, {0.0, 1.0, 1.0}, {1.0, 1.0, 1.0}}},
			{[3]Vec3d{{0.0, 1.0, 0.0}, {1.0, 1.0, 1.0}, {1.0, 1.0, 0.0}}},
			// Bottom
			{[3]Vec3d{{1.0, 0.0, 1.0}, {0.0, 0.0, 1.0}, {0.0, 0.0, 0.0}}},
			{[3]Vec3d{{1.0, 0.0, 1.0}, {0.0, 0.0, 0.0}, {1.0, 0.0, 0.0}}},
		},
	}

	return tempCube
}