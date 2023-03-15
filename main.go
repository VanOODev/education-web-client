package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
)

func main() {
	l := list.NewList()
	l.Add(100)
	l.Add(20)
	l.Add(10)
	l.Add(40)
	l.Add(60)
	l.Add(50)
	l.SortIncrease()
	//l.Delete(4)

	fmt.Println(l.String())
}
