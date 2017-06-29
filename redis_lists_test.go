package redis

import (
	"fmt"
)

func ExampleRedis_LPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lpop"
	r.Del(key)

	r.RPush(key, "a", "b", "c")
	fmt.Println(r.LPop(key))

	// Output:
	// a
}

func ExampleRedis_RPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rpop"
	r.Del(key)

	r.RPush(key, "a", "b", "c")
	fmt.Println(r.RPop(key))

	// Output:
	// c
}

func ExampleRedis_LIndex() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lindex"
	r.Del(key)

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
	r.Del(key)

	r.RPush(key, "Hello", "World")
	fmt.Println(r.LInsert(key, "before", "World", "There"))

	// Output:
	// 3
}

func ExampleRedis_LLen() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-llen"
	r.Del(key)

	r.LPush(key, "Hello", "World")
	fmt.Println(r.LLen(key))

	// Output:
	// 2
}

func ExampleRedis_LPushX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lpushx"
	r.Del(key)

	r.LPush(key, "World")
	fmt.Println(r.LPushX(key, "Hello"))
	fmt.Println(r.LPushX("nonexisting", "Hello"))

	// Output:
	// 2
	// 0
}

func ExampleRedis_RPushX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rpushx"
	r.Del(key)

	r.RPush(key, "Hello")
	fmt.Println(r.RPushX(key, "World"))
	fmt.Println(r.RPushX("nonexisting", "World"))

	// Output:
	// 2
	// 0
}

func ExampleRedis_RPopLPush() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rpoplpush"
	key2 := "test-rpoplpush2"
	r.Del(key)
	r.Del(key2)

	r.RPush(key, "one", "two", "three")
	fmt.Println(r.RPopLPush(key, key2))
	fmt.Println(r.LRange(key, 0, -1))
	fmt.Println(r.LRange(key2, 0, -1))

	// Output:
	// three
	// [one two]
	// [three]
}

func ExampleRedis_LRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lrange"
	r.Del(key)

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
	r.Del(key)

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
	r.Del(key)

	r.RPush(key, "one", "two", "three")
	r.LSet(key, 0, "four")
	r.LSet(key, -2, "five")
	fmt.Println(r.LRange(key, 0, -1))

	// Output:
	// [four five three]
}

func ExampleRedis_LTrim() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-ltrim"
	r.Del(key)

	r.RPush(key, "one", "two", "three")
	r.LTrim(key, 1, -1)
	fmt.Println(r.LRange(key, 0, -1))

	// Output:
	// [two three]
}

func ExampleRedis_BLPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-blpop"
	r.Del(key)

	r.RPush(key, "one", "two", "three")
	fmt.Println(r.BLPop(key, 0))

	// Output:
	// [test-blpop one]
}

func ExampleRedis_BRPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-brpop"
	r.Del(key)

	r.RPush(key, "one", "two", "three")
	fmt.Println(r.BRPop(key, 0))

	// Output:
	// [test-brpop three]
}

func ExampleRedis_BRPopLPush() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-brpoplpush"
	key2 := "test-brpoplpush2"
	r.Del(key)
	r.Del(key2)

	r.RPush(key, "one", "two", "three")
	fmt.Println(r.BRPopLPush(key, key2, 0))
	fmt.Println(r.LRange(key, 0, -1))
	fmt.Println(r.LRange(key2, 0, -1))

	// Output:
	// three
	// [one two]
	// [three]
}
