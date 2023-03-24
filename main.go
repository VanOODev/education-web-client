package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
)

func main() {
	l := list.NewList()
	l.Add(1)
	l.Add(2)
	l.Add(4)
	l.Add(3)
	//l.Add(10)
	//l.Add(40)
	//l.Add(60)
	//l.Add(50)
	l.SortIncrease2()
	l.Add(20)
	l.Add(19)
	l.Add(18)
	l.Add(17)
	fmt.Println(l.String())
	fmt.Println()
	l.SortIncrease2()

	//l.Delete(4)

	fmt.Println(l.String())
}
