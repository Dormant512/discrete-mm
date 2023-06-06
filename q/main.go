package main

import (
	"fmt"
	"os"
	"time"
)

// metrics
// 1) Average waiting time in queue
// 2) Probability that a customer has to wait
// 3) Probability of an Idle server
// 4) Average service time (theoretical 3.2)
// 5) Average time between arrivals (theoretical 4.5)
// 6) Average waiting time for those who wait
// 7) Average time a customer spends in the system

type Customer struct {
	Id                         int
	PeopleBefore               int
	MetricCreatedTime          time.Time
	MetricArrivedToQueueTime   time.Time
	MetricArrivedToCashierTime time.Time
	MetricLeftTime             time.Time
}

func main() {
	fCli, err := os.Create(fmt.Sprintf("./logs/clients_dynamic_%t_%d.csv", DYNAMIC, MINWORKERS))
	if err != nil {
		panic(err)
	}
	defer func() {
		closeErr := fCli.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	fServ, err := os.Create(fmt.Sprintf("./logs/servers_dynamic_%t_%d.csv", DYNAMIC, MINWORKERS))
	if err != nil {
		panic(err)
	}
	defer func() {
		closeErr := fServ.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	fCli.WriteString("Customer,Cashier,Created,ArrivedQueue,ArrivedCashier,Left,PeopleBefore\n")
	fServ.WriteString("Server,Action,Time\n")

	users := make(chan Customer, CUSTOMERS)
	var qs [MAXWORKERS]*chan Customer
	for idx := 0; idx < MINWORKERS; idx++ {
		curChan := make(chan Customer, CUSTOMERS)
		qs[idx] = &curChan
	}

	if DYNAMIC {
		Dynamic(users, qs, fCli, fServ)
	} else {
		Static(users, qs, fCli, fServ)
	}
}
