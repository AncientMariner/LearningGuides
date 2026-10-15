package behavioral

import "fmt"

// State: changes behaviour based on internal state.
// Use an interface-per-state for complex transitions; iota+switch for simple ones.

type TrafficLight struct {
	state LightState
}

type LightState interface {
	Next(t *TrafficLight)
	String() string
}

type RedState struct{}
type GreenState struct{}
type YellowState struct{}

func (s *RedState) Next(t *TrafficLight)    { t.state = &GreenState{} }
func (s *GreenState) Next(t *TrafficLight)  { t.state = &YellowState{} }
func (s *YellowState) Next(t *TrafficLight) { t.state = &RedState{} }

func (s *RedState) String() string    { return "RED" }
func (s *GreenState) String() string  { return "GREEN" }
func (s *YellowState) String() string { return "YELLOW" }

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{state: &RedState{}}
}

func (t *TrafficLight) Change() {
	fmt.Printf("Light: %s\n", t.state)
	t.state.Next(t)
}

func ExampleState() {
	light := NewTrafficLight()
	light.Change()
	light.Change()
	light.Change()
	// Output:
	// Light: RED
	// Light: GREEN
	// Light: YELLOW
}
