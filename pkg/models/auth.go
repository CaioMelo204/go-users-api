package models

type Auth struct {
	Id    uint64 `json:"id"`
	Token string `json:"token"`
}
