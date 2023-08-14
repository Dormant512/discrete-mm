package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func CreateUsers(customers chan<- Customer, capacity int) {
	for i := 0; i < capacity; i++ {
		newDude := Customer{
			Id:                         i,
			PeopleBefore:               0,
			MetricCreatedTime:          time.Now(),
			MetricArrivedToQueueTime:   time.Now(),
			MetricArrivedToCashierTime: time.Now(),
			MetricLeftTime:             time.Now(),
		}
		customers <- newDude
	}
}

func Cashier(qs [MAXWORKERS]*chan Customer, id, maxTime int, fCli, fServ *os.File) {
	servMsg := fmt.Sprintf("%d,%s,%d\n", id, "CREATED", time.Now().UnixMilli())
	fServ.WriteString(servMsg)
	lastWorked := time.Now()
	for {
		// Cashier idle
		if time.Now().Sub(lastWorked).Seconds() >= MAXCHILL && id >= MINWORKERS {
			qs[id] = nil
			servMsg = fmt.Sprintf("%d,%s,%d\n", id, "DELETED", time.Now().UnixMilli())
			fServ.WriteString(servMsg)
			return
		}
		cur, ok := <-*qs[id]
		if ok {
			cur.MetricArrivedToCashierTime = time.Now()
			//workingTime := rand.Intn(maxTime)
			//time.Sleep(time.Duration(workingTime) * time.Millisecond)
			cur.MetricLeftTime = time.Now()
			// Customer, Cashier, Created, ArrivedQueue, ArrivedCashier, Left
			msg := fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d\n",
				cur.Id,
				id,
				cur.MetricCreatedTime.UnixMilli(),
				cur.MetricArrivedToQueueTime.UnixMilli(),
				cur.MetricArrivedToCashierTime.UnixMilli(),
				cur.MetricLeftTime.UnixMilli(),
				cur.PeopleBefore,
			)
			fCli.WriteString(msg)
			lastWorked = time.Now()
			servMsg = fmt.Sprintf("%d,%s,%d\n", id, "STARTED", cur.MetricArrivedToCashierTime.UnixMilli())
			fServ.WriteString(servMsg)
			servMsg = fmt.Sprintf("%d,%s,%d\n", id, "FINISHED", cur.MetricLeftTime.UnixMilli())
			fServ.WriteString(servMsg)
		} else {
			return
		}
	}
}

func FanoutUsersStatic(customers chan Customer, qs [MAXWORKERS]*chan Customer, maxTime int) {
	for user := range customers {
		userDest := rand.Intn(MINWORKERS)
		minLen := len(*qs[userDest])

		//shoppingTime := rand.Intn(maxTime)
		//time.Sleep(time.Duration(shoppingTime) * time.Millisecond)

		for i := 0; i < MINWORKERS; i++ {
			if len(*qs[i]) < minLen {
				minLen = len(*qs[i])
				userDest = i
			}
		}
		user.PeopleBefore = minLen
		user.MetricArrivedToQueueTime = time.Now()
		*qs[userDest] <- user
	}
	for _, q := range qs {
		if q == nil {
			continue
		}
		close(*q)
	}
}

func FanoutUsersDynamic(customers chan Customer, qs [MAXWORKERS]*chan Customer, maxTime int, wg *sync.WaitGroup, fCli, fServ *os.File) {
	for user := range customers {
		available := make([]int, 0)
		for i, q := range qs {
			if q == nil {
				continue
			}
			available = append(available, i)
		}
		userDest := rand.Intn(len(available))
		minLen := len(*qs[available[userDest]])

		//shoppingTime := rand.Intn(maxTime)
		//time.Sleep(time.Duration(shoppingTime) * time.Millisecond)

		// user looks at all the q's and considers, to wait or not
		emptyQ := false
		for _, q := range qs {
			if q == nil {
				continue
			}
			if len(*q) == 0 {
				emptyQ = true
				break
			}
		}

		isWaiting := true
		if !emptyQ {
			isWaiting = rand.Intn(2) == 0
			//fmt.Println("isWaiting", isWaiting)
		}

		if !isWaiting {
			// create new cashier, assign there
			newIdx := -1
			for idx, val := range qs {
				if val != nil {
					continue
				}
				// this one is nil
				curChan := make(chan Customer, CUSTOMERS)
				qs[idx] = &curChan
				newIdx = idx
				break
			}

			if newIdx != -1 {
				wg.Add(1)
				go func(idx int) {
					defer wg.Done()
					//fmt.Println("Dynamic cashier", idx, "started.")
					Cashier(qs, idx, WORKTIME, fCli, fServ)
				}(newIdx)
				user.PeopleBefore = 0
				user.MetricArrivedToQueueTime = time.Now()
				*qs[newIdx] <- user
				continue
			}
			//fmt.Println("User", user.Id, "couldn't create a new queue")
		}

		for i, q := range qs {
			if q == nil {
				continue
			}
			if len(*q) < minLen {
				minLen = len(*q)
				userDest = i
			}
		}
		user.PeopleBefore = minLen
		user.MetricArrivedToQueueTime = time.Now()
		*qs[userDest] <- user
		//fmt.Println("User", user.Id, "went to queue", userDest)
	}
	for _, q := range qs {
		if q == nil {
			continue
		}
		close(*q)
	}
}
