package utils

// BBC 异或校验和
func BBC(data []byte) uint8 {
	var result uint8 = 0
	for _, d := range data {
		result ^= d
	}
	return result
}

// Sum2 累加校验和
func Sum2(data []byte) uint16 {
	var sum uint16
	for _, c := range data {
		sum += uint16(c)
	}
	return sum
}
