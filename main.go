package main

import (
	"fmt"
	"shopping-list/handlers"
	"shopping-list/models"
	"shopping-list/storage"
)

var menu = map[string]func(*[]models.Item, *int){
	"1": handlers.AddNameList,
	"2": handlers.ShowList,
	"3": handlers.DeleteItem,
	"4": handlers.MarkItemAsChecked,
	"5": handlers.ClearList,
}

func main() {
	vault := storage.JsonDB{
		FileName: "shopping_list.json",
	}
	list, err := vault.Load()
	if err != nil {
		fmt.Println("Ошибка при загрузке:", err)
		list = []models.Item{}
	}
	fmt.Println("🛒 Добро пожаловать в список покупок!")
	idCounter := 1
	for _, item := range list {
		if item.ID >= idCounter {
			idCounter = item.ID + 1
		}
	}
Menu:
	for {
		variant := promptData(
			"1. Добавить товар в список",
			"2. Показать список покупок",
			"3. Удалить товар из списка",
			"4. Отметить купленный товар в списке",
			"5. Очистить список",
			"6. Выйти из программы",
			"Выберите действие (введите номер)",
		)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(&list, &idCounter)
		vault.Save(&list)
	}
}

func promptData(prompt ...any) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var input string
	fmt.Scanln(&input)
	return input
}
