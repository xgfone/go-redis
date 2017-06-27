package redis

import "strings"

// Set executes the redis command SET.
//
// Panic if an error occurs.
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
// Panic if an error occurs.
//
// Return "" if the key does not exist.
func (r *Redis) Get(key string) string {
	if reply, err := r.Do("GET", key); err != nil {
		panic(err)
	} else if reply != nil {
		return string(reply.([]byte))
	}
	return ""
}

// Append executes the redis command APPEND.
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
