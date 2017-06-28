package redis

// HSet executes the redis command HSET.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) HSet(key, field, value string) bool {
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
