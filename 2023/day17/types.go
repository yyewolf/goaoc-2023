package main

type direction struct {
	row, col int
}

var (
	left  = direction{row: 0, col: -1}
	right = direction{row: 0, col: 1}
	up    = direction{row: -1, col: 0}
	down  = direction{row: 1, col: 0}
)

var rotations = map[direction][]direction{
	left:  {up, down},
	right: {up, down},
	up:    {left, right},
	down:  {left, right},
}

var reverse = map[direction]direction{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type state struct {
	row   int
	col   int
	dir   direction
	moves int
}

type item struct {
	heatLoss int
	state    state
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
