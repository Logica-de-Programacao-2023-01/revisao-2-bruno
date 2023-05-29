package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func mergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	mergedData := make(map[string]Student)

	for key, value := range studentData1 {
		mergedData[key] = value
	}

	for key, value := range studentData2 {
		if _, exists := mergedData[key]; exists {
			for subject, grade := range value.Subjects {
				mergedData[key].Subjects[subject] = grade
			}
		} else {
			mergedData[key] = value
		}
	}

	return mergedData
}

func printStudentData(data map[string]Student) {
	for _, student := range data {
		subjects := make([]string, 0, len(student.Subjects))
		for subject := range student.Subjects {
			subjects = append(subjects, subject)
		}

		fmt.Printf("Name: %s\n", student.Name)
		fmt.Printf("Age: %d\n", student.Age)
		fmt.Printf("Subjects: %s\n", strings.Join(subjects, ", "))
		fmt.Println("-----------------------")
	}
}

func main() {
	studentData1 := map[string]Student{
		"John": Student{
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Math":    8.5,
				"Science": 7.8,
			},
		},
		"Alice": Student{
			Name: "Alice",
			Age:  22,
			Subjects: map[string]float64{
				"History": 9.2,
			},
		},
	}

	studentData2 := map[string]Student{
		"John": Student{
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Science": 8.0,
				"English": 7.5,
			},
		},
		"Bob": Student{
			Name: "Bob",
			Age:  21,
			Subjects: map[string]float64{
				"Math":    6.7,
				"Physics": 8.9,
			},
		},
	}

	result := mergeStudentData(studentData1, studentData2)
	printStudentData(result)
}
