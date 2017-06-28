package redis

// PFAdd executes the redis command PFADD.
//
// For the returned value, true is 1 and false is 0. Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) PFAdd(key, element string, elements ...string) bool {
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
// Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) PFCount(key string, keys ...string) int64 {
	args := make([]interface{}, len(keys)+1)
	args[0] = key
	for i, v := range keys {
		args[i+1] = v
	}
	return r.doToInt("PFCOUNT", args...)
}

// PFMerge executes the redis command PFMERGE.
//
// Panic if an error occurs.
//
// New in redis version 2.8.9.
func (r *Redis) PFMerge(dstKey, srcKey string, srcKeys ...string) {
	args := make([]interface{}, len(srcKeys)+2)
	args[0] = dstKey
	args[1] = srcKey
	for i, v := range srcKeys {
		args[i+2] = v
	}
	r.do("PFMERGE", args...)
}
