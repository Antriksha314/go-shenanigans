package model

import "fmt"

type DurationType float32 // in hours
type Course struct {
	Id         int
	Name       string
	Duration   DurationType
	Instructor Instructor
}

func (c Course) SignUp() bool {
	return true
}
func (c Course) String() string {
	return fmt.Sprintf("%v %v %v %v", c.Id, c.Name, c.Duration, c.Instructor)
}
