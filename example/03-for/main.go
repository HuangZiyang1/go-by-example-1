package main

import "fmt"

func main() {

	i := 1
	// for循环，这个只运行了一次，因为是无限循环，但是有break
	// 如果注释掉break，会一直打印loop，并且程序会提示警告信息
	for {
		fmt.Println("loop")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
