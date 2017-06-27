package redis

// Keys executes the redis command KEYS
func (r *Redis) Keys(pattern string) []string {
	reply, err := r.Do("KEYS", pattern)
	if err != nil {
		return nil
	}
	_results := reply.([]interface{})
	if len(_results) == 0 {
		return nil
	}

	results := make([]string, len(_results))
	for i, _r := range _results {
		results[i] = string(_r.([]byte))
	}

	return results
}
