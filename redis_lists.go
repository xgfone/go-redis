package redis

import (
	"strings"
)

// LPush executes the redis command LPUSH.
//
// New in redis version 1.0.0.
func (r *Redis) LPush(key string, value string, values ...string) (int64, error) {
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
// New in redis version 1.0.0.
func (r *Redis) RPush(key string, value string, values ...string) (int64, error) {
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
// New in redis version 1.0.0.
func (r *Redis) LPop(key string) (string, error) {
	return r.doToString("LPOP", key)
}

// RPop executes the redis command RPOP.
//
// New in redis version 1.0.0.
func (r *Redis) RPop(key string) (string, error) {
	return r.doToString("RPOP", key)
}

// LIndex executes the redis command LINDEX.
//
// New in redis version 1.0.0.
func (r *Redis) LIndex(key string, index int) (string, error) {
	return r.doToString("LINDEX", key, index)
}

// LInsert executes the redis command LINSERT.
//
// New in redis version 2.2.0.
func (r *Redis) LInsert(key, ba, pivot, value string) (int64, error) {
	ba = strings.ToUpper(ba)
	if ba != "BEFORE" && ba != "AFTER" {
		return 0, ErrInvalidArgs
	}

	return r.doToInt("LINSERT", key, ba, pivot, value)
}

// LLen executes the redis command LLEN.
//
// New in redis version 1.0.0.
func (r *Redis) LLen(key string) (int64, error) {
	return r.doToInt("LLEN", key)
}

// LPushX executes the redis command LPUSHX.
//
// New in redis version 2.2.0.
func (r *Redis) LPushX(key, value string) (int64, error) {
	return r.doToInt("LPUSHX", key, value)
}

// RPushX executes the redis command RPUSHX.
//
// New in redis version 2.2.0.
func (r *Redis) RPushX(key, value string) (int64, error) {
	return r.doToInt("RPUSHX", key, value)
}

// RPopLPush executes the redis command RPOPLPUSH.
//
// New in redis version 1.2.0.
func (r *Redis) RPopLPush(src, dst string) (string, error) {
	return r.doToString("RPOPLPUSH", src, dst)
}

// LRange executes the redis command LRANGE.
//
// New in redis version 1.0.0.
func (r *Redis) LRange(key string, start, stop int) ([]string, error) {
	return r.doToStringSlice("LRANGE", key, start, stop)
}

// LRem executes the redis command LREM.
//
// New in redis version 1.0.0.
func (r *Redis) LRem(key string, count int, value string) (int64, error) {
	return r.doToInt("LREM", key, count, value)
}

// LSet executes the redis command LSET.
//
// New in redis version 1.0.0.
func (r *Redis) LSet(key string, index int, value string) error {
	return r.do("LSET", key, index, value)
}

// LTrim executes the redis command LTRIM.
//
// New in redis version 1.0.0.
func (r *Redis) LTrim(key string, start, stop int) error {
	return r.do("LTRIM", key, start, stop)
}

func (r *Redis) bpop(cmd, key string, keys ...interface{}) ([]string, error) {
	_len := len(keys)
	if _len < 1 {
		return nil, ErrInvalidArgs
	}

	if _, ok := keys[_len-1].(int); !ok {
		return nil, ErrInvalidArgs
	}

	args := make([]interface{}, _len+1)
	args[0] = key
	for i, v := range keys {
		args[i+1] = v
	}
	return r.doToStringSlice(cmd, args...)
}

// BLPop executes the redis command BLPOP.
//
// Notice: The argument keys has one element, which is the timeout,
// and the type of which must be int. If keys has more then one element,
// the timeout is the last.
//
// New in redis version 2.0.0.
func (r *Redis) BLPop(key string, keys ...interface{}) ([]string, error) {
	return r.bpop("BLPOP", key, keys...)
}

// BRPop executes the redis command BRPOP.
//
// Notice: The argument keys has one element, which is the timeout,
// and the type of which must be int. If keys has more then one element,
// the timeout is the last.
//
// New in redis version 2.0.0.
func (r *Redis) BRPop(key string, keys ...interface{}) ([]string, error) {
	return r.bpop("BRPOP", key, keys...)
}

// BRPopLPush executes the redis command BRPOPLPUSH.
//
// New in redis version 2.2.0.
func (r *Redis) BRPopLPush(src, dst string, timeout int) (string, error) {
	return r.doToString("BRPOPLPUSH", src, dst, timeout)
}
