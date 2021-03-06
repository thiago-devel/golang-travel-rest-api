package dtstructs

// Implementação customizada de Árvore Heap p/ tratar
// para representar as rotas entre os aeroportos

import (
	hp "container/heap" // importa a implementação padrão de Árvore Heap do Golang
)

type Route struct {
	Price float32
	Nodes []string
}

type minorPriceRoute []Route

func (h minorPriceRoute) Len() int {
	return len(h)
}

func (h minorPriceRoute) Less(i, j int) bool {
	return h[i].Price < h[j].Price
}

func (h minorPriceRoute) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minorPriceRoute) Push(x interface{}) {
	*h = append(*h, x.(Route))
}

func (h *minorPriceRoute) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type heap struct {
	prices *minorPriceRoute
}

// NewHeap constructs Heap that encapsulates and abstracts
// the internal function of the heap. It only exposes the
// Push and Pop function
func NewHeap() *heap {
	return &heap{prices: &minorPriceRoute{}}
}

func (h *heap) Push(r Route) {
	hp.Push(h.prices, r)
}

func (h *heap) Pop() Route {
	i := hp.Pop(h.prices)

	return i.(Route)
}

func (h heap) Len() int {
	return h.prices.Len()
}
