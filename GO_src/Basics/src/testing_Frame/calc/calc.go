package calc

// 函数一，正确
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

// 函数二，错误
func getSubtract(n, n1 int) int {
	return n + n1
}
