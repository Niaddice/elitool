package main

import (
	"elitool/structs"
	"fmt"
)

type stru struct {
	Name string
	Age  int
}

func main() {
	st := stru{Name: "姓名", Age: 1}
	m := structs.Struct2map(st, "Name", "")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
