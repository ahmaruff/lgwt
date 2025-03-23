package ptrs

import (
	"testing"
)

func TestWallet(t *testing.T) {

    assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()


        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }  



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
}
