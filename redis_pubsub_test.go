package redis

import (
	"fmt"
)

func ExampleRedis_PUnsubscribe() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-punsubscribe*"
	r.PSubscribe(key)
	fmt.Println("Subscribe")

	r.PUnsubscribe(key)
	fmt.Println("Unsubscribe")

	// Output:
	// Subscribe
	// Unsubscribe
}
