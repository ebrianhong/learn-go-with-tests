package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{5}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
