package redis

import "strconv"

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

func (r *Redis) do(cmd string, args ...interface{}) error {
	_, err := r.Do(cmd, args...)
	return err

}

func (r *Redis) doToInt(cmd string, args ...interface{}) (int64, error) {
	_r, err := r.Do(cmd, args...)
	if err != nil {
		return 0, err
	}
	return _r.(int64), nil

}

func (r *Redis) doToBool(cmd string, args ...interface{}) (bool, error) {
	_r, err := r.Do(cmd, args...)
	if err != nil {
		return false, err
	}
	return toBool(_r), nil
}

func (r *Redis) doToString(cmd string, args ...interface{}) (string, error) {
	_r, err := r.Do(cmd, args...)
	if err != nil {
		return "", err
	} else if _r != nil {
		switch _r.(type) {
		case string:
			return _r.(string), nil
		case []byte:
			return string(_r.([]byte)), nil
		default:
			return "", ErrInvalidResult
		}
	}
	return "", nil
}

func (r *Redis) doToFloat(cmd string, args ...interface{}) (float64, error) {
	if _r, err := r.Do(cmd, args...); err != nil {
		return 0, err
	} else if _r != nil {
		var v string
		switch _r.(type) {
		case []byte:
			v = string(_r.([]byte))
		case string:
			v = _r.(string)
		default:
			return 0, ErrInvalidResult
		}
		_v, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return _v, nil
	}
	return 0, ErrNotExist
}

func (r *Redis) doToStringSlice(cmd string, args ...interface{}) ([]string, error) {
	if _r, err := r.Do(cmd, args...); err != nil {
		return nil, err
	} else if _r != nil {
		if bs, ok := _r.([]byte); ok {
			return []string{string(bs)}, nil
		}

		vs := _r.([]interface{})
		results := make([]string, len(vs))
		for i, v := range vs {
			if v == nil {
				results[i] = ""
			} else {
				results[i] = string(v.([]byte))
			}
		}
		return results, nil
	}
	return nil, nil
}
