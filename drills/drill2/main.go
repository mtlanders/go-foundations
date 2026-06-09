/*
=============================================================

COVERS:

- Guard ordering — check preconditions before operating on data
- 3-index slicing mechanics — both the syntax and the defensive max == high pattern
- Positional struct literals — named fields only
- Variable shadowing — naming a local the same as an enclosing function
- Unnecessary error re-wrapping — returning err directly vs. errors.New(err.Error())
- Spec fidelity under complexity — multiple requirements in play simultaneously

SCENARIO:
  A game server tracks player sessions. Each session has a
  player name, a score history (slice of int), and a status
  string ("active" or "inactive").

REQUIREMENTS:
  1. Define a Session struct with fields: Name (string),
     Scores ([]int), Status (string).
     Instantiate all Session values using named fields only.

  2. Write a function `recentScores(s Session, n int) ([]int, error)`
     that returns the last n scores from s.Scores as a 3-index
     sub-slice with capacity capped at n.
     - Return an error if len(s.Scores) < n. The guard must
       execute before any slice operation.
     - Error strings: lowercase, no trailing punctuation.

  3. Write a function `activeSessions(sessions []Session) []Session`
     that returns a new slice containing only sessions with
     Status == "active". Do not shadow the function name with
     a local variable.

  4. Write a function `recentForActive(sessions []Session, n int) ([][]int, error)`
     that calls recentScores on each active session (use activeSessions
     internally) and collects the results. Return the first error
     encountered directly — do not re-wrap with errors.New.

  5. In main:
     - Declare a map[string]Session keyed by player name containing
       exactly 4 players: 2 active, 2 inactive. At least one active
       player must have fewer scores than the value of n you pass to
       recentForActive, so the error path is exercised.
     - Build a []Session from the map values and pass it to
       recentForActive.
     - Use comma-ok to look up one player by name before the
       recentForActive call and print their status.
     - Print the results of recentForActive, or the error if one
       is returned.
     - Handle all errors; do not discard with _.

ACCEPTANCE CRITERIA:
  - Guard in recentScores fires before the slice operation.
  - 3-index sub-slice syntax is correct; capacity is capped at n.
  - All struct literals use named fields.
  - No local variable in activeSessions shadows the function name.
  - Errors from recentScores are returned directly, not re-wrapped.
  - Comma-ok used on map lookup in main.
  - No errors discarded with _.
  - The error path in recentForActive is reachable given the data
    in main.
=============================================================
*/

/*
	Developer Notes:

	This exercise was poorly structured and was a best attempt at making it workable.
	Certain acceptance criteria directly contradicts certain requirements making them
	untestable.
*/

package main

import (
	"errors"
	"fmt"
)

// Requirement 1
type Session struct {
	Name   string
	Scores []int
	Status string
}

// Requirement 2
func recentScores(s Session, n int) ([]int, error) {
	if len(s.Scores) < n {
		errStr := "too few scores for player " + s.Name
		return nil, errors.New(errStr)
	}
	return s.Scores[(len(s.Scores) - n):(len(s.Scores)):(len(s.Scores))], nil
}

// Requirement 3
func activeSessions(sessions []Session) []Session {
	active := make([]Session, 0)
	for _, n := range sessions {
		if n.Status == "active" {
			active = append(active, n)
		}
	}
	return active
}

// Requirement 4
func recentForActive(sessions []Session, n int) ([][]int, error) {
	if len(sessions) == 0 {
		return nil, errors.New("no player sessions provided")
	}

	active := activeSessions(sessions)
	if len(active) == 0 {
		return nil, errors.New("no active sessions found")
	}

	sessionScores := make([][]int, 0)
	for _, x := range active {
		recScore, err := recentScores(x, n)
		if err != nil {
			return nil, err
		}
		sessionScores = append(sessionScores, recScore)
	}

	return sessionScores, nil
}

func main() {

	sessMap := make(map[string]Session)
	sessMap["Joe"] = Session{Name: "Joe", Scores: []int{77, 88, 99, 55, 44, 2}, Status: "active"}
	sessMap["John"] = Session{Name: "John", Scores: []int{11, 22, 33, 66}, Status: "inactive"}
	sessMap["Jack"] = Session{Name: "Jack", Scores: []int{12, 34, 56, 78, 90}, Status: "active"}
	sessMap["Jason"] = Session{Name: "Jason", Scores: []int{21, 43, 65, 78, 91, 36, 59, 24}, Status: "inactive"}

	if sess, ok := sessMap["Joe"]; ok {
		// The map access for name is a bit unnecessary but I wanted to do it anyway
		fmt.Printf("Player %s's session status: %s\n", sess.Name, sess.Status)
	}

	sessions := make([]Session, 0)
	for _, n := range sessMap {
		sessions = append(sessions, n)
	}

	var (
		err       error
		recScores [][]int
	)

	recScores, err = recentForActive(sessions, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i, o := range recScores {
			fmt.Printf("Active player %d scores: %d\n", i, o)
		}
	}
}
