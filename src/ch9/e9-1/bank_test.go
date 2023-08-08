// Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
// program. The result should indicate whether the transaction succeeded or failed due to insuf-
// ficient funds. The message sent to the monitor goroutine must contain both the amount to
// withdraw and a new channel over which the monitor goroutine can send the boolean result
// back to Withdraw.

package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		Withdraw(100)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(300)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	if got, want := Balance(), 400; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if succeeded, bal := Withdraw(600), 400; succeeded == true || bal != Balance() {
		t.Errorf("Withdraw broke invariants; succeeded: %t; bal = %d", succeeded, bal)
	}
}
