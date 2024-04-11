package model

func NewInstructor(id int, firstName string, lastName string) Instructor {
	return Instructor{Id: id, FirstName: firstName, LastName: lastName}
}
