package redis

import (
	"fmt"
)

func ExampleRedis_Keys() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-keys"
	r.Del(key)

	r.Set(key, key)
	keys, _ := r.Keys(key + "*")
	fmt.Printf("len=%d, key=%s\n", len(keys), keys[0])
	v, _ := r.Keys("test-not-keys*")
	fmt.Printf("len=%d\n", len(v))

	// Output:
	// len=1, key=test-keys
	// len=0
}

func ExampleRedis_Del() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-del1"
	key2 := "test-del2"
	r.Del(key1)
	r.Del(key2)

	r.Set(key1, key1)
	r.Set(key2, key2)
	v, _ := r.Del(key1, key2)
	fmt.Println(v)

	// Output:
	// 2
}

func ExampleRedis_Exists() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-exists"
	key2 := "test-nonexisting"
	r.Del(key1)
	r.Del(key2)

	r.Set(key1, key1)
	v, _ := r.Exists(key1, key2)
	fmt.Println(v)

	// Output:
	// true
}

func ExampleRedis_Expire() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-expire"
	r.Del(key)

	r.Set(key, key)
	v, _ := r.Expire(key, 1)
	fmt.Println(v)
	v, _ = r.Expire("nonexisting", 1)
	fmt.Println(v)

	// Output:
	// true
	// false
}

func ExampleRedis_PExpire() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pexpire"
	r.Del(key)

	r.Set(key, key)
	v, _ := r.PExpire(key, 1)
	fmt.Println(v)
	v, _ = r.PExpire("nonexisting", 1)
	fmt.Println(v)

	// Output:
	// true
	// false
}

func ExampleRedis_ExpireAt() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-expireat"
	r.Del(key)

	r.Set(key, key)
	v, _ := r.ExpireAt(key, 1293840000)
	fmt.Println(v)
	v, _ = r.ExpireAt("nonexisting", 1)
	fmt.Println(v)

	// Output:
	// true
	// false
}

func ExampleRedis_PExpireAt() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pexpireat"
	r.Del(key)

	r.Set(key, key)
	v, _ := r.PExpireAt(key, 1293840000)
	fmt.Println(v)
	v, _ = r.PExpireAt("nonexisting", 1)
	fmt.Println(v)

	// Output:
	// true
	// false
}

func ExampleRedis_Persist() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-persist"
	r.Del(key)

	r.SetEX(key, 5, key)
	v, _ := r.Persist(key)
	fmt.Println(v)
	v, _ = r.Persist("nonexisting")
	fmt.Println(v)

	// Output:
	// true
	// false
}

func ExampleRedis_TTL() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-ttl"
	r.Del(key)

	r.SetEX(key, 5, key)
	v, _ := r.TTL(key)
	fmt.Println(v)

	// Output:
	// 5
}

func ExampleRedis_PTTL() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pttl"
	r.Del(key)

	r.SetEX(key, 1, key)
	v, _ := r.PTTL(key)
	fmt.Println(v > 990)

	// Output:
	// true
}

func ExampleRedis_Rename() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-rename"
	newKey := key + "-new"
	r.Del(key)
	r.Del(newKey)

	r.Set(key, key)
	r.Rename(key, newKey)
	v, _ := r.Get(newKey)
	fmt.Println(v)

	// Output:
	// test-rename
}

func ExampleRedis_RenameNX() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-renamenx"
	newKey := key + "-new"
	r.Del(key)
	r.Del(newKey)

	r.Set(key, key)
	r.Set(newKey, newKey)
	b, _ := r.RenameNX(key, newKey)
	fmt.Println(b)
	v, _ := r.Get(newKey)
	fmt.Println(v)

	// Output:
	// false
	// test-renamenx-new
}

func ExampleRedis_Type() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-type"
	r.Del(key)

	r.Set(key, key)
	v, _ := r.Type(key)
	fmt.Println(v)

	// Output:
	// string
}
