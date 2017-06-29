package redis

import "fmt"

func ExampleRedis_ZRange() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrange"
	r.Del(key)
	fmt.Println(r.ZAdd(key, 1, "one", 2, "two", 3, "three"))
	fmt.Println(r.ZRange(key, 0, -1))
	fmt.Println(r.ZRange(key, 0, 1, true))

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
	fmt.Println(r.ZCard(key))

	// Output:
	// 3
}

func ExampleRedis_ZCount() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zcount"
	r.Del(key)
	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	fmt.Println(r.ZCount(key, "-inf", "+inf"))
	fmt.Println(r.ZCount(key, "(1", 3))

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
	fmt.Println(r.ZIncrBy(key, 3, "one"))
	fmt.Println(r.ZRange(key, 0, -1, true))

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

	fmt.Println(r.ZInterStore(key, 2, key1, key2, "WEIGHTS", 2, 3))
	fmt.Println(r.ZRange(key, 0, -1, true))

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

	fmt.Println(r.ZLexCount(key, "-", "+"))
	fmt.Println(r.ZLexCount(key, "[b", "[f"))

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

	fmt.Println(r.ZRangeByLex(key, "-", "(c"))
	fmt.Println(r.ZRangeByLex(key, "[aaa", "(g"))

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
	fmt.Println(r.ZRangeByScore(key, "-inf", "+inf"))
	fmt.Println(r.ZRangeByScore(key, 0, 2))
	fmt.Println(r.ZRangeByScore(key, "(1", 2))
	fmt.Println(r.ZRangeByScore(key, "(1", "(2"))

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
	fmt.Println(r.ZRank(key, "three"))
	fmt.Println(r.ZRank(key, "four"))

	// Output:
	// 2
	// -1
}

func ExampleRedis_ZRem() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-zrem"
	r.Del(key)

	r.ZAdd(key, 1, "one", 2, "two", 3, "three")
	fmt.Println(r.ZRem(key, "one", "two"))
	fmt.Println(r.ZRange(key, 0, -1, true))

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

	fmt.Println(r.ZRange(key, 0, -1))
	fmt.Println(r.ZRemRangeByLex(key, "[b", "[d"))
	fmt.Println(r.ZRange(key, 0, -1))

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
	fmt.Println(r.ZRemRangeByRank(key, 0, 1))
	fmt.Println(r.ZRange(key, 0, -1, true))

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
	fmt.Println(r.ZRemRangeByScore(key, "-inf", "(2"))
	fmt.Println(r.ZRange(key, 0, -1, true))

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
	fmt.Println(r.ZRevRange(key, 0, -1, true))
	fmt.Println(r.ZRevRange(key, 2, 3))

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

	fmt.Println(r.ZRevRangeByLex(key, "[c", "-"))
	fmt.Println(r.ZRevRangeByLex(key, "(c", "-"))
	fmt.Println(r.ZRevRangeByLex(key, "(g", "[aaa"))

	// Output:
	// [c b a]
	// [b a]
	// [f e d c b]
}
