package list

import (
	"testing"
	"bytes"
	"fmt"
)

type byteList struct {
	Data []byte
	Slice
}

func TestSlice(t *testing.T) {
	data := []byte{'h','e','l','l','o'}
	list := NewSliceList(&byteList{Data: data}).(*byteList)
	for iter := list.Iterator(0); iter.Next(); {
		if list.Data[iter.Pos()] == 'h' {
			iter.Remove()
		} else if list.Data[iter.Pos()] == 'l' {
			list.Data[iter.Insert()] = 'l'
			iter.Next()
		} else if list.Data[iter.Pos()] == 'o' {
			list.Data[list.Append()] = 's'
		}
	}
	t.Logf("list len = %d cap = %d", len(list.Data), cap(list.Data))
	if bytes.Compare(list.Data, []byte("elllos")) != 0 {
		t.Fatalf("unexpected list value %s ", list.Data)
	}
}

type stringList struct {
	Data []string
	Slice
}

func TestInterfaceList(t *testing.T) {
	list := NewSliceList(&stringList{}).(*stringList)

	list.Data[list.Append()] = "Alpha"
	list.Data[list.Append()] = "hello world"
	list.Data[list.Append()] = "testdata testdata"

	list.Data[list.Insert(1)] = "XY"

	list.Remove(3)

	found := false
	for iter := list.Iterator(list.Pos(0)); iter.Next(); {
		if list.Data[iter.Pos()] == "hello world" {
			found = true
		}
		t.Logf(list.Data[iter.Pos()])
	}
	if !found {
		t.Fatal("unexpected to contain hello world")
	}
}

func ExampleIterator() {
	list := NewSliceList(&stringList{}).(*stringList)

	list.Data[list.Append()] = "Alpha"
	list.Data[list.Append()] = "testdata"
	list.Data[list.Append()] = "testdata2"

	for iter := list.Iterator(list.Pos(0)); iter.Next(); {
		list.Data[iter.Insert()] = "separator"
		if list.Data[iter.Pos()] == "testdata" {
			iter.Remove()
		}
	}
	for iter := list.Iterator(list.Pos(0)); iter.Next(); {
		fmt.Printf("%s\n", list.Data[iter.Pos()])
	}
	// Output:
	// separator
	// Alpha
	// separator
	// separator
	// testdata2
}


type interfaceList struct {
	Data []interface{}
	Slice
}

type fixedStack struct {
	length int
	*interfaceList
}

func (f *fixedStack) Push(element interface{}) {
	fmt.Printf("push %v\n", element)
	f.Data[f.Insert(0)] = element
	f.Remove(f.length)

	for x := f.Iterator(0); x.Next(); {
		fmt.Printf("stack %v\n", f.Data[x.Pos()])
	}

}

func useStack(s *fixedStack) {
	s.Push(1)
	s.Push(0)
	s.Push(2)
	s.Push(1)
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(1)
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(1)
	s.Push(0)

	l := s.interfaceList
	l.Remove(0)
	l.Remove(0)
	l.Remove(0)
	l.Data[l.Append()] = 10
	fmt.Printf("%v", l.Data[l.Pos(0)])
}

func ExampleStack() {
	s := &fixedStack{3, NewSliceList(&interfaceList{}).(*interfaceList)}
	useStack(s)
	// Output:
	// push 1
	// stack 1
	// push 0
	// stack 0
	// stack 1
	// push 2
	// stack 2
	// stack 0
	// stack 1
	// push 1
	// stack 1
	// stack 2
	// stack 0
	// push 0
	// stack 0
	// stack 1
	// stack 2
	// push 1
	// stack 1
	// stack 0
	// stack 1
	// push 2
	// stack 2
	// stack 1
	// stack 0
	// push 1
	// stack 1
	// stack 2
	// stack 1
	// push 0
	// stack 0
	// stack 1
	// stack 2
	// push 1
	// stack 1
	// stack 0
	// stack 1
	// push 2
	// stack 2
	// stack 1
	// stack 0
	// push 1
	// stack 1
	// stack 2
	// stack 1
	// push 0
	// stack 0
	// stack 1
	// stack 2
	// 10
}
