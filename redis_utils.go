package redis

import (
	"strconv"
)

func toBool(v interface{}) bool {
	switch v.(type) {
	case int64:
		_v := v.(int64)
		if _v == 0 {
			return false
		}
		return true
	default:
		return false
	}
}

func (r *Redis) do(cmd string, args ...interface{}) {
	if _, err := r.Do(cmd, args...); err != nil {
		panic(err)
	}
}

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

func (r *Redis) doToStringSlice(cmd string, args ...interface{}) []string {
	if _r, err := r.Do(cmd, args...); err != nil {
		panic(err)
	} else if _r != nil {
		vs := _r.([]interface{})
		results := make([]string, len(vs))
		for i, v := range vs {
			if v == nil {
				results[i] = ""
			} else {
				results[i] = string(v.([]byte))
			}
		}
		return results
	}
	return nil
}
