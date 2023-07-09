package main

import (
	"container/heap"
	"fmt"
	"math"
)

const INF = math.MaxInt

type Item struct {
	vertex   int
	distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(graph map[int]map[int]int, start int, end int) int {
	distances := make(map[int]int)
	for vertex := range graph {
		distances[vertex] = INF
	}
	distances[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{vertex: start, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		currentVertex := item.vertex
		currentDistance := item.distance

		if currentVertex == end {
			return distances[currentVertex]
		}

		if currentDistance > distances[currentVertex] {
			continue
		}

		for neighbor, weight := range graph[currentVertex] {
			newDistance := currentDistance + weight

			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				heap.Push(&pq, &Item{vertex: neighbor, distance: newDistance})
			}
		}
	}

	return -1
}

func main() {
	var n, m, start, end int
	fmt.Scan(&n, &m)

	graph := make(map[int]map[int]int)
	for i := 1; i <= n; i++ {
		graph[i] = make(map[int]int)
	}

	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		graph[u][v] = w
	}

	fmt.Scan(&start, &end)

	shortestDistance := dijkstra(graph, start, end)
	fmt.Println(shortestDistance)
}
