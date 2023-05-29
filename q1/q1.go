package main

import (
	"fmt"
)

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	result := make(map[string]Student)

	// Primeiro, copiamos os dados do studentData1 para o resultado
	for name, student := range studentData1 {
		result[name] = student
	}

	// Em seguida, atualizamos ou adicionamos os dados do studentData2
	for name, student := range studentData2 {
		if existingStudent, ok := result[name]; ok {
			// O aluno já existe no resultado, precisamos atualizar as matérias e notas
			for subject, grade := range student.Subjects {
				existingStudent.Subjects[subject] = grade
			}
			result[name] = existingStudent
		} else {
			// O aluno não existe no resultado, adicionamos todas as informações
			result[name] = student
		}
	}

	return result
}

func main() {
	// Exemplo de uso
	studentData1 := map[string]Student{
		"John": {
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Math":    8.5,
				"Science": 7.8,
			},
		},
		"Alice": {
			Name: "Alice",
			Age:  22,
			Subjects: map[string]float64{
				"History": 9.2,
			},
		},
	}

	studentData2 := map[string]Student{
		"John": {
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Science": 8.0,
				"English": 7.5,
			},
		},
		"Bob": {
			Name: "Bob",
			Age:  21,
			Subjects: map[string]float64{
				"Math":    6.7,
				"Physics": 8.9,
			},
		},
	}

	result := MergeStudentData(studentData1, studentData2)
	fmt.Println(result)
}
