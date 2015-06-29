package binary_tree

func GetIndexes(sliceToIndex []interface{}) ([]int) {
	sequentialIndexes := sequentialIndexesForLength(len(sliceToIndex))
	return getEachLevel(buildInitialSlices(len(sliceToIndex), sequentialIndexes))
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

func sequentialIndexesForLength(length int) (sequentialIndexes []int) {
	for index := 0; index  < length; index++ {
		sequentialIndexes = append(sequentialIndexes, index)
	}
	return
}

func buildInitialSlices(length int, sequentialIndexes []int) (allIndexes []int, currentIndexes [][]int) {
	rootIndex := getRoot(length)
	allIndexes = append(allIndexes, rootIndex)
	currentIndexes = append(currentIndexes, sequentialIndexes[:rootIndex])
	currentIndexes = append(currentIndexes, sequentialIndexes[rootIndex + 1:])
	return
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
