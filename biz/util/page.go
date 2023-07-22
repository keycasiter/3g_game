package util

func PageNoToOffset(pageNo int64, pageSize int64) int {
	return int(pageNo) * int(pageSize)
}
