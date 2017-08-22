// Package fsm contains a state machine implemenation.
//
// See the examples directory for detailed usage examples.
package fsm

import "errors"

// Delegate is responsible for handling actions whenever a transition has one.
//
// The actual state change of the transition might happen inside a delegate.
type Delegate interface {
	// Handle handles transition actions.
	//
	// For now handling an action cannot result in an error.
	// This is aligned with the nature of a state machine:
	// every probable input needs to have a valid input.
	//
	// Eg. invalid input should prevent a state change.
	// Erroneous cases are recommended to be logged.
	//
	// Note: this might change in the future.
	Handle(action string, fromState string, toState string, args []interface{})
}

// Transition represents a state transition.
type Transition struct {
	FromState string
	Event     string
	ToState   string
	Action    string
}

// StateMachine handles state transitions when an event is fired and calls the underlying delegate.
type StateMachine struct {
	delegate    Delegate
	transitions []Transition
}

// NewStateMachine returns a new StateMachine.
func NewStateMachine(delegate Delegate, transitions []Transition) *StateMachine {
	return &StateMachine{
		delegate,
		transitions,
	}
}

// Trigger fires an event and calls the underlying delegate.
func (sm *StateMachine) Trigger(currentState string, event string, args ...interface{}) error {
	t := sm.findTransition(currentState, event)
	if t == nil {
		return errors.New("invalid transition")
	}

	if t.Action != "" {
		sm.delegate.Handle(t.Action, t.FromState, t.ToState, args)
	}

	return nil
}

// findTransition returns a transition if there is one for the state-event pair.
func (sm *StateMachine) findTransition(fromState string, event string) *Transition {
	for _, t := range sm.transitions {
		if t.FromState == fromState && t.Event == event {
			return &t
		}
	}

	return nil
}

// Subject represents a stateful structure exposing it's current state.
type Subject interface {
	State() string
}

// TriggerSubject triggers an event using the Subject's current state.
//
// It also passes the subject as the first argument.
func (sm *StateMachine) TriggerSubject(subject Subject, event string, args ...interface{}) {
	args = append([]interface{}{subject}, args...)

	sm.Trigger(subject.State(), event, args...)
}
