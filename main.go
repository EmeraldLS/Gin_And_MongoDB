package main

import (
	"fmt"

	router "github.com/EmeraldLS/Gin_And_MongoDB/routers"
)

func main() {
	var message string = `


		Development Server started at port :8080
		Below are the endpoints and their methods
		
		/api/students - GetAllStudent, DeleteAllStudents, AddNewStudent
		/api/student/Id - GetSingleStudent, DeleteOneStudent

		Below is an example of the JSON format to send

		{
			matric_no : int,
			first_name: string,
			last_name: string,
			faculty: string
			Department: string,
			age: int,
			cgpa: float64,
		}

		Happy Coding!!!

		localhost:8080
	`
	fmt.Print("\n\n\n")
	fmt.Println(message)
	fmt.Print("\n\n\n")
	router.Router(":8080")

}
