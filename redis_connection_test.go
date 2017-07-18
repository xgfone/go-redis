package redis

import (
	"fmt"
)

func ExampleRedis_Auth() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.Auth(""); err != nil {
		fmt.Println(err)
	}

	// Output:
	// ERR Client sent AUTH, but no password is set
}

func ExampleRedis_Echo() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-echo"
	v, _ := r.Echo(key)
	fmt.Println(v)

	// Output:
	// test-echo
}

func ExampleRedis_Ping() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-ping"
	v, _ := r.Ping(key)
	fmt.Println(v)
	v, _ = r.Ping()
	fmt.Println(v)

	// Output:
	// [test-ping]
	// PONG
}

func ExampleRedis_Select() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.Select(1); err != nil {
		fmt.Println(err)
	}
	if err := r.Select(100); err != nil {
		fmt.Println(err)
	}

	// Output:
	// ERR invalid DB index
}
