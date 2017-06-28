package redis

import (
	"fmt"
	"time"
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

func ExampleRedis_BitPos() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-bitpos"
	r.Set(key, "\xff\xf0\x00")
	fmt.Println(r.BitPos(key, false))

	r.Set(key, "\x00\xff\xf0")
	fmt.Println(r.BitPos(key, true, 0))
	fmt.Println(r.BitPos(key, true, 2))

	r.Set(key, "\x00\x00\x00")
	fmt.Println(r.BitPos(key, true))

	// Output:
	// 12
	// 8
	// 16
	// -1
}

func ExampleRedis_Decr() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	key := "test-decr"
	r.Set(key, "10")
	fmt.Println(r.Decr(key))

	r.Set(key, "234293482390480948029348230948")
	fmt.Println(r.Decr(key))

	// Output:
	// 9
	// ERR value is not an integer or out of range
}

func ExampleRedis_Incr() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-incr"
	r.Set(key, "10")
	fmt.Println(r.Incr(key))
	fmt.Println(r.Get(key))

	// Output:
	// 11
	// 11
}

func ExampleRedis_DecrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-decrby"
	r.Set(key, "10")
	fmt.Println(r.DecrBy(key, 5))

	// Output:
	// 5
}

func ExampleRedis_IncrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-incrby"
	r.Set(key, "10")
	fmt.Println(r.IncrBy(key, 5))

	// Output:
	// 15
}

func ExampleRedis_IncrByFloat() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-incrbyfloat"
	r.Set(key, "10.50")
	fmt.Println(r.IncrByFloat(key, 0.1))

	// Output:
	// 10.6
}

func ExampleRedis_GetBit() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-getbit"
	fmt.Println(r.SetBit(key, 7, true))
	fmt.Println(r.GetBit(key, 0))
	fmt.Println(r.GetBit(key, 7))
	fmt.Println(r.GetBit(key, 100))

	// Output:
	// 0
	// 0
	// 1
	// 0
}

func ExampleRedis_GetRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-getrange"
	r.Set(key, "This is a string")
	fmt.Println(r.GetRange(key, 0, 3))
	fmt.Println(r.GetRange(key, -3, -1))
	fmt.Println(r.GetRange(key, 0, -1))
	fmt.Println(r.GetRange(key, 10, 100))

	// Output:
	// This
	// ing
	// This is a string
	// string
}

func ExampleRedis_GetSet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-getset"
	r.Set(key, "1")
	fmt.Println(r.GetSet(key, "2"))
	fmt.Println(r.Get(key))

	// Output:
	// 1
	// 2
}

func ExampleRedis_MGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-mget1"
	key2 := "test-mget2"
	r.Set(key1, "Hello")
	r.Set(key2, "World")
	fmt.Println(r.MGet(key1, key2, "nonexisting"))

	// Output:
	// [Hello World ]
}

func ExampleRedis_MSet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-mset1"
	key2 := "test-mset2"
	r.MSet(key1, "Hello", key2, "World")
	fmt.Println(r.Get(key1))
	fmt.Println(r.Get(key2))

	// Output:
	// Hello
	// World
}

func ExampleRedis_MSetNX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-msetnx1"
	key2 := "test-msetnx2"
	key3 := "test-msetnx3"
	fmt.Println(r.MSetNX(key1, "Hello", key2, "World"))
	fmt.Println(r.MSetNX(key2, "there", key3, "there"))
	fmt.Println(r.Get(key1))
	fmt.Println(r.Get(key2))
	fmt.Println(r.Get(key3))

	// Output:
	// true
	// false
	// Hello
	// World
	//
}

func ExampleRedis_SetNX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-setnx"
	fmt.Println(r.SetNX(key, "Hello"))
	fmt.Println(r.SetNX(key, "World"))
	fmt.Println(r.Get(key))

	// Output:
	// true
	// false
	// Hello
}

func ExampleRedis_SetEX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-setex"
	r.SetEX(key, 2, "Hello")
	fmt.Println(r.Get(key))
	time.Sleep(2 * time.Second)
	fmt.Println(r.Get(key))

	// Output:
	// Hello
	//
}
