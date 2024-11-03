package models

type Blog struct {
	ID      uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title   string
	Content string
	Author  string
}
type ListResponse struct {
	Total int
	List  []Blog
}
