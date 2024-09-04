package models

type Task struct {
	ID          uint   `gorm: "primaryKey;autoIncrement json:"id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Due_date    string `json: "due_date"`
	Created_at  string `json: "created_at"`
	Updated_at  string `json: "updated_at"`
}
