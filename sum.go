package utils

func BBC(data []byte) uint8 {
	var result uint8 = 0
	for _, d := range data {
		result ^= d
	}
	return result
}
