// сохранение/загрузка JSON

package storage

import (
	"encoding/json"
	"os"
	"shopping-list/models"
)

type JsonDB struct {
	FileName string
}

func (db *JsonDB) Load() ([]models.Item, error) {
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return nil, err
	}
	var list []models.Item
	err = json.Unmarshal(data, &list)
	return list, err
}

func (db *JsonDB) Save(data *[]models.Item) error {
	file, err := os.Create(db.FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = file.Write(dataBytes)
	if err != nil {
		return err
	}
	return nil
}
