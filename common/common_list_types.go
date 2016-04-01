package common

import "github.com/timob/sindex"

type ByteList struct {
	Data []byte
	sindex.List
}

type IntList struct {
	Data []int
	sindex.List
}

type StringList struct {
	Data []string
	sindex.List
}

type InterfaceList struct {
	Data []interface{}
	sindex.List
}
