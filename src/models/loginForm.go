package models

type LoginForm struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func (login *LoginForm) Validate() bool {
	if login.Login == "" || login.Password == "" {
		return false
	}
	return true
}
