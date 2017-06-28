package redis

// PSubscribe executes the redis command PSUBSCRIBE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) PSubscribe(pattern string, patterns ...string) {
	args := make([]interface{}, len(patterns)+1)
	args[0] = pattern
	for i, v := range patterns {
		args[i] = v
	}
	r.do("PSUBSCRIBE", args...)
}

// PUnsubscribe executes the redis command PUNSUBSCRIBE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) PUnsubscribe(patterns ...string) {
	args := make([]interface{}, len(patterns))
	for i, v := range patterns {
		args[i] = v
	}
	r.do("PUNSUBSCRIBE", args...)
}

// Subscribe executes the redis command SUBSCRIBE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) Subscribe(channel string, channels ...string) {
	args := make([]interface{}, len(channels)+1)
	args[0] = channel
	for i, v := range channels {
		args[i] = v
	}
	r.do("SUBSCRIBE", args...)
}

// Unsubscribe executes the redis command UNSUBSCRIBE.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) Unsubscribe(channels ...string) {
	args := make([]interface{}, len(channels))
	for i, v := range channels {
		args[i] = v
	}
	r.do("UNSUBSCRIBE", args...)
}
