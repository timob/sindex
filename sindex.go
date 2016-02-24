package sindex

import (
	"errors"
)

var IndexOutOfBounds = errors.New("List index out of bounds")

type IteratorInterface interface {
	Next() bool
	Prev() bool
	Pos() int
	Insert() int
	Remove()
	//	Err() (err error)
}

type Interface interface {
	Append() int
	Remove(int)
	Insert(int) int
	Pos(int) int
	Len() int
	Iterator(int) IteratorInterface
	Clear()
	//	Err() (err error)
}

type Option struct{}

func (o *Option) listOption() {}

type OptionInterface interface {
	listOption()
}
