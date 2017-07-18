package redis

import (
	"fmt"
)

func ExampleRedis_HGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hget"
	r.Del(key)

	b, _ := r.HSet(key, "field1", "foo")
	fmt.Println(b)
	v, _ := r.HGet(key, "field1")
	fmt.Println(v)
	v, _ = r.HGet(key, "field2")
	fmt.Println(v)

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

	b, _ := r.HSetNX(key, "field", "Hello")
	fmt.Println(b)
	b, _ = r.HSetNX(key, "field", "World")
	fmt.Println(b)
	v, _ := r.HGet(key, "field")
	fmt.Println(v)

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
	v, _ := r.HDel(key, "field1")
	fmt.Println(v)
	v, _ = r.HDel(key, "field2")
	fmt.Println(v)

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
	v, _ := r.HExists(key, "field1")
	fmt.Println(v)
	v, _ = r.HExists(key, "field2")
	fmt.Println(v)

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
	v, _ := r.HGetAll(key)
	fmt.Println(v)

	// Output:
	// [field1 foo field2 bar]
}

func ExampleRedis_HIncrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hinceby"
	r.Del(key)

	r.HSet(key, "field", 5)
	v, _ := r.HIncrBy(key, "field", 1)
	fmt.Println(v)
	v, _ = r.HIncrBy(key, "field", -1)
	fmt.Println(v)
	v, _ = r.HIncrBy(key, "field", -10)
	fmt.Println(v)

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
	v, _ := r.HIncrByFloat(key, "field", 0.1)
	fmt.Println(v)

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
	v, _ := r.HKeys(key)
	fmt.Println(v)

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
	v, _ := r.HLen(key)
	fmt.Println(v)

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
	v, _ := r.HMGet(key, "field1", "field2", "nofield")
	fmt.Println(v)

	// Output:
	// [Hello World ]
}

func ExampleRedis_HMSet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-hmset"
	r.Del(key)

	r.HMSet(key, "field1", "Hello", "field2", "World")
	v, _ := r.HGet(key, "field1")
	fmt.Println(v)
	v, _ = r.HGet(key, "field2")
	fmt.Println(v)

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
	v, _ := r.HStrLen(key, "f1")
	fmt.Println(v)
	v, _ = r.HStrLen(key, "f2")
	fmt.Println(v)
	v, _ = r.HStrLen(key, "f3")
	fmt.Println(v)

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
	v, _ := r.HVals(key)
	fmt.Println(v)

	// Output:
	// [Hello World]
}
