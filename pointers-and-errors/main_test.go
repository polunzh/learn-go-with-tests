package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("want an error but didn't get one")
		}

		if want.Error() != got.Error() {
			t.Errorf("got %s, want %s", got.Error(), want)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Fatal()
		}

		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient founds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))
		want := Bitcoin(20)

		assertBalance(t, wallet, want)
		assertError(t, err, ErrInsufficientFunds)
	})
}
