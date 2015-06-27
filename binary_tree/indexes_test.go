package binary_tree_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/TopShelfBullard/binary_tree/binary_tree"
)

func TestIndexesAreProperlyAssociatedForAPerfectBinaryTree(t *testing.T) {
	Convey("Given a slice of numbers 1 through 63", t, func() {
		var numbers []interface{}
		number := 1

		for number <= 63 {
			numbers = append(numbers, number)
			number++
		}

		So(numbers[0], ShouldEqual, 1)
		So(numbers[62], ShouldEqual, 63)
		So(len(numbers), ShouldEqual, 63)

		Convey("When a slice is built to hold the indexes in the proper order for the creating of a binary tree", func() {
			binaryTreeIndexes := binary_tree.GetIndexes(numbers)
			So(len(binaryTreeIndexes), ShouldEqual, 63)

			Convey("The indexes are in the proper order for the creation of a binary tree", func() {
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
