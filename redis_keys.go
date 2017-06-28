package redis

// Keys executes the redis command KEYS。
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Keys(pattern string) []string {
	reply, err := r.Do("KEYS", pattern)
	if err != nil {
		panic(err)
	}
	_results := reply.([]interface{})
	if len(_results) == 0 {
		return nil
	}

	results := make([]string, len(_results))
	for i, _r := range _results {
		results[i] = string(_r.([]byte))
	}

	return results
}

// Del executes the redis command DEL.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Del(key string, keys ...string) int64 {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToInt("DEL", args...)
}

// Exists executes the redis command EXISTS.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Exists(key string, keys ...string) bool {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToBool("EXISTS", args...)
}

// Expire executes the redis command EXPIRE.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Expire(key string, timeout int) bool {
	return r.doToBool("EXPIRE", key, timeout)
}

// PExpire executes the redis command PEXPIRE.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) PExpire(key string, timeout int) bool {
	return r.doToBool("PEXPIRE", key, timeout)
}

// ExpireAt executes the redis command EXPIREAT.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ExpireAt(key string, timestamp int) bool {
	return r.doToBool("EXPIREAT", key, timestamp)
}

// PExpireAt executes the redis command PEXPIREAT.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) PExpireAt(key string, timestamp int) bool {
	return r.doToBool("PEXPIREAT", key, timestamp)
}

// Move executes the redis command MOVE.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Move(key string, db int) bool {
	return r.doToBool("MOVE", key, db)
}

// Persist executes the redis command PERSIST.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) Persist(key string) bool {
	return r.doToBool("PERSIST", key)
}

// TTL executes the redis command TTL.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) TTL(key string) int64 {
	return r.doToInt("TTL", key)
}

// PTTL executes the redis command PTTL.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) PTTL(key string) int64 {
	return r.doToInt("PTTL", key)
}
