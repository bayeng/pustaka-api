package user

import "time"

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	CreateAt time.Time
}
