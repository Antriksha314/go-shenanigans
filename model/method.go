package model

import "fmt"

// Methods of Instructor struct
func (inst Instructor) InstructorDetails() string {
	return fmt.Sprintf("%v %v", inst.FirstName, inst.LastName)
}
