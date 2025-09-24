package main
import (
     "fmt"
	 "errors"
)

type Student struct {
	ID string `json:"id"`
    Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}

func (s *Student) IsHoner() bool{
    return s.GPA >= 3.50
}

func (s *Student) Validate() error{
	if s.Name == ""{
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4{
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4{
		return errors.New("gpa must be between 0-4")
	}
	return nil
}

func main() {
     // var st Student = Student{ID:"1", Name:"thikhamphon", Email:"thikhamphon@hotmail.com", Year:3, GPA:3.13}

	 // st := Student = Student({ID:"1", Name:"thikhamphon", Email:"thikhamphon@hotmail.com", Year:3, GPA:3.13})

     students := []Student{
		{ID:"1", Name:"thikhamphon", Email:"thikhamphon@hotmail.com", Year:3, GPA:3.13},
		{ID:"2", Name:"alice", Email:"alice@hotmail.com", Year:4, GPA:2.75},
	 }

	 newstudent := Student{ID:"3", Name:"trudy", Email:"trudy@hotmail.com", Year:4, GPA:3.50}
	 students = append(students, newstudent)

	 for i, student:= range students {
	 fmt.Printf("%d Honor = %v\n", i, student.IsHoner())
     fmt.Printf("%d Validation = %v\n", i, student.Validate())
	 }

}