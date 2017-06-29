package redis

import (
	"fmt"
)

func ExampleRedis_HGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hget"
	r.Del(key)

	fmt.Println(r.HSet(key, "field1", "foo"))
	fmt.Println(r.HGet(key, "field1"))
	fmt.Println(r.HGet(key, "field2"))

	// Output:
	// true
	// foo
	//
}

func ExampleRedis_HSetNX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hsetnx"
	r.Del(key)

	fmt.Println(r.HSetNX(key, "field", "Hello"))
	fmt.Println(r.HSetNX(key, "field", "World"))
	fmt.Println(r.HGet(key, "field"))

	// Output:
	// true
	// false
	// Hello
}

func ExampleRedis_HDel() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hdel"
	r.Del(key)

	r.HSet(key, "field1", "foo")
	fmt.Println(r.HDel(key, "field1"))
	fmt.Println(r.HDel(key, "field2"))

	// Output:
	// 1
	// 0
}

func ExampleRedis_HExists() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hexists"
	r.Del(key)

	r.HSet(key, "field1", "foo")
	fmt.Println(r.HExists(key, "field1"))
	fmt.Println(r.HExists(key, "field2"))

	// Output:
	// true
	// false
}

func ExampleRedis_HGetAll() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hgetall"
	r.Del(key)

	r.HSet(key, "field1", "foo")
	r.HSet(key, "field2", "bar")
	fmt.Println(r.HGetAll(key))

	// Output:
	// [field1 foo field2 bar]
}

func ExampleRedis_HIncrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hinceby"
	r.Del(key)

	r.HSet(key, "field", 5)
	fmt.Println(r.HIncrBy(key, "field", 1))
	fmt.Println(r.HIncrBy(key, "field", -1))
	fmt.Println(r.HIncrBy(key, "field", -10))

	// Output:
	// 6
	// 5
	// -5
}

func ExampleRedis_HIncrByFloat() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hincebyfloat"
	r.Del(key)

	r.HSet(key, "field", 10.50)
	fmt.Println(r.HIncrByFloat(key, "field", 0.1))

	// Output:
	// 10.6
}

func ExampleRedis_HKeys() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hkeys"
	r.Del(key)

	r.HSet(key, "field1", "Hello")
	r.HSet(key, "field2", "World")
	fmt.Println(r.HKeys(key))

	// Output:
	// [field1 field2]
}

func ExampleRedis_HLen() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hlen"
	r.Del(key)

	r.HSet(key, "field1", "Hello")
	r.HSet(key, "field2", "World")
	fmt.Println(r.HLen(key))

	// Output:
	// 2
}

func ExampleRedis_HMGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hmget"
	r.Del(key)

	r.HSet(key, "field1", "Hello")
	r.HSet(key, "field2", "World")
	fmt.Println(r.HMGet(key, "field1", "field2", "nofield"))

	// Output:
	// [Hello World ]
}

func ExampleRedis_HMSet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hmset"
	r.Del(key)

	r.HMSet(key, "field1", "Hello", "field2", "World")
	fmt.Println(r.HGet(key, "field1"))
	fmt.Println(r.HGet(key, "field2"))

	// Output:
	// Hello
	// World
}

func ExampleRedis_HStrLen() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hstrlen"
	r.Del(key)

	r.HMSet(key, "f1", "HelloWorld", "f2", 99, "f3", -256)
	fmt.Println(r.HStrLen(key, "f1"))
	fmt.Println(r.HStrLen(key, "f2"))
	fmt.Println(r.HStrLen(key, "f3"))

	// Output:
	// 10
	// 2
	// 4
}

func ExampleRedis_HVals() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hvals"
	r.Del(key)

	r.HMSet(key, "f1", "Hello", "f2", "World")
	fmt.Println(r.HVals(key))

	// Output:
	// [Hello World]
}
