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
	a := 0
	b := 0
	var l Storage = list.NewList()
	var sl Storage = slice.NewSliceStorage()

	for isTrue := true; isTrue != false; {
		fmt.Println("1: Односвязный список")
		fmt.Println("2: Срезы")
		fmt.Println("0: Выйти из программы")
		fmt.Scan(&a)
		switch a {

		case 1:
			for isTrue1 := true; isTrue1 != false; {
				fmt.Println("1: Добавить в элемент в лист")
				fmt.Println("2: Удалить элемент из листа")
				fmt.Println("3: Показать элемент")
				fmt.Println("0: Выйти из подпрограммы")
				fmt.Scan(&b)
				switch b {

				case 1:
					fmt.Println("Введите значение")
					var temp int64 = 0
					fmt.Scan(&temp)
					l.Add(temp)
				case 2:
					fmt.Println("Введите индекс")
					var temp int64 = 0
					fmt.Scan(&temp)
					l.Delete(temp)
				case 3:
					fmt.Println("Введите индекс")
					var temp int64 = 0
					fmt.Scan(&temp)
					fmt.Println(l.Get(temp))
				case 0:
					isTrue1 = false
				}
			}

		case 2:
			for isTrue1 := true; isTrue1 != false; {
				fmt.Println("1: Добавить в элемент в срез")
				fmt.Println("2: Удалить элемент из среза")
				fmt.Println("3: Показать элемент")
				fmt.Println("0: Выйти из подпрограммы")
				fmt.Scan(&b)
				switch b {

				case 1:
					fmt.Println("Введите значение")
					var temp int64 = 0
					fmt.Scan(&temp)
					sl.Add(temp)
				case 2:
					fmt.Println("Введите индекс")
					var temp int64 = 0
					fmt.Scan(&temp)
					sl.Delete(temp)
				case 3:
					fmt.Println("Введите индекс")
					var temp int64 = 0
					fmt.Scan(&temp)
					fmt.Println(sl.Get(temp))

				case 0:
					isTrue1 = false
				}
			}

		case 0:
			isTrue = false
		}
	}
	fmt.Println("До свидания!")
}
