package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// GetDashboardTodoLists returns 5 todo lists
func GetDashboardTodoLists(id string) []models.TodoList {
	var todoLists []models.TodoList

	res := gorm.DBCon().Limit(5).Where("created_by = ?", id).Find(&todoLists, &models.TodoList{})
	if res.RecordNotFound() {
		panic(nil)
	}

	return todoLists
}

// GetTotals returns row totals
func GetTotals(id string) models.DashboardCounts {
	var todoLists []models.TodoList
	var todoListCount int16

	gorm.DBCon().Where("created_by = ?", id).Find(&todoLists, &models.TodoList{}).Count(&todoListCount)

	dashboardCounts := models.DashboardCounts{
		TodoCount: todoListCount,
	}

	return dashboardCounts
}
