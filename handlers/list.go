// функции: добавить, удалить, показать

package handlers

import (
	"fmt"
	"shopping-list/models"
)

func AddNameList(list *[]models.Item, idCounter *int) {
	var nameProduct string
	fmt.Println("Введите название продукта: ")
	fmt.Scan(&nameProduct)
	item := models.Item{
		ID:   *idCounter,
		Name: nameProduct,
	}
	*list = append(*list, item)
	*idCounter++
	fmt.Printf("Продукт '%s' добавлен в список покупок.\n", nameProduct)
}

func ShowList(list *[]models.Item, idCounter *int) {
	if len(*list) == 0 {
		fmt.Println("Список покупок пуст.")
		return
	}
	fmt.Println("Список покупок:")
	for _, item := range *list {
		status := "не куплен"
		if item.Checked {
			status = "куплен"
		}
		fmt.Printf("ID: %d, Название: %s, Статус: %s\n", item.ID, item.Name, status)
	}
}

func DeleteItem(list *[]models.Item, idCounter *int) {
	var id int
	fmt.Print("Введите ID товара, который хотите удалить: ")
	fmt.Scan(&id)

	for i, item := range *list {
		if item.ID == id {
			*list = append((*list)[:i], (*list)[i+1:]...)
			fmt.Printf("Продукт с ID %d удален из списка покупок.\n", id)
			return
		}
	}
	fmt.Printf("Продукт с ID %d не найден в списке покупок.\n", id)
}

func MarkItemAsChecked(list *[]models.Item, idCounter *int) {
	var id int
	fmt.Print("Введите ID товара, который хотите отметить как купленный: ")
	fmt.Scan(&id)

	for i, item := range *list {
		if item.ID == id {
			(*list)[i].Checked = true
			fmt.Printf("Продукт с ID %d отмечен как купленный.\n", id)
			return
		}
	}
	fmt.Printf("Продукт с ID %d не найден в списке покупок.\n", id)
}

func ClearList(list *[]models.Item, idCounter *int) {
	fmt.Println("Список очищен. ID сброшен.")
	*list = []models.Item{}
	*idCounter = 1
}
