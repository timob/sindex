package sindex

type ByteList struct {
	Data []byte
	List
}

type StringList struct {
	Data []string
	List
}

type InterfaceList struct {
	Data []interface{}
	List
}
