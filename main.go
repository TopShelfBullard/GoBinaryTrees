package main

import (
	"fmt"
	"github.com/TopShelfBullard/binary_tree/binary_tree"
)

func main() {
	var numbers []interface{}
	number := 0

	for number < 1000 {
		numbers = append(numbers, number)
		number++
	}

	indexes := binary_tree.GetIndexes(numbers)
	fmt.Println(len(numbers))
	fmt.Println(len(indexes))
}