/*
=============================================================
CAPSTONE EXERCISE — Go Days 1–4 Review
=============================================================

COVERS: Modules 1-4

SCENARIO
You are building a simple student grade tracker. Students are
stored in a map by student name. Each student has a name and a
recorded list of exam scores.

REQUIREMENTS

1. Define a Student struct with the following fields:
   - Name     (string)
   - Scores   ([]float64)

2. Write a function addStudent that takes the student map, a
   string name, and a *Student, and inserts it. If the name already
   exists, return an error.

3. Write a function recordScore that takes the student map, a
   string name, and a float64 score, and appends the score to that
   student's Scores slice. If the name does not exist, return an
   error.

4. Write a method average on *Student that returns the mean of
   all scores as a float64 and an error if the Scores slice is
   empty.

5. Write a variadic function topStudents that takes the student
   map and a variadic list of string names, and returns a slice of
   *Student containing only the students whose average is at or
   above 80.0. Students whose name does not exist in the map are
   silently skipped. If no names are passed, return an error.

6. Write a function scoreRange that takes a *Student and returns
   a sub-slice of Scores using 3-index slicing from index 1 to
   the second-to-last element (exclusive of the first and last
   score). Return an error if the slice has fewer than 3 elements.

7. In main, demonstrate all of the above with at least 3 students
   and varied score sets. Print results to verify correctness.

ACCEPTANCE CRITERIA
- Map is not passed as a pointer (it is already a reference type)
- Scores slice is not passed as a pointer where avoidable
- error strings are lowercase with no trailing punctuation
- 3-index slicing is used in scoreRange
- comma-ok is used for all map key existence checks
- average and topStudents handle empty/nil edge cases correctly
=============================================================
*/

package main

import (
	"errors"
	"fmt"
)

const topAverage float64 = 80.0

// Requirement 1
type Student struct {
	Name   string
	Scores []float64
}

// Requirement 2
func addStudent(students map[string]*Student, name string, student *Student) error {
	if _, ok := students[name]; ok {
		return errors.New("student is already registered")
	}
	students[name] = student
	return nil
}

// Requirement 3
func recordScore(students map[string]*Student, name string, score float64) error {
	if _, ok := students[name]; !ok {
		return errors.New("no such student in registry")
	}

	students[name].Scores = append(students[name].Scores, score)
	return nil
}

// Requirement 4
func (student *Student) average() (float64, error) {
	if len(student.Scores) == 0 {
		return -1.0, errors.New("student has no scores to average")
	}

	total := 0.0
	for _, n := range student.Scores {
		total += n
	}
	return (total / float64(len(student.Scores))), nil
}

// Requirement 5
func topStudents(students map[string]*Student, names ...string) ([]*Student, error) {
	if len(names) == 0 {
		return nil, errors.New("no names of students to check")
	}

	aboveAvg := make([]*Student, 0)
	for _, n := range names {
		if _, ok := students[n]; ok {
			avg, _ := students[n].average()
			if avg >= topAverage {
				aboveAvg = append(aboveAvg, students[n])
			}
		}
	}
	return aboveAvg, nil
}

// Requirement 6
func scoreRange(student *Student) ([]float64, error) {
	if len(student.Scores) < 3 {
		return nil, errors.New("score range slice contains fewer than 3 elements")
	}

	scoreRng := student.Scores[1:(len(student.Scores) - 1):(len(student.Scores) - 1)]
	return scoreRng, nil
}

func main() {

	// 3 Students required
	stud1 := Student{Name: "Joe", Scores: []float64{75.0, 72.0, 67.0, 81.0, 55.0}}
	stud2 := Student{Name: "Haley", Scores: []float64{82.0, 86.0, 74.0, 93.0, 97.0}}
	stud3 := Student{Name: "Alex", Scores: []float64{80.0, 85.0, 84.0, 88.0, 89.0}}

	registry := make(map[string]*Student, 0)

	// Requirement 7
	addStudent(registry, stud1.Name, &stud1)
	addStudent(registry, stud2.Name, &stud2)
	addStudent(registry, stud3.Name, &stud3)
	fmt.Printf("Number of students: %d\n", len(registry))

	// Requirement 7
	recordScore(registry, stud1.Name, 42.0)
	recordScore(registry, stud2.Name, 91.0)
	recordScore(registry, stud3.Name, 87.0)

	// Requirement 7
	studAvg, _ := stud1.average()
	fmt.Printf("%s grade avg: %f\n", stud1.Name, studAvg)
	studAvg, _ = stud2.average()
	fmt.Printf("%s grade avg: %f\n", stud2.Name, studAvg)
	studAvg, _ = stud3.average()
	fmt.Printf("%s grade avg: %f\n", stud3.Name, studAvg)

	// Requirement 7 - Does not make sense to call on "at least 3 students"
	topStuds, _ := topStudents(registry, stud1.Name, stud3.Name, stud2.Name, "Michael", "Jennifer")
	fmt.Printf("Number of students at or exceeding 80.0 average: %d\n", len(topStuds))
	gradeRng, _ := scoreRange(&stud1)

	// Requirement 7
	fmt.Printf("%s grade range (excl. 0 and N-1): %f\n", stud1.Name, gradeRng)
	gradeRng, _ = scoreRange(&stud2)
	fmt.Printf("%s grade range (excl. 0 and N-1): %f\n", stud2.Name, gradeRng)
	gradeRng, _ = scoreRange(&stud3)
	fmt.Printf("%s grade range (excl. 0 and N-1): %f\n", stud3.Name, gradeRng)
}
