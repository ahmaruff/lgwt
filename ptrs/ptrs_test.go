package ptrs

import (
	"testing"
)

func TestWallet(t *testing.T) {
    wallet := Wallet{}

    var btc Bitcoin = 10

    wallet.Deposit(btc)

    got := wallet.Balance()

    var want Bitcoin = 10

    if got != want {
		t.Errorf("got %s want %s", got, want)
	}


    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        var btc Bitcoin = 10

        wallet.Deposit(btc)
        got := wallet.Balance()
        var want Bitcoin = 10

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })

    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(10)}

        wallet.Withdraw(Bitcoin(5))
        got := wallet.Balance()
        want := Bitcoin(5)
        

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })
}
