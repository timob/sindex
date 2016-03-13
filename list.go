package sindex

/*
 List
 Slice contains the list in order, the length of the slice is the length of the list.
 The slice header (length/capacity) should *NOT* be modified directly.

 Uses reflection for allocation when growing
 Uses unsafe copy for remove/insert
*/

import (
	"reflect"
)

const reserveSize = 10
const growthFactor = 2

type ListInterface interface {
	Interface
	getListStruct() *List
}

/*
 List
 cap(slice) == capLen >= len(slice) == listLen

 After append/insert when listLen == capLen
 cap(slice) == capLen == (growthFactor * listLen) + reserveSize >= len(slice) == listLen
*/
type List struct {
	listLen int
	capLen  int
	sliceV  reflect.Value
	unsafeSlice
}

func NewList(slicePointer interface{}, options ...OptionInterface) *List {
	pv := reflect.ValueOf(slicePointer)
	if pv.Kind() != reflect.Ptr || pv.Elem().Kind() != reflect.Slice {
		panic("slicePtr argument is not a pointer to a slice")
	}

	ls := &List{}
	ls.sliceV = pv.Elem()
	setUnsafeSliceBase(ls)
	ls.listLen = ls.sliceV.Len()
	ls.capLen = ls.sliceV.Cap()
	return ls
}

func InitList(structPointer ListInterface, options ...OptionInterface) (structPointerRet interface{}) {
	structPointerRet = structPointer
	ls := structPointer.getListStruct()
	pv := reflect.ValueOf(structPointer)
	if pv.Kind() != reflect.Ptr || pv.Elem().Kind() != reflect.Struct || ls == nil {
		return
	}
	lv := pv.Elem()
	var sv reflect.Value
	for i := 0; i < lv.NumField(); i++ {
		fv := lv.Field(i)
		if fv.Kind() == reflect.Slice {
			sv = fv
			break
		}
	}
	if (sv == reflect.Value{}) {
		return
	}

	nl := NewList(sv.Addr().Interface())
	*ls = *nl
	return
}

func (s *List) getListStruct() *List {
	return s
}

// set slice capacity
func (s *List) SetCap(size int) {
	if size > s.listLen {
		if size <= s.capLen {
			s.sliceV.SetCap(size)
		} else {
			newSlice := reflect.MakeSlice(s.sliceV.Type(), s.listLen, size)
			reflect.Copy(newSlice, s.sliceV)
			s.sliceV.Set(newSlice)
			setUnsafeSliceBase(s)
		}
		s.capLen = size
	}
}

// reset slice to zero length. which means list len is zero too
func (s *List) Clear() {
	s.sliceV.SetLen(0)
	s.listLen = 0
}

func (s *List) Append() (i int) {
	i = s.listLen
	if s.listLen == s.capLen {
		s.SetCap(s.listLen*growthFactor + reserveSize)
	}
	s.listLen++
	setSliceLen(s, s.listLen)
	return
}

func (s *List) Remove(pos int) {
	iter := s.Iterator(pos)
	if iter.Next() {
		iter.Remove()
	}
}

func (s *List) Insert(pos int) int {
	iter := s.Iterator(pos)
	if iter.Next() {
		return iter.Insert()
	} else if pos == 0 {
		return s.Append()
	}
	return 0
}

func (s *List) Pos(pos int) int {
	return pos
}

func (s *List) Len() int {
	return s.listLen
}

func (s *List) Iterator(pos int) IteratorInterface {
	if pos >= s.listLen {
		return newEmptyIterator(&listIteratorAdapter{s, pos})
	} else {
		return newIterator(&listIteratorAdapter{s, pos})
	}
}

type listIteratorAdapter struct {
	list *List
	pos  int
}

func (s *listIteratorAdapter) AtLastElement() bool {
	return s.pos == s.list.listLen-1
}

func (s *listIteratorAdapter) AtFirstElement() bool {
	return s.pos == 0
}

func (s *listIteratorAdapter) MoveForward() {
	s.pos++
}

func (s *listIteratorAdapter) MoveBack() {
	s.pos--
}

func (s *listIteratorAdapter) RemoveElement(relPos int) {
	pos := s.pos + relPos
	copySlice(s.list, pos, pos+1, s.list.listLen-1-pos)
	if relPos == prev {
		s.pos--
	}
	s.list.listLen--
	setSliceLen(s.list, s.list.listLen)
}

func (s *listIteratorAdapter) InsertElement() int {
	if s.list.listLen == s.list.capLen {
		s.list.SetCap(s.list.listLen*growthFactor + reserveSize)
	}
	setSliceLen(s.list, s.list.listLen+1)
	copySlice(s.list, s.pos+1, s.pos, s.list.listLen-s.pos)

	s.pos++
	s.list.listLen++
	return s.pos - 1
}

func (s *listIteratorAdapter) Pos() int {
	return s.pos
}
