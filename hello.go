package main

import (
	"fmt"

	"example/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("olleh"))
	fmt.Println(cmp.Diff("hello world", "hello go"))
}

