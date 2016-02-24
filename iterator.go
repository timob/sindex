package sindex

type IteratorAdapter interface {
	AtLastElement() bool
	AtFirstElement() bool
	MoveForward()
	MoveBack()
	RemoveElement(relPos int)
	InsertElement() int
	Pos() int
}

type Iterator struct {
	adapter IteratorAdapter

	valid bool

	moveOnNext bool
	moveOnPrev bool

	empty bool
}

func NewIterator(adapter IteratorAdapter) IteratorInterface {
	return &Iterator{adapter: adapter}
}

func NewEmptyIterator(adapter IteratorAdapter) IteratorInterface {
	return &Iterator{adapter: adapter, empty: true}
}

func (i *Iterator) Next() bool {
	if !i.empty {
		if i.moveOnNext {
			if !i.adapter.AtLastElement() {
				i.adapter.MoveForward()
				i.valid = true
				return true
			} else {
				i.moveOnPrev = false
				i.valid = false
				return false
			}
		} else {
			i.valid = true
			i.moveOnNext = true
			i.moveOnPrev = true
			return true
		}
	} else {
		return false
	}
}

func (i *Iterator) Prev() bool {
	if !i.empty {
		if i.moveOnPrev {
			if !i.adapter.AtFirstElement() {
				i.adapter.MoveBack()
				i.valid = true
				return true
			} else {
				i.moveOnNext = false
				i.valid = false
				return false
			}
		} else {
			i.valid = true
			i.moveOnPrev = true
			i.moveOnNext = true
			return true
		}
	} else {
		return false
	}
}

func (i *Iterator) Insert() (index int) {
	if i.valid {
		return i.adapter.InsertElement()
	} else {
		return 0
	}
}

const (
	prev int = -1
	cur int = 0
	next int = +1
)

func (i *Iterator) Remove() {
	if !i.valid {
		return
	}

	i.valid = false
	if i.adapter.AtLastElement() {
		if i.adapter.AtFirstElement() {
			i.empty = true
			i.adapter.RemoveElement(cur)
		} else {
			i.moveOnNext = true
			i.moveOnPrev = false
			i.adapter.MoveBack()
			i.adapter.RemoveElement(next)
		}
	} else  {
		i.moveOnPrev = true
		i.moveOnNext = false
		i.adapter.MoveForward()
		i.adapter.RemoveElement(prev)
	}
}

func (i *Iterator) Pos() int {
	return i.adapter.Pos()
}