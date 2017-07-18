package redis

// PFAdd executes the redis command PFADD.
//
// For the returned value, true is 1 and false is 0.
//
// New in redis version 2.8.9.
func (r *Redis) PFAdd(key, element string, elements ...string) (bool, error) {
	args := make([]interface{}, len(elements)+2)
	args[0] = key
	args[1] = element
	for i, v := range elements {
		args[i+2] = v
	}
	return r.doToBool("PFADD", args...)
}

// PFCount executes the redis command PFCOUNT.
//
// New in redis version 2.8.9.
func (r *Redis) PFCount(key string, keys ...string) (int64, error) {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, v := range keys {
		args[i+1] = v
	}
	return r.doToInt("PFCOUNT", args...)
}

// PFMerge executes the redis command PFMERGE.
//
// New in redis version 2.8.9.
func (r *Redis) PFMerge(dstKey, srcKey string, srcKeys ...string) error {
	args := make([]interface{}, len(srcKeys)+2)
	args[0] = dstKey
	args[1] = srcKey
	for i, v := range srcKeys {
		args[i+2] = v
	}
	return r.do("PFMERGE", args...)
}
