package main

import "fmt"

func main() {

}

// 常数阶
func ConstantTime(i, j int) int {
	var a, b int
	if i%2 == 0 {
		a = i + j
	}
	if j%2 == 1 {
		b = i + j
	}
	return a + b
}

// 对数阶
func LogarithmicTime(n int) {
	i := 1
	for i < n {
		i = i * 2
	}
}

// 线性阶
func LinearTime(n int) {
	// 几百万行代码
	for i := 0; i < n; i++ {
		// 几百万行代码
	}
	// 几百万行代码
}

// 线性对数阶
func LinearithmicTime(m, n int) {
	for i := 0; i < m; i++ {
		j := 1
		for j < n {
			j = j * 2
		}
	}
}

// 平方阶
func QuadraticTime(n int)  {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d\t%d", i, j)
		}
		fmt.Println()
	}
	
}
