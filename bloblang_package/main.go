package main

import (
	"fmt"

	"github.com/warpstreamlabs/bento/public/bloblang"

	// for the bloblang methods zip amongst others
	_ "github.com/warpstreamlabs/bento/public/components/pure"
)

func main() {
	s := "hello world"
	fmt.Println(s)

	mapping := `root = this.uppercase()`

	exe, err := bloblang.Parse(mapping)
	if err != nil {
		panic(err)
	}

	res, err := exe.Query(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
