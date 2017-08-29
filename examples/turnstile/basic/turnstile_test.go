package basic_test

import (
	"github.com/goph/fsm/examples/turnstile"
	"github.com/goph/fsm/examples/turnstile/basic"
)

func Example_insertACoinAndPass() {
	t := basic.New()
	stateMachine := turnstile.NewStateMachine()

	stateMachine.Trigger(t.GetState(), "coin_inserted", t)
	stateMachine.Trigger(t.GetState(), "pushed", t)

	// Output:
	// Coin inserted, you shall pass
	// Passed the gate, coin please
}

func Example_cannotPassWhenLocked() {
	t := basic.New()
	stateMachine := turnstile.NewStateMachine()

	stateMachine.Trigger(t.GetState(), "pushed", t)

	// Output:
	// You shall not pass
}

func Example_insertCoinsAndPassOnce() {
	t := basic.New()
	stateMachine := turnstile.NewStateMachine()

	stateMachine.TriggerSubject(t, "coin_inserted")
	stateMachine.TriggerSubject(t, "coin_inserted")
	stateMachine.TriggerSubject(t, "pushed")
	stateMachine.TriggerSubject(t, "pushed")

	// Output:
	// Coin inserted, you shall pass
	// Coin inserted, you shall pass
	// Passed the gate, coin please
	// You shall not pass
}
