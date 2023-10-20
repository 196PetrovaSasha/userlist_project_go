package models

type UserStore struct {
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}
