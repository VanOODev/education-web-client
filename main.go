package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/slice"
)

func main() {
	l := list.NewList()
	in := l.Add(23)
	fmt.Println(in)
}
