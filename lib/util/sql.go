package util

import "math"

func CreatePageCount(count, pageSize int) (page int64) {
	page = int64(math.Ceil(float64(float64(count) / float64(pageSize))))
	return
}
