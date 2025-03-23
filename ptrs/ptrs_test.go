package ptrs

import (
	"testing"
)

func TestWallet(t *testing.T) {
    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        var btc Bitcoin = 10

        wallet.Deposit(btc)

        assertBalance(t, wallet, btc)
    })

    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(10)}
        wallet.Withdraw(Bitcoin(5))

        assertBalance(t, wallet, Bitcoin(5))
    })


    t.Run("withdraw insuffient funds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{balance: startingBalance}

        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsuffientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}



func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}  

func assertError(t testing.TB, got error, want error) {
    t.Helper()

    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
