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

    fmt.Printf("address of balance in test is %p \n", &wallet.balance)

    var want Bitcoin = 10

    if(got != want) {
        t.Errorf("got %d want %d", got, want)
    }
}
