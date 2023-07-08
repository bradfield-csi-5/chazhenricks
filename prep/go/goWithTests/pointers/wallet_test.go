package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("%#v - got:'%s' - want:'%s'", wallet, got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted err but didnt get one")
		}

		if got != want {
			t.Errorf("got: '%q' - want:'%q'", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, 15)
	})

	t.Run("insuffucent funds withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		startingBalance := wallet.Balance()
		err := wallet.Withdraw(Bitcoin(15))

		assertError(t, err, ErrInsuffucentFunds)
		assertBalance(t, wallet, startingBalance)

	})
}
