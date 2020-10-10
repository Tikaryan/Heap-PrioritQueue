package main

import (
	"fmt"
)

//Data struct for array to save name with priority
type Data struct {
	name     string
	priority int
}

type priorityQueue []Data

var pq = []Data{
	{name: "Rojap", priority: 5},
	{name: "Hirdaya", priority: 6},
	{name: "Tinkur", priority: 8},
	{name: "Popat", priority: 6},
	{name: "Enola", priority: 10},
}

func main() {
	push("Hindu", 3)
	maxHeapifyBottomUp()
	for len(pq) > 0 {
		item := pop()
		fmt.Printf("Name = %s Priority = %d\n", item.name, item.priority)
	}

}
func pop() Data {
	n := len(pq)
	pq[0], pq[n-1] = pq[n-1], pq[0]

	item := Data{
		name:     pq[n-1].name,
		priority: pq[n-1].priority,
	}
	pq = pq[:n-1]
	maxHeapifyBottomUp()
	return item
}
func push(val string, prior int) {
	item := Data{
		name:     val,
		priority: prior,
	}
	pq = append(pq, item)
	maxHeapifyBottomUp()
}
func maxHeapifyBottomUp() {
	n := len(pq)
	for i := n - 1; i >= 0; i-- {
		p := (i - 1) / 2 //parent index
		if p < 0 {
			break
		} else {
			if pq[i].priority > pq[p].priority {
				pq = swap(p, i)
			}
		}
	}
}
func swap(p, c int) []Data {
	//fmt.Printf("before swapping of %d = %v\n", p, pq)
	pq[p], pq[c] = pq[c], pq[p]
	//fmt.Printf("after swapping %v\n", pq)
	return pq
}
