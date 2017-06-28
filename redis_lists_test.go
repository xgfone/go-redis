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

func ExampleRedis_LIndex() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lindex"
	r.RPush(key, "a", "b", "c")
	fmt.Println(r.LIndex(key, 1))
	fmt.Println(r.LIndex(key, -1))
	fmt.Println(r.LIndex(key, 4))

	// Output:
	// b
	// c
	//
}

func ExampleRedis_LInsert() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-linsert"
	r.RPush(key, "Hello", "World")
	fmt.Println(r.LInsert(key, "before", "World", "There"))

	// Output:
	// 3
}

func ExampleRedis_LLen() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-llen"
	r.LPush(key, "Hello", "World")
	fmt.Println(r.LLen(key))

	// Output:
	// 2
}

func ExampleRedis_LPushX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lpushx"
	r.LPush(key, "World")
	fmt.Println(r.LPushX(key, "Hello"))
	fmt.Println(r.LPushX("nonexisting", "Hello"))

	// Output:
	// 2
	// 0
}

func ExampleRedis_LRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lrange"
	r.RPush(key, "one", "two", "three")
	fmt.Println(r.LRange(key, 0, 0))
	fmt.Println(r.LRange(key, -3, 2))
	fmt.Println(r.LRange(key, -100, 100))
	fmt.Println(r.LRange(key, 5, 10))

	// Output:
	// [one]
	// [one two three]
	// [one two three]
	// []
}

func ExampleRedis_LRem() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lrem"
	r.RPush(key, "hello", "hello", "foo", "hello")
	fmt.Println(r.LRem(key, -2, "hello"))
	fmt.Println(r.LRange(key, 0, -1))

	// Output:
	// 2
	// [hello foo]
}

func ExampleRedis_LSet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lset"
	r.RPush(key, "one", "two", "three")
	r.LSet(key, 0, "four")
	r.LSet(key, -2, "five")
	fmt.Println(r.LRange(key, 0, -1))

	// Output:
	// [four five three]
}
