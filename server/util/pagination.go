package util

func GetLimitAndOffset(limit *int, offset *int) (int, int) {
	var l = 10
	if limit != nil {
		l = *limit
	}

	var o = 10
	if offset != nil {
		o = *offset
	}
	return l, o
}
