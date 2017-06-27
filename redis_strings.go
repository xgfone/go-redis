package redis

// Set executes the redis command SET.
func (r *Redis) Set(key string, value interface{}, args ...interface{}) error {
	if len(args) == 0 {
		_, err := r.Do("SET", key, value)
		return err
	}

	_args := make([]interface{}, len(args)+2)
	_args[0] = key
	_args[1] = value
	for i, v := range args {
		_args[i+2] = v
	}
	_, err := r.Do("SET", _args...)
	return err
}

// Get executes the redis command GET.
//
// Return "" if the key does not exist.
func (r *Redis) Get(key string) string {
	if reply, err := r.Do("GET", key); err == nil && reply != nil {
		return string(reply.([]byte))
	}
	return ""
}

// Append executes the redis command APPEND.
func (r *Redis) Append(key, value string) error {
	_, err := r.Do("APPEND", key, value)
	return err
}

// BitCount executes the redis command BITCOUNT.
//
// Return -1 if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) BitCount(key string, args ...int) int64 {
	_len := len(args)
	if _len != 0 && _len != 2 {
		panic("The number of arguments is not right.")
	}

	var reply interface{}
	var err error
	if _len == 0 {
		reply, err = r.Do("BITCOUNT", key)
	} else {
		reply, err = r.Do("BITCOUNT", key, args[0], args[1])
	}
	if err != nil {
		return -1
	}
	return reply.(int64)
}
