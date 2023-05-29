package q1

import (
	"fmt"
)

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func mergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	result := make(map[string]Student)

	// Adiciona os alunos do studentData1 ao resultado
	for name, student := range studentData1 {
		result[name] = student
	}

	// Atualiza as notas dos alunos com base no studentData2
	for name, student := range studentData2 {
		if _, ok := result[name]; ok {
			// O aluno já existe no resultado, atualiza as matérias e notas
			for subject, grade := range student.Subjects {
				result[name].Subjects[subject] = grade
			}
		} else {
			// O aluno não existe no resultado, adiciona-o
			result[name] = student
		}
	}

	return result
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
	fmt.Println(result)
}
