package redis

import (
	"fmt"
	"time"
)

func ExampleRedis_Set() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-set-get"
	r.Del(key)

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
	r.Del(key)

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
	r.Del(key)

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
	key1 := "test-bitop1"
	key2 := "test-bitop2"
	r.Del(key)
	r.Del(key1)
	r.Del(key2)

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
	r.Del(key)

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
	r.Del(key)

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
	r.Del(key)

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
	r.Del(key)

	r.Set(key, "10")
	fmt.Println(r.DecrBy(key, 5))

	// Output:
	// 5
}

func ExampleRedis_IncrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-incrby"
	r.Del(key)

	r.Set(key, "10")
	fmt.Println(r.IncrBy(key, 5))

	// Output:
	// 15
}

func ExampleRedis_IncrByFloat() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-incrbyfloat"
	r.Del(key)

	r.Set(key, "10.50")
	fmt.Println(r.IncrByFloat(key, 0.1))

	// Output:
	// 10.6
}

func ExampleRedis_GetBit() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-getbit"
	r.Del(key)

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
	r.Del(key)

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

func ExampleRedis_SetRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-setrange"
	r.Del(key)

	r.Set(key, "Hello World")
	fmt.Println(r.SetRange(key, 6, "Redis"))
	fmt.Println(r.Get(key))

	key2 := "test-setrange2"
	r.Del(key2)

	fmt.Println(r.SetRange(key2, 6, "Redis"))

	v := []byte(r.Get(key2))
	fmt.Println(v[:6])
	fmt.Println(string(v[6:]))

	// Output:
	// 11
	// Hello Redis
	// 11
	// [0 0 0 0 0 0]
	// Redis
}

func ExampleRedis_GetSet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-getset"
	r.Del(key)

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
	r.Del(key1)
	r.Del(key2)

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
	r.Del(key1)
	r.Del(key2)

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
	r.Del(key1)
	r.Del(key2)
	r.Del(key3)

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
	r.Del(key)

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
	r.Del(key)

	r.SetEX(key, 1, "Hello")
	fmt.Println(r.Get(key))
	time.Sleep(1200 * time.Millisecond)
	fmt.Println(r.Get(key))

	// Output:
	// Hello
	//
}

func ExampleRedis_PSetEX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-psetex"
	r.Del(key)

	r.PSetEX(key, 1000, "Hello")
	fmt.Println(r.Get(key))
	time.Sleep(1200 * time.Millisecond)
	fmt.Println(r.Get(key))

	// Output:
	// Hello
	//
}

func ExampleRedis_StrLen() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-strlen"
	r.Del(key)

	r.Set(key, "Hello World")
	fmt.Println(r.StrLen(key))
	fmt.Println(r.StrLen("nonexisting"))

	// Output:
	// 11
	// 0
}
