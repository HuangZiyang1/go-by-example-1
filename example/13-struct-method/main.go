package main

import "fmt"

type user struct {
	name     string
	password string
}

// user 结构体的方法，检查密码是否正确，返回布尔值
// 第一次见这种这么写的方法
func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}
