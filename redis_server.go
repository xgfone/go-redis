package redis

import (
	"strconv"
	"strings"
)

// BGRewriteAOF executes the redis command BGREWRITEAOF.
//
// New in redis version 1.0.0.
func (r *Redis) BGRewriteAOF() error {
	return r.do("BGREWRITEAOF")
}

// BGSave executes the redis command BGSAVE.
//
// New in redis version 1.0.0.
func (r *Redis) BGSave() error {
	return r.do("BGSAVE")
}

// ClientGetName executes the redis command CLIENT GETNAME.
//
// New in redis version 2.6.9.
func (r *Redis) ClientGetName() (string, error) {
	return r.doToString("CLIENT", "GETNAME")
}

// ClientSetName executes the redis command CLIENT SETNAME.
//
// New in redis version 2.6.9.
func (r *Redis) ClientSetName(name string) error {
	return r.do("CLIENT", "SETNAME", name)
}

// ClientKill executes the redis command CLIENT KILL.
//
// Return nil or a int64.
//
// New in redis version 2.4.0.
func (r *Redis) ClientKill(args ...interface{}) (interface{}, error) {
	_args := make([]interface{}, len(args)+1)
	_args[0] = "KILL"
	for i, v := range args {
		_args[i+1] = v
	}

	if _r, err := r.Do("CLIENT", _args...); err != nil {
		return nil, err
	} else if _, ok := _r.(int64); ok {
		return _r, nil
	}
	return nil, nil
}

// ClientList executes the redis command CLIENT LIST.
//
// New in redis version 2.4.0.
func (r *Redis) ClientList() ([]map[string]string, error) {
	ss, err := r.doToString("CLIENT", "LIST")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(ss), "\n")
	results := make([]map[string]string, len(lines))
	for i, line := range lines {
		items := strings.Split(line, " ")
		sm := make(map[string]string, len(items))
		for _, item := range items {
			tmp := strings.Split(item, "=")
			sm[tmp[0]] = tmp[1]
		}
		results[i] = sm
	}

	return results, nil
}

// ClientPause executes the redis command CLIENT PAUSE.
//
// New in redis version 2.9.50.
func (r *Redis) ClientPause(timeout int) error {
	return r.do("CLIENT", "PAUSE", timeout)
}

// ClientReply executes the redis command CLIENT REPLY.
//
// New in redis version 3.2.0.
func (r *Redis) ClientReply(arg string) error {
	arg = strings.ToUpper(arg)
	switch arg {
	case "ON", "OFF", "SKIP":
	default:
		return ErrInvalidArgs
	}
	return r.do("CLIENT", "REPLY", arg)
}

// CommandCount executes the redis command COMMAND COUNT.
//
// New in redis version 2.8.13.
func (r *Redis) CommandCount() (int64, error) {
	return r.doToInt("COMMAND", "COUNT")
}

// CommandGetKeys executes the redis command COMMAND GETKEYS.
//
// New in redis version 2.8.13.
func (r *Redis) CommandGetKeys(command string, commands ...interface{}) ([]string, error) {
	args := make([]interface{}, len(commands)+2)
	args[0] = "GETKEYS"
	args[1] = command
	for i, v := range commands {
		args[i+2] = v
	}
	return r.doToStringSlice("COMMAND", args...)
}

// ConfigGet executes the redis command CONFIG GET.
//
// Notice the difference of the returned value between the Linux Redis server
// and the Windows Redis server. For instance, r.ConfigGet("save") returns
// ["save", "jd 900 jd 300"] for Windows, and ["save", "900", "1", "300", "10"]
// for Linux.
//
// New in redis version 2.0.0.
func (r *Redis) ConfigGet(parameter string) ([]string, error) {
	return r.doToStringSlice("CONFIG", "GET", parameter)
}

// ConfigResetStat executes the redis command CONFIG RESETSTAT.
//
// New in redis version 2.0.0.
func (r *Redis) ConfigResetStat() error {
	return r.do("CONFIG", "RESETSTAT")
}

// ConfigRewrite executes the redis command CONFIG REWRITE.
//
// New in redis version 2.8.0.
func (r *Redis) ConfigRewrite() error {
	return r.do("CONFIG", "REWRITE")
}

// ConfigSet executes the redis command CONFIG SET.
//
// New in redis version 2.0.0.
func (r *Redis) ConfigSet(parameter, value string) error {
	return r.do("CONFIG", "SET", parameter, value)
}

// DBSize executes the redis command DBSIZE.
//
// New in redis version 1.0.0.
func (r *Redis) DBSize() (int64, error) {
	return r.doToInt("DBSIZE")
}

// FlushAll executes the redis command FLUSHALL.
//
// New in redis version 1.0.0.
func (r *Redis) FlushAll() error {
	return r.do("FLUSHALL")
}

// FlushDB executes the redis command FLUSHDB.
//
// New in redis version 1.0.0.
func (r *Redis) FlushDB() error {
	return r.do("FLUSHDB")
}

// Info executes the redis command INFO.
//
// New in redis version 1.0.0.
func (r *Redis) Info(section ...string) (map[string]string, error) {
	var ss string
	var err error
	if len(section) == 0 {
		ss, err = r.doToString("INFO")
	} else {
		ss, err = r.doToString("INFO", section[0])
	}
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(ss), "\n")
	results := make(map[string]string, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}

		items := strings.Split(line, ":")
		results[items[0]] = items[1]
	}

	return results, nil
}

// LastSave executes the redis command LASTSAVE.
//
// New in redis version 1.0.0.
func (r *Redis) LastSave() (int64, error) {
	return r.doToInt("LASTSAVE")
}

// Save executes the redis command SAVE.
//
// New in redis version 1.0.0.
func (r *Redis) Save() error {
	return r.do("SAVE")
}

// Shutdown executes the redis command SHUTDOWN.
//
// All the clients will be halted when executed this command,
// and the connection will be closed.
//
// New in redis version 1.0.0.
func (r *Redis) Shutdown(save ...string) error {
	if len(save) > 1 {
		return ErrInvalidArgs
	} else if len(save) == 1 {
		_save := strings.ToUpper(save[0])
		switch _save {
		case "NOSAVE", "SAVE":
		default:
			return ErrInvalidArgs
		}
		r.do("SHUTDOWN", _save)
		return nil
	}
	return r.do("SHUTDOWN")
}

// SlaveOf executes the redis command SLAVEOF.
//
// If the parameters are either NO ONE or host port, which the port must be
// the type of int and between 1 and b5535.
//
// New in redis version 1.0.0.
func (r *Redis) SlaveOf(host string, port interface{}) error {
	switch port.(type) {
	case string:
		host = strings.ToUpper(host)
		_port := strings.ToUpper(port.(string))
		if host != "NO" || _port != "ONE" {
			return ErrInvalidArgs
		}
		port = _port
	case int:
		_port := port.(int)
		if _port < 1 || _port > 65535 {
			return ErrInvalidArgs
		}
	default:
		return ErrInvalidArgs
	}
	return r.do("SLAVEOF", host, port)
}

// Time executes the redis command TIME.
//
// New in redis version 2.6.0.
func (r *Redis) Time() (seconds, microseconds int64, err error) {
	vs, err := r.doToStringSlice("TIME")
	if err != nil {
		return
	}

	if seconds, err = strconv.ParseInt(vs[0], 10, 64); err != nil {
		return
	}
	if microseconds, err = strconv.ParseInt(vs[1], 10, 64); err != nil {
		return
	}
	return
}
