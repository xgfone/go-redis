package redis

// Auth executes the redis command AUTH.
//
// New in redis version 1.0.0.
func (r *Redis) Auth(passwd string) error {
	return r.do("AUTH", passwd)
}

// Echo executes the redis command ECHO.
//
// New in redis version 1.0.0.
func (r *Redis) Echo(msg string) (string, error) {
	return r.doToString("ECHO", msg)
}

// Ping executes the redis command PING.
//
// New in redis version 1.0.0.
func (r *Redis) Ping(msg ...string) (string, error) {
	if len(msg) == 0 {
		return r.doToString("PING")
	}
	return r.doToString("PING", msg)
}

// Quit executes the redis command QUIT.
//
// New in redis version 1.0.0.
func (r *Redis) Quit() error {
	return r.do("QUIT")
}

// Select executes the redis command SELECT.
//
// New in redis version 1.0.0.
func (r *Redis) Select(index int) error {
	return r.do("SELECT", index)
}

// SwapDB executes the redis command SWAPDB.
//
// New in redis version 4.0.0.
func (r *Redis) SwapDB(index1, index2 int) error {
	return r.do("SWAPDB", index1, index2)
}
