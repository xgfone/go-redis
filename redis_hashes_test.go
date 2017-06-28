package redis

import (
	"fmt"
)

func ExampleRedis_HGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hget"
	r.HSet(key, "field1", "foo")
	fmt.Println(r.HGet(key, "field1"))
	fmt.Println(r.HGet(key, "field2"))

	// Output:
	// foo
	//
}
