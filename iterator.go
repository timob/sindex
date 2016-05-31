package sindex

type iteratorAdapter interface {
	Append() int
	AtLastElement() bool
	AtFirstElement() bool
	MoveForward()
	MoveBack()
	RemoveElement(relPos int)
	InsertElement() int
	Pos() int
}

type iterator struct {
	adapter iteratorAdapter

	valid bool

	moveOnNext bool
	moveOnPrev bool

	empty bool
}

func newIterator(adapter iteratorAdapter) IteratorInterface {
	return &iterator{adapter: adapter}
}

func newEmptyIterator(adapter iteratorAdapter) IteratorInterface {
	return &iterator{adapter: adapter, empty: true}
}

func (i *iterator) Next() bool {
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

func (i *iterator) Prev() bool {
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

func (i *iterator) Insert() (index int) {
	if i.valid {
		return i.adapter.InsertElement()
	} else {
		var pos int
		if i.empty {
			i.moveOnPrev = false
			i.moveOnNext = true
			i.empty = false
			pos = i.adapter.Append()
		} else if i.adapter.AtLastElement() {
			pos = i.adapter.Append()
			i.adapter.MoveForward()
		}
		return pos
	}
}

const (
	prev int = -1
	cur  int = 0
	next int = +1
)

func (i *iterator) Remove() {
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
	} else {
		i.moveOnPrev = true
		i.moveOnNext = false
		i.adapter.MoveForward()
		i.adapter.RemoveElement(prev)
	}
}

func (i *iterator) Pos() int {
	return i.adapter.Pos()
}
