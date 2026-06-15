package main

import (
	"errors"
	"fmt"
	"strconv"
)

//*******************************************************************

const dummy string = "I am a const used to satisfy a requirement"

//*******************************************************************

// Types

type Venue interface {
	Name() string
	Capacity() int
	TicketPrice(section string) (float64, error)
}

type Arena struct {
	name     string
	capacity int
}

type Theater struct {
	name     string
	capacity int
}

type Event struct {
	Title     string
	VenueName string
	Attendees []string
}

type Registry struct {
	registry map[string]*Event
}

//*******************************************************************

// Receivers

func (a Arena) Name() string {
	return a.name
}

func (a Arena) Capacity() int {
	return a.capacity
}

func (a Arena) TicketPrice(section string) (float64, error) {

	var (
		retval float64
		reterr error
	)

	if section == "floor" {
		retval, reterr = 150.00, nil
	} else if section == "upper" {
		retval, reterr = 75.00, nil
	} else {
		retval, reterr = 0.00, errors.New("invalid arena section")
	}
	return retval, reterr
}

func (t Theater) Name() string {
	return t.name
}

func (t Theater) Capacity() int {
	return t.capacity
}

func (t Theater) TicketPrice(section string) (float64, error) {

	var (
		retval float64
		reterr error
	)

	if section == "orchestra" {
		retval, reterr = 200.00, nil
	} else if section == "balcony" {
		retval, reterr = 90.00, nil
	} else {
		retval, reterr = 0.00, errors.New("invalid theater section")
	}
	return retval, reterr
}

func (e *Event) MiddleAttendees() ([]string, error) {
	if len(e.Attendees) < 3 {
		return nil, errors.New("fewer than 3 attendees in event")
	}
	return e.Attendees[1 : len(e.Attendees)-1 : len(e.Attendees)-1], nil
}

//*******************************************************************

func PrintVenueInfo(v Venue) {
	fmt.Printf("\nName: %s\n", v.Name())
	fmt.Printf("Capacity: %d\n", v.Capacity())

	var (
		sec1Price float64
		sec2Price float64
		err       error
	)

	switch v.(type) {
	case Arena:
		sec1Price, err = v.TicketPrice("floor")
		if err != nil {
			fmt.Println(err)
		}
		sec2Price, err = v.TicketPrice("upper")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Ticket prices: $%.2f (floor), $%.2f (upper)\n", sec1Price, sec2Price)
	case Theater:
		sec1Price, err = v.TicketPrice("orchestra")
		if err != nil {
			fmt.Println(err)
		}
		sec2Price, err = v.TicketPrice("balcony")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Ticket prices: $%.2f (orchestra), $%.2f (balcony)\n", sec1Price, sec2Price)
	default:
		fmt.Println("unknown venue type")
	}
}

//*******************************************************************

// Free functions

func RegisterEvent(r Registry, id string, event *Event) error {
	if _, ok := r.registry[id]; ok {
		return errors.New("event already registered")
	}
	r.registry[id] = event
	return nil
}

func AddAttendee(r Registry, id string, attendee string) error {
	if _, ok := r.registry[id]; !ok {
		return errors.New("no such event in registry")
	}
	r.registry[id].Attendees = append(r.registry[id].Attendees, attendee)
	return nil
}

func SummarizeEvents(r Registry, ids ...string) error {
	if len(ids) == 0 {
		return errors.New("no IDs passed for summary")
	}

	for _, n := range ids {
		if _, ok := r.registry[n]; ok {
			fmt.Printf("\nTitle: %s\n", r.registry[n].Title)
			fmt.Printf("Name: %s\n", r.registry[n].VenueName)
			fmt.Printf("Attendees: %d\n", len(r.registry[n].Attendees))

			atts, err := r.registry[n].MiddleAttendees()
			if err != nil {
				fmt.Println(err) // Not sure what to do with this one, so just printing
			} else {
				fmt.Printf("Middle Attendees: %s\n", atts)
			}
		}
	}
	return nil
}

func Describe(v any) string {
	var retval string
	switch x := v.(type) {
	case string:
		retval = "text: " + x
	case int:
		retval = "number: " + strconv.Itoa(x)
	case bool:
		retval = "flag: " + strconv.FormatBool(x)
	case Venue:
		retval = "venue: " + x.Name()
	default:
		retval = "unknown"
	}
	return retval
}

//*******************************************************************

func main() {

	venue1 := Arena{name: "O2", capacity: 20000}
	venue2 := Theater{name: "Ryman", capacity: 2362}

	PrintVenueInfo(venue1)
	PrintVenueInfo(venue2)

	event1 := Event{Title: "Def Leppard", VenueName: "O2", Attendees: make([]string, 0)}
	event2 := Event{Title: "U2", VenueName: "O2", Attendees: make([]string, 0)}
	event3 := Event{Title: "Ryan Adams", VenueName: "Ryman", Attendees: make([]string, 0)}

	reg := Registry{registry: make(map[string]*Event)}

	err := RegisterEvent(reg, event1.Title, &event1)
	if err != nil {
		fmt.Println(err)
	}

	err = RegisterEvent(reg, event2.Title, &event2)
	if err != nil {
		fmt.Println(err)
	}

	err = RegisterEvent(reg, event3.Title, &event3)
	if err != nil {
		fmt.Println(err)
	}

	// Add Def Leppard attendees
	err = AddAttendee(reg, "Def Leppard", "John Smith")
	if err != nil {
		fmt.Println(err)
	}
	err = AddAttendee(reg, "Def Leppard", "Alex Johnson")
	if err != nil {
		fmt.Println(err)
	}
	err = AddAttendee(reg, "Def Leppard", "Steve Wilson")
	if err != nil {
		fmt.Println(err)
	}
	err = AddAttendee(reg, "Def Leppard", "Katie Jackson")
	if err != nil {
		fmt.Println(err)
	}
	err = AddAttendee(reg, "Def Leppard", "Jack Edwards")
	if err != nil {
		fmt.Println(err)
	}

	// Add U2 attendees
	err = AddAttendee(reg, "U2", "Peter Jordan")
	if err != nil {
		fmt.Println(err)
	}
	err = AddAttendee(reg, "U2", "Ivan Ellis")
	if err != nil {
		fmt.Println(err)
	}

	// Add Ryan Adams attendees
	err = AddAttendee(reg, "Ryan Adams", "Aiden Paxton")
	if err != nil {
		fmt.Println(err)
	}

	err = SummarizeEvents(reg, "Def Leppard", "AC/DC", "Metallica", "Ryan Adams", "U2")
	if err != nil {
		fmt.Println(err)
	}

	str := Describe("hello")
	fmt.Printf("\n%s\n", str)
	str = Describe(1234)
	fmt.Println(str)
	str = Describe(false)
	fmt.Println(str)
	str = Describe(venue2)
	fmt.Println(str)
	str = Describe(567.8)
	fmt.Println(str)

	fmt.Printf("\nConst Val: %s\n", dummy)
}
