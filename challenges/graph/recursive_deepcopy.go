package main

import (
	"fmt"

	"github.com/kris-nova/job2021/cs/graph"
)

// Given a connected undirectional graph write a method
// deepcopy() to make a new copy of the original graph.
// Print the name of each node as it is visited.

func main() {
	fmt.Println("<----")
	g1 := graph.NewDiGraph()
	g2 := deepCopy(g1)
	fmt.Println("Want(true): ", graph.AssertGraphEquals(g1, g2, "Asserting 2 identical graphs"))
	fmt.Println("---->")
}

// deepCopyFunc will recursively copy a graph to a new graph
func deepCopyFunc(g *graph.Graph, f func(vertex *graph.Vertex) *graph.Vertex) *graph.Graph {
	newG := &graph.Graph{
		Root: recursiveDeepCopy(g.Root),
	}
	return newG
}

// deepCopy will recursively copy a graph to a new graph
func deepCopy(g *graph.Graph) *graph.Graph {
	newG := &graph.Graph{
		Root: recursiveDeepCopy(g.Root),
	}
	return newG
}

func recursiveDeepCopy(v *graph.Vertex) *graph.Vertex {
	vCopy := copyVertex(v)
	// Recursion logic goes here
	for _, neighbor := range v.Neighbors {
		neighborCopy := recursiveDeepCopy(neighbor)
		vCopy.Neighbors = append(vCopy.Neighbors, neighborCopy)
	}
	return vCopy
}

func copyVertex(v *graph.Vertex) *graph.Vertex {
	vNew := &graph.Vertex{
		Name: v.Name,
		// Ignore Neighbors
		Value: v.Value,
	}
	return vNew
}
