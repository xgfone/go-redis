package redis

// HSet executes the redis command HSET.
//
// New in redis version 2.0.0.
func (r *Redis) HSet(key, field string, value interface{}) (bool, error) {
	return r.doToBool("HSET", key, field, value)
}

// HSetNX executes the redis command HSETNX.
//
// New in redis version 2.0.0.
func (r *Redis) HSetNX(key, field string, value interface{}) (bool, error) {
	return r.doToBool("HSETNX", key, field, value)
}

// HGet executes the redis command HGET.
//
// New in redis version 2.0.0.
func (r *Redis) HGet(key, field string) (string, error) {
	return r.doToString("HGET", key, field)
}

// HDel executes the redis command HDEL.
//
// New in redis version 2.0.0.
func (r *Redis) HDel(key, field string, fields ...string) (int64, error) {
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
// New in redis version 2.0.0.
func (r *Redis) HExists(key, field string) (bool, error) {
	return r.doToBool("HEXISTS", key, field)
}

// HGetAll executes the redis command HGETALL.
//
// New in redis version 2.0.0.
func (r *Redis) HGetAll(key string) ([]string, error) {
	return r.doToStringSlice("HGETALL", key)
}

// HIncrBy executes the redis command HINCRBY.
//
// New in redis version 2.0.0.
func (r *Redis) HIncrBy(key, field string, n int64) (int64, error) {
	return r.doToInt("HINCRBY", key, field, n)
}

// HIncrByFloat executes the redis command HINCRBYFLOAT.
//
// New in redis version 2.6.0.
func (r *Redis) HIncrByFloat(key, field string, n float64) (float64, error) {
	return r.doToFloat("HINCRBYFLOAT", key, field, n)
}

// HKeys executes the redis command HKEYS.
//
// New in redis version 2.0.0.
func (r *Redis) HKeys(key string) ([]string, error) {
	return r.doToStringSlice("HKEYS", key)
}

// HLen executes the redis command HLEN.
//
// New in redis version 2.0.0.
func (r *Redis) HLen(key string) (int64, error) {
	return r.doToInt("HLEN", key)
}

// HMGet executes the redis command HMGet.
//
// New in redis version 2.0.0.
func (r *Redis) HMGet(key, field string, fields ...string) ([]string, error) {
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
// New in redis version 2.0.0.
func (r *Redis) HMSet(key, field string, value interface{}, fields ...interface{}) error {
	_len := len(fields)
	if _len%2 != 0 {
		return ErrInvalidArgs
	}

	args := make([]interface{}, _len+3)
	args[0] = key
	args[1] = field
	args[2] = value
	for i, a := range fields {
		args[i+3] = a
	}

	if _, err := r.Do("HMSET", args...); err != nil {
		return err
	}
	return nil
}

// HStrLen executes the redis command HSTRLEN.
//
// New in redis version 3.2.0.
func (r *Redis) HStrLen(key, field string) (int64, error) {
	return r.doToInt("HSTRLEN", key, field)
}

// HVals executes the redis command HVALS.
//
// New in redis version 2.0.0.
func (r *Redis) HVals(key string) ([]string, error) {
	return r.doToStringSlice("HVALS", key)
}
