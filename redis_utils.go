package redis

import (
	"strconv"
)

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

func (r *Redis) doToFloat(cmd string, args ...interface{}) float64 {
	if _r, err := r.Do(cmd, args...); err != nil {
		panic(err)
	} else if _r != nil {
		var v string
		switch _r.(type) {
		case []byte:
			v = string(_r.([]byte))
		case string:
			v = _r.(string)
		default:
			return 0.0
		}
		_v, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		return _v
	}
	return 0.0
}
