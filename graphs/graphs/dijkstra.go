package graphs

import (
	"fmt"
	"math"

	PQ "github.com/jupp0r/go-priority-queue"
)


func DijkstraPath(graph map[int]map[int]int, s, e int) {
	var visitedD [5]bool
	var prev []int
	var dist [len(visitedD)]int
	for i, _ := range visitedD{
		dist[i] = math.MaxInt
	}
	dist[s] = 0
	pq := PQ.New()
	pq.Insert(s, 0)
	for pq.Len() != 0{
		cur,_ := pq.Pop()
		index := cur.(int)
		visitedD[index] = true
		for to, weight := range graph[index]{
			if visitedD[to]{
				continue
			}
			newDist := dist[index] + weight
			if newDist < dist[to]{
				prev = append(prev, index)
				
				dist[to] = newDist
				pq.Insert(to, float64(newDist))
			}
		}
	}
	fmt.Println("path: ", prev)
}