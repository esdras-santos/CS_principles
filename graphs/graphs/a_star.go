package graphs

import (
	PQ "github.com/jupp0r/go-priority-queue"
)

type NodeAS struct{
	pather *NodeAS
	Children []*NodeAS
	Weights map[*NodeAS]int 
	checked bool
	local  int
	global int
	Label string
}



func AStarPath(s ,target  *NodeAS){
	heuristic := map[string]int{"A": 6, "B": 3, "C": 2, "D": 0, "E": 5, "F": 1}
	pq := PQ.New()
	s.global = 6
	pq.Insert(s, float64(s.global))
	for pq.Len() != 0{
		aux,_ := pq.Pop()
		cur := aux.(*NodeAS)
		if(cur == target){
			continue
		}
		cur.checked = true
		for _, child := range cur.Children{
			// if child.checked{
			// 	continue
			// }
			newLocal := cur.local + cur.Weights[child]
			if newLocal < child.local{
				child.local = newLocal
				child.global = newLocal + heuristic[child.Label]
				child.pather = cur
				pq.Insert(child, float64(child.global))
			}
		}
	}
} 