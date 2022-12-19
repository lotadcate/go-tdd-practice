package pointerserrors
import (
  "fmt"
  "errors"
  )

type Stringer interface {
  String() string
}

type Bitcoin int

func(b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
  balance Bitcoin
}

func(w *Wallet) Deposit(amount Bitcoin) {
  w.balance += amount 
}

func(w *Wallet) Balance() Bitcoin {
  return w.balance
}

var ErrInsufficientfunds = errors.New("cannot withdraw, insufficient funds")
func(w *Wallet) WithDraw(amount Bitcoin) error {
  if amount > w.balance {
    return ErrInsufficientfunds
  }
  w.balance -= amount
  return nil
}
