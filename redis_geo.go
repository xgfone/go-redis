package redis

// GeoAdd executes the redis command GEOADD.
//
// Panic if an error occurs.
//
// New in redis version 3.2.0.
func (r *Redis) GeoAdd(key string, longitude, latitude interface{},
	member string, others ...interface{}) int64 {
	_len := len(others)
	if _len%3 != 0 {
		panic(ErrInvalidArgs)
	}

	args := make([]interface{}, _len+4)
	args[0] = key
	args[1] = longitude
	args[2] = latitude
	args[3] = member
	for i, v := range others {
		args[i+4] = v
	}
	return r.doToInt("GEOADD", args...)
}

// GeoDist executes the redis command GEODIST.
//
// Return -1.0 if one or both the members are missing. Panic if an error occurs.
//
// New in redis version 3.2.0.
func (r *Redis) GeoDist(key, member1, member2 string, unit ...string) float64 {
	if len(unit) == 0 {
		return r.doToFloat("GEODIST", key, member1, member2)
	}

	switch unit[0] {
	case "m", "km", "mi", "ft":
	default:
		panic(ErrInvalidArgs)
	}

	return r.doToFloat("GEODIST", key, member1, member2, unit[0])
}

// GeoRadius executes the redis command GEORADIUS.
//
// The type of the returned value []string, or [][]string If WITHCOORD, WITHDIST
// or WITHHASH options are specified. Panic if an error occurs.
//
// New in redis version 3.2.0.
func (r *Redis) GeoRadius(key string, longitude, latitude, radius interface{},
	unit string, others ...interface{}) []interface{} {
	switch unit {
	case "m", "km", "mi", "ft":
	default:
		panic(ErrInvalidArgs)
	}

	args := make([]interface{}, len(others)+5)
	args[0] = key
	args[1] = longitude
	args[2] = latitude
	args[3] = radius
	args[4] = unit
	for i, v := range others {
		args[i+5] = v
	}

	if _r, err := r.Do("GEORADIUS", args...); err != nil {
		panic(err)
	} else if _r != nil {
		vs := _r.([]interface{})
		for i, v := range vs {
			switch v.(type) {
			case []byte:
				vs[i] = string(v.([]byte))
			case []interface{}:
				_vs := v.([]interface{})
				for j, _v := range _vs {
					_vs[j] = string(_v.([]byte))
				}
				vs[i] = _vs
			}
		}
		return vs
	}
	return nil
}

// GeoHash executes the redis command GEOHASH.
//
// Panic if an error occurs.
//
// New in redis version 3.2.0.
func (r *Redis) GeoHash(key, member string, members ...string) []string {
	args := make([]interface{}, len(members)+2)
	args[0] = key
	args[1] = member
	for i, v := range members {
		args[i+2] = v
	}

	return r.doToStringSlice("GEOHASH", args...)
}

// GeoPos executes the redis command GEOPOS.
//
// Panic if an error occurs.
//
// New in redis version 3.2.0.
func (r *Redis) GeoPos(key, member string, members ...string) [][]string {
	args := make([]interface{}, len(members)+2)
	args[0] = key
	args[1] = member
	for i, v := range members {
		args[i+2] = v
	}

	if _r, err := r.Do("GEOPOS", args...); err != nil {
		panic(err)
	} else if _r != nil {
		vs := _r.([]interface{})
		results := make([][]string, len(vs))
		for i, v := range vs {
			switch v.(type) {
			case []interface{}:
				_vs := v.([]interface{})
				_results := make([]string, len(_vs))
				for j, _v := range _vs {
					_results[j] = string(_v.([]byte))
				}
				results[i] = _results
			default:
				results[i] = nil
			}
		}
		return results
	}
	return nil
}

// GeoRadiusByMember executes the redis command GEORADIUSBYMEMBER.
//
// The type of the returned value []string, or [][]string If WITHCOORD, WITHDIST
// or WITHHASH options are specified. Panic if an error occurs.
//
// New in redis version 3.2.0.
func (r *Redis) GeoRadiusByMember(key, member string, radius interface{},
	unit string, others ...interface{}) []interface{} {
	switch unit {
	case "m", "km", "mi", "ft":
	default:
		panic(ErrInvalidArgs)
	}

	args := make([]interface{}, len(others)+4)
	args[0] = key
	args[1] = member
	args[2] = radius
	args[3] = unit
	for i, v := range others {
		args[i+4] = v
	}

	if _r, err := r.Do("GEORADIUSBYMEMBER", args...); err != nil {
		panic(err)
	} else if _r != nil {
		vs := _r.([]interface{})
		for i, v := range vs {
			switch v.(type) {
			case []byte:
				vs[i] = string(v.([]byte))
			case []interface{}:
				_vs := v.([]interface{})
				for j, _v := range _vs {
					_vs[j] = string(_v.([]byte))
				}
				vs[i] = _vs
			}
		}
		return vs
	}
	return nil
}
