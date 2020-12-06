## SIndex


Slice Indexing library for Golang.

#### Documentation
https://godoc.org/github.com/timob/sindex

#### Example
```go
func ExampleList() {
	bytes := []byte("helloworld")
	bl := NewList(&bytes)
	bytes[bl.Insert(5)] = ' '
	bytes[bl.Append()] = '!'

	fmt.Println(string(bytes))

	for iter := bl.Iterator(0); iter.Next(); {
		fmt.Print(string(bytes[iter.Pos()]))
	}
	// Output:
	// hello world!
	// hello world!
}
```

#### Projects using this library
* GNU Ls clone for win/unix https://github.com/timob/ls
