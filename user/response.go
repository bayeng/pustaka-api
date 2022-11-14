package user

import "time"

type UserResponse struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
}
