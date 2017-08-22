// Package pool embeds a state machine pool into the subject.
//
// Although using a state machine should be safe for concurrent usage, delegates are out of control of this library and by using pools one can make sure that the state machine is in fact concurrent safe.
package pool

import (
	"sync"

	"github.com/goph/fsm"
	"github.com/goph/fsm/examples/turnstile"
)

// Turnstile is a a mechanical gate consisting of revolving horizontal arms fixed to a vertical post, allowing only one person at a time to pass through if the entry fee is paid.
type Turnstile struct {
	state string

	pool sync.Pool
}

// New returns a new Turnstile.
func New() *Turnstile {
	return &Turnstile{
		state: turnstile.Locked,

		pool: sync.Pool{
			New: func() interface{} {
				return turnstile.NewStateMachine()
			},
		},
	}
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

// InsertCoin fires a coin_inserted event for the turnstile.
func (t *Turnstile) InsertCoin() {
	sm := t.pool.Get().(*fsm.StateMachine)

	sm.TriggerSubject(t, "coin_inserted")

	t.pool.Put(sm)
}

// Push fires a pushed event for the turnstile.
func (t *Turnstile) Push() {
	sm := t.pool.Get().(*fsm.StateMachine)

	sm.TriggerSubject(t, "pushed")

	t.pool.Put(sm)
}
