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

func ExampleRedis_PExpire() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pexpire"
	r.Set(key, key)
	fmt.Println(r.PExpire(key, 1))
	fmt.Println(r.PExpire("nonexisting", 1))

	// Output:
	// true
	// false
}

func ExampleRedis_ExpireAt() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-expireat"
	r.Set(key, key)
	fmt.Println(r.ExpireAt(key, 1293840000))
	fmt.Println(r.ExpireAt("nonexisting", 1))

	// Output:
	// true
	// false
}

func ExampleRedis_PExpireAt() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pexpireat"
	r.Set(key, key)
	fmt.Println(r.PExpireAt(key, 1293840000))
	fmt.Println(r.PExpireAt("nonexisting", 1))

	// Output:
	// true
	// false
}

func ExampleRedis_Persist() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-persist"
	r.SetEX(key, 5, key)
	fmt.Println(r.Persist(key))
	fmt.Println(r.Persist("nonexisting"))

	// Output:
	// true
	// false
}

func ExampleRedis_TTL() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-ttl"
	r.SetEX(key, 5, key)
	fmt.Println(r.TTL(key))

	// Output:
	// 5
}

func ExampleRedis_PTTL() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pttl"
	r.SetEX(key, 1, key)
	fmt.Println(r.PTTL(key))

	// Output:
	// 1000
}

func ExampleRedis_Rename() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rename"
	newKey := key + "-new"
	r.Set(key, key)
	r.Rename(key, newKey)
	fmt.Println(r.Get(newKey))

	// Output:
	// test-rename
}
