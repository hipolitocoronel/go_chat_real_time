package userdtos

import "go_real_time_chat/internal/auth/domain"

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

//fron entity

func (r *UserResponse) FromEntity(user domain.User) {
	r.ID = user.ID
	r.Name = user.Name
	r.Username = user.Username
	r.PhoneNumber = user.PhoneNumber
	r.Email = user.Email
	r.Password = user.Password
}
