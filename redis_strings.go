package redis

import "strings"

// Set executes the redis command SET.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Set(key string, value interface{}, args ...interface{}) {
	var err error

	if len(args) == 0 {
		_, err = r.Do("SET", key, value)
	} else {
		_args := make([]interface{}, len(args)+2)
		_args[0] = key
		_args[1] = value
		for i, v := range args {
			_args[i+2] = v
		}
		_, err = r.Do("SET", _args...)
	}

	if err != nil {
		panic(err)
	}
}

// SetEX executes the redis command SETEX.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) SetEX(key string, timeout int, value string) {
	if _, err := r.Do("SETEX", key, timeout, value); err != nil {
		panic(err)
	}
}

// PSetEX executes the redis command PSETEX.
//
// Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) PSetEX(key string, timeout int, value string) {
	if _, err := r.Do("PSETEX", key, timeout, value); err != nil {
		panic(err)
	}
}

// SetNX executes the redis command SETNX.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) SetNX(key string, value string) bool {
	if _r, err := r.Do("SETNX", key, value); err != nil {
		panic(err)
	} else {
		return toBool(_r)
	}
}

// Get executes the redis command GET.
//
// Panic if an error occurs. Return "" if the key does not exist.
//
// New in redis version 1.0.0.
func (r *Redis) Get(key string) string {
	return r.doToString("GET", key)
}

// Append executes the redis command APPEND.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) Append(key, value string) {
	if _, err := r.Do("APPEND", key, value); err != nil {
		panic(err)
	}
}

// BitCount executes the redis command BITCOUNT.
//
// Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) BitCount(key string, args ...int) int64 {
	_len := len(args)
	var _args []interface{}
	if _len == 0 {
		_args = []interface{}{key}
	} else if _len == 2 {
		_args = []interface{}{key, args[0], args[1]}
	} else {
		panic(ErrInvalidArgs)
	}

	return r.doToInt("BITCOUNT", _args...)
}

// BitOp executes the redis command BITOP.
//
// Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) BitOp(op, dest, src string, srcs ...string) int64 {
	op = strings.ToUpper(op)
	switch op {
	case "AND", "OR", "NOT", "XOR":
	default:
		panic(ErrInvalidArgs)
	}

	args := make([]interface{}, len(srcs)+3)
	args[0] = op
	args[1] = dest
	args[2] = src
	for i, s := range srcs {
		args[i+3] = s
	}

	return r.doToInt("BITOP", args...)
}

// BitPos executes the redis command BITPOS.
//
// For the argument, bit, true is 1 and false is 0.
//
// Panic if an error occurs.
//
// New in redis version 2.8.7.
func (r *Redis) BitPos(key string, bit bool, args ...int) int64 {
	_args := make([]interface{}, len(args)+2)
	_args[0] = key
	if bit {
		_args[1] = 1
	} else {
		_args[1] = 0
	}
	for i, a := range args {
		_args[i+2] = a
	}

	return r.doToInt("BITPOS", _args...)
}

// Decr executes the redis command DECR.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Decr(key string) int64 {
	return r.doToInt("DECR", key)
}

// Incr executes the redis command INCR.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Incr(key string) int64 {
	return r.doToInt("INCR", key)
}

// DecrBy executes the redis command DECRBY.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) DecrBy(key string, n int) int64 {
	return r.doToInt("DECRBY", key, n)
}

// IncrBy executes the redis command INCRBY.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) IncrBy(key string, n int) int64 {
	return r.doToInt("INCRBY", key, n)
}

// IncrByFloat executes the redis command INCRBYFloat.
//
// Panic if an error occurs.
//
// New in redis version 2.6.0.
func (r *Redis) IncrByFloat(key string, n float64) float64 {
	return r.doToFloat("INCRBYFLOAT", key, n)
}

// GetBit executes the redis command GETBIT.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) GetBit(key string, offset int) int64 {
	return r.doToInt("GETBIT", key, offset)
}

// SetBit executes the redis command SETBIT.
//
// For the argument, value, true is 1 and false is 0.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) SetBit(key string, offset int, value bool) int64 {
	var v int8
	if value {
		v = 1
	} else {
		v = 0
	}

	return r.doToInt("SETBIT", key, offset, v)
}

// GetRange executes the redis command GETRANGE.
//
// Panic if an error occurs.
//
// New in redis version 2.4.0.
func (r *Redis) GetRange(key string, start, end int) string {
	return r.doToString("GETRANGE", key, start, end)
}

// SetRange executes the redis command SETRANGE.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) SetRange(key string, offset int, value string) int64 {
	return r.doToInt("SETRANGE", key, offset, value)
}

// GetSet executes the redis command GETSET.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) GetSet(key, value string) string {
	return r.doToString("GETSET", key, value)
}

// MGet executes the redis command MGET.
//
// If a certain key does not exist, this value is "". Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) MGet(key string, keys ...string) []string {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, a := range keys {
		args[i+1] = a
	}
	return r.doToStringSlice("MGET", args...)
}

func (r *Redis) mSet(cmd, key string, value interface{}, kvs ...interface{}) interface{} {
	_len := len(kvs)
	if _len%2 != 0 {
		panic(ErrInvalidArgs)
	}

	args := make([]interface{}, _len+2)
	args[0] = key
	args[1] = value
	for i, a := range kvs {
		args[i+2] = a
	}

	if _r, err := r.Do(cmd, args...); err != nil {
		panic(err)
	} else {
		return _r
	}
}

// MSet executes the redis command MSET.
//
// Panic if an error occurs.
//
// New in redis version 1.0.1.
func (r *Redis) MSet(key string, value interface{}, kvs ...interface{}) {
	r.mSet("MSET", key, value, kvs...)
}

// MSetNX executes the redis command MSETNX.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 1.0.1.
func (r *Redis) MSetNX(key string, value interface{}, kvs ...interface{}) bool {
	return toBool(r.mSet("MSETNX", key, value, kvs...))
}

// StrLen executes the redis command STRLEN.
//
// Panic if an error occurs.
//
// New in redis version 2.2.0.
func (r *Redis) StrLen(key string) int64 {
	return r.doToInt("STRLEN", key)
}
