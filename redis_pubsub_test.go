package redis

import (
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
)

func ExampleRedis_PubSub() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	psc, _ := r.PubSub()

	var wg sync.WaitGroup
	wg.Add(2)

	// This goroutine receives and prints pushed notifications from the server.
	// The goroutine exits when the connection is unsubscribed from all
	// channels or there is an error.
	go func() {
		defer wg.Done()
		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
			case redis.PMessage:
				fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
			case redis.Subscription:
				fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
				if n.Count == 0 {
					return
				}
			case error:
				fmt.Printf("error: %v\n", n)
				return
			}
		}
	}()

	// This goroutine manages subscriptions for the connection.
	go func() {
		defer wg.Done()

		psc.Subscribe("example")
		psc.PSubscribe("p*")

		r.Publish("example", "hello")
		r.Publish("example", "world")
		r.Publish("example", "foo")
		r.Publish("example", "bar")
		r.Publish("publish", "pattern")

		// Unsubscribe from all connections. This will cause the receiving
		// goroutine to exit.
		psc.Unsubscribe()
		psc.PUnsubscribe()
	}()

	wg.Wait()

	// Output:
	// Subscription: subscribe example 1
	// Subscription: psubscribe p* 2
	// Message: example hello
	// Message: example world
	// Message: example foo
	// Message: example bar
	// PMessage: p* publish pattern
	// Subscription: unsubscribe example 1
	// Subscription: punsubscribe p* 0
}
