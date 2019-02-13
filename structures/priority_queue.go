package structures

type PriorityQueue []*PQItem

type PQItem struct {
	priority int
	val      interface{}
}

func MakeHeap(items []interface{}) PriorityQueue {
	pq := PriorityQueue{}

	for _, item := range items {
		pqItem := item.(*PQItem)
		pq.Push(pqItem)
	}

	return pq
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Push(item interface{}) {
	pqItem := item.(*PQItem)
	*pq = append(*pq, pqItem)
	pq.bubbleUp(pq.Len())
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	var min *PQItem
	old := *pq
	min = old[0]
	*pq = old[1:]

	pq.bubbleDown(1)

	return min
}

func (pq PriorityQueue) parent(n int) int {
	if n == 1 {
		return -1
	}
	return n / 2
}

func (pq PriorityQueue) child(n int) int {
	return n * 2
}

func (pq PriorityQueue) bubbleUp(n int) {
	if pq.parent(n) == -1 {
		return
	}
	if pq[pq.parent(n)-1].priority > pq[n-1].priority {
		pq[n-1], pq[pq.parent(n)-1] = pq[pq.parent(n)-1], pq[n-1]
		pq.bubbleUp(pq.parent(n))
	}
}

func (pq PriorityQueue) bubbleDown(n int) {
	child := pq.child(n) - 1
	minIdx := n - 1

	for i := 0; i <= 1; i++ {
		if child+i < pq.Len() {
			if pq[minIdx].priority > pq[child+i].priority {
				minIdx = child + i
			}
		}
	}
	if minIdx != n-1 {
		pq[n-1], pq[minIdx] = pq[minIdx], pq[n-1]
		pq.bubbleDown(minIdx)
	}
}
