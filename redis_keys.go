package redis

// Keys executes the redis command KEYS。
//
// New in redis version 1.0.0.
func (r *Redis) Keys(pattern string) ([]string, error) {
	reply, err := r.Do("KEYS", pattern)
	if err != nil {
		return nil, err
	}
	_results := reply.([]interface{})
	if len(_results) == 0 {
		return nil, nil
	}

	results := make([]string, len(_results))
	for i, _r := range _results {
		results[i] = string(_r.([]byte))
	}

	return results, nil
}

// Del executes the redis command DEL.
//
// New in redis version 1.0.0.
func (r *Redis) Del(key string, keys ...string) (int64, error) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToInt("DEL", args...)
}

// Exists executes the redis command EXISTS.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 1.0.0.
func (r *Redis) Exists(key string, keys ...string) (bool, error) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToBool("EXISTS", args...)
}

// Expire executes the redis command EXPIRE.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 1.0.0.
func (r *Redis) Expire(key string, timeout int) (bool, error) {
	return r.doToBool("EXPIRE", key, timeout)
}

// PExpire executes the redis command PEXPIRE.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 2.6.0.
func (r *Redis) PExpire(key string, timeout int) (bool, error) {
	return r.doToBool("PEXPIRE", key, timeout)
}

// ExpireAt executes the redis command EXPIREAT.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 1.2.0.
func (r *Redis) ExpireAt(key string, timestamp int) (bool, error) {
	return r.doToBool("EXPIREAT", key, timestamp)
}

// PExpireAt executes the redis command PEXPIREAT.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 2.6.0.
func (r *Redis) PExpireAt(key string, timestamp int) (bool, error) {
	return r.doToBool("PEXPIREAT", key, timestamp)
}

// Move executes the redis command MOVE.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 1.0.0.
func (r *Redis) Move(key string, db int) (bool, error) {
	return r.doToBool("MOVE", key, db)
}

// Persist executes the redis command PERSIST.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 2.2.0.
func (r *Redis) Persist(key string) (bool, error) {
	return r.doToBool("PERSIST", key)
}

// TTL executes the redis command TTL.
//
// New in redis version 1.0.0.
func (r *Redis) TTL(key string) (int64, error) {
	return r.doToInt("TTL", key)
}

// PTTL executes the redis command PTTL.
//
// New in redis version 2.6.0.
func (r *Redis) PTTL(key string) (int64, error) {
	return r.doToInt("PTTL", key)
}

// RandomKey executes the redis command RANDOMKEY.
//
// Return "" if no key exist.
//
// New in redis version 1.0.0.
func (r *Redis) RandomKey() (string, error) {
	return r.doToString("RANDOMKEY")
}

// Rename executes the redis command RENAME.
//
// New in redis version 1.0.0.
func (r *Redis) Rename(oldKey, newKey string) error {
	return r.do("RENAME", oldKey, newKey)
}

// RenameNX executes the redis command RENAMENX.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 1.0.0.
func (r *Redis) RenameNX(oldKey, newKey string) (bool, error) {
	return r.doToBool("RENAMENX", oldKey, newKey)
}

// Type executes the redis command TYPE.
//
// Return "" if the key does not exist.
//
// New in redis version 1.0.0.
func (r *Redis) Type(key string) (string, error) {
	return r.doToString("TYPE", key)
}
