package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
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
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("%#v - got:'%s' - want:'%s'", wallet, got, want)
	}
}

// asserts that we did not get an error
// if an error is thrown, this test fails
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didnt get one")
	}
}

// assertion for when we do in fact expect an error
// if we expect an error and dont get one, this test fails
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted err but didnt get one")
	}

	if got != want {
		t.Errorf("got: '%q' - want:'%q'", got, want)
	}
}
