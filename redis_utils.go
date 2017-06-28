package redis

func (r *Redis) doToInt(cmd string, args ...interface{}) int64 {
	if _r, err := r.Do(cmd, args...); err != nil {
		panic(err)
	} else {
		return _r.(int64)
	}
}
