package redis

// HSet executes the redis command HSET.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HSet(key, field string, value interface{}) bool {
	return r.doToBool("HSET", key, field, value)
}

// HGet executes the redis command HGET.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HGet(key, field string) string {
	return r.doToString("HGET", key, field)
}

// HDel executes the redis command HDEL.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HDel(key, field string, fields ...string) int64 {
	args := make([]interface{}, len(fields)+2)
	args[0] = key
	args[1] = field
	for i, v := range fields {
		args[i+2] = v
	}
	return r.doToInt("HDEL", args...)
}

// HExists executes the redis command HEXISTS.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HExists(key, field string) bool {
	return r.doToBool("HEXISTS", key, field)
}

// HGetAll executes the redis command HGETALL.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HGetAll(key string) []string {
	return r.doToStringSlice("HGETALL", key)
}

// HIncrBy executes the redis command HINCRBY.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HIncrBy(key, field string, n int64) int64 {
	return r.doToInt("HINCRBY", key, field, n)
}

// HIncrByFloat executes the redis command HINCRBYFLOAT.
//
// Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) HIncrByFloat(key, field string, n float64) float64 {
	return r.doToFloat("HINCRBYFLOAT", key, field, n)
}

// HKeys executes the redis command HKEYS.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HKeys(key string) []string {
	return r.doToStringSlice("HKEYS", key)
}

// HLen executes the redis command HLEN.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HLen(key string) int64 {
	return r.doToInt("HLEN", key)
}

// HMGet executes the redis command HMGet.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HMGet(key, field string, fields ...string) []string {
	args := make([]interface{}, len(fields)+2)
	args[0] = key
	args[1] = field
	for i, v := range fields {
		args[i+2] = v
	}
	return r.doToStringSlice("HMGET", args...)
}

// HMSet executes the redis command HMSet.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HMSet(key, field string, value interface{}, fields ...interface{}) {
	_len := len(fields)
	if _len%2 != 0 {
		panic(ErrInvalidArgs)
	}

	args := make([]interface{}, _len+3)
	args[0] = key
	args[1] = field
	args[2] = value
	for i, a := range fields {
		args[i+3] = a
	}

	if _, err := r.Do("HMSET", args...); err != nil {
		panic(err)
	}
}
