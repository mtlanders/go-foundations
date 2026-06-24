package main

import "fmt"

//*******************************************************************

// Requirement 3
type Validator interface {
	Valid() bool
}

//*******************************************************************

// Requirement 4
type Config struct {
	ValidConfig bool
}

//*******************************************************************

// Requirement 4
func (c *Config) Valid() bool {
	return c.ValidConfig
}

//*******************************************************************

func GetValidator(present bool) Validator {
	var c *Config = nil
	if present {
		c = &Config{ValidConfig: true}
	}
	return c // Returning typed nil
}

//*******************************************************************

// Classify function - Req 1
func Classify(v any) string {
	if _, ok := v.(string); ok { // Req 2
		return "string"
	} else if _, ok := v.(int); ok { // Req 2
		return "int"
	}
	return "unknown" // Req 2
}

//*******************************************************************

func main() {

	str := Classify("raw str")
	fmt.Println(str)
	str = Classify(42)
	fmt.Println(str)
	str = Classify(3.14159)
	fmt.Println(str)

	config := GetValidator(false)

	/*
		Dev Note: Expect not nil. I'm returning an explicitly typed
		pointer whose value may be nil, but that is not how Go will
		see it. My interpretation of this is that the successful
		result is to see the typed nil trap in action, not avoid it.
	*/
	if config == nil {
		fmt.Println("config is nil (typed)") // Won't print
	}

	// Adding real nil just to emphasize difference between typed and un-typed nil
	config = nil
	if config == nil {
		fmt.Println("config is nil (un-typed)") // Will print
	}

	// Repeating with true just to exercise that path
	config = GetValidator(true)
	if config != nil {
		fmt.Println("config is not nil")
	}
}

//*******************************************************************
