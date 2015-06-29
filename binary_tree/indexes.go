package binary_tree

func GetIndexes(sliceToIndex []interface{}) ([]int) {
	var currentIndexes [][]int
	var allIndexes []int
	var indexes []int

	for n := 0; n  < len(sliceToIndex); n++ {
		indexes = append(indexes, n)
	}

	rootIndex := getRoot(len(sliceToIndex))
	allIndexes = append(allIndexes, rootIndex)

	currentIndexes = append(currentIndexes, indexes[:rootIndex])
	currentIndexes = append(currentIndexes, indexes[rootIndex + 1:])

	return getEachLevel(allIndexes, currentIndexes)
}

func getEachLevel(indexes []int, currentIndexes [][]int) ([]int) {
	if len(currentIndexes) == 0 {
		return indexes
	}

	var newIndexes [][]int

	for _, current := range currentIndexes {
		length := len(current)
		midpoint := length / 2

		if length > 0 {
			indexes = append(indexes, current[midpoint])
		}

		if length > 1 {
			newIndexes = append(newIndexes, current[:midpoint])
		}

		if length > 2 {
			newIndexes = append(newIndexes, current[midpoint + 1:])
		}
	}

	return getEachLevel(indexes, newIndexes)
}

func isEven(number int) bool {
	return (number / 2) + (number / 2) == number
}

func getRoot(number int) int {
	if isEven(number) {
		return (number / 2) - 1
	}
	return (number / 2)
}
