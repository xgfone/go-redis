package redis

// SAdd executes the redis command SADD.
//
// New in redis version 1.0.0.
func (r *Redis) SAdd(key string, member string, members ...string) (int64, error) {
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
// New in redis version 1.0.0.
func (r *Redis) SMembers(key string) ([]string, error) {
	return r.doToStringSlice("SMEMBERS", key)
}

// SCard executes the redis command SCARD.
//
// New in redis version 1.0.0.
func (r *Redis) SCard(key string) (int64, error) {
	return r.doToInt("SCARD", key)
}

// SDiff executes the redis command SDIFF.
//
// New in redis version 1.0.0.
func (r *Redis) SDiff(key string, keys ...string) ([]string, error) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToStringSlice("SDIFF", args...)
}

// SDiffStore executes the redis command SDIFFSTORE.
//
// New in redis version 1.0.0.
func (r *Redis) SDiffStore(dest, key string, keys ...string) (int64, error) {
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
// New in redis version 1.0.0.
func (r *Redis) SInter(key string, keys ...string) ([]string, error) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}
	return r.doToStringSlice("SINTER", args...)
}

// SInterStore executes the redis command SINTERSTORE.
//
// New in redis version 1.0.0.
func (r *Redis) SInterStore(dest, key string, keys ...string) (int64, error) {
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
// For the returned value, ture is 1 and false is 0.
//
// New in redis version 1.0.0.
func (r *Redis) SIsMember(key, member string) (bool, error) {
	return r.doToBool("SISMEMBER", key, member)
}

// SMove executes the redis command SMOVE.
//
// For the returned value, ture is 1 and false is 0.
//
// New in redis version 1.0.0.
func (r *Redis) SMove(src, dst, member string) (bool, error) {
	return r.doToBool("SMOVE", src, dst, member)
}

// SPop executes the redis command SPOP.
//
// Return nil if the key does not exist.
//
// New in redis version 1.0.0.
// Changed: Adding count from 3.2.
func (r *Redis) SPop(key string, count ...int) ([]string, error) {
	if len(count) == 0 {
		return r.doToStringSlice("SPOP", key)
	}

	if count[0] < 1 {
		return nil, ErrInvalidArgs
	}
	return r.doToStringSlice("SPOP", key, count[0])
}

// SRandMember executes the redis command SRANDMEMBER.
//
// Return nil if the key does not exist.
//
// New in redis version 1.0.0.
// Changed: Adding count from 2.6.
func (r *Redis) SRandMember(key string, count ...int) ([]string, error) {
	if len(count) == 0 {
		v, err := r.doToString("SRANDMEMBER", key)
		if err != nil {
			return nil, err
		}
		if v != "" {
			return []string{v}, nil
		}
		return nil, nil
	}

	return r.doToStringSlice("SRANDMEMBER", key, count[0])
}

// SRem executes the redis command SREM.
//
// New in redis version 1.0.0.
func (r *Redis) SRem(key, member string, members ...string) (int64, error) {
	args := make([]interface{}, len(members)+2)
	args[0] = key
	args[1] = member
	for i, m := range members {
		args[i+2] = m
	}
	return r.doToInt("SREM", args...)
}

// SUnion executes the redis command SUNION.
//
// Return nil if the key does not exist.
//
// New in redis version 1.0.0.
func (r *Redis) SUnion(key string, keys ...string) ([]string, error) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, k := range keys {
		args[i+1] = k
	}

	return r.doToStringSlice("SUNION", args...)
}

// SUnionStore executes the redis command SUNIONSTORE.
//
// New in redis version 1.0.0.
func (r *Redis) SUnionStore(dest, key string, keys ...string) (int64, error) {
	args := make([]interface{}, len(keys)+2)
	args[0] = dest
	args[1] = key
	for i, k := range keys {
		args[i+2] = k
	}
	return r.doToInt("SUNIONSTORE", args...)
}
