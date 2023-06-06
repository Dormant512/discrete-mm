package main

import (
	"os"
	"sync"
)

func Dynamic(users chan Customer, qs [MAXWORKERS]*chan Customer, fCli, fServ *os.File) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(users)
		CreateUsers(users, CUSTOMERS)
	}()

	for i := 0; i < MINWORKERS; i++ {
		wg.Add(1)
		idx := i
		go func(idx int) {
			defer wg.Done()
			//fmt.Println("Cashier", idx, "started.")
			Cashier(qs, idx, WORKTIME, fCli, fServ)
		}(idx)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		FanoutUsersDynamic(users, qs, ARRIVETIME, &wg, fCli, fServ)
	}()

	wg.Wait()
}
