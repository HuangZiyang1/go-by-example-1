package main

import "fmt"

// 两个add没区别，只是在参数的类型写法上有区别
func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

// 返回两个值，一个是值，一个是是否存在，如果存在返回true，否则返回false
// 返回值的类型写在括号外面
// 返回值可以命名，也可以不命名
// 返回值可以有多个
// 返回值可以是任意类型
// 返回值可以是函数
// 参数的类型写在括号里面，参数可以有多个，参数的类型可以是任意类型
func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}
