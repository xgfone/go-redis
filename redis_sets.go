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

// SDiffStore executes the redis command SDIFFSTORE.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SDiffStore(dest, key string, keys ...string) int64 {
	args := make([]interface{}, len(keys)+2)
	args[0] = dest
	args[1] = key
	for i, k := range keys {
		args[i+2] = k
	}
	return r.doToInt("SDIFFSTORE", args...)
}

// SInter executes the redis command SINTER.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SInter(key string, keys ...string) []string {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToStringSlice("SINTER", args...)
}

// SInterStore executes the redis command SINTERSTORE.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SInterStore(dest, key string, keys ...string) int64 {
	args := make([]interface{}, len(keys)+2)
	args[0] = dest
	args[1] = key
	for i, k := range keys {
		args[i+2] = k
	}
	return r.doToInt("SINTERSTORE", args...)
}

// SIsMember executes the redis command SISMEMBER.
//
// For the returned value, ture is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SIsMember(key, member string) bool {
	return r.doToBool("SISMEMBER", key, member)
}
