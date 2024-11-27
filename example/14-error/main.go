package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	// 这个_是什么意思？
	// 这个_是占位符，表示不关心这个值
	// 这个_是一个特殊的变量，不能被赋值
	// 这个_是一个只写变量，不能被读取
	// 这个_是一个只读变量，不能被赋值
	// 这是_是索引，表示索引
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
