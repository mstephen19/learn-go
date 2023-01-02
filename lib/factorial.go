package lib

import "errors"

func Factorial(num int) int {
	if (num == 0) {
		return 1
	}

	return num * Factorial(num - 1)
}

// Find the total number of possible combinations in a list given the size of the
// list and the group count
func NumberOfCombinations(itemCount int, groupSize int) (result int, err error) {
	// If they are equal, the answer will be one
	if (groupSize == itemCount) {
		return 1, nil
	}

	// If the group size is larger, incorrect answers will result, so we must error
	if (groupSize > itemCount) {
		return -1, errors.New("group size must be less than or equal to item count")
	}

	return Factorial(itemCount) / (Factorial(itemCount - groupSize) * Factorial(groupSize)), nil
}