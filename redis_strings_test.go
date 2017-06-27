package redis

import (
	"fmt"
)

func ExampleRedis_Set() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-set-get"
	err := r.Set(key, key, "EX", 10)
	if err == nil {
		fmt.Println(r.Get(key))
	}

	// Output:
	// test-set-get
	//
}
