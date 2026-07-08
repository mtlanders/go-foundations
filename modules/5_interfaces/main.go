/*
=============================================================
SCENARIO
=============================================================
You're building a simple notification system. Different
notification types — email and SMS — need to be sent through
a unified interface.

=============================================================
REQUIREMENTS
=============================================================
1. Declare a Notifier interface with a single method:
   Send(message string) error

2. Implement two concrete types: EmailNotifier and
   SMSNotifier, each with a To field (string)

3. EmailNotifier.Send should return an error if message is
   empty; otherwise return nil

4. SMSNotifier.Send should return an error if message exceeds
   160 characters; otherwise return nil

5. Write a function Notify(n Notifier, message string) error
   that calls Send on the notifier and returns the error

6. In main, create one of each notifier and call Notify on
   both with valid and invalid messages, printing results

=============================================================
ACCEPTANCE CRITERIA
=============================================================
- Notifier is declared as an interface type at package scope
- Both concrete types satisfy Notifier implicitly
- Error strings are lowercase, no trailing punctuation
- Notify accepts a Notifier interface, not a concrete type
- All errors are handled, none discarded with _
- Use value receivers on both Send methods
=============================================================
*/

/*
	Submission Corrections:

	Idiomatic Go has interface names and related functions as upper-case
	- e.g., send -> Send, notify -> Notify

	- emptyEmailSize constant unnecessary

	- Using strings.Repeat() for repeated string, not loop
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

const maxSmsSize int = 160

// Requirement 1
type Notifier interface {
	Send(message string) error
}

// Requirement 2
type SMSNotifier struct {
	To string
}

// Requirement 2
type EmailNotifier struct {
	To string
}

func (s SMSNotifier) Send(message string) error {
	if len(message) > maxSmsSize { // Requirement 4
		return errors.New("message too big to send - exceeds 160 characters")
	}
	fmt.Printf("\nSMSNotifier sending message to: %s\n", s.To)
	fmt.Printf("***message: %s\n\n", message)
	return nil
}

func (e EmailNotifier) Send(message string) error {
	if len(message) == 0 { // Requirement 3
		return errors.New("no message to send - 0-length message")
	}
	fmt.Printf("\nEmailNotifier sending message to: %s\n", e.To)
	fmt.Printf("***message: %s\n\n", message)
	return nil
}

func Notify(n Notifier, message string) error {
	return n.Send(message)
}

func main() {

	sms := SMSNotifier{To: "123-4567-8910"}
	email := EmailNotifier{To: "johnsmith@gmail.com"}

	err := Notify(sms, "Hello, friend. How are you?")
	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
	}

	badSms := strings.Repeat(" ", (maxSmsSize + 5))
	err = Notify(sms, badSms)
	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
	}

	err = Notify(email, "All, this is an email. Regards, Matt")
	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
	}

	err = Notify(email, "")
	if err != nil {
		fmt.Printf("error: %s\n\n", err.Error())
	}
}
