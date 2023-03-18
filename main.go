package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
	"github.com/VanOODev/education-web-client/storages/slice"
)

type Storage interface {
	Add(data int64) (index int64)
	Get(index int64) (data int64)
	Delete(index int64)
	String() string
}

func main() {
	var l Storage = list.NewList()
	var sl Storage = slice.NewSliceStorage()
	l.Add(10)
	l.Add(30)
	sl.Add(20)
	sl.Add(50)
	//	fmt.Println(l.Get(3))
	fmt.Println(l.String())
	fmt.Print(sl.String())

}
