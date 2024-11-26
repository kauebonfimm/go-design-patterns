package command

import "fmt"

type DispatchEvent struct {
	EventName  string
	Parameters any
}

func (d *DispatchEvent) Send() error {
	fmt.Println(d.EventName)
	fmt.Println("Event dispatch")
	return nil
}
