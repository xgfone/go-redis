package redis

// PSubscribe executes the redis command PSUBSCRIBE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) PSubscribe(pattern string) {
	r.do("PSUBSCRIBE", pattern)
}

// PUnsubscribe executes the redis command PUNSUBSCRIBE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) PUnsubscribe(patterns ...string) {
	args := make([]interface{}, len(patterns))
	for i, p := range patterns {
		args[i] = p
	}
	r.do("PUNSUBSCRIBE", args...)
}
