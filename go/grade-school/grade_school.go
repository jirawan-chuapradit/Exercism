package school

import (
	"log"
	"sort"
)


type Grade struct {
	number int
	Names  []string
}

type School struct {
	Grades []Grade
}

func New() *School {
	s := School{}
	return &s
}

// Add a student's name to the roster for a grade
func (s *School) Add(name string, grade int) {
	// Exits grade
	for i, v := range s.Grades {
		if v.number == grade {
			s.Grades[i].Names = append(s.Grades[i].Names , name)
			return
		}
	}


	// new grade
	newGrade := Grade{
		number: grade,
		Names:  []string{name},
	}
	s.Grades = append(s.Grades, newGrade)

}

// Get a list of all students enrolled in a grade
func (s *School) Grade(grade int) []string {
	for _, v := range s.Grades {
		if v.number == grade {
			return v.Names
		}
	}
	return []string{}
}

// Get a sorted list of all students in all grades.
func (s *School) Enrollment() []Grade {
	log.Println("grade: ", s.Grades)
	 tempGrade := make(map[int]interface{})
	var grades []int
	for _, v := range s.Grades {
		grades = append(grades, v.number)
		sort.Strings(v.Names)
		tempGrade[v.number] = v.Names


	}




	sort.Ints(grades)

	var ans []Grade
	for _, v := range grades {
		g := Grade{
			number: v,
			Names: tempGrade[v].([]string),
		}
		ans = append(ans, g)
	}

	return ans
}

