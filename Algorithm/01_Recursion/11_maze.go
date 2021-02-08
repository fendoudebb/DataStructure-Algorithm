package main

import "fmt"

// 迷宫
func main() {
	maze := make([][7]int, 8)
	//Print(maze)

	for i := 0; i < 8; i++ {
		maze[i][0] = 1
		maze[i][6] = 1
		if i < 7 {
			maze[0][i] = 1
			maze[7][i] = 1
		}
	}

	//for j := 0; j < 7; j++ {
	//	maze[0][j] = 1
	//	maze[7][j] = 1
	//}

	maze[3][1] = 1
	maze[3][2] = 1

	maze[3][3] = 1
	maze[2][3] = 1
	maze[1][3] = 1


	Print(maze)

	SetWay(maze, 1, 1)

	//Print(maze)

}

func SetWay(maze [][7]int, i, j int) bool {
	if maze[6][5] == 2 {
		return true
	} else {
		if maze[i][j] == 0 {
			maze[i][j] = 2
			Print(maze)
			// 下 右 上 左
			if SetWay(maze, i+1, j) { // 往下
				return true
			} else if SetWay(maze, i, j+1) { // 往右
				return true
			} else if SetWay(maze, i-1, j) { // 往上
				return true
			} else if SetWay(maze, i, j-1) { // 往左
				return true
			} else {
				maze[i][j] = 3
				Print(maze)
				return false
			}
		} else {
			// 数值为1, 2, 3的情况都返回false
			return false
		}
	}
}

func Print(maze [][7]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d\t", maze[i][j])
		}
		fmt.Println()
	}

	fmt.Println("-------------------------")
}
