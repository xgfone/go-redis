package redis

import (
	"fmt"
)

func ExampleRedis_Exec() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-exec"
	r.Multi()
	r.Set(key, "123")
	r.Get(key)
	rs := r.Exec()
	fmt.Println(rs[0])
	fmt.Println(string(rs[1].([]byte)))

	// Output:
	// OK
	// 123
}
