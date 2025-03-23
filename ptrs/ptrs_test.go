package ptrs

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
    wallet := Wallet{}

    var btc Bitcoin = 10

    wallet.Deposit(btc)

    got := wallet.Balance()

    var want Bitcoin = 1

    if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
