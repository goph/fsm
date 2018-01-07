// Package fsm contains a state machine implemenation.
//
// See the examples directory for detailed usage examples.
package fsm

import "fmt"

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
	Handle(action string, fromState string, toState string, args []interface{}) error
}

// StateMachineAwareDelegate plays a role when a state transition itself requires another state transition to happen.
//
// For example: some kind of business validation fails and a workflow needs to be terminated immediately.
//
// Since this kind of behaviour requires a circular dependency between the state machine and the delegate,
// regular "constructor" injection doesn't work in this case.
type StateMachineAwareDelegate interface {
	// SetStateMachine sets the current state machine in the delegate
	SetStateMachine(sm *StateMachine)
}

// Transition represents a state transition.
type Transition struct {
	FromState string
	Event     string
	ToState   string
	Action    string
}

// InvalidTransitionError is returned when a transition is invalid.
type InvalidTransitionError struct {
	currentState string
	event        string
	args         []interface{}
}

// Error returns the formatted error message.
func (e *InvalidTransitionError) Error() string {
	return fmt.Sprintf("cannot transition from %q state triggered by %q event", e.currentState, e.event)
}

// CurrentState returns the current state.
func (e *InvalidTransitionError) CurrentState() string {
	return e.currentState
}

// Event returns the current state.
func (e *InvalidTransitionError) Event() string {
	return e.event
}

// Arguments returns the current state.
func (e *InvalidTransitionError) Arguments() []interface{} {
	return e.args
}

type DelegateError struct {
	err          error
	currentState string
	event        string
	args         []interface{}
}

// Cause implements the causer interface from github.com/pkg/errors.
func (e *DelegateError) Cause() error {
	return e.err
}

// Error returns the formatted error message.
func (e *DelegateError) Error() string {
	return fmt.Sprintf(
		"delegate reported an error during transition from %q state triggered by %q event: %s",
		e.currentState,
		e.event,
		e.err.Error(),
	)
}

// CurrentState returns the current state.
func (e *DelegateError) CurrentState() string {
	return e.currentState
}

// Event returns the current state.
func (e *DelegateError) Event() string {
	return e.event
}

// Arguments returns the current state.
func (e *DelegateError) Arguments() []interface{} {
	return e.args
}

// StateMachine handles state transitions when an event is fired and calls the underlying delegate.
type StateMachine struct {
	delegate    Delegate
	transitions []Transition
}

// NewStateMachine returns a new StateMachine.
func NewStateMachine(delegate Delegate, transitions []Transition) *StateMachine {
	stateMachine := &StateMachine{
		transitions: transitions,
	}

	if smaDelegate, ok := delegate.(StateMachineAwareDelegate); ok {
		smaDelegate.SetStateMachine(stateMachine)
	}

	stateMachine.delegate = delegate

	return stateMachine
}

// Trigger fires an event and calls the underlying delegate.
func (sm *StateMachine) Trigger(currentState string, event string, args ...interface{}) error {
	t := sm.findTransition(currentState, event)
	if t == nil {
		return &InvalidTransitionError{
			currentState: currentState,
			event:        event,
			args:         args,
		}
	}

	if t.Action != "" {
		err := sm.delegate.Handle(t.Action, t.FromState, t.ToState, args)
		if err != nil {
			return &DelegateError{
				err:          err,
				currentState: currentState,
				event:        event,
				args:         args,
			}
		}
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
	GetState() string
}

// TriggerSubject triggers an event using the Subject's current state.
//
// It also passes the subject as the first argument.
func (sm *StateMachine) TriggerSubject(subject Subject, event string, args ...interface{}) error {
	args = append([]interface{}{subject}, args...)

	return sm.Trigger(subject.GetState(), event, args...)
}
