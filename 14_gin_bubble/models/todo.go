package models

import (
	"gin_web_demo/14_gin_bubble/dao"
)

/*
	Todo Model
 */
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

/*
	创建待办事项
 */
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

/*
	查看所有待办事项
 */
func GetAllTodo() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

/*
	通过 id 获取待办事项
 */
func GetTodoById(id string) (todo *Todo, err error) {
	todo =  new(Todo)
	err = dao.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return
}

/*
	更新待办事项
 */
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

/*
	删除待办事项
 */
func DeleteATodo(id string) (err error) {
	var todo = new(Todo)
	err = dao.DB.Where("id=?", id).Delete(&todo).Error
	return
}
