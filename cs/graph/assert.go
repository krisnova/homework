package graph

// AssertGraphEquals will compare g1 to g2 and return the
// boolean value true if they are copies of each other with
// different locations in memory.
// TODO Check for different order and same values
func AssertGraphEquals(g1, g2 *Graph, msg string) bool {
	if *g1 == *g2 {
		return false //  This is the same graph
	}
	g1str := recursiveString(g1.Root, 0, false)
	g2str := recursiveString(g2.Root, 0, false)
	// Here we do a SLOW string compare but it should* work

	//return (g1str == g2str)
	return g1str == g2str
	//if g1str == g2str {
	//	return true
	//}
	//return false
}
