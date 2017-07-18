package redis

import "fmt"

func ExampleRedis_ZRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrange"
	r.Del(key)
	v, _ := r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	fmt.Println(v)

	ss, _ := r.ZRange(key, 0, -1)
	fmt.Println(ss)
	ss, _ = r.ZRange(key, 0, 1, true)
	fmt.Println(ss)

	// Output:
	// 3
	// [one two three]
	// [one 1 two 2]
}

func ExampleRedis_ZCard() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zcard"
	r.Del(key)
	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZCard(key)
	fmt.Println(v)

	// Output:
	// 3
}

func ExampleRedis_ZCount() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zcount"
	r.Del(key)
	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZCount(key, "-inf", "+inf")
	fmt.Println(v)
	v, _ = r.ZCount(key, "(1", 3)
	fmt.Println(v)

	// Output:
	// 3
	// 2
}

func ExampleRedis_ZIncrBy() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zincrby"
	r.Del(key)
	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZIncrBy(key, 3, "one")
	fmt.Println(v)
	ss, _ := r.ZRange(key, 0, -1, true)
	fmt.Println(ss)

	// Output:
	// 4
	// [two 2 three 3 one 4]
}

func ExampleRedis_ZInterStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zinterstore"
	key1 := "test-zinterstore1"
	key2 := "test-zinterstore2"
	r.Del(key)
	r.Del(key1)
	r.Del(key2)

	r.ZAdd(key1, 1, "one", 2, "two")
	r.ZAdd(key2, 1, "one", 2, "two", 3, "three")

	v, _ := r.ZInterStore(key, 2, key1, key2, "WEIGHTS", 2, 3)
	fmt.Println(v)
	ss, _ := r.ZRange(key, 0, -1, true)
	fmt.Println(ss)

	// Output:
	// 2
	// [one 5 two 10]
}

func ExampleRedis_ZLexCount() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zlexcount"
	r.Del(key)

	r.ZAdd(key, 0, "a", 0, "b", 0, "c", 0, "d", 0, "e")
	r.ZAdd(key, 0, "f", 0, "g")

	v, _ := r.ZLexCount(key, "-", "+")
	fmt.Println(v)
	v, _ = r.ZLexCount(key, "[b", "[f")
	fmt.Println(v)

	// Output:
	// 7
	// 5
}

func ExampleRedis_ZRangeByLex() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrangebylex"
	r.Del(key)

	r.ZAdd(key, 0, "a", 0, "b", 0, "c", 0, "d", 0, "e")
	r.ZAdd(key, 0, "f", 0, "g")

	v, _ := r.ZRangeByLex(key, "-", "(c")
	fmt.Println(v)
	v, _ = r.ZRangeByLex(key, "[aaa", "(g")
	fmt.Println(v)

	// Output:
	// [a b]
	// [b c d e f]
}

func ExampleRedis_ZRangeByScore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrangebyscore"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRangeByScore(key, "-inf", "+inf")
	fmt.Println(v)
	v, _ = r.ZRangeByScore(key, 0, 2)
	fmt.Println(v)
	v, _ = r.ZRangeByScore(key, "(1", 2)
	fmt.Println(v)
	v, _ = r.ZRangeByScore(key, "(1", "(2")
	fmt.Println(v)

	// Output:
	// [one two three]
	// [one two]
	// [two]
	// []
}

func ExampleRedis_ZRank() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrank"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRank(key, "three")
	fmt.Println(v)
	v, _ = r.ZRank(key, "four")
	fmt.Println(v)

	// Output:
	// 2
	// 0
}

func ExampleRedis_ZRem() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrem"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRem(key, "one", "two")
	fmt.Println(v)
	ss, _ := r.ZRange(key, 0, -1, true)
	fmt.Println(ss)

	// Output:
	// 2
	// [three 3]
}

func ExampleRedis_ZRemRangeByLex() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zremrangebylex"
	r.Del(key)

	r.ZAdd(key, 0, "a", 0, "b", 0, "c", 0, "d", 0, "e")
	r.ZAdd(key, 0, "f", 0, "g")

	ss, _ := r.ZRange(key, 0, -1)
	fmt.Println(ss)
	v, _ := r.ZRemRangeByLex(key, "[b", "[d")
	fmt.Println(v)
	ss, _ = r.ZRange(key, 0, -1)
	fmt.Println(ss)

	// Output:
	// [a b c d e f g]
	// 3
	// [a e f g]
}

func ExampleRedis_ZRemRangeByRank() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zremrangebyrank"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRemRangeByRank(key, 0, 1)
	fmt.Println(v)
	ss, _ := r.ZRange(key, 0, -1, true)
	fmt.Println(ss)

	// Output:
	// 2
	// [three 3]
}

func ExampleRedis_ZRemRangeByScore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zremrangebyscore"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRemRangeByScore(key, "-inf", "(2")
	fmt.Println(v)
	ss, _ := r.ZRange(key, 0, -1, true)
	fmt.Println(ss)

	// Output:
	// 1
	// [two 2 three 3]
}

func ExampleRedis_ZRevRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrevrange"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRevRange(key, 0, -1, true)
	fmt.Println(v)
	v, _ = r.ZRevRange(key, 2, 3)
	fmt.Println(v)

	// Output:
	// [three 3 two 2 one 1]
	// [one]
}

func ExampleRedis_ZRevRangeByLex() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrevrangebylex"
	r.Del(key)

	r.ZAdd(key, 0, "a", 0, "b", 0, "c", 0, "d", 0, "e")
	r.ZAdd(key, 0, "f", 0, "g")

	v, _ := r.ZRevRangeByLex(key, "[c", "-")
	fmt.Println(v)
	v, _ = r.ZRevRangeByLex(key, "(c", "-")
	fmt.Println(v)
	v, _ = r.ZRevRangeByLex(key, "(g", "[aaa")
	fmt.Println(v)

	// Output:
	// [c b a]
	// [b a]
	// [f e d c b]
}

func ExampleRedis_ZRevRangeByScore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrevrangebyscore"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRevRangeByScore(key, "+inf", "-inf")
	fmt.Println(v)
	v, _ = r.ZRevRangeByScore(key, 2, 1)
	fmt.Println(v)
	v, _ = r.ZRevRangeByScore(key, 2, "(1")
	fmt.Println(v)
	v, _ = r.ZRevRangeByScore(key, "(2", "(1")
	fmt.Println(v)

	// Output:
	// [three two one]
	// [two one]
	// [two]
	// []
}

func ExampleRedis_ZRevRank() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrevrank"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZRevRank(key, "one")
	fmt.Println(v)
	v, _ = r.ZRevRank(key, "four")
	fmt.Println(v)

	// Output:
	// 2
	// 0
}

func ExampleRedis_ZScore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zscore"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	v, _ := r.ZScore(key, "one")
	fmt.Println(v)

	// Output:
	// 1
}

func ExampleRedis_ZUnionStore() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zunionstore"
	key1 := "test-zunionstore1"
	key2 := "test-zunionstore2"
	r.Del(key)
	r.Del(key1)
	r.Del(key2)

	r.ZAdd(key1, 1, "one", 2, "two")
	r.ZAdd(key2, 1, "one", 2, "two", 3, "three")

	v, _ := r.ZUnionStore(key, 2, key1, key2, "WEIGHTS", 2, 3)
	fmt.Println(v)
	ss, _ := r.ZRange(key, 0, -1, true)
	fmt.Println(ss)

	// Output:
	// 3
	// [one 5 three 9 two 10]
}
