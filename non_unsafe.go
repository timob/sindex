// +build js

package list

import (
	"reflect"
)

type unsafeSlice struct {}

func setUnsafeSliceBase(s *Slice) {}

func copySlice(s *Slice, dstI, srcI, len int) {
	reflect.Copy(s.sliceV.Slice(dstI, dstI + len), s.sliceV.Slice(srcI, srcI+len))
}

func setSliceLen(s *Slice, len int) {
	s.sliceV.SetLen(len)
}