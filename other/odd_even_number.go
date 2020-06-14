package other

// 判断奇偶数：通过模2的方法
func IsEvenNumber(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

// 判断奇偶数：通过与1做与运算的方法
func IsEvenNumber1(n int) bool {
	if (n & 0x1) == 1 {
		return false
	} else {
		return true
	}
}
