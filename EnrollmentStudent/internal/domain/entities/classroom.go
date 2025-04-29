package entities

type ClassRoom struct {
	ClassId string
	Students []Student
}

func NewClassRoom(classroom_id string, students []Student) ClassRoom{
	return ClassRoom{ClassId: classroom_id, Students: students}
}