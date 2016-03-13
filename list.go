package sindex

/*
 List
 Slice contains the list in order, the length of the slice is the length of the list.
 The slice header (length/capacity) should *NOT* be modified directly.

 Uses reflection for allocation when growing
 Uses unsafe copy for remove/insert (when supported)
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

// List satisfies this interface.

// Maps list positions directly to slice indexes. The length of the list is the length of the slice.
// The slice will grow if at capacity when calling Append/Insert. Satisfies ListInterface.
type List struct {
	listLen int
	capLen  int
	sliceV  reflect.Value
	unsafeSlice
}

/*
 cap(slice) == capLen >= len(slice) == listLen
 After append/insert when listLen == capLen:
 cap(slice) == capLen == (growthFactor * listLen) + reserveSize >= len(slice) == listLen
*/

// Returns a new List that manages the slice pointed to by slicePointer. slicePointer should be a pointer to a slice.
// The slice pointed to by slicePointer will not be changed, the new List will be of the same length.
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

// Initialize struct that contains a List and slice field which is pointed to by structPointer. structPointer is passed through as return value. Calls NewList to have List field manage slice field.
func InitList(structPointer ListInterface, options ...OptionInterface) (structPointerRet ListInterface) {
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

// Set slice capacity, size must be greater than list length.
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

// Clear the list to be empty. Slice capacity remains the same.
func (s *List) Clear() {
	s.sliceV.SetLen(0)
	s.listLen = 0
}

// Returns the index of where to put newly appended list item.
func (s *List) Append() (index int) {
	index = s.listLen
	if s.listLen == s.capLen {
		s.SetCap(s.listLen*growthFactor + reserveSize)
	}
	s.listLen++
	setSliceLen(s, s.listLen)
	return
}

// Remove item at given position. Slice capacity remains the same.
func (s *List) Remove(pos int) {
	iter := s.Iterator(pos)
	if iter.Next() {
		iter.Remove()
	}
}

// Insert item before given position. Returns index of where to put newly inserted list item.
func (s *List) Insert(pos int) (index int) {
	iter := s.Iterator(pos)
	if iter.Next() {
		return iter.Insert()
	} else if pos == 0 {
		return s.Append()
	}
	return 0
}

// Return the index of the given position. O(1) operation.
func (s *List) Pos(pos int) (index int) {
	return pos
}

// Return length of list.
func (s *List) Len() int {
	return s.listLen
}

// Return iterator. It is initially invalid. The given position will be the position of the iterator after a Next() or Prev() call.
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
