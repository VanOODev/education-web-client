package main

import (
	"context"
	"fmt"
	"github.com/VanOODev/education-web-client/storages/list"
	"github.com/VanOODev/education-web-client/storages/slice"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/jackc/pgx/v4"
	"log"
)

type Storage interface {
	Add(data int64) (index int64)
	Get(index int64) (data int64)
	Delete(index int64)
	String() string
}

type Config struct {
	Databasest struct {
		DatabaseUrl string `hcl:"Database_url"`
	} `hcl:"Database,block"`
	StorageTypest struct {
		StorageType string `hcl:"Storage_type"`
	} `hcl:"Storagetype,block"'`
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

	switch cfg.StorageTypest.StorageType {

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
		conn, err := pgx.Connect(context.Background(), cfg.Databasest.DatabaseUrl)
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
		defer conn.Close(context.Background())

		var exists bool
		err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "mytable").Scan(&exists)
		if err != nil {
			panic(err)
		}

		if !exists {
			// Создаем таблицу, если она не существует
			_, err := conn.Exec(context.Background(), "CREATE TABLE mytable (id SERIAL PRIMARY KEY, data integer NOT NULL)")
			if err != nil {
				panic(err)
			}
			fmt.Println("Table created successfully")
		} else {
			fmt.Println("Table already exists")
		}

		for isTrue1 := true; isTrue1 != false; {
			fmt.Println("1: Добавить в элемент в базу данных")
			fmt.Println("2: Удалить элемент из базы данных")
			fmt.Println("3: Показать элемент")
			fmt.Println("0: Выйти из подпрограммы")
			fmt.Scan(&b)
			switch b {

			case 1:
				fmt.Println("Введите значение")
				var data int
				fmt.Scan(&data)
				stmt := `INSERT INTO mytable (data) VALUES ($1)`
				_, err = conn.Exec(context.Background(), stmt, data)
				if err != nil {
					fmt.Println("Error inserting item:", err)
					return
				}
				fmt.Println("Item inserted successfully")

			case 2:
				fmt.Println("Введите индекс")
				var id int
				fmt.Scan(&id)
				stmt := `DELETE FROM mytable WHERE id = $1`
				_, err = conn.Exec(context.Background(), stmt, id)
				if err != nil {
					fmt.Println("Error deleting:", err)
					return
				}
				fmt.Println("Deleted")

			case 3:
				fmt.Println("Введите индекс")
				var id, data int
				fmt.Scan(&id)
				stmt := `SELECT data FROM mytable WHERE id = $1`
				err = conn.QueryRow(context.Background(), stmt, id).Scan(&data)
				if err != nil {
					fmt.Println("Error getting:", err)
					return
				}
				fmt.Println(data)

			case 0:
				isTrue1 = false
			}
		}

		fmt.Println("До свидания!")
	}
}
