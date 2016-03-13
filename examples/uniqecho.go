// print command line arguments in the order they were given removing duplicates.
package main

import (
	"github.com/timob/sindex"
	"fmt"
	"os"
)

func main() {
	args := sindex.InitList(&sindex.StringList{Data:append([]string{}, os.Args...)}).(*sindex.StringList)
	args.Remove(0)
	argSet := make(map[string]struct{})
	for iter := args.Iterator(0); iter.Next(); {
		v := args.Data[iter.Pos()]
		if _, ok := argSet[v] ; ok {
			iter.Remove()
		} else {
			argSet[v] = struct{}{}
			if iter.Pos() != 0 {
				args.Data[iter.Insert()] = " "
			}
		}
	}
	for _, v := range args.Data {
		fmt.Print(v)
	}
	fmt.Println()
}
