package models

type Student struct {
	Username	string
	Name		string
	//Mail		string
	Password	string
	ID		string
	Class		string
}

type Class struct {
	Id		int
	Name		string
}
