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
