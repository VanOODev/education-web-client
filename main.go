package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
)

func main() {
	l := list.NewList()
	l.Add(10)
	l.Add(20)
	l.Add(30)
	l.Add(40)
	l.Add(50)
	//l.Delete(5)
	//l.Add(60)
	fmt.Println(l.String())
}
