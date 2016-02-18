package utils

func RemoveIntFromInts(i int64, ints []int64) []int64 {
	index := 0
	endIndex := len(ints) - 1
	var result = make([]int64, 0)
	for k, v := range ints {
		if v == i {
			result = append(result, ints[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, ints[index:endIndex+1]...)
		}
	}
	return result
}
func InInts(i int64, ints []int64) bool {
	for _, v := range ints {
		if i == v {
			return true
		}
	}
	return false
}
