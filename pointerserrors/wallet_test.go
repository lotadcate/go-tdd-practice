package pointerserrors
import (
  "testing"
)

func TestWallet (t *testing.T) {
  assertBalance := func(t testing.TB, got, want Bitcoin) {
    t.Helper()
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  }
  assertError := func(t testing.TB, err error, msg error){
    t.Helper()
    if err == nil {
      t.Fatal("didn't get an error but wanted one")
    }
    if err != msg {
      t.Errorf("err %q, msg %q", err, msg)
    }
  }

  t.Run("Deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))

    got := wallet.Balance()
    want := Bitcoin(10)
    assertBalance(t, got, want)
  })

  t.Run("WithDraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20)}
    wallet.WithDraw(Bitcoin(10))

    got := wallet.Balance()
    want := Bitcoin(10)
    assertBalance(t, got, want)
  })

  t.Run("WithDraw insufficient funds", func(t *testing.T) {
    rest := Bitcoin(20)
    wallet := Wallet{rest}
    err := wallet.WithDraw(Bitcoin(30))
    
    got := wallet.Balance()
    want := rest
    assertBalance(t, got, want)
    assertError(t, err, ErrInsufficientfunds)
  })
}








