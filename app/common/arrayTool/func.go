package arrayTool

func InArray[T comparable](search T, array []T) bool {
	_, ok := Find(search, array)
	return ok
}

// Find获取一个切片并在其中查找元素。如果找到它，它将返回它的密钥，否则它将返回-1和一个错误的bool。
func Find[T comparable](search T, array []T) (int, bool) {
	for index, value := range array {
		if value == search {
			return index, true
		}
	}
	return -1, false
}
