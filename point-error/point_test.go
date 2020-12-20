package pointAndError

import "testing"

func TestWallet(t *testing.T) {
	// wallet := Wallet{}
	// wallet.Deposit(Bitcoin(10))
	// got := wallet.Balance()
	// want := Bitcoin(20)

	// if got != want {
	// 	t.Errorf("got %s want %s", got, want)
	// }

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		// got := wallet.Balance()
		// want := Bitcoin(10)
		assertBalance(t, wallet, Bitcoin(10))
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		// got := wallet.Balance()
		// want := Bitcoin(10)
		assertBalance(t, wallet, Bitcoin(10))
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertNoError(t, err)
	})

	t.Run("Withdraw is over balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}
