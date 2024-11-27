package main

import "fmt"

func main() {

	var a [5]int
	a[4] = 100
	// 获取值，默认为0
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))
	fmt.Println(a)

	// 声明并初始化，用大括号包裹
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 二维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
