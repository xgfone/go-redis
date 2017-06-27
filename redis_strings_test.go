package redis

import (
	"fmt"
)

func ExampleRedis_Set() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-set-get"
	r.Set(key, key, "EX", 10)
	fmt.Println(r.Get(key))

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

func ExampleRedis_BitOp() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-bitop"
	r.Set("key1", "foobar")
	r.Set("key2", "abcdef")
	fmt.Println(r.BitOp("AND", key, "key1", "key2"))
	fmt.Println(r.Get(key))

	// Output:
	// 6
	// `bc`ab
}
