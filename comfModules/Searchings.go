package search


func binarrySearch(num int, array []int, offset int) (int, bool) {
	if len(array) == 0 {
		return -1, false
	}
	middle := len(array) / 2
	if num == array[middle] {
		return middle + offset, true
	}
	if num > array[middle] {
		return binarrySearch(num, array[middle+1:], middle+offset+1)
	}
	return binarrySearch(num, array[:middle], offset)
}
