package heap

import "fmt"

type Heap struct {
	size  int
	arr   []int
	isMin bool
}

func NewHeap(arrInput []int, isMin bool) *Heap {
	size := len(arrInput)
	arr := []int{1}
	arr = append(arr, arrInput...)
	h := &Heap{size: size, arr: arr, isMin: isMin}
	for i := (h.size / 2); i > 0; i-- {
		h.procalateDown(i)
	}
	return h
}

func (h *Heap) procalateDown(parent int) {
	lChild := 2 * parent
	rChild := lChild + 1
	small := -1
	if lChild <= h.size {
		small = lChild
	}
	if rChild <= h.size && h.comp(lChild, rChild) {
		small = rChild
	}
	if small != -1 && h.comp(parent, small) {
		h.swap(parent, small)
		h.procalateDown(small)
	}
}

func NewHeap2(isMin bool) *Heap {
	arr := []int{1}
	h := &Heap{size: 0, arr: arr, isMin: isMin}
	return h
}

func (h *Heap) comp(i, j int) bool {
	if h.isMin == true {
		return h.arr[i] > h.arr[j]
	}
	return h.arr[i] < h.arr[j]
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Empty() bool {
	return (h.size == 0)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() (int, bool) {
	if h.Empty() {
		fmt.Println("heapemptyerror")
		return 0, false
	}
	return h.arr[1], true
}

func (h *Heap) procalateUp(child int) {
	parent := child / 2
	if parent == 0 {
		return
	}

	if h.comp(parent, child) {
		h.swap(parent, child)
		h.procalateUp(parent)
	}

}

func (h *Heap) Add(value int) {
	h.size++
	h.arr = append(h.arr, value)
	h.procalateUp(h.size)
}

func (h *Heap) Remove() (int, bool) {
	if h.Empty() {
		fmt.Println("empty error")
		return 0, false
	}

	value := h.arr[1]
	h.arr[1] = h.arr[h.size]
	h.size--
	h.procalateDown(1)
	h.arr = h.arr[0 : h.size+1]
	return value, true
}

func HeapSort(arrInput []int) {
	hp := NewHeap(arrInput, true)
	n := len(arrInput)
	for i := 0; i < n; i++ {
		arrInput[i], _ = hp.Remove()
	}
}

func InitHeap() {
	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	hp := NewHeap(nil, true)
	n := len(a)
	for i := 0; i < n; i++ {
		hp.Add(a[i])
	}

}
