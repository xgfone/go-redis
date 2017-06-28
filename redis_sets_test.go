package redis

import (
	"fmt"
)

func ExampleRedis_SAdd() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-sadd"
	r.SAdd(key, "1", "2", "3")
	fmt.Println(r.SMembers(key))

	// Output:
	// [1 2 3]
}

func ExampleRedis_SCard() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-scard"
	r.SAdd(key, "1", "2", "3")
	fmt.Println(r.SCard(key))

	// Output:
	// 3
}

func ExampleRedis_SDiff() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sdiff1"
	key2 := "test-sdiff2"
	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SDiff(key1, key2))

	// Output:
	// [a b]
}
