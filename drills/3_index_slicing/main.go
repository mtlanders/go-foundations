/*
=============================================================
EXERCISE — Targeted Review: 3-Index Slicing & Guard Ordering
=============================================================

COVERS: 3-index slicing, guard ordering, error strings,
        pointer receivers, comma-ok, variadic functions

SCENARIO
You are building a simple quiz result processor. Each quiz
has a name and a recorded list of float64 question scores.

REQUIREMENTS

1. Define a Quiz struct with the following fields:
   - Name    (string)
   - Scores  ([]float64)

2. Write a method trimmed on *Quiz that returns a sub-slice
   of Scores using 3-index slicing, excluding the first and
   last elements. Return an error if Scores has fewer than
   3 elements. The guard check must come before the slice
   operation.

3. Write a method highest on *Quiz that returns the single
   highest score in Scores as a float64 and an error if
   Scores is empty.

4. Write a variadic function summarize that takes a map of
   string to *Quiz and a variadic list of quiz names. For
   each name that exists in the map, print the quiz name,
   its highest score, and its trimmed score slice. Silently
   skip names that do not exist. Return an error if no names
   are passed.

5. In main, create at least 3 Quiz values, store them in a
   map, and call summarize with a mix of valid and invalid
   names.

ACCEPTANCE CRITERIA
- Guard checks precede slice operations in all functions
- 3-index slicing used correctly in trimmed
- comma-ok used for all map key existence checks
- error strings are lowercase with no trailing punctuation
- Map is not passed as a pointer
- Pointer receiver used on trimmed and highest
=============================================================
*/

package main

import (
	"errors"
	"fmt"
)

// Requirement 1
type Quiz struct {
	Name   string
	Scores []float64
}

// Requirement 2
func (quiz *Quiz) trimmed() ([]float64, error) {
	if len(quiz.Scores) < 3 {
		return nil, errors.New("fewer than 3 quizzes provided")
	}
	return quiz.Scores[1:(len(quiz.Scores) - 1):(len(quiz.Scores) - 1)], nil
}

// Requirement 3
func (quiz *Quiz) highest() (float64, error) {
	if len(quiz.Scores) == 0 {
		return -1.0, errors.New("no grade available for this quiz")
	}

	highVal := 0.0
	for _, n := range quiz.Scores {
		if n > highVal {
			highVal = n
		}
	}
	return highVal, nil
}

// Requirement 4
func summarize(quizzes map[string]*Quiz, names ...string) error {
	if len(names) == 0 {
		return errors.New("no names provided to summarize")
	}

	for _, n := range names {
		if _, ok := quizzes[n]; ok {
			highVal, err := quizzes[n].highest()
			if err != nil {
				return err
			}

			trim, err := quizzes[n].trimmed()
			if err != nil {
				return err
			}

			fmt.Printf("Quiz %s: Highest Score: %f\tTrimmed Slice: %f\n", quizzes[n].Name, highVal, trim)
		}
	}
	return nil
}

func main() {

	quizMap := make(map[string]*Quiz)

	quiz1 := Quiz{Name: "math", Scores: []float64{10.0, 21.0, 76.0, 92.0, 63.0}}
	quiz2 := Quiz{Name: "history", Scores: []float64{67.0, 53.0, 98.0, 75.0, 100.0}}
	quiz3 := Quiz{Name: "science", Scores: []float64{73.0, 94.0, 81.0, 85.0, 44.0}}

	quizMap[quiz1.Name] = &quiz1
	quizMap[quiz2.Name] = &quiz2
	quizMap[quiz3.Name] = &quiz3

	err := summarize(quizMap, "math", "literature", "driver's ed", "science", "english", "history", "health")
	if err != nil {
		fmt.Println(err.Error())
	}
}
