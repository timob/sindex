package list
/*
 Slice list
 Uses reflection for allocation when growing
 Uses unsafe copy for remove/insert
 Slice contains the list in order, the length of the slice is the length of the list.

 The slice header (length/capacity) should *NOT* be modified directly.
 */

import (
	"reflect"
)

const reserveSize = 10
const growthFactor = 2

type SliceListInterface interface {
	ListInterface
	getListSliceStruct() *Slice
}

/*
 Slice List
 cap(slice) == capLen >= len(slice) == listLen

 After append/insert when listLen == capLen
 cap(slice) == capLen == (growthFactor * listLen) + reserveSize >= len(slice) == listLen
 */
type Slice struct {
	listLen int
	capLen int
	sliceV reflect.Value
	unsafeSlice
}

func NewSliceList(sliceList SliceListInterface, options ...OptionInterface) (ret interface{}) {
	ret = sliceList
	lv := reflect.ValueOf(sliceList).Elem()
	if lv.Kind() != reflect.Struct {
		return
	}
	ls := sliceList.getListSliceStruct()
	for i := 0 ; i < lv.NumField(); i++ {
		fv := lv.Field(i)
		if fv.Kind() == reflect.Slice {
			ls.sliceV = fv
			break
		}
	}
	setUnsafeSliceBase(ls)

	if (ls.sliceV == reflect.Value{}) {
		return
	}
	ls.listLen = ls.sliceV.Len()
	ls.capLen = ls.sliceV.Cap()
	return
}

func (s *Slice) getListSliceStruct() *Slice {
	return s
}

// set slice capacity
func (s *Slice) SetCap(size int) {
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
func (s *Slice) Clear() {
	s.sliceV.SetLen(0)
	s.listLen = 0
}

func (s *Slice) Append() (i int) {
	i = s.listLen
	if s.listLen == s.capLen {
		s.SetCap(s.listLen * growthFactor + reserveSize)
	}
	s.listLen++
	setSliceLen(s, s.listLen)
	return
}

func (s *Slice) Remove(pos int) {
	iter := s.Iterator(pos)
	if iter.Next() {
		iter.Remove()
	}
}

func (s *Slice) Insert(pos int) int {
	iter := s.Iterator(pos)
	if iter.Next() {
		return iter.Insert()
	} else if pos == 0 {
		return s.Append()
	}
	return 0
}

func (s *Slice) Pos(pos int) int {
	return pos
}

func (s *Slice) Len() int {
	return s.listLen
}

func (s *Slice) Iterator(pos int) IteratorInterface {
	if pos >= s.listLen {
		return NewEmptyIterator(&sliceIteratorAdapter{s, pos})
	} else {
		return NewIterator(&sliceIteratorAdapter{s, pos})
	}
}

type sliceIteratorAdapter struct {
	list *Slice
	pos int
}

func (s *sliceIteratorAdapter) AtLastElement() bool {
	return s.pos == s.list.listLen - 1
}

func (s *sliceIteratorAdapter) AtFirstElement() bool {
	return s.pos == 0
}

func (s *sliceIteratorAdapter) MoveForward() {
	s.pos++
}

func (s *sliceIteratorAdapter) MoveBack() {
	s.pos--
}

func (s *sliceIteratorAdapter) RemoveElement(relPos int) {
	pos := s.pos + relPos
	copySlice(s.list, pos, pos + 1, s.list.listLen - 1 - pos)
	if relPos == prev {
		s.pos--
	}
	s.list.listLen--
	setSliceLen(s.list, s.list.listLen)
}

func (s *sliceIteratorAdapter) InsertElement() int {
	if s.list.listLen == s.list.capLen {
		s.list.SetCap(s.list.listLen * growthFactor + reserveSize)
	}
	setSliceLen(s.list, s.list.listLen + 1)
	copySlice(s.list, s.pos+1, s.pos, s.list.listLen - s.pos)

	s.pos++
	s.list.listLen++
	return s.pos-1
}

func (s *sliceIteratorAdapter) Pos() int {
	return s.pos
}