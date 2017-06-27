package redis

import (
	"fmt"
)

func ExampleRedis_Keys() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-keys"
	r.Set(key, key)
	keys := r.Keys(key + "*")
	fmt.Printf("len=%d, key=%s\n", len(keys), keys[0])
	fmt.Printf("len=%d\n", len(r.Keys("test-not-keys*")))

	// Output:
	// len=1, key=test-keys
	// len=0
}
