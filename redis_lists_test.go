package redis

import (
	"fmt"
)

func ExampleRedis_LPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lpop"
	r.RPush(key, "a", "b", "c")
	fmt.Println(r.LPop(key))

	// Output:
	// a
}

func ExampleRedis_RPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rpop"
	r.RPush(key, "a", "b", "c")
	fmt.Println(r.RPop(key))

	// Output:
	// c
}
