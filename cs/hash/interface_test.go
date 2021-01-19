package hash

import (
	"testing"

	"github.com/kris-nova/job2021/cs/graph"
)

func TestHashControlHappy(t *testing.T) {
	g1 := graph.NewDiGraph()
	g2 := graph.NewDiGraph()
	h1, err1 := InterfaceHash(g1)
	if err1 != nil {
		t.Errorf("error during hash1: %v", err1)
	}
	h2, err2 := InterfaceHash(g2)
	if err2 != nil {
		t.Errorf("error during hash2: %v", err1)
	}
	t.Logf("Hash1: %s\n", h1)
	t.Logf("Hash2: %s\n", h2)
	if h1 != h2 {
		t.Errorf("Have h2(%s): Want h1(%s)", h2, h1)
	}
	// Win
}

func TestHashControlSad(t *testing.T) {
	g1 := graph.NewDiGraph()
	g2 := graph.NewDiGraph()
	g2.Root.Name = "NewNameThatWillFail"
	h1, err1 := InterfaceHash(g1)
	if err1 != nil {
		t.Errorf("error during hash1: %v", err1)
	}
	h2, err2 := InterfaceHash(g2)
	if err2 != nil {
		t.Errorf("error during hash2: %v", err1)
	}
	t.Logf("Hash1: %s\n", h1)
	t.Logf("Hash2: %s\n", h2)
	if h1 == h2 {
		t.Errorf("Have h2(%s): Want h1(%s)", h2, h1)
	}
	// Win
}
