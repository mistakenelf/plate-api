package models

// DashboardCounts model for total counts
type DashboardCounts struct {
	TodoCount int16 `json:"todoCount"`
}

// Dashboard model for dashboard
type Dashboard struct {
	DashboardCounts DashboardCounts `json:"dashboardCounts"`
	TodoLists       []TodoList      `json:"todoLists"`
}
