/*
Package turnstile provides an example for the state machine library.

A turnstile is a a mechanical gate consisting of revolving horizontal arms fixed to a vertical post, allowing only one person at a time to pass through if the entry fee is paid.

This example is separated into three parts:

Basic: explains the basic usage of a state machine

Embedded: embeds the state machine into the subject (turnstile) and exposes commands hiding the state machine

Pool: embeds a state machine pool into the subject. Although using a state machine should be safe for concurrent usage, delegates are out of control of this library and by using pools one can make sure that the state machine is in fact concurrent safe.
*/
package turnstile

import (
	"fmt"

	"github.com/goph/fsm"
)

const (
	// Locked represents the locked turnstile state.
	Locked = "locked"

	// Unlocked represents the unlocked turnstile state.
	Unlocked = "unlocked"
)

// Turnstile is a a mechanical gate consisting of revolving horizontal arms fixed to a vertical post, allowing only one person at a time to pass through if the entry fee is paid.
type Turnstile interface {
	// SetState sets the state of the Turnstile.
	//
	// Note: this is generally a huge mistake and leat, DO NOT DO IT.
	// This method is purely here for the sake of the examples.
	SetState(state string)
}

// NewStateMachine returns a new StateMachine for a turnstile.
func NewStateMachine() *fsm.StateMachine {
	return fsm.NewStateMachine(
		fsm.NewActionMuxDelegate(map[string]fsm.Delegate{
			"coin":   &coinAction{},
			"pass":   &passAction{},
			"nopass": &noPassAction{},
		}),
		[]fsm.Transition{
			fsm.Transition{
				FromState: Locked,
				Event:     "coin_inserted",
				ToState:   Unlocked,
				Action:    "coin",
			},
			fsm.Transition{
				FromState: Unlocked,
				Event:     "coin_inserted",
				ToState:   Unlocked,
				Action:    "coin",
			},
			fsm.Transition{
				FromState: Unlocked,
				Event:     "pushed",
				ToState:   Locked,
				Action:    "pass",
			},
			fsm.Transition{
				FromState: Locked,
				Event:     "pushed",
				ToState:   Locked,
				Action:    "nopass",
			},
		},
	)
}

// coinAction is the delegate called when a coin is placed in the machine.
type coinAction struct{}

// Handle unlocks the turnstile.
func (d *coinAction) Handle(action string, fromState string, toState string, args []interface{}) {
	if turnstile, ok := args[0].(Turnstile); ok {
		turnstile.SetState(toState)

		fmt.Println("Coin inserted, you shall pass")
	}
}

// passAction is the delegate called when the turnstile is pushed.
type passAction struct{}

// Handle unlocks the turnstile.
func (d *passAction) Handle(action string, fromState string, toState string, args []interface{}) {
	if turnstile, ok := args[0].(Turnstile); ok {
		turnstile.SetState(toState)

		fmt.Println("Passed the gate, coin please")
	}
}

// noPassAction is the delegate called when the turnstile is pushed.
type noPassAction struct{}

// Handle unlocks the turnstile.
func (d *noPassAction) Handle(action string, fromState string, toState string, args []interface{}) {
	fmt.Println("You shall not pass")
}
