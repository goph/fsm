package pool_test

import "github.com/goph/fsm/examples/turnstile/pool"

func Example_insertACoinAndPass() {
	t := pool.New()

	t.InsertCoin()
	t.Push()

	// Output:
	// Coin inserted, you shall pass
	// Passed the gate, coin please
}

func Example_cannotPassWhenLocked() {
	t := pool.New()

	t.Push()

	// Output:
	// You shall not pass
}

func Example_insertCoinsAndPassOnce() {
	t := pool.New()

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
