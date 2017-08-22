package embedded_test

import "github.com/goph/fsm/examples/turnstile/embedded"

func Example_insertACoinAndPass() {
	t := embedded.New()

	t.InsertCoin()
	t.Push()

	// Output:
	// Coin inserted, you shall pass
	// Passed the gate, coin please
}

func Example_cannotPassWhenLocked() {
	t := embedded.New()

	t.Push()

	// Output:
	// You shall not pass
}

func Example_insertCoinsAndPassOnce() {
	t := embedded.New()

	t.InsertCoin()
	t.InsertCoin()
	t.Push()
	t.Push()

	// Output:
	// Coin inserted, you shall pass
	// Coin inserted, you shall pass
	// Passed the gate, coin please
	// You shall not pass
}
