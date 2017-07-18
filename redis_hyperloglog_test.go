package redis

import (
	"fmt"
)

func ExampleRedis_PFAdd() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-pfadd"
	r.Del(key)

	b, _ := r.PFAdd(key, "a", "b", "c", "d", "e", "f", "g")
	fmt.Println(b)
	v, _ := r.PFCount(key)
	fmt.Println(v)

	// Output:
	// true
	// 7
}

func ExampleRedis_PFMerge() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	dst := "test-pfmerge-dst"
	src1 := "test-pfmerge-src1"
	src2 := "test-pfmerge-src2"
	r.Del(dst)
	r.Del(src1)
	r.Del(src2)

	b, _ := r.PFAdd(src1, "foo", "bar", "zap", "a")
	fmt.Println(b)
	b, _ = r.PFAdd(src2, "a", "b", "c", "foo")
	fmt.Println(b)
	r.PFMerge(dst, src1, src2)
	v, _ := r.PFCount(dst)
	fmt.Println(v)

	// Output:
	// true
	// true
	// 6
}
