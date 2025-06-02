package main

import (
	"fmt"
	"github.com/epicseven-cup/leetcode-cli/internal/leetcode"
)

func main() {
	fmt.Println("Hello world")
	lc := leetcode.Leetcode{}
	lc.GetDaily()
}
