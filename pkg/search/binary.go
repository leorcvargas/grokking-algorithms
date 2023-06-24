package search

const BinarySearchNotFound = -1

func BinarySearch(list []int, target int) int {
	startPosition := 0
	endPosition := len(list) - 1

	for startPosition <= endPosition {
		centerPosition := (startPosition + endPosition) / 2

		try := list[centerPosition]

		if try == target {
			return centerPosition
		}

		if try > target {
			endPosition = centerPosition - 1
		} else {
			startPosition = centerPosition + 1
		}
	}

	return BinarySearchNotFound
}
