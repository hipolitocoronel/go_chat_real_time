package userdtos

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Email string `json:"email"`
}
