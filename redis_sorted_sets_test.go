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
