package redis

// Auth executes the redis command AUTH.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Auth(passwd string) {
	r.do("AUTH", passwd)
}

// Echo executes the redis command ECHO.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Echo(msg string) string {
	return r.doToString("ECHO", msg)
}

// Ping executes the redis command PING.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Ping(msg ...string) string {
	if len(msg) == 0 {
		return r.doToString("PING")
	}
	return r.doToString("PING", msg)
}

// Quit executes the redis command QUIT.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Quit() {
	r.do("QUIT")
}

// Select executes the redis command SELECT.
//
// Return an error if failed, or nil. Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Select(index int) {
	r.doToString("SELECT", index)
}
