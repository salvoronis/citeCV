package models

type Student struct {
	Username	string
	Name		string
	Password	string
	ID		int
	Class		string
}

type Class struct {
	Id		int
	Name		string
}
