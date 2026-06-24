package main

import (
	"fmt"
)

//*********************************************************

// Requirement 1 - Stage with Run([]int) []int
type Stage interface {
	Run(data []int) []int
}

//*********************************************************

// Requirement 2 - Composed interface
type NamedStage interface {
	Stage
	Name() string
}

//*********************************************************

// Requirement 3 - Concrete Types
type DoubleProcessor struct {
	StageName string
}

type NegativeProcessor struct {
	StageName string
}

type ThresholdPocessor struct {
	StageName string
	Threshold int
	Exec      func(data []int, thresh int) []int // Requirement 5 - Closure function
}

// Requirement 8 - Non-NamedStage type for branching
type SubtractProcessor struct {
}

//*********************************************************

// Requirement 3 - Receivers/interface satisfaction

// Using value receivers since the incoming slice is not a field of
// the slice and thus internal mutation is not a factor
func (d DoubleProcessor) Run(data []int) []int {
	if len(data) == 0 {
		return nil
	}

	newData := make([]int, 0)
	for _, d := range data {
		newData = append(newData, (d * 2))
	}
	return newData
}

// Using value receivers since the incoming slice is not a field of
// the slice and thus internal mutation is not a factor
func (n NegativeProcessor) Run(data []int) []int {
	if len(data) == 0 {
		return nil
	}

	newData := make([]int, 0)

	// Remove all negative values
	for _, d := range data {
		if d >= 0 {
			newData = append(newData, d)
		}
	}
	return newData
}

// Requirement 8 - Type which satisfies Stage but not NamedStage
func (s SubtractProcessor) Run(data []int) []int {
	return nil
}

// Using value receivers since the incoming slice is not a field of
// the slice and thus internal mutation is not a factor
func (t ThresholdPocessor) Run(data []int) []int {
	return t.Exec(data, t.Threshold)
}

// Using value receivers since the name copy is trivial and StageName is not mutated
func (d DoubleProcessor) Name() string {
	return d.StageName
}

// Using value receivers since the name copy is trivial StageName and is not mutated
func (n NegativeProcessor) Name() string {
	return n.StageName
}

// Using value receivers since the name copy is trivial StageName and is not mutated
func (t ThresholdPocessor) Name() string {
	return t.StageName
}

//*********************************************************

// Requirement 4 - RunPipeline()
func RunPipeline(stages []NamedStage, data []int) []int {
	// Disregarding overhead in exchange for added safety at each stage
	if len(stages) == 0 || len(data) == 0 {
		return nil
	}

	finalSlice := data
	for _, s := range stages {
		fmt.Printf("*** Running pipeline stage: %s - %d\n", s.Name(), finalSlice)
		finalSlice = s.Run(finalSlice)
	}
	return finalSlice
}

//*********************************************************

func isNamedStage(s Stage) bool {
	if _, ok := s.(NamedStage); ok {
		return true
	}
	return false
}

//*********************************************************

func main() {
	dp := DoubleProcessor{StageName: "double_values"}
	np := NegativeProcessor{StageName: "filter_negatives"}
	tp := ThresholdPocessor{StageName: "filter_lower", Threshold: 20,
		Exec: func(data []int, thresh int) []int {
			if len(data) == 0 {
				return nil
			}
			newData := make([]int, 0)
			for _, d := range data {
				if d > thresh {
					newData = append(newData, d)
				}
			}
			return newData
		}}

	vals := []int{-1, 5, 12, -4, 6, 34, -37, 18, 9}
	stages := []NamedStage{dp, np, tp}

	// Expected final slice: {24, 68, 36}
	vals = RunPipeline(stages, vals)

	fmt.Printf("*** Final Slice: %d\n", vals)

	// Requirement 8 - Type switching/branching
	sp := SubtractProcessor{}
	if isNamedStage(sp) {
		fmt.Println("\nsp implements NamedStage") // Won't execute because sp is not NamedStage
	} else {
		fmt.Println("\nsp does not implement NamedStage")
	}
}
