package redis

const redisTestConnURL = "redis://127.0.0.1:6379/0"

func ExampleRedis_Keys() {
	r := NewRedis(redisTestConnURL, 1)
	defer r.Close()

	keys := r.Keys("test_keys_pattern*")
	if len(keys) != 0 {
		for _, key := range keys {
			println(key)
		}
	}

	// Output:
	//
}

func ExampleRedis_Set() {
	r := NewRedis(redisTestConnURL, 1)
	defer r.Close()

	err := r.Set("test_key", "test_value", "EX", 10)
	if err != nil {
		println(err)
	} else {
		println("OK")
	}

	// Output:
	//
}

func ExampleRedis_Get() {
	r := NewRedis(redisTestConnURL, 1)
	defer r.Close()

	value := r.Get("test_key")
	println(value)

	// Output:
	//
}
