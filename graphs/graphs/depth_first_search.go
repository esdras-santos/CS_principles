package graphs

import "fmt"

type Node struct {
	Value    int
	Children []*Node
}

// visited array should have the length as the number of nodes
var visited [12]bool

// insert the start node
func DFS(s *Node) {
	visited[s.Value] = true
	for _, child := range s.Children {
		if visited[child.Value] == false{
			
			fmt.Printf("node %d visited\n", child.Value)
			DFS(child)
		}
	}
}
