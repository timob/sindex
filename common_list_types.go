package list

type ByteSlice struct {
	Data []byte
	Slice
}

type StringSlice struct {
	Data []string
	Slice
}

type InterfaceSlice struct {
	Data []interface{}
	Slice
}