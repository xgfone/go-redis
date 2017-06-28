package redis

// SAdd executes the redis command SADD.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SAdd(key string, member string, members ...string) int64 {
	args := make([]interface{}, len(members)+2)
	args[0] = key
	args[1] = member
	for i, m := range members {
		args[i+2] = m
	}
	return r.doToInt("SADD", args...)
}

// SMembers executes the redis command SMEMBERS.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SMembers(key string) []string {
	return r.doToStringSlice("SMEMBERS", key)
}

// SCard executes the redis command SCARD.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SCard(key string) int64 {
	return r.doToInt("SCARD", key)
}

// SDiff executes the redis command SDIFF.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SDiff(key string, keys ...string) []string {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToStringSlice("SDIFF", args...)
}
