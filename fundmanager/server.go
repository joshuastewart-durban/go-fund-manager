package fundmanager

import "fmt"

type FundServer struct {
	Commands chan interface{}
	fund     *Fund
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make() creates builtins like channels, maps, and slices
		Commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	// The built-in "range" clause can iterate over channels,
	// amongst other things
	for command := range s.Commands {

		switch command.(type) {
		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}

	}
}
