package redis

import (
	"fmt"
)

func ExampleRedis_Set() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-set-get"
	err := r.Set(key, key, "EX", 10)
	if err == nil {
		fmt.Println(r.Get(key))
	}

	// Output:
	// test-set-get
	//
}

func ExampleRedis_Append() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-append"
	r.Set(key, key)
	r.Append(key, "1")
	fmt.Println(r.Get(key))

	// Output:
	// test-append1
}

func ExampleRedis_BitCount() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-bitcount"
	r.Set(key, "foobar")
	fmt.Println(r.BitCount(key))
	fmt.Println(r.BitCount(key, 0, 0))
	fmt.Println(r.BitCount(key, 1, 1))

	// Output:
	// 26
	// 4
	// 6
}
