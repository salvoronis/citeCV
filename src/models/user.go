package models

type User struct {
	Id uint `json:-`
	Login string `json:"login"`
	Password string `json:"password"`
	Class uint `json:"class"`
	FirstName string `json:"firstName"`
	SecondName string `json:"secondName"`
	TermOfUse bool `json:"termOfUse"`
	Email string `json:"email"`
}

func (user *User)Validate() bool{
	if	user.Login == "" ||
		user.Password == "" ||
		!user.TermOfUse ||
		user.Email == "" {
		return false
	}
	return true
}
