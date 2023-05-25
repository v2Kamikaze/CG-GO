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
		s           = geo.Copy(shape)
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
					}

					pi = pf
				}

				pf := shape.Vertices[0]

				if point, _, _ := Intersection(pi, pf, edgePoint[0], edgePoint[1]); point.X > -1 {
					s.Vertices = InsertAt(s.Vertices, 1, point)
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
						fmt.Println("Chegou aqui", s.Vertices[i])
					}
				}

				fmt.Println("Após o clipping: ", clipVertices)

				s.Vertices = clipVertices
			}
		}

	}

	return s
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

/*
	---------------------------
	|                          |
	|                          |
	|                          |
	|                          |
	|                          |
	---------------------------
*/

/*
check := func() {
	// Checando se intersecta a parte superior da viewport
	point, _, _ = Intersection(pi, pf, topLeft, topRight)
	if point.X > 0 {
		 fmt.Printf("Se chocou com a parte superior nos pontos: %+v %+v\n", pi, point)
		scan.ScanlineBasic(mem, geo.NewRect(5, 5, point), colors.ColorRed)
		fmt.Printf("[%d] Ocorreu uma intersecção na aresta %v - %v no ponto %v\n", i, pi, pf, point)
		newVertices = append(newVertices, point)
	}

	// Checando se intersecta a parte lateral direita da viewport
	point, _, _ = Intersection(pi, pf, topRight, bottomRight)
	if point.X > 0 {
		fmt.Printf("Se chocou com a parte lateral direita %+v %+v\n", point, pf)
		scan.ScanlineBasic(mem, geo.NewRect(5, 5, point), colors.ColorRed)
		newVertices = append(newVertices, point)
		newVertices = append(newVertices, pf)

	}
	// Checando se intersecta a parte inferior da viewport
	point, _, _ = Intersection(pi, pf, bottomRight, bottomLeft)
	if point.X > 0 {
		fmt.Printf("Se chocou com a parte inferior %+v %+v\n", point, pf)
		scan.ScanlineBasic(mem, geo.NewRect(5, 5, point), colors.ColorRed)
		newVertices = append(newVertices, point)
		newVertices = append(newVertices, pf)

	}
	// Checando se intersecta a parte lateral esquerda da viewport
	point, _, _ = Intersection(pi, pf, bottomLeft, topLeft)
	if point.X > 0 {
		fmt.Printf("Se chocou com a parte lateral esquerda %+v %+v\n", point, pf)
		scan.ScanlineBasic(mem, geo.NewRect(5, 5, point.Plus(vec.NewVec2D(10, 0))), colors.ColorRed)
		newVertices = append(newVertices, point)
		newVertices = append(newVertices, pf)
	}

	newVertices = append(newVertices, pi)
} */
