package userdtos

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	Token string `json:"token"`
}
