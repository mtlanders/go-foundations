package main

import (
	"fmt"
	"strconv"
)

//*******************************************************************

// Entry interface - Req 1
type Entry interface {
	Describe() string
}

// Settler interfafce - Req 2
type Settler interface {
	Settle()
}

// SettleableEntry interface - Req 3
type SettleableEntry interface {
	Entry
	Settler
}

//*******************************************************************

// Concrete types - Req 4
type Invoice struct {
	Amount float64
	Paid   bool
}

type Note struct {
	Text string
}

//*******************************************************************

// Receiver functions - Req 5
// Functions mutate the structure internals, so pointer receiver
// is required here.
func (i *Invoice) Settle() {
	i.Amount = 0.0 // Actually resolving the real-world issue by paying the amount (spec drift)
	i.Paid = true
}

// No mutation occurs here, and copy considerations are trivial, so
// using value receiver
func (i Invoice) Describe() string {
	return string("Invoice: $" + strconv.FormatFloat(i.Amount, 'f', 2, 64) +
		", paid: " + strconv.FormatBool(i.Paid))
}

// Using value receiver to satisfy Entry - no mutation, so no pointer receiver
func (n Note) Describe() string {
	return n.Text
}

//*******************************************************************

// Modified Req 6 - making two functions to exercise SettleableEntry
// and Entry interfaces
func SettleEntries(entries []SettleableEntry) {
	for _, e := range entries {
		e.Settle()
	}
}

func DescribeEntries(entries []Entry) {
	for _, e := range entries {
		fmt.Println(e.Describe())
	}
}

//*******************************************************************

func main() {
	inv1 := Invoice{Amount: 150.46, Paid: false}
	inv2 := Invoice{Amount: 75.23, Paid: false}
	inv3 := Invoice{Amount: 300.92, Paid: false}
	not1 := Note{Text: "Cleaning service invoice"}
	not2 := Note{Text: "Grocery delivery invoice"}
	not3 := Note{Text: "Utilities invoice"}

	// Only passing references to invoices because of receiver mutation
	allEntries := []Entry{not1, &inv1, not2, &inv2, not3, &inv3}
	fmt.Println("\n*** Original Entries:\n---------------------")
	DescribeEntries(allEntries)

	settleableEntries := make([]SettleableEntry, 0)
	for _, a := range allEntries {
		switch v := a.(type) {
		case *Invoice:
			settleableEntries = append(settleableEntries, v)
		default:
			// Do nothing
		}
	}

	SettleEntries(settleableEntries)

	fmt.Println("\n*** Settled Entries:\n--------------------")
	DescribeEntries(allEntries)
}

//*******************************************************************
