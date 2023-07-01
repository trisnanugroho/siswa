package models

type Students struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Class int    `json:"class"`
	Sex   string `json:"sex"`
	City  string `json:"city"`
}
