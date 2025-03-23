package ptrs

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
    wallet := Wallet{}

    wallet.Deposit(20)

    got := wallet.Balance()

    fmt.Printf("address of balance in test is %p \n", &wallet.balance)

    want := 20

    if(got != want) {
        t.Errorf("got %d want %d", got, want)
    }
}
