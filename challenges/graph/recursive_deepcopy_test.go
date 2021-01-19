package main

import (
	"testing"

	"github.com/kris-nova/job2021/cs/graph"
)

func TestHashControlHappy(t *testing.T) {
	g1 := graph.NewDiGraph()
	g2 := deepCopy(g1)
	h1 := graph.SprettyPrint(g1)
	h2 := graph.SprettyPrint(g2)
	//t.Logf("Hash1: %s\n", h1)
	//t.Logf("Hash2: %s\n", h2)
	if h1 != h2 {
		t.Errorf("Have h2(%s): Want h1(%s)", h2, h1)
	}
	// Win
}
