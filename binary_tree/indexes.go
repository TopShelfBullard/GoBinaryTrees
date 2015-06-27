package binary_tree

import (
	"math"
)

func GetIndexes(sliceToIndex []interface{}) (allIndexes []int) {
	var currentIndexes []int
	rootNodeIndex := getRoot(len(sliceToIndex))
	currentIndexes = append(currentIndexes, rootNodeIndex)
	allIndexes = append(allIndexes, rootNodeIndex)

	for len(currentIndexes) > 0 {
		currentIndexes = getNext(currentIndexes)
		if len(currentIndexes) > 0 {
			for _, index := range currentIndexes {
				allIndexes = append(allIndexes, index)
			}
		}
	}

	allIndexes = unOffByOneify(allIndexes, len(sliceToIndex))
	return
}

func unOffByOneify(wrongIndexes []int, expectedLength int) (rightIndexes []int) {
	if len(wrongIndexes) < expectedLength {
		wrongIndexes = append(wrongIndexes, expectedLength)
	}

	for _, wrongIndex := range wrongIndexes {
		rightIndexes = append(rightIndexes, wrongIndex - 1)
	}
	return
}

func getNext(indexes []int) (newIndexes []int) {
	var low, high, previousLow, previousHigh int
	diff := getMiddle(indexes[0])

	for _, index := range indexes {
		low = index - diff
		high = index + diff

		if diff > 1 || (diff == 1 && thereWillNotBeDuplicates(low, high, previousLow, previousHigh)) {
			newIndexes = append(newIndexes, low)
			newIndexes = append(newIndexes, high)
		}

		previousLow = low
		previousHigh = high
	}

	return
}

func thereWillNotBeDuplicates(low int, high int, previousLow int, previousHigh int) bool {
	return low != previousLow && low != previousHigh && high != previousLow && high != previousHigh
}

func getMiddle(number int) int {
	return int(math.Floor(float64(number / 2)))
}

func getRoot(number int) int {
	return int(math.Floor(float64(number / 2)) + 1)
}
