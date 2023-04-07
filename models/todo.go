package models

import (
	"time"
)

type Todo struct {
	Todo_Id           int64     `json:"id" gorm:"column:todo_id;primaryKey"`
	Activity_Group_Id int64     `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title             string    `json:"title" gorm:"column:title"`
	Priority          string    `json:"priority" gorm:"column:priority"`
	Created_At        time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time
}

type Test struct {
	Counter int64
}

type PostTodo_Req struct {
	Activity_Group_Id int64  `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title             string `json:"title" gorm:"column:title"`
	Priority          string `json:"priority" gorm:"column:priority"`
	Is_Active         bool   `json:"is_active" gorm:"column:is_active"`
}
