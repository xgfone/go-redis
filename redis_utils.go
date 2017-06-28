package redis

func (r *Redis) doToInt(cmd string, args ...interface{}) int64 {
	if _r, err := r.Do(cmd, args...); err != nil {
		panic(err)
	} else {
		return _r.(int64)
	}
}

func (r *Redis) doToString(cmd string, args ...interface{}) string {
	if _r, err := r.Do(cmd, args...); err != nil {
		panic(err)
	} else if _r != nil {
		return string(_r.([]byte))
	}
	return ""
}
