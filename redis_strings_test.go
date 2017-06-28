package redis

import (
	"fmt"
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

func ExampleRedis_DecrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-decrby"
	r.Set(key, "10")
	fmt.Println(r.DecrBy(key, 5))

	// Output:
	// 5
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
