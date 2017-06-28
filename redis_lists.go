package redis

import (
	"strings"
)

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

// LIndex executes the redis command LINDEX.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) LIndex(key string, index int) string {
	return r.doToString("LINDEX", key, index)
}

// LInsert executes the redis command LINSERT.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) LInsert(key, ba, pivot, value string) int64 {
	ba = strings.ToUpper(ba)
	if ba != "BEFORE" && ba != "AFTER" {
		panic(ErrInvalidArgs)
	}

	return r.doToInt("LINSERT", key, ba, pivot, value)
}

// LLen executes the redis command LLEN.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) LLen(key string) int64 {
	return r.doToInt("LLEN", key)
}

// LPushX executes the redis command LPUSHX.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) LPushX(key, value string) int64 {
	return r.doToInt("LPUSHX", key, value)
}
