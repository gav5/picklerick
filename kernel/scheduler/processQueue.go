package scheduler

import (
	"log"
)

type processQueue struct {
	base   []uint8
	logger *log.Logger
}

func (pq processQueue) Len() int {
	return len(pq.base)
}

func (pq *processQueue) Push(x interface{}) {
	pNum := x.(uint8)
	(*pq).base = append(pq.base, pNum)
	pq.logger.Printf("Pushed process %d", pNum)
	pq.logger.Printf("Queue now %v", pq.base)
}

func (pq *processQueue) Pop() interface{} {
	old := pq.base
	n := len(old)
	x := old[n-1]
	(*pq).base = old[0 : n-1]

	pq.logger.Printf("Popped process %d", x)
	pq.logger.Printf("Queue now %v", pq.base)
	return x
}

func (pq processQueue) Peek() uint8 {
	return pq.base[len(pq.base)-1]
}
