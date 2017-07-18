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
	v, _ := r.LPop(key)
	fmt.Println(v)

	// Output:
	// a
}

func ExampleRedis_RPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rpop"
	r.Del(key)

	r.RPush(key, "a", "b", "c")
	v, _ := r.RPop(key)
	fmt.Println(v)

	// Output:
	// c
}

func ExampleRedis_LIndex() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lindex"
	r.Del(key)

	r.RPush(key, "a", "b", "c")
	v, _ := r.LIndex(key, 1)
	fmt.Println(v)
	v, _ = r.LIndex(key, -1)
	fmt.Println(v)
	v, _ = r.LIndex(key, 4)
	fmt.Println(v)

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
	v, _ := r.LInsert(key, "before", "World", "There")
	fmt.Println(v)

	// Output:
	// 3
}

func ExampleRedis_LLen() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-llen"
	r.Del(key)

	r.LPush(key, "Hello", "World")
	v, _ := r.LLen(key)
	fmt.Println(v)

	// Output:
	// 2
}

func ExampleRedis_LPushX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-lpushx"
	r.Del(key)

	r.LPush(key, "World")
	v, _ := r.LPushX(key, "Hello")
	fmt.Println(v)
	v, _ = r.LPushX("nonexisting", "Hello")
	fmt.Println(v)

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
	v, _ := r.RPushX(key, "World")
	fmt.Println(v)
	v, _ = r.RPushX("nonexisting", "World")
	fmt.Println(v)

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
	v, _ := r.RPopLPush(key, key2)
	fmt.Println(v)
	ss, _ := r.LRange(key, 0, -1)
	fmt.Println(ss)
	ss, _ = r.LRange(key2, 0, -1)
	fmt.Println(ss)

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
	v, _ := r.LRange(key, 0, 0)
	fmt.Println(v)
	v, _ = r.LRange(key, -3, 2)
	fmt.Println(v)
	v, _ = r.LRange(key, -100, 100)
	fmt.Println(v)
	v, _ = r.LRange(key, 5, 10)
	fmt.Println(v)

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
	v, _ := r.LRem(key, -2, "hello")
	fmt.Println(v)
	ss, _ := r.LRange(key, 0, -1)
	fmt.Println(ss)

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
	v, _ := r.LRange(key, 0, -1)
	fmt.Println(v)

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
	v, _ := r.LRange(key, 0, -1)
	fmt.Println(v)

	// Output:
	// [two three]
}

func ExampleRedis_BLPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-blpop"
	r.Del(key)

	r.RPush(key, "one", "two", "three")
	v, _ := r.BLPop(key, 0)
	fmt.Println(v)

	// Output:
	// [test-blpop one]
}

func ExampleRedis_BRPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-brpop"
	r.Del(key)

	r.RPush(key, "one", "two", "three")
	v, _ := r.BRPop(key, 0)
	fmt.Println(v)

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
	v, _ := r.BRPopLPush(key, key2, 0)
	fmt.Println(v)
	ss, _ := r.LRange(key, 0, -1)
	fmt.Println(ss)
	ss, _ = r.LRange(key2, 0, -1)
	fmt.Println(ss)

	// Output:
	// three
	// [one two]
	// [three]
}
