package main

import (
	"fmt"
	"github.com/akolybelnikov/codewars/go/beeramid"
)

func main() {
	num := beeramid.Beeramid(1500, 2)
	fmt.Printf("level: %d", num)
}
