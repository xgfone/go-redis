package redis

// Discard executes the redis command DISCARD.
//
// New in redis version 2.0.0.
func (r *Redis) Discard() error {
	return r.do("DISCARD")
}

// Exec executes the redis command EXEC.
//
// New in redis version 1.2.0.
func (r *Redis) Exec() ([]interface{}, error) {
	_r, err := r.Do("EXEC")
	if err != nil {
		return nil, err
	}
	return _r.([]interface{}), nil
}

// Multi executes the redis command MULTI.
//
// New in redis version 1.2.0.
func (r *Redis) Multi() error {
	return r.do("MULTI")
}

// Unwatch executes the redis command UNWATCH.
//
// New in redis version 2.2.0.
func (r *Redis) Unwatch() error {
	return r.do("UNWATCH")
}

// Watch executes the redis command WATCH.
//
// New in redis version 2.2.0.
func (r *Redis) Watch(key string, keys ...string) error {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, v := range keys {
		args[i+1] = v
	}
	return r.do("WATCH", args...)
}
