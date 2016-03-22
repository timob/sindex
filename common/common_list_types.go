package common

import sindex ".."

// List of byte. Use with InitList.
type ByteList struct {
	Data []byte
	sindex.List
}

// List of int. Use with InitList.
type IntList struct {
	Data []int
	sindex.List
}

// List of string. Use with InitList.
type StringList struct {
	Data []string
	sindex.List
}

// List of interface. Use with InitList.
type InterfaceList struct {
	Data []interface{}
	sindex.List
}

func NewByteList(s *[]byte) *sindex.List {
	return sindex.NewList(s)
}

func NewIntList(s *[]int) *sindex.List {
	return sindex.NewList(s)
}

func NewStringList(s *[]string) *sindex.List {
	return sindex.NewList(s)
}

func NewInterfaceList(s *[]interface{}) *sindex.List {
	return sindex.NewList(s)
}
