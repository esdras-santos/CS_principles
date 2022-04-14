package graphs

import "fmt"

type NodeBFS struct {
	Value    int
	Children []*NodeBFS
}


// visited array should have the length as the number of nodes
var visitedBFS [12]bool

// insert the start node
func BFS(s *NodeBFS) {
	queue := []*NodeBFS{s}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visitedBFS[current.Value] == false {
			visitedBFS[current.Value] = true
			fmt.Printf("node %d visited\n", current.Value)	
			for _, child := range current.Children {
				if visitedBFS[child.Value] == false {
					queue = append(queue, child)
					
				}
				
			}
		}
		
	}

}
