package heap

type PQueue struct {
	items []*item
	Count int
	isMin bool
}

type item struct {
	value    interface{}
	priority int
}

func newItem(value interface{}, priority int) *item {
	return &item{
		value:    value,
		priority: priority,
	}
}
