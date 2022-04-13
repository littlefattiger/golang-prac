package models

import (
	"les25/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {

	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	} else {
		return
	}
}

func GetAllList() (todoList []*Todo, err error) {

	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	} else {
		return todoList, nil
	}
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}