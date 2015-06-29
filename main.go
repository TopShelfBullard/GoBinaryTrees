package main

import (
	"fmt"
	"github.com/TopShelfBullard/binary_tree/binary_tree"
)

func main() {
	var numbers []interface{}
	for number := 1; number <= 1000000; number++ {
		numbers = append(numbers, number)
	}

	perfectBinaryTreeIndexes := binary_tree.GetIndexes(numbers)

	fmt.Println(len(numbers), ":", len(perfectBinaryTreeIndexes))
	fmt.Println(perfectBinaryTreeIndexes)
}