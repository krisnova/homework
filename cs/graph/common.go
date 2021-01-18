package graph

import (
	"fmt"
)

// Graph represents a connected undirectional graph
// and is used as the main starting point of the graph.
type Graph struct {
	Name string
	Root *Vertex
}

type Vertex struct {
	Name      string
	Value     int
	Neighbors []*Vertex // Vertices
}

type Edge struct {
	Left  *Vertex
	Right *Vertex
}

// NewDiGraph (Directional Graph)
func NewDiGraph() *Graph {
	graph := &Graph{
		Name: "ConnectedGraph",
		Root: &Vertex{
			Value: 1,
			Name:  "RootVertex",
			Neighbors: []*Vertex{
				{
					Value: 16,
					Name:  "Alice",
					Neighbors: []*Vertex{
						{
							Value: 1024,
							Name:  "Jessica",
							Neighbors: []*Vertex{
								{
									Value: 512,
									Name:  "John",
									Neighbors: []*Vertex{
										{
											Value: 513,
											Name:  "Ringo",
										},
										{
											Value: 514,
											Name:  "George",
										},
										{
											Value: 515,
											Name:  "Paul",
										},
									},
								},
							},
						},
						{
							Value: 4,
							Name:  "Eileen",
							Neighbors: []*Vertex{
								{
									Value: 12,
									Name:  "Tigey",
								},
							},
						},
						{
							Value: 164,
							Name:  "Charlie",
							Neighbors: []*Vertex{
								{
									Value: 256,
									Name:  "Aubrey",
									Neighbors: []*Vertex{
										{
											Value: 654,
											Name:  "Emily",
										},
									},
								},
							},
						},
					},
				},
				{
					Name:  "Bob", // Marley
					Value: 101,
				},
				{
					Name:  "Bob", // Dylan
					Value: 102,
				},
				{
					Name:  "Bob", // Ross
					Value: 103,
				},
			},
		},
	}
	return graph
}

// PrettyPrint will print the contents of a *Graph
// to STDOUT in the terminal
func PrettyPrint(graph *Graph) {
	fmt.Println(graph.Name)
	fmt.Println("")
	str := recursiveString(graph.Root, 0, true)
	fmt.Printf("%s", str)
}

const level string = " "

func recursiveString(v *Vertex, depth int, showPointer bool) string {
	str := ""
	str = fmt.Sprintf("%s%s", str, p(v, depth, showPointer))
	for _, neighbor := range v.Neighbors {
		str = fmt.Sprintf("%s%s", str, recursiveString(neighbor, depth+1, showPointer))
	}
	return str
}

// p will print a *Vertex in a uniform way
func p(v *Vertex, depth int, showPointer bool) string {
	str := ""
	levels := ""
	// Calculate our depth string
	// levels := strings.Repeat(" ", depth)
	for i := depth; i >= 0; i-- {
		levels = fmt.Sprintf("%s%s", levels, " ")
	}
	str = fmt.Sprintf("%s\n", str)
	str = fmt.Sprintf("%s%sDepth     :  %d\n", str, levels, depth)
	str = fmt.Sprintf("%s%sName      :  %s\n", str, levels, v.Name)
	str = fmt.Sprintf("%s%sValue     :  %d\n", str, levels, v.Value)
	if showPointer == true {
		str = fmt.Sprintf("%s%sLocation  :  %p\n", str, levels, v)
	}
	str = fmt.Sprintf("%s\n", str)
	return str
}
