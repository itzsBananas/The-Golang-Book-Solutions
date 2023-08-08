// Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
// program. The result should indicate whether the transaction succeeded or failed due to insuf-
// ficient funds. The message sent to the monitor goroutine must contain both the amount to
// withdraw and a new channel over which the monitor goroutine can send the boolean result
// back to Withdraw.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan invoice)

type invoice struct {
	amount  int
	success chan<- bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	suc := make(chan bool)
	withdrawals <- invoice{amount, suc}
	return <-suc
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount

		case invoice := <-withdrawals:
			if invoice.amount > balance {
				invoice.success <- false
				continue
			}
			balance -= invoice.amount
			invoice.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
