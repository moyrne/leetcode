package tools

type (
	Heap struct {
		items []*Item
	}
	Item struct {
		Val int
	}
)

func (h *Heap) Len() int { return len(h.items) }

func (h *Heap) Less(i int, j int) bool { return h.items[i].Val < h.items[j].Val }

func (h *Heap) Push(x any) { h.items = append(h.items, x.(*Item)) }

func (h *Heap) Swap(i int, j int) { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *Heap) Pop() any {
	n := h.Len()
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}
