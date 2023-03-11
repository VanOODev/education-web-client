package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
)

func main() {
	l := list.NewList()
	in := l.Add(23)
	fmt.Println(in)
}
