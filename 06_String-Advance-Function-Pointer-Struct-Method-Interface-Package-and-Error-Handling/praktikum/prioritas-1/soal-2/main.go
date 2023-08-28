package main

import (
	"fmt"
	"math"
)

type Student struct{
	name string
	score float64
}

type StudentList []Student

func (s *StudentList) AddStudent(name string, score float64) {
	student := Student{name, score}
	*s = append(*s, student)
}

func (s StudentList) AverageScore() float64 {
	totalScore := 0.0
	for _, student := range s {
		totalScore += student.score
	}
	return totalScore / float64(len(s))
}

func (s StudentList) MinScore() Student {
	minScore := math.Inf(1)
	var minStudent Student

	for _, student := range s {
		if student.score < minScore {
			minScore = student.score
			minStudent = student
		}
	}
	return minStudent
}

func (s StudentList) MaxScore() Student {
	maxScore := math.Inf(-1)
	var maxStudent Student

	for _, student := range s {
		if student.score > maxScore {
			maxScore = student.score
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	var students StudentList

	for i := 0; i < 5; i++ {
		var name string
		var score float64

		fmt.Printf("Enter name for student %d: ", i+1)
		fmt.Scanln(&name)

		fmt.Printf("Enter score for student %d: ", i+1)
		fmt.Scanln(&score)

		students.AddStudent(name, score)
	}

	average := students.AverageScore()
	minStudent := students.MinScore()
	maxStudent := students.MaxScore()

	fmt.Printf("Average score: %.2f\n", average)
	fmt.Printf("Student with minimum score: %s with score %.2f\n", minStudent.name, minStudent.score)
	fmt.Printf("Student with maximum score: %s with score %.2f\n", maxStudent.name, maxStudent.score)
}