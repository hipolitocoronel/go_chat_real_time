package domain

type UserGoogle struct {
	Name     string `json:"name"`
	Sub      string `json:"sub"`
	Email    string `json:"email"`
	Verified bool   `json:"email_verified"`
	Picture  string `json:"picture"`
}
