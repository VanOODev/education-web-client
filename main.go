package main

import (
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
	"github.com/VanOODev/education-web-client/storages/slice"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Storage interface {
	Add(data int64) (index int64)
	Get(index int64) (data int64)
	Delete(index int64)
	String() string
}

type Config struct {
	Database struct {
		DBName   string `hcl:"dbname"`
		User     string `hcl:"user"`
		Password string `hcl:"password"`
		Host     string `hcl:"host"`
		Port     int    `hcl:"port"`
	} `hcl:"database,block"`
	StorageType struct {
		Storage_Type string `hcl:"storage_type"`
	} `hcl:"storagetype,block"'`
}

func main() {
	var cfg Config
	err := hclsimple.DecodeFile("config.hcl", nil, &cfg)
	if err != nil {
		fmt.Printf("Ошибка загрузки файла конфигурации: %v\n", err)
		return
	}
	b := 0
	var l Storage = list.NewList()
	var sl Storage = slice.NewSliceStorage()

	switch cfg.StorageType.Storage_Type {

	case "Лист":
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

	case "Срез":
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
	case "База данных":

		for isTrue1 := true; isTrue1 != false; {
			fmt.Println("1: Добавить в элемент в базу данных")
			fmt.Println("2: Удалить элемент из базы данных")
			fmt.Println("3: Показать элемент")
			fmt.Println("0: Выйти из подпрограммы")
			fmt.Scan(&b)
			switch b {

			case 1:
				fmt.Println("Введите значение")

			case 2:
				fmt.Println("Введите индекс")

			case 3:
				fmt.Println("Введите индекс")

			case 0:
				isTrue1 = false
			}

		}
		fmt.Println("До свидания!")
	}
}
