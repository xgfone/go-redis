package redis

import (
	"fmt"
)

func ExampleRedis_Echo() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-echo"
	fmt.Println(r.Echo(key))

	// Output:
	// test-echo
}

func ExampleRedis_Ping() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-ping"
	fmt.Println(r.Ping(key))
	fmt.Println(r.Ping())

	// Output:
	// [test-ping]
	// PONG
}
