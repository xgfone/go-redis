package redis

// Discard executes the redis command DISCARD.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) Discard() {
	r.do("DISCARD")
}

// Exec executes the redis command EXEC.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) Exec() []interface{} {
	if _r, err := r.Do("EXEC"); err != nil {
		panic(err)
	} else {
		return _r.([]interface{})
	}

}

// Multi executes the redis command MULTI.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) Multi() {
	r.do("MULTI")
}

// Unwatch executes the redis command UNWATCH.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) Unwatch() {
	r.do("UNWATCH")
}

// Watch executes the redis command WATCH.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) Watch(key string, keys ...string) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, v := range keys {
		args[i+1] = v
	}
	r.do("WATCH", args...)
}
