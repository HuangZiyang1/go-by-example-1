package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	// 创建一个结构体，然后初始化
	// 有两种初始化方式
	// 全写
	a := user{name: "wang", password: "1024"}
	// 按顺序写
	b := user{"wang", "1024"}
	// 也可以只初始化部分字段
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}
