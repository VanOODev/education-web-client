package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
)

func main() {
	l := list.NewList()
	l.Add(23)
	l.Add(35)
	l.Add(12)
	l.Add(8)
	l.Add(1)
	fmt.Println(l)
}
