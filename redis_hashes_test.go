package redis

import (
	"fmt"
)

func ExampleRedis_HGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hget"
	fmt.Println(r.HSet(key, "field1", "foo"))
	fmt.Println(r.HGet(key, "field1"))
	fmt.Println(r.HGet(key, "field2"))

	// Output:
	// true
	// foo
	//
}

func ExampleRedis_HDel() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hdel"
	r.HSet(key, "field1", "foo")
	fmt.Println(r.HDel(key, "field1"))
	fmt.Println(r.HDel(key, "field2"))

	// Output:
	// 1
	// 0
}
