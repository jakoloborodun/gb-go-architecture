package models

type Item struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}
