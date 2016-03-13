sindex
======

Slice indexing library for Golang.

####Documentation
https://godoc.org/github.com/timob/sindex

####Example
```go
func ExampleList() {
	bytes := []byte("helloworld")
	bl := NewList(&bytes)
	bytes[bl.Insert(5)] = ' '
	bytes[bl.Append()] = '!'
	for iter := bl.Iterator(0); iter.Next(); {
		fmt.Print(string(bytes[iter.Pos()]))
	}
	// Output: hello world!
}
```

####Projects using this library
* GNU Ls clone for win/unix https://github.com/timob/ls
