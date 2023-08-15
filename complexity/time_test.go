package complexity

// 时间复杂度
// https://www.hello-algo.com/chapter_computational_complexity/time_complexity/#224

// 常数阶
func timeConstant(n int) {
	count := 0
	for i := 0; i < 5; i++ {
		count++
	}
}

// 线性阶
func timeLinear(n int) {
	count := 0
	for i := 0; i < n; i++ {
		count++
	}
}

// 平方阶
func timeQuadratic(n int) {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count++
		}
	}
}

/* 指数阶（循环实现）*/
func timeExponential(n int) {
	count := 0
	base := 1
	for i := 0; i < n; i++ {
		for j := 0; j < base; j++ {
			count++
		}
		base *= 2
	}
}

/* 指数阶（递归实现）*/
func timeExpRecur(n int) int {
	if n <= 1 {
		return 1
	}
	return timeExpRecur(n-1) + timeExpRecur(n-1)
}

/* 对数阶（循环实现）*/
func timeLogarithmic(n int) {
	for n > 1 {
		n /= 2
	}
}

/* 对数阶（递归实现）*/
func timeLogRecur(n int) {
	if n <= 1 {
		return
	}
	timeLogRecur(n / 2)
}

/* 线性对数阶 快速排序、归并排序、堆排序*/
func timeLinearLogRecur(n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
	}
	timeLinearLogRecur(n / 2)
}

/* 阶乘阶（递归实现） */
func timeFactorialRecur(n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		timeFactorialRecur(n - 1)
	}
}
