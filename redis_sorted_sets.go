package redis

import (
	"strconv"
)

// ZAdd executes the redis command ZADD.
//
// Return a float64 if giving the option INCR, or a int64.
//
// New in redis version 1.2.0.
// Adding from 3.0.2: XX, NX, CH, INCR.
func (r *Redis) ZAdd(key string, values ...interface{}) (interface{}, error) {
	args := make([]interface{}, len(values)+1)
	args[0] = key
	for i, v := range values {
		args[i+1] = v
	}

	_r, err := r.Do("ZADD", args...)
	if err != nil {
		return nil, err
	}

	if v, ok := _r.(int64); ok {
		return v, nil
	}

	var s string
	switch _r.(type) {
	case []byte:
		s = string(_r.([]byte))
	case string:
		s = _r.(string)
	default:
		return nil, nil
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// ZRange executes the redis command ZRANGE.
//
// New in redis version 1.2.0.
func (r *Redis) ZRange(key string, start, stop int, WITHSCORES ...bool) ([]string, error) {
	if len(WITHSCORES) > 0 && WITHSCORES[0] {
		return r.doToStringSlice("ZRANGE", key, start, stop, "WITHSCORES")
	}
	return r.doToStringSlice("ZRANGE", key, start, stop)
}

// ZCard executes the redis command ZCARD.
//
// New in redis version 1.2.0.
func (r *Redis) ZCard(key string) (int64, error) {
	return r.doToInt("ZCARD", key)
}

// ZCount executes the redis command ZCOUNT.
//
// New in redis version 2.0.0.
func (r *Redis) ZCount(key string, min, max interface{}) (int64, error) {
	return r.doToInt("ZCOUNT", key, min, max)
}

// ZIncrBy executes the redis command ZINCRBY.
//
// New in redis version 1.2.0.
func (r *Redis) ZIncrBy(key string, incr float64, member string) (float64, error) {
	return r.doToFloat("ZINCRBY", key, incr, member)
}

// ZInterStore executes the redis command ZINTERSTORE.
//
// New in redis version 2.0.0.
func (r *Redis) ZInterStore(dstKey string, num int, key string, others ...interface{}) (int64, error) {
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
// New in redis version 2.8.9.
func (r *Redis) ZLexCount(key string, min, max interface{}) (int64, error) {
	return r.doToInt("ZLEXCOUNT", key, min, max)
}

// ZRangeByLex executes the redis command ZRANGEBYLEX.
//
// New in redis version 2.8.9.
func (r *Redis) ZRangeByLex(key string, min, max interface{}, limit ...interface{}) ([]string, error) {
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
// New in redis version 1.0.5.
func (r *Redis) ZRangeByScore(key string, min, max interface{}, others ...interface{}) ([]string, error) {
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
// Return the rank which is 0-based of member if member exists in the sorted set.
//
// New in redis version 2.0.0.
func (r *Redis) ZRank(key, member string) (int64, error) {
	_r, err := r.Do("ZRANK", key, member)
	if err != nil {
		return 0, err
	}
	if _r == nil {
		return 0, ErrNotExist
	}
	return _r.(int64), nil
}

// ZRem executes the redis command ZREM.
//
// New in redis version 1.2.0.
func (r *Redis) ZRem(key, member string, members ...string) (int64, error) {
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
// New in redis version 2.8.9.
func (r *Redis) ZRemRangeByLex(key string, min, max interface{}) (int64, error) {
	return r.doToInt("ZREMRANGEBYLEX", key, min, max)
}

// ZRemRangeByRank executes the redis command ZREMRANGEBYRank.
//
// New in redis version 2.0.0.
func (r *Redis) ZRemRangeByRank(key string, start, stop int) (int64, error) {
	return r.doToInt("ZREMRANGEBYRank", key, start, stop)
}

// ZRemRangeByScore executes the redis command ZREMRANGEBYSCORE.
//
// New in redis version 1.2.0.
func (r *Redis) ZRemRangeByScore(key string, min, max interface{}) (int64, error) {
	return r.doToInt("ZREMRANGEBYSCORE", key, min, max)
}

// ZRevRange executes the redis command ZREVRANGE.
//
// New in redis version 1.2.0.
func (r *Redis) ZRevRange(key string, start, stop int, WITHSCORES ...bool) ([]string, error) {
	if len(WITHSCORES) > 0 && WITHSCORES[0] {
		return r.doToStringSlice("ZREVRANGE", key, start, stop, "WITHSCORES")
	}
	return r.doToStringSlice("ZREVRANGE", key, start, stop)
}

// ZRevRangeByLex executes the redis command ZREVRANGEBYLEX.
//
// New in redis version 2.8.9.
func (r *Redis) ZRevRangeByLex(key string, min, max interface{}, limit ...interface{}) ([]string, error) {
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
// New in redis version 1.0.5.
func (r *Redis) ZRevRangeByScore(key string, min, max interface{}, others ...interface{}) ([]string, error) {
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
// Return the rank which is 0-based of member if member exists in the sorted set.
//
// New in redis version 2.0.0.
func (r *Redis) ZRevRank(key, member string) (int64, error) {
	_r, err := r.Do("ZREVRANK", key, member)
	if err != nil {
		return 0, err
	}
	if _r == nil {
		return 0, ErrNotExist
	}
	return _r.(int64), nil
}

// ZScore executes the redis command ZSCORE.
//
// New in redis version 1.2.0.
func (r *Redis) ZScore(key, member string) (float64, error) {
	return r.doToFloat("ZSCORE", key, member)
}

// ZUnionStore executes the redis command ZUNIONSTORE.
//
// New in redis version 2.0.0.
func (r *Redis) ZUnionStore(dstKey string, num int, key string, others ...interface{}) (int64, error) {
	args := make([]interface{}, len(others)+3)
	args[0] = dstKey
	args[1] = num
	args[2] = key
	for i, v := range others {
		args[i+3] = v
	}

	return r.doToInt("ZUNIONSTORE", args...)
}
