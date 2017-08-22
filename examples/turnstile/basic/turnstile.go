// Package basic explains the basic usage of a state machine.
package basic

import "github.com/goph/fsm/examples/turnstile"

// Turnstile is a a mechanical gate consisting of revolving horizontal arms fixed to a vertical post, allowing only one person at a time to pass through if the entry fee is paid.
type Turnstile struct {
	state string
}

// State returns the current state of the turnstile.
func (t *Turnstile) State() string {
	return t.state
}

// SetState sets the state of the Turnstile.
//
// Note: this is generally a huge mistake and leat, DO NOT DO IT.
// This method is purely here for the sake of the examples.
func (t *Turnstile) SetState(state string) {
	t.state = state
}

// New returns a new Turnstile.
func New() *Turnstile {
	return &Turnstile{
		state: turnstile.Locked,
	}
}
