package models

type User struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Likes []int64 `json:"-"`
}
type Relationship struct {
	Id    int64  `json:"id"`
	State string `json:"state"`
	Type  string `json:"type"`
}
