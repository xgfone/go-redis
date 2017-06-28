package redis

import (
	"fmt"
)

func ExampleRedis_Keys() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-keys"
	r.Set(key, key)
	keys := r.Keys(key + "*")
	fmt.Printf("len=%d, key=%s\n", len(keys), keys[0])
	fmt.Printf("len=%d\n", len(r.Keys("test-not-keys*")))

	// Output:
	// len=1, key=test-keys
	// len=0
}

func ExampleRedis_Del() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-del1"
	key2 := "test-del2"
	r.Set(key1, key1)
	r.Set(key2, key2)
	fmt.Println(r.Del(key1, key2))

	// Output:
	// 2
}

func ExampleRedis_Exists() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-exists"
	key2 := "test-nonexisting"
	r.Set(key1, key1)
	fmt.Println(r.Exists(key1, key2))

	// Output:
	// true
}

func ExampleRedis_Expire() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-expire"
	r.Set(key, key)
	fmt.Println(r.Expire(key, 1))
	fmt.Println(r.Expire("nonexisting", 1))

	// Output:
	// true
	// false
}
