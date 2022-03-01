package csp

import (
	"context"
	"fmt"
	"sync"

	"github.com/raspiantoro/go-actor/payload"
	"github.com/raspiantoro/go-actor/random"
)

type ctxKey string

var (
	counterKey ctxKey = "counter"
)

type message struct {
	id  int
	val uint64
}

func responder(ctx context.Context, req chan message, resp chan message) {
	ct := ctx.Value(counterKey).(*payload.Counter)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("aggregator done")
			return
		case msg := <-req:
			random.Sleep()
			ct.Count += msg.val
			msg.val = ct.Count
			resp <- msg
			fmt.Printf("reply has been sent to goroutine %d with total count: %d\n", msg.id, ct.Count)
		}
	}
}

func ExecuteGoroutine() {
	ct := &payload.Counter{
		Count: 50,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, counterKey, ct)
	ctx, cancel := context.WithCancel(ctx)

	requestorNums := 10
	wg := sync.WaitGroup{}
	wg.Add(requestorNums)

	reqChan := make(chan message, 10)
	respChan := make(chan message, 10)

	go func() {
		wg.Wait()
		cancel()
	}()

	for i := 0; i < requestorNums; i++ {
		go func(id int) {
			reqChan <- message{
				id:  id,
				val: 1,
			}

			fmt.Printf("goroutine %d already sent message\n", id)

			random.Sleep()

			msg := <-respChan

			fmt.Printf("goroutine %d receive from %d, current total count: %d\n", id, msg.id, msg.val)

			wg.Done()
		}(i)
	}

	responder(ctx, reqChan, respChan)

	fmt.Println("total count is: ", ct.Count)
}
