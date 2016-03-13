package sindex

// List of byte. Use with InitList.
type ByteList struct {
	Data []byte
	List
}

// List of string. Use with InitList.
type StringList struct {
	Data []string
	List
}

// List of interface. Use with InitList.
type InterfaceList struct {
	Data []interface{}
	List
}
