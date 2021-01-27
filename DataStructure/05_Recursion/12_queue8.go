package main

import (
	"fmt"
	"math"
)

func main() {

	// [0, 1, 2, 3, 4, 5, 6, 7]
	// 数组的索引i表示第i+1行皇后，索引从0开始，表示第1行
	// 元素的值arr[i]=value，value表示第value+1列
	queue := make([]int, 8)
	Check(queue, 0)

	fmt.Println(count)
	fmt.Println(countJudge)

}

func Check(queue []int, n int) {
	if n == 8 {
		PrintQueue(queue)
		return
	}
	for i := 0; i < 8; i++ {
		queue[n] = i
		if Judge(queue, n) {
			Check(queue, n+1)
		}
	}
}


var countJudge = 0

// 判断摆放的皇后和已经摆放的皇后是否有冲突
func Judge(queue []int, n int) bool {
	countJudge++
	// 之前摆放的皇后依此判断
	for i := 0; i < n; i++ {
		// queue[i] == queue[n] 判断是否在同一列
		// math.Abs(float64(n-i)) == math.Abs(float64(queue[n]-queue[i]) 判断是否在同一斜线
		if queue[i] == queue[n] || math.Abs(float64(n-i)) == math.Abs(float64(queue[n]-queue[i])) {
			return false
		}
	}
	return true
}

var count = 0

func PrintQueue(queue []int) {
	count++
	for _, value := range queue {
		fmt.Printf("%d\t", value)
	}
	fmt.Println()
}
