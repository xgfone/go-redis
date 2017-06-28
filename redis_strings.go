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

// Get executes the redis command GET.
//
// Panic if an error occurs. Return "" if the key does not exist.
//
// New in redis version 1.0.0.
func (r *Redis) Get(key string) string {
	if reply, err := r.Do("GET", key); err != nil {
		panic(err)
	} else if reply != nil {
		return string(reply.([]byte))
	}
	return ""
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
	if _len != 0 && _len != 2 {
		panic(ErrInvalidArgs)
	}

	var reply interface{}
	var err error
	if _len == 0 {
		reply, err = r.Do("BITCOUNT", key)
	} else {
		reply, err = r.Do("BITCOUNT", key, args[0], args[1])
	}
	if err != nil {
		panic(err)
	}
	return reply.(int64)
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

	if _r, err := r.Do("BITOP", args...); err != nil {
		panic(err)
	} else {
		return _r.(int64)
	}
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

	if _r, err := r.Do("BITPOS", _args...); err != nil {
		panic(err)
	} else {
		return _r.(int64)
	}
}

// Decr executes the redis command DECR.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Decr(key string) int64 {
	if _r, err := r.Do("DECR", key); err != nil {
		panic(err)
	} else {
		return _r.(int64)
	}
}

// DecrBy executes the redis command DECRBY.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) DecrBy(key string, n int) int64 {
	if _r, err := r.Do("DECRBY", key, n); err != nil {
		panic(err)
	} else {
		return _r.(int64)
	}
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
