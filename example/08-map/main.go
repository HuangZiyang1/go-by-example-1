package main

import "fmt"

func main() {
	// 创建一个新的 map
	m := make(map[string]int)

	// 向 map 中添加键值对
	m["one"] = 1
	m["two"] = 2

	// 打印 map
	fmt.Println(m) // map[one:1 two:2]

	// 打印 map 的长度
	fmt.Println(len(m)) // 2

	// 通过键访问值
	fmt.Println(m["one"]) // 1

	// 访问不存在的键
	fmt.Println(m["unknow"]) // 0

	// 检查键是否存在
	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false

	// 删除键值对
	delete(m, "one")

	// 初始化带有值的 map
	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}

	// 打印初始化的 map
	fmt.Println(m2, m3)
}
