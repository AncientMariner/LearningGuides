package behavioral

import "fmt"

// Command: encapsulates a request as an object.
// In Go, a plain func() is often sufficient; use an interface when
// you need undo, queuing, or logging.

type Command interface {
	Execute()
	Undo()
}

type Light struct{ on bool }

func (l *Light) TurnOn()  { l.on = true; fmt.Println("Light on") }
func (l *Light) TurnOff() { l.on = false; fmt.Println("Light off") }

type TurnOnCommand struct{ light *Light }
type TurnOffCommand struct{ light *Light }

func (c *TurnOnCommand) Execute()  { c.light.TurnOn() }
func (c *TurnOnCommand) Undo()     { c.light.TurnOff() }
func (c *TurnOffCommand) Execute() { c.light.TurnOff() }
func (c *TurnOffCommand) Undo()    { c.light.TurnOn() }

type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) Press(cmd Command) {
	cmd.Execute()
	r.history = append(r.history, cmd)
}

func (r *RemoteControl) UndoLast() {
	n := len(r.history)
	if n == 0 {
		return
	}
	r.history[n-1].Undo()
	r.history = r.history[:n-1]
}

func ExampleCommand() {
	light := &Light{}
	remote := &RemoteControl{}
	remote.Press(&TurnOnCommand{light: light})
	remote.Press(&TurnOffCommand{light: light})
	remote.UndoLast()
	// Output:
	// Light on
	// Light off
	// Light on
}
