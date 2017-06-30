# go-redis
Supply the high-level API interface based on https://github.com/garyburd/redigo. It uses a pool to handle the redis connection in the underlying implementation, but you don't have to care about it, and it should be transparent to the user.

## About the API

For the meaning of APIs, please see the corresponding redis command [doc](https://redis.io/commands).

#### Name

The API name is the name of the redis command, but the API name is the format of `CamelCase`. For instance,

|  Redis Command  |  API Name
|-----------------|------------
|  SET            | Set
|  SETNX          | SetNX
|  SETRANGE       | SetRange
|  CONFIG GET     | ConfigGet

#### Arguments

The arguments of API is the same as the redis command. See the redis [doc](https://redis.io/commands).

#### Return

The correspondences is as follows.

| Redis Return | API Return
|---------------|--------------
| String("OK")  | NO RETURN VALUE
| Integer       | `int64`
| Integer(0/1)  | `bool`
| Float String  | `float64`
| Bulk String   | `string`
| Array         | `[]string`
| Maybe Two Types | `interface{}`
| Maybe Two Slice Types | `[]interface{}`

**Notice:**

1. If the redis connection has an error, or the redis server returns an error, the API will panic with the error.
2. For the specail commands of `INFO` and `CLIENT LIST`, the corresponding APIs, `Info` and `ClientList`, will return the parsed Key-Values, that's, `map[string]string`.


## Implemented Commands Table

|                   |   Implemented Commands   |  Unimplemented Commands
|-------------------|--------------------------|---------------------------
| **Connection**    | AUTH, ECHO, PING, QUIT, SELECT, SWAPDB(required redis 4.0) |
| **Geo**           | GEOADD, GEODIST, GEOHASH, GEOPOS, GEORADIUS, GEORADIUSBYMEMBER |
| **Hashes**        | HDEL, HEXISTS, HGET, HGETALL, HINCRBY, HINCRBYFLOAT, HKEYS, HLEN, HMGET, HMSET, HSET, HSETNX, HSTRLEN, HVALS | HSCAN
| **HyperLogLog**   | PFADD, PFCOUNT, PFMERGE |
| **Keys**          | DEL, DUMP, EXISTS, EXPIRE, EXPIREAT, KEYS, MOVE, PERSIST, PEXPIRE, PEXPIREAT, PTTL, RANDOMKEY, RENAME, RENAMENX, TTL, TYPE | MIGRATE, OBJECT, RESTORE, SORT, WAIT, SCAN
| **Lists**         | BLPOP, BRPOP, BRPOPLPUSH, LINDEX, LINSERT, LLEN, LPOP, LPUSH, LPUSHX, LRANGE, LREM, LSET, LTRIM, RPOP, RPOPLPUSH, RPUSH, RPUSHX |
| **Pub/Sub**       | PSUBSCRIBE, PUNSUBSCRIBE, SUBSCRIBE, UNSUBSCRIBE, PUBLISH, PUBSUB |
| **Server**        | BGREWRITEAOF, BGSAVE, CLIENT GETNAME, CLIENT SETNAME, CLIENT KILL, CLIENT LIST, CLIENT PAUSE, CLIENT REPLY, COMMAND COUNT, COMMAND GETKEYS, CONFIG GET, CONFIG SET, CONFIG RESETSTAT, CONFIG REWRITE, DBSIZE, FLUSHALL, FLUSHDB, INFO, LASTSAVE, SAVE, SHUTDOWN, SLAVEOF, TIME | COMMAND, COMMAND INFO, DEBUG OBJECT, DEBUG SEGFAULT, MONITOR, ROLE, SLOWLOG, SYNC
| **Sets**          | SADD, SCARD, SDIFF, SDIFFSTORE, SINTER, SINTERSTORE, SISMEMBER, SMEMBERS, SMOVE, SPOP, SRANDMEMBER, SREM, SUNION, SUNIONSTORE | SSCAN
| **Sorted Sets**   | ZADD, ZCARD, ZCOUNT, ZINCRBY, ZINTERSTORE, ZLEXCOUNT, ZRANGE, ZRANGEBYLEX, ZRANGEBYSCORE, ZRANK, ZREM, ZREMRANGEBYLEX, ZREMRANGEBYRANK, ZREMRANGEBYSCORE, ZREVRANGE, ZREVRANGEBYLEX, ZREVRANGEBYSCORE, ZREVRANK, ZSCORE, ZUNIONSTORE | ZSCAN
| **String**        | APPEND, BITCOUNT, BITOP, BITPOS, DECR, DECRBY, GET, GETBIT, GETRANGE, GETSET, INCR, INCRBY, INCRBYFLOAT, MGET, MSET, MSETNX, PSETEX, SET, SETBIT, SETEX, SETNX, SETRANGE, STRLEN | BITFIELD
| **Transactions**  | DISCARD, EXEC, MULTI, WATCH, UNWATCH |
| **Cluster**       | In future |
| **Scripting**     | In future |


### Notice

For the unimplemented commands, you can use the Do method,

    func (r *Redis) Do(cmd string, args ...interface{}) (reply interface{}, err error)

to implement it.

Welcome to give me a pull request.
