package redis

import (
	"strconv"
)

// ZAdd executes the redis command ZADD.
//
// Return a float64 if giving the option INCR, or a int64.
// Panic if an error occurs.
//
// New in redis version 1.2.0.
// Adding from 3.0.2: XX, NX, CH, INCR.
func (r *Redis) ZAdd(key string, values ...interface{}) interface{} {
	args := make([]interface{}, len(values)+1)
	args[0] = key
	for i, v := range values {
		args[i+1] = v
	}

	if _r, err := r.Do("ZADD", args...); err != nil {
		panic(err)
	} else {
		if v, ok := _r.(int64); ok {
			return v
		}

		var s string
		switch _r.(type) {
		case []byte:
			s = string(_r.([]byte))
		case string:
			s = _r.(string)
		default:
			return nil
		}
		if v, err := strconv.ParseFloat(s, 64); err != nil {
			panic(err)
		} else {
			return v
		}
	}
}

// ZRange executes the redis command ZRANGE.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZRange(key string, start, stop int, WITHSCORES ...bool) []string {
	if len(WITHSCORES) > 0 && WITHSCORES[0] {
		return r.doToStringSlice("ZRANGE", key, start, stop, "WITHSCORES")
	}
	return r.doToStringSlice("ZRANGE", key, start, stop)
}

// ZCard executes the redis command ZCARD.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZCard(key string) int64 {
	return r.doToInt("ZCARD", key)
}

// ZCount executes the redis command ZCOUNT.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ZCount(key string, min, max interface{}) int64 {
	return r.doToInt("ZCOUNT", key, min, max)
}

// ZIncrBy executes the redis command ZINCRBY.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZIncrBy(key string, incr float64, member string) float64 {
	return r.doToFloat("ZINCRBY", key, incr, member)
}

// ZInterStore executes the redis command ZINTERSTORE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ZInterStore(dstKey string, num int, key string, others ...interface{}) int64 {
	args := make([]interface{}, len(others)+3)
	args[0] = dstKey
	args[1] = num
	args[2] = key
	for i, v := range others {
		args[i+3] = v
	}

	return r.doToInt("ZINTERSTORE", args...)
}

// ZLexCount executes the redis command ZLEXCOUNT.
//
// Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) ZLexCount(key string, min, max interface{}) int64 {
	return r.doToInt("ZLEXCOUNT", key, min, max)
}

// ZRangeByLex executes the redis command ZRANGEBYLEX.
//
// Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) ZRangeByLex(key string, min, max interface{}, limit ...interface{}) []string {
	args := make([]interface{}, len(limit)+3)
	args[0] = key
	args[1] = min
	args[2] = max
	for i, v := range limit {
		args[i+3] = v
	}
	return r.doToStringSlice("ZRANGEBYLEX", args...)
}

// ZRangeByScore executes the redis command ZRANGEBYSCORE.
//
// Panic if an error occurs.
//
// New in redis version 1.0.5.
func (r *Redis) ZRangeByScore(key string, min, max interface{}, others ...interface{}) []string {
	args := make([]interface{}, len(others)+3)
	args[0] = key
	args[1] = min
	args[2] = max
	for i, v := range others {
		args[i+3] = v
	}
	return r.doToStringSlice("ZRANGEBYSCORE", args...)
}

// ZRank executes the redis command ZRANK.
//
// Return -1 if member does not exist in the sorted set or key does not exist.
// Return the rank which is 0-based of member if member exists in the sorted set.
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ZRank(key, member string) int64 {
	if _r, err := r.Do("ZRANK", key, member); err != nil {
		panic(err)
	} else {
		if _r == nil {
			return -1
		}
		return _r.(int64)
	}
}

// ZRem executes the redis command ZREM.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZRem(key, member string, members ...string) int64 {
	args := make([]interface{}, len(members)+2)
	args[0] = key
	args[1] = member
	for i, v := range members {
		args[i+2] = v
	}

	return r.doToInt("ZREM", args...)
}

// ZRemRangeByLex executes the redis command ZREMRANGEBYLEX.
//
// Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) ZRemRangeByLex(key string, min, max interface{}) int64 {
	return r.doToInt("ZREMRANGEBYLEX", key, min, max)
}

// ZRemRangeByRank executes the redis command ZREMRANGEBYRank.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ZRemRangeByRank(key string, start, stop int) int64 {
	return r.doToInt("ZREMRANGEBYRank", key, start, stop)
}

// ZRemRangeByScore executes the redis command ZREMRANGEBYSCORE.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZRemRangeByScore(key string, min, max interface{}) int64 {
	return r.doToInt("ZREMRANGEBYSCORE", key, min, max)
}

// ZRevRange executes the redis command ZREVRANGE.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZRevRange(key string, start, stop int, WITHSCORES ...bool) []string {
	if len(WITHSCORES) > 0 && WITHSCORES[0] {
		return r.doToStringSlice("ZREVRANGE", key, start, stop, "WITHSCORES")
	}
	return r.doToStringSlice("ZREVRANGE", key, start, stop)
}

// ZRevRangeByLex executes the redis command ZREVRANGEBYLEX.
//
// Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) ZRevRangeByLex(key string, min, max interface{}, limit ...interface{}) []string {
	args := make([]interface{}, len(limit)+3)
	args[0] = key
	args[1] = min
	args[2] = max
	for i, v := range limit {
		args[i+3] = v
	}
	return r.doToStringSlice("ZREVRANGEBYLEX", args...)
}

// ZRevRangeByScore executes the redis command ZREVRANGEBYSCORE.
//
// Panic if an error occurs.
//
// New in redis version 1.0.5.
func (r *Redis) ZRevRangeByScore(key string, min, max interface{}, others ...interface{}) []string {
	args := make([]interface{}, len(others)+3)
	args[0] = key
	args[1] = min
	args[2] = max
	for i, v := range others {
		args[i+3] = v
	}
	return r.doToStringSlice("ZREVRANGEBYSCORE", args...)
}

// ZRevRank executes the redis command ZREVRANK.
//
// Return -1 if member does not exist in the sorted set or key does not exist.
// Return the rank which is 0-based of member if member exists in the sorted set.
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ZRevRank(key, member string) int64 {
	if _r, err := r.Do("ZREVRANK", key, member); err != nil {
		panic(err)
	} else {
		if _r == nil {
			return -1
		}
		return _r.(int64)
	}
}

// ZScore executes the redis command ZSCORE.
//
// Panic if an error occurs.
//
// New in redis version 1.2.0.
func (r *Redis) ZScore(key, member string) float64 {
	return r.doToFloat("ZSCORE", key, member)
}
