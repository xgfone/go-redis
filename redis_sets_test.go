package redis

import (
	"fmt"
)

func ExampleRedis_SAdd() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-sadd"
	r.Del(key)

	r.SAdd(key, "1", "2", "3")
	fmt.Println(r.SMembers(key))

	// Output:
	// [1 2 3]
}

func ExampleRedis_SCard() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-scard"
	r.Del(key)

	r.SAdd(key, "1", "2", "3")
	fmt.Println(r.SCard(key))

	// Output:
	// 3
}

func ExampleRedis_SDiff() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sdiff1"
	key2 := "test-sdiff2"
	r.Del(key1)
	r.Del(key2)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SDiff(key1, key2))

	// Output:
	// [a b]
}

func ExampleRedis_SDiffStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sdiffstore1"
	key2 := "test-sdiffstore2"
	dest := "test-sdiffstore-dest"
	r.Del(key1)
	r.Del(key2)
	r.Del(dest)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SDiffStore(dest, key1, key2))
	fmt.Println(r.SMembers(dest))

	// Output:
	// 2
	// [a b]
}
func ExampleRedis_SInter() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sinter1"
	key2 := "test-sinter2"
	r.Del(key1)
	r.Del(key2)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SInter(key1, key2))

	// Output:
	// [c]
}

func ExampleRedis_SInterStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sinterstore1"
	key2 := "test-sinterstore2"
	dest := "test-sinterstore-dest"
	r.Del(key1)
	r.Del(key2)
	r.Del(dest)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SInterStore(dest, key1, key2))
	fmt.Println(r.SMembers(dest))

	// Output:
	// 1
	// [c]
}

func ExampleRedis_SIsMember() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-sismember"
	r.Del(key)

	r.SAdd(key, "a", "b", "c")
	fmt.Println(r.SIsMember(key, "a"))
	fmt.Println(r.SIsMember(key, "z"))

	// Output:
	// true
	// false
}

func ExampleRedis_SMove() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	src := "test-smove-src"
	dst := "test-smove-dest"
	r.Del(src)
	r.Del(dst)

	r.SAdd(src, "a", "b", "c")
	r.SAdd(dst, "d")
	fmt.Println(r.SMove(src, dst, "c"))
	fmt.Println(r.SMembers(src))
	fmt.Println(r.SMembers(dst))

	// Output:
	// true
	// [b a]
	// [d c]
}

func ExampleRedis_SPop() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-spop"
	r.Del(key)

	r.SAdd(key, "a", "b", "c")
	fmt.Println(len(r.SPop(key)))
	fmt.Println(len(r.SPop(key, 3)))

	// Output:
	// 1
	// 2
}

func ExampleRedis_SRandMember() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-srandmember"
	r.Del(key)

	r.SAdd(key, "a", "b", "c")
	fmt.Println(len(r.SRandMember(key)))
	fmt.Println(len(r.SRandMember(key, 2)))
	fmt.Println(len(r.SRandMember(key, -5)))

	// Output:
	// 1
	// 2
	// 5
}

func ExampleRedis_SRem() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-srem"
	r.Del(key)

	r.SAdd(key, "a", "b", "c", "d")
	fmt.Println(r.SRem(key, "a", "b"))
	fmt.Println(r.SRem(key, "z"))
	fmt.Println(r.SMembers(key))

	// Output:
	// 2
	// 0
	// [d c]
}

func ExampleRedis_SUnion() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sunion1"
	key2 := "test-sunion2"
	r.Del(key1)
	r.Del(key2)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SUnion(key1, key2))

	// Output:
	// [c a b d e]
}

func ExampleRedis_SUnionStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sunionstore1"
	key2 := "test-sunionstore2"
	dest := "test-sunionstore-dest"
	r.Del(key1)
	r.Del(key2)
	r.Del(dest)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	fmt.Println(r.SUnionStore(dest, key1, key2))
	fmt.Println(r.SMembers(dest))

	// Output:
	// 5
	// [c a b d e]
}
