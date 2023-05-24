package api

import (
	"github.com/Pavel-Hrdina/todo/backend/config/database"
	"gorm.io/gorm"
)

// defiens the todo model
type Todo struct {
	gorm.Model
	Task        string `gorm:"not null"`
	Completed   bool   `gorm:"default:false"`
	Id          *uint  `gorm:"not null" gorm:"index"`
	Description string `gorm:"not null"`
}

// creates a todo entery
func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(todo)
}

// find a todo by given conditions
// Použít wire místo conditions a where
func FindTodo(dest interface{}, conditions ...interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Take(dest, conditions...)
}

// find all todos
func FindTodos(dest interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest)
}

// find todo by id
// BROKEN
func FindTodoByID(dest interface{}, todoID interface{}) *gorm.DB {
	return FindTodo(dest, "id = ?", todoID)
}

// delets a todo with a given id
func DeleteTodo(todoID interface{}) *gorm.DB {
	return database.DB.Unscoped().Delete(&Todo{}, "id = ?", todoID)
}

// updates a todo with a given id
func UpdateTodo(totoID interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", totoID).Updates(data)
}
