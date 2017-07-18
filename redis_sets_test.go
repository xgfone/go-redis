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
	v, _ := r.SMembers(key)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleRedis_SCard() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-scard"
	r.Del(key)

	r.SAdd(key, "1", "2", "3")
	v, _ := r.SCard(key)
	fmt.Println(v)

	// Output:
	// 3
}

func ExampleRedis_SDiff() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sdiff1"
	key2 := "test-sdiff2"
	if _, err := r.Del(key1, key2); err != nil {
		fmt.Println(err)
	}

	r.SAdd(key1, "a")
	r.SAdd(key1, "b")
	r.SAdd(key1, "c")
	r.SAdd(key2, "c")
	r.SAdd(key2, "d")
	r.SAdd(key2, "e")
	v, _ := r.SDiff(key1, key2)
	fmt.Println(len(v))

	// Output:
	// 2
}

func ExampleRedis_SDiffStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sdiffstore1"
	key2 := "test-sdiffstore2"
	dest := "test-sdiffstore-dest"
	r.Del(key1, key2, dest)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	v, _ := r.SDiffStore(dest, key1, key2)
	fmt.Println(v)
	ss, _ := r.SMembers(dest)
	fmt.Println(len(ss) == 2)

	// Output:
	// 2
	// true
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
	v, _ := r.SInter(key1, key2)
	fmt.Println(v)

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
	v, _ := r.SInterStore(dest, key1, key2)
	fmt.Println(v)
	ss, _ := r.SMembers(dest)
	fmt.Println(ss)

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
	v, _ := r.SIsMember(key, "a")
	fmt.Println(v)
	v, _ = r.SIsMember(key, "z")
	fmt.Println(v)

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
	b, _ := r.SMove(src, dst, "c")
	fmt.Println(b)
	v, _ := r.SMembers(src)
	fmt.Println(v)
	v, _ = r.SMembers(dst)
	fmt.Println(v)

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
	v, _ := r.SPop(key)
	fmt.Println(len(v))
	v, _ = r.SPop(key, 3)
	fmt.Println(len(v))

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
	v, _ := r.SRandMember(key)
	fmt.Println(len(v))
	v, _ = r.SRandMember(key, 2)
	fmt.Println(len(v))
	v, _ = r.SRandMember(key, -5)
	fmt.Println(len(v))

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
	v, _ := r.SRem(key, "a", "b")
	fmt.Println(v)
	v, _ = r.SRem(key, "z")
	fmt.Println(v)
	ss, _ := r.SMembers(key)
	fmt.Println(ss)

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
	v, _ := r.SUnion(key1, key2)
	fmt.Println(len(v) == 5)

	// Output:
	// true
}

func ExampleRedis_SUnionStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key1 := "test-sunionstore1"
	key2 := "test-sunionstore2"
	dest := "test-sunionstore-dest"
	r.Del(key1, key2, dest)

	r.SAdd(key1, "a", "b", "c")
	r.SAdd(key2, "c", "d", "e")
	v, _ := r.SUnionStore(dest, key1, key2)
	fmt.Println(v)
	ss, _ := r.SMembers(dest)
	fmt.Println(len(ss) == 5)

	// Output:
	// 5
	// true
}
