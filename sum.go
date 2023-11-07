package utils

// BBC 异或校验和
func BBC(data []byte) []byte {
	var result uint8 = 0
	for _, d := range data {
		result ^= d
	}
	return []byte{result}
}

// Sum2BE 累加校验和 大段
func Sum2BE(data []byte) []byte {
	var sum uint16
	for _, c := range data {
		sum += uint16(c)
	}
	return []byte{uint8(sum / 256), uint8(sum % 256)}
}

// Sum2LE 累加校验和 小段
func Sum2LE(data []byte) []byte {
	var sum uint16
	for _, c := range data {
		sum += uint16(c)
	}
	return []byte{uint8(sum % 256), uint8(sum / 256)}
}
