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