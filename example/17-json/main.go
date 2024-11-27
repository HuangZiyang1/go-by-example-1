package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	// a is an instance of the userInfo struct, initialized with the following values:
	// Name: "wang"
	// Age: 18
	// Hobby: a slice containing "Golang" and "TypeScript"
	// main.go 文件展示了如何在 Go 语言中使用 JSON。
	// 该示例定义了一个名为 userInfo 的结构体，并创建了一个包含用户信息的实例。
	// 结构体包含用户的姓名、年龄和爱好列表。
	// 在这个示例中，创建了一个名为 a 的 userInfo 实例，
	// 其中 Name 字段为 "wang"，Age 字段为 18，Hobby 字段为一个包含 "Golang" 和 "TypeScript" 的切片。
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)         // [123 34 78 97...]
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}
