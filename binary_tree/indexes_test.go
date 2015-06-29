package binary_tree_test

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/TopShelfBullard/binary_tree/binary_tree"
)

func TestIndexesAreProperlyAssociatedForAPerfectBinaryTree(t *testing.T) {
	Convey("Given a slice with a length of sixty-three", t, func() {
		var numbers []interface{}
		for number := 1; number <= 63; number++ {
			numbers = append(numbers, number)
		}
		fmt.Println(numbers)
		So(len(numbers), ShouldEqual, 63)

		Convey("When the indexes for creating a perfect binary tree are calculated", func() {
			binaryTreeIndexes := binary_tree.GetIndexes(numbers)
			So(len(binaryTreeIndexes), ShouldEqual, 63)

			Convey("The indexes are in the proper order", func() {
				expected := []int{31,15,47,7,23,39,55,3,11,19,27,35,43,51,59,1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,
					0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,52,54,56,58,60,62}
				actual := binaryTreeIndexes
				number := 0
				for number < len(expected) {
					So(actual[number], ShouldEqual, expected[number])
					number++
				}
			})
		})
	})
}

func TestIndexesAreProperlyAssociatedForANearPerfectBinaryTree(t *testing.T) {
	Convey("Given a slice with a length of forty", t, func() {
		var numbers []interface{}
		for number := 1; number <= 40; number++ {
			numbers = append(numbers, number)
		}
		fmt.Println(numbers)
		So(len(numbers), ShouldEqual, 40)

		Convey("When the indexes for creating a perfect binary tree are calculated", func() {
			binaryTreeIndexes := binary_tree.GetIndexes(numbers)
			So(len(binaryTreeIndexes), ShouldEqual, 40)

			Convey("The indexes are in the proper order", func() {
				expected := []int{19, 9, 30, 4, 14, 25, 35, 2, 7, 12, 17, 22, 28, 33, 38, 1, 3, 6, 8, 11, 13, 16, 18, 21,
					24, 27, 29, 32, 34, 37, 39, 0, 5, 10, 15, 20, 23, 26, 31, 36}
				actual := binaryTreeIndexes
				number := 0
				for number < len(expected) {
					So(actual[number], ShouldEqual, expected[number])
					number++
				}
			})
		})
	})
}

func TestTheAlgorithmCanHandleASmallSlice(t *testing.T) {
	Convey("Given a slice with a length of ten", t, func() {
		var numbers []interface{}
		for number := 1; number <= 10; number++ {
			numbers = append(numbers, number)
		}
		fmt.Println(numbers)
		So(len(numbers), ShouldEqual, 10)

		Convey("When the indexes for creating a perfect binary tree are calculated", func() {
			binaryTreeIndexes := binary_tree.GetIndexes(numbers)
			So(len(binaryTreeIndexes), ShouldEqual, 10)

			Convey("The indexes are in the proper order", func() {
				expected := []int{4, 2, 7, 1, 3, 6, 9, 0, 5, 8}
				actual := binaryTreeIndexes
				number := 0
				for number < len(expected) {
					So(actual[number], ShouldEqual, expected[number])
					number++
				}
			})
		})
	})
}

func TestTheAlgorithmCanHandleALargeSlice(t *testing.T) {
	Convey("Given a slice with a length of one million", t, func() {
		var numbers []interface{}
		for number := 1; number <= 1000000; number++ {
			numbers = append(numbers, number)
		}
		So(len(numbers), ShouldEqual, 1000000)

		Convey("When the indexes for creating a perfect binary tree are calculated, the length is as expected", func() {
			perfectBinaryTreeIndexes := binary_tree.GetIndexes(numbers)
			So(len(perfectBinaryTreeIndexes), ShouldEqual, 1000000)
		})
	})
}