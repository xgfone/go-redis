package redis

// LPush executes the redis command LPUSH.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) LPush(key string, value string, values ...string) int64 {
	args := make([]interface{}, len(values)+2)
	args[0] = key
	args[1] = value
	for i, v := range values {
		args[i+2] = v
	}
	return r.doToInt("LPUSH", args...)
}

// RPush executes the redis command RPUSH.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) RPush(key string, value string, values ...string) int64 {
	args := make([]interface{}, len(values)+2)
	args[0] = key
	args[1] = value
	for i, v := range values {
		args[i+2] = v
	}
	return r.doToInt("RPUSH", args...)
}

// LPop executes the redis command LPOP.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) LPop(key string) string {
	return r.doToString("LPOP", key)
}

// RPop executes the redis command RPOP.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) RPop(key string) string {
	return r.doToString("RPOP", key)
}
