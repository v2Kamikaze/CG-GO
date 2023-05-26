package window

import (
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/vec"
	"fmt"
)

type Border int

func (e Border) String() string {
	switch e {
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	case Left:
		return "Left"
	default:
		return "Right"
	}
}

type EdgePoints = [2]vec.Vec2D
type Edge = [2]vec.Vec2D

const (
	Top Border = iota
	Right
	Bottom
	Left
)

func SutherlandHodgeman(mem memory.Memory, shape *geo.GeometricShape, vp *Viewport) *geo.GeometricShape {

	if IsFullInsideViewport(shape, vp) {
		return shape
	}

	var (
		topLeft     = vp.pi
		topRight    = vec.NewVec2D(vp.pf.X, vp.pi.Y)
		bottomLeft  = vec.NewVec2D(vp.pi.X, vp.pf.Y)
		bottomRight = vp.pf

		s = geo.Copy(shape)
	)

	edges := []EdgePoints{
		{topLeft, topRight},
		{topRight, bottomRight},
		{bottomRight, bottomLeft},
		{bottomLeft, topLeft},
	}

	for edge, edgePoint := range edges {
		for _, vertice := range shape.Vertices {

			if !IsInside(vertice, Border(edge), edgePoint[0]) {

				pi := shape.Vertices[0]

				for i := 1; i < len(shape.Vertices); i++ {
					pf := shape.Vertices[i]

					if point, _, _ := Intersection(pi, pf, edgePoint[0], edgePoint[1]); point.X > -1 {
						s.Vertices = InsertAt(s.Vertices, i, point)
						fmt.Println("Inseriu: ", point, i)
						if !IsInside(pi, Border(edge), edgePoint[0]) {
							s.Vertices = RemoveAt(s.Vertices, i-1)
						}
					}

					pi = pf
				}

				pf := shape.Vertices[0]

				if point, _, _ := Intersection(pi, pf, edgePoint[0], edgePoint[1]); point.X > -1 {
					s.Vertices = InsertAt(s.Vertices, 1, point)
					fmt.Println("Inseriu: ", point, 1)
					if !IsInside(pi, Border(edge), edgePoint[0]) {
						s.Vertices = RemoveAt(s.Vertices, len(s.Vertices)-1)
					}
				}

				fmt.Println("Vertices originais: ", shape.Vertices)
				fmt.Println("Vertices da cópia: ", s.Vertices)

				// Removendo as arestas que não estão dentro da borda atual
				clipVertices := make([]vec.Vec2D, len(s.Vertices))

				copy(clipVertices, s.Vertices)
				fmt.Println("Antes do clipping: ", clipVertices)

				fmt.Println("Edge: ", edgePoint, Border(edge))
				for i := range s.Vertices {
					if !IsInside(s.Vertices[i], Border(edge), edgePoint[0]) {
						clipVertices = RemoveElement(clipVertices, s.Vertices[i])
						fmt.Println("Vertice removido:  ", s.Vertices[i])
					}
				}

				fmt.Println("Após o clipping: ", clipVertices)

				s.Vertices = clipVertices
			}
		}

	}

	return s
}

func Contains(vertices []vec.Vec2D, point vec.Vec2D) (contains bool) {

	for _, vertex := range vertices {
		if vertex == point {
			contains = true
			return
		}
	}

	return
}

func InsertAt(vertices []vec.Vec2D, pos int, point vec.Vec2D) []vec.Vec2D {
	newVertices := make([]vec.Vec2D, len(vertices)+1)
	copy(newVertices[:pos], vertices[:pos])
	newVertices[pos] = point
	copy(newVertices[pos+1:], vertices[pos:])
	return newVertices
}

func RemoveAt(vertices []vec.Vec2D, pos int) []vec.Vec2D {
	newVertices := make([]vec.Vec2D, len(vertices)-1)
	copy(newVertices[:pos], vertices[:pos])
	copy(newVertices[pos:], vertices[pos+1:])
	return newVertices
}

func RemoveElement(vertices []vec.Vec2D, point vec.Vec2D) []vec.Vec2D {
	index := -1
	for i, vertex := range vertices {
		if vertex.X == point.X && vertex.Y == point.Y {
			index = i
			break
		}
	}

	if index == -1 {
		return vertices
	}

	return RemoveAt(vertices, index)
}

func CreateEdges(shape *geo.GeometricShape) []Edge {
	var (
		pi    vec.Vec2D
		pf    vec.Vec2D
		edges []Edge
	)

	pi = shape.Vertices[0]

	for i := range shape.Vertices {
		pf = shape.Vertices[i]
		edges = append(edges, Edge{pi, pf})
		pi = pf
	}

	pf = shape.Vertices[0]
	edges = append(edges, Edge{pi, pf})

	return edges
}

func IsInside(point vec.Vec2D, edge Border, edgePoint vec.Vec2D) bool {
	switch edge {
	case Top:
		return point.Y >= edgePoint.Y
	case Bottom:
		return point.Y <= edgePoint.Y
	case Left:
		return point.X >= edgePoint.X
	case Right:
		return point.X <= edgePoint.X
	}

	return false
}

func IsFullInsideViewport(shape *geo.GeometricShape, vp *Viewport) bool {
	totalInside := 0

	for _, vertice := range shape.Vertices {
		if InsideViewport(vertice, vp) {
			totalInside += 1
		}
	}

	return totalInside == len(shape.Vertices)
}

func InsideViewport(point vec.Vec2D, vp *Viewport) bool {
	return point.X >= vp.pi.X && point.X <= vp.pf.X && point.Y >= vp.pi.Y && point.Y <= vp.pf.Y
}

func lerp(a, b, t float64) float64 {
	return a + (a-b)*t
}
