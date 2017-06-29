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
