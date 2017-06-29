package redis

import (
	"strings"
)

// BGRewriteAOF executes the redis command BGREWRITEAOF.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) BGRewriteAOF() {
	r.do("BGREWRITEAOF")
}

// BGSave executes the redis command BGSAVE.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) BGSave() {
	r.do("BGSAVE")
}

// ClientGetName executes the redis command CLIENT GETNAME.
//
// Panic if an error occurs.
//
// New in redis version 2.6.9.
func (r *Redis) ClientGetName() string {
	return r.doToString("CLIENT", "GETNAME")
}

// ClientSetName executes the redis command CLIENT SETNAME.
//
// Panic if an error occurs.
//
// New in redis version 2.6.9.
func (r *Redis) ClientSetName(name string) {
	r.do("CLIENT", "SETNAME", name)
}

// ClientKill executes the redis command CLIENT KILL.
//
// Return nil or a int64. Panic if an error occurs.
//
// New in redis version 2.4.0.
func (r *Redis) ClientKill(args ...interface{}) interface{} {
	_args := make([]interface{}, len(args)+1)
	_args[0] = "KILL"
	for i, v := range args {
		_args[i+1] = v
	}

	if _r, err := r.Do("CLIENT", _args...); err != nil {
		panic(err)
	} else if _, ok := _r.(int64); ok {
		return _r
	}
	return nil
}

// ClientList executes the redis command CLIENT LIST.
//
// Panic if an error occurs.
//
// New in redis version 2.4.0.
func (r *Redis) ClientList() []map[string]string {
	ss := r.doToString("CLIENT", "LIST")

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

	return results
}

// ClientPause executes the redis command CLIENT PAUSE.
//
// Panic if an error occurs.
//
// New in redis version 2.9.50.
func (r *Redis) ClientPause(timeout int) {
	r.do("CLIENT", "PAUSE", timeout)
}

// CommandCount executes the redis command COMMAND COUNT.
//
// Panic if an error occurs.
//
// New in redis version 2.8.13.
func (r *Redis) CommandCount() int64 {
	return r.doToInt("COMMAND", "COUNT")
}

// CommandGetKeys executes the redis command COMMAND GETKEYS.
//
// Panic if an error occurs.
//
// New in redis version 2.8.13.
func (r *Redis) CommandGetKeys() []string {
	return r.doToStringSlice("COMMAND", "GETKEYS")
}

// ConfigGet executes the redis command CONFIG GET.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ConfigGet(parameter string) []string {
	return r.doToStringSlice("CONFIG", "GET", parameter)
}

// ConfigResetStat executes the redis command CONFIG RESETSTAT.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ConfigResetStat() {
	r.do("CONFIG", "RESETSTAT")
}

// ConfigRewrite executes the redis command CONFIG REWRITE.
//
// Panic if an error occurs.
//
// New in redis version 2.8.0.
func (r *Redis) ConfigRewrite() {
	r.do("CONFIG", "REWRITE")
}

// ConfigSet executes the redis command CONFIG SET.
//
// Panic if an error occurs.
//
// New in redis version 2.0.0.
func (r *Redis) ConfigSet(parameter, value string) {
	r.do("CONFIG", "SET", parameter, value)
}

// DBSize executes the redis command DBSIZE.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) DBSize() int64 {
	return r.doToInt("DBSIZE")
}

// FlushAll executes the redis command FLUSHALL.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) FlushAll() {
	r.do("FLUSHALL")
}

// FlushDB executes the redis command FLUSHDB.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) FlushDB() {
	r.do("FLUSHDB")
}

// Info executes the redis command INFO.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Info(section ...string) map[string]string {
	var ss string
	if len(section) == 0 {
		ss = r.doToString("INFO")
	} else {
		ss = r.doToString("INFO", section[0])
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

	return results
}
