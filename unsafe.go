// +build !js

package list

import (
	"reflect"
	"unsafe"
)

type unsafeSlice struct {
	size uintptr
	base unsafe.Pointer
	header unsafe.Pointer
}

func setUnsafeSliceBase(s *Slice) {
	s.size = s.sliceV.Type().Elem().Size()
	if s.sliceV.Cap() > 0 {
		oldLen := s.sliceV.Len()
		s.sliceV.SetLen(1)
		s.base = unsafe.Pointer(s.sliceV.Index(0).Addr().Pointer())
		s.sliceV.SetLen(oldLen)
	}
	s.header = unsafe.Pointer(s.sliceV.Addr().Pointer())
}

func copySlice(s *Slice, dstI, srcI, len int) {
	lenBytes := int(s.size * uintptr(len))

	data := uintptr(s.base) + s.size*uintptr(dstI)
	dst := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{Data: data, Len: lenBytes, Cap: lenBytes}))

	data = uintptr(s.base) + s.size*uintptr(srcI)
	src := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{Data: data, Len: lenBytes, Cap: lenBytes}))

	copy(dst, src)
}

func setSliceLen(s *Slice, len int) {
	header := (*reflect.SliceHeader)(s.header)
	header.Len = len
}