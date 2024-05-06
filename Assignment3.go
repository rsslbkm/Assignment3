package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID           uint
	Name         string
	Age          int
	DepartmentID uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Enrollments  []Enrollment   `gorm:"many2many:student_enrollments;"`
}

type Course struct {
	ID           uint
	Name         string
	InstructorID uint
	Students     []Student `gorm:"many2many:student_courses;"` // Define the relationship with foreign key tag
}

type Department struct {
	ID   uint
	Name string
}

type Instructor struct {
	ID      uint
	Name    string
	Courses []Course
}

type Enrollment struct {
	ID             uint
	StudentID      uint
	CourseID       uint
	EnrollmentDate time.Time
}

func main() {
	connStr := "user=postgres dbname=ass2 password=1234 host=localhost sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db", err)
	}
	db.AutoMigrate(&Student{}, &Course{}, &Department{}, &Instructor{}, &Enrollment{})

	var students []Student
	var departmentID int = 1
	db.Where("department_id = ?", departmentID).Find(&students)
	fmt.Println("Students in department", departmentID)
	for _, student := range students {
		fmt.Println(student)
	}
	var courses []Course
	var instructorName string
	db.Joins("JOIN instructors ON courses.instructor_id = instructors.id").
		Where("instructors.name = ?", instructorName).
		Find(&courses)
	for _, course := range courses {
		fmt.Println(course)
	}
	var enrollments []Enrollment
	var studentID uint
	db.Where("student_id = ?", studentID).Find(&enrollments)
	for _, enrollment := range enrollments {
		fmt.Println(enrollment)
	}
	db.Transaction(func(tx *gorm.DB) error {
		student := Student{Name: "John", Age: 20, DepartmentID: 1}
		course := Course{Name: "Math", InstructorID: 1}

		if err := tx.Create(&student).Error; err != nil {
			return err
		}

		if err := tx.Create(&course).Error; err != nil {
			return err
		}

		enrollment := Enrollment{StudentID: student.ID, CourseID: course.ID, EnrollmentDate: time.Now()}
		if err := tx.Create(&enrollment).Error; err != nil {
			return err
		}

		return nil
	})
}

func AddStudent(db *gorm.DB, student *Student) {
	currentTime := time.Now()
	student.CreatedAt = currentTime
	student.UpdatedAt = currentTime
	result := db.Create(student)
	if result.Error != nil {
		fmt.Println("failed to create student: ", result.Error.Error())
	}
	fmt.Println("Student created successfully")
}
func AddCourse(db *gorm.DB, course *Course) {
	result := db.Create(course)
	if result.Error != nil {
		fmt.Println("failed to create course: ", result.Error.Error())
	}
	fmt.Println("Course created successfully")
}
func AddInstructor(db *gorm.DB, instructor *Instructor) {
	result := db.Create(instructor)
	if result.Error != nil {
		fmt.Println("failed to create instructor: ", result.Error.Error())
	}
	fmt.Println("instucrtor created successfully")
}
func AddDepartment(db *gorm.DB, department *Department) {
	result := db.Create(department)
	if result.Error != nil {
		fmt.Println("failed to create department: ", result.Error.Error())
	}
	fmt.Println("Department created successfully")
}

func UpdateStudent(db *gorm.DB, student *Student) {
	student.UpdatedAt = time.Now()
	result := db.Save(student)
	if result.Error != nil {
		fmt.Println("failed to update instructor: ", result.Error.Error())
	}
	fmt.Println("Student updated successfully")
}

func UpdateInstructor(db *gorm.DB, instructor *Instructor) {
	result := db.Save(instructor)
	if result.Error != nil {
		fmt.Println("failed to update instructor: ", result.Error.Error())
	}
	fmt.Println("Instructor updated successfully")
}

func UpdateCourse(db *gorm.DB, course *Course) {
	result := db.Save(course)
	if result.Error != nil {
		fmt.Println("failed to update course: ", result.Error.Error())
	}
	fmt.Println("Course updated successfully")
}

func UpdateDepartment(db *gorm.DB, department *Department) {
	result := db.Save(department)
	if result.Error != nil {
		fmt.Println("failed to update department: ", result.Error.Error())
	}
	fmt.Println("Department updated successfully")
}

func RetrieveCourse(db *gorm.DB, courseID uint) Course {
	var course Course
	result := db.First(&course, courseID)
	if result.Error != nil {
		fmt.Println("failed to retrieve course: ", result.Error.Error())
	}
	return course
}

func RetrieveStudent(db *gorm.DB, studentID uint) Student {
	var student Student
	result := db.First(&student, studentID)
	if result.Error != nil {
		fmt.Println("failed to retrieve student: ", result.Error.Error())
	}
	return student
}

func RetrieveDepartment(db *gorm.DB, departmentID uint) Department {
	var department Department
	result := db.First(&department, departmentID)
	if result.Error != nil {
		fmt.Println("failed to retrieve deparment: ", result.Error.Error())
	}
	return department
}

func RetrieveInstructor(db *gorm.DB, instructorID uint) Instructor {
	var instructor Instructor
	result := db.First(&instructor, instructorID)
	if result.Error != nil {
		fmt.Println("failed to retrieve instructor: ", result.Error.Error())
	}
	return instructor
}

func DeleteDepartment(db *gorm.DB, department *Department) {
	result := db.Delete(department)
	if result.Error != nil {
		fmt.Println("failed to delete department: ", result.Error.Error())
	}
	fmt.Println("Department deleted successfully")
}

func DeleteStudent(db *gorm.DB, student *Student) {
	result := db.Delete(student)
	if result.Error != nil {
		fmt.Println("failed to delete student: ", result.Error.Error())
	}
	fmt.Println("Student deleted successfully")
}
func DeleteCourse(db *gorm.DB, course *Course) {
	result := db.Delete(course)
	if result.Error != nil {
		fmt.Println("failed to delete course: ", result.Error.Error())
	}
	fmt.Println("Course deleted successfully")
}
func DeleteInstructor(db *gorm.DB, instructor *Instructor) {
	result := db.Delete(instructor)
	if result.Error != nil {
		fmt.Println("failed to delete instructor: ", result.Error.Error())
	}
	fmt.Println("Instructor deleted successfully")
}
