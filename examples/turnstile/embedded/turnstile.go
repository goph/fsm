// Package embedded embeds the state machine into the subject (turnstile) and exposes commands hiding the state machine.
package embedded

import (
	"github.com/goph/fsm"
	"github.com/goph/fsm/examples/turnstile"
)

// Turnstile is a a mechanical gate consisting of revolving horizontal arms fixed to a vertical post, allowing only one person at a time to pass through if the entry fee is paid.
type Turnstile struct {
	state string

	stateMachine *fsm.StateMachine
}

// New returns a new Turnstile.
func New() *Turnstile {
	return &Turnstile{
		state: turnstile.Locked,

		stateMachine: turnstile.NewStateMachine(),
	}
}

// GetState returns the current state of the turnstile.
func (t *Turnstile) GetState() string {
	return t.state
}

// SetState sets the state of the Turnstile.
//
// Note: this is generally a huge mistake and leat, DO NOT DO IT.
// This method is purely here for the sake of the examples.
func (t *Turnstile) SetState(state string) {
	t.state = state
}

// InsertCoin fires a coin_inserted event for the turnstile.
func (t *Turnstile) InsertCoin() {
	t.stateMachine.TriggerSubject(t, "coin_inserted")
}

// Push fires a pushed event for the turnstile.
func (t *Turnstile) Push() {
	t.stateMachine.TriggerSubject(t, "pushed")
}
