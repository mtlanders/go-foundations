// Module 4 Exercise
// Create a new directory 4_maps in your dev environment.

// Scenario
// You're building a simple inventory system for a warehouse.
// Requirements

// Define a Product struct with the following fields:

// Name (string)
// Quantity (int)
// Price (float64)

// Write a function add_product that takes a map[string]*Product, a key (string), and a *Product, and inserts it into the map.
// If the key already exists, return an error. If not, insert and return nil.
//
// Write a function restock that takes the inventory map and a key, and increments the Quantity of the product
// at that key by a caller-supplied amount. If the key doesn't exist, return an error.
//
// Write a function total_value that takes the inventory map and returns the total value of all inventory
// (Quantity * Price summed across all products) as a float64.
//
// Write a function remove_product that takes the inventory map and a key, and deletes it. If the key doesn't exist, return an error.
//
// In main, demonstrate all four functions with at least 2 products. Print results to verify correctness.

// Criteria

// Correct use of comma-ok for key existence checks
// Error strings lowercase, no trailing punctuation
// Pointer receiver or pointer map value used appropriately — no unnecessary copies
// 3-index slice not required here, but if you use any slices, apply it where it makes sense
// Correct behavior on missing keys for all functions that check existence

/* Observations:
 *
 * I forgot that a map is inherently a reference type, so I originally passed a pointer to
 * the map, which is not idiomatic and thus creates noise. I changed it back to a pass-by-value
 */

package main

import (
	"errors"
	"fmt"
)

// Product struct with specified fields - Requirement 1
type Product struct {
	Name     string
	Quantity int
	Price    float64
}

// addProduct func which adds if not present, errors if present - Requirement 2
func addProduct(prodMap map[string]*Product, key string, prod *Product) error {
	// Existence check good - Criteria 1
	if _, ok := prodMap[key]; ok { // Don't need first value, so '_' to discard
		return errors.New("product already exists in inventory")
	}
	prodMap[key] = prod
	return nil
}

// restock func which restocks existing product by supply, errors if not present - Requirement 3
func restock(prodMap map[string]*Product, key string, supply int) error {
	// Existence check good - Criteria 1
	if _, ok := prodMap[key]; !ok {
		return errors.New("no such product in inventory") // Lower-case errors - Criteria 2
	}
	prodMap[key].Quantity += supply
	return nil
}

// totalValue func which returns total of all product quantities * their values - Requirement 4
func totalValue(prodMap map[string]*Product) float64 {
	retval := 0.0
	for _, n := range prodMap {
		retval += float64(n.Quantity) * n.Price
	}
	return retval
}

// removeProduct func which deletes the product at an existing key, errs if not present - Requirement 5
func removeProduct(prodMap map[string]*Product, key string) error {
	if _, ok := prodMap[key]; !ok {
		return errors.New("cannot delete non-existent product")
	}
	delete(prodMap, key)
	return nil
}

func main() {
	pmap := make(map[string]*Product) // Create empty map - avoid panics for nil insertion

	prod1 := Product{"soda", 10, .99}
	prod2 := Product{"chips", 20, 1.99}

	// Call addProduct on 2 products - Requirement 5
	err := addProduct(pmap, prod1.Name, &prod1)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = addProduct(pmap, prod2.Name, &prod2)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Call restock on 2 products - Requirement 5
	err = restock(pmap, prod1.Name, 10)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = restock(pmap, prod2.Name, 5)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Call totalValue - technically Requirement 5
	totalVal := totalValue(pmap)
	fmt.Printf("Total value of inventory: $%f\n", totalVal)

	// Call remove product on 2 products - Requirement 5
	err = removeProduct(pmap, prod1.Name)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = removeProduct(pmap, prod2.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
}
