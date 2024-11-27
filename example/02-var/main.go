package main

import (
	"fmt"
	"math"
)

func main() {

	// 同js
	var a = "initial"
	// 使用var后，指定了b，c的类型
	var b, c int = 1, 2
	// 使用var后
	var d = true
	// 使用var后，指定了e的类型
	var e float64
	// 使用:=，省略了var，类型自动判断
	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	// 常量
	const s string = "constant"
	const h = 500000000
	// 一个数字可以使用科学计数法
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
