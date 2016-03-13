package sindex

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	list := NewList(&data)
	for iter := list.Iterator(0); iter.Next(); {
		if data[iter.Pos()] == 'h' {
			iter.Remove()
		} else if data[iter.Pos()] == 'l' {
			data[iter.Insert()] = 'l'
			iter.Next()
		} else if data[iter.Pos()] == 'o' {
			data[list.Append()] = 's'
		}
	}
	t.Logf("list len = %d cap = %d", len(data), cap(data))
	if string(data) != "elllos" {
		t.Fatalf("unexpected list value %s ", data)
	}
}

type stringList struct {
	Data []string
	List
}

func TestInterfaceList(t *testing.T) {
	list := InitList(&stringList{}).(*stringList)

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

func TestSimpleIterator(t *testing.T) {
	list := InitList(&stringList{}).(*stringList)

	list.Data[list.Append()] = "aaa"
	list.Data[list.Append()] = "bbb"
	iter := list.Iterator(1)
	iter.Next()
	if list.Data[iter.Pos()] != "bbb" {
		t.Fatal("fail")
	}
	iter.Prev()
	if list.Data[iter.Pos()] != "aaa" {
		t.Fatalf("fail %v", list.Data[iter.Pos()])
	}
}

func ExampleIterator() {
	list := InitList(&stringList{}).(*stringList)

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
	List
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
	s := &fixedStack{3, InitList(&interfaceList{}).(*interfaceList)}
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
