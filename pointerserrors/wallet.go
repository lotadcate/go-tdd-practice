package pointerserrors
import "fmt"

type Wallet struct {
  balance int
}

func(w *Wallet) Deposit(amount int) {
  w.balance += amount 
}

func(w *Wallet) Balance() int {
  fmt.Printf("address of balance in test is %v \n", &w.balance)
  return w.balance
}
