# go-redis
Supply the high-level API interface based on https://github.com/garyburd/redigo

## Implemented Commands Table

|                   |   Implemented Commands   |  Unimplemented Commands
|-------------------|--------------------------|---------------------------
| **Cluster**       | In future |
| **Connection**    | AUTH, ECHO, PING, QUIT, SELECT, SWAPDB(from redis 4.0) |
| **Geo**           | In future |
| **Hashes**        | In future |
| **HyperLogLog**   | In future |
| **Keys**          | DEL, DUMP, EXISTS, EXPIRE, EXPIREAT, KEYS, MOVE, PERSIST, PEXPIRE, PEXPIREAT, PTTL, RANDOMKEY, RENAME, RENAMENX, TTL, TYPE | MIGRATE, OBJECT, RESTORE, SORT, WAIT, SCAN
| **Lists**         | BLPOP, BRPOP, BRPOPLPUSH, LINDEX, LINSERT, LLEN, LPOP, LPUSH, LPUSHX, LRANGE, LREM, LSET, LTRIM, RPOP, RPOPLPUSH, RPUSH, RPUSHX |
| **Pub/Sub**       | PSUBSCRIBE, PUNSUBSCRIBE, SUBSCRIBE, UNSUBSCRIBE, PUBLISH, PUBSUB |
| **Scripting**     | In future |
| **Server**        | In future |
| **Sets**          | SADD, SCARD, SDIFF, SDIFFSTORE, SINTER, SINTERSTORE, SISMEMBER, SMEMBERS, SMOVE, SPOP, SRANDMEMBER, SREM, SUNION, SUNIONSTORE | SSCAN
| **Sorted Sets**   | In future |
| **String**        | APPEND, BITCOUNT, BITOP, BITPOS, DECR, DECRBY, GET, GETBIT, GETRANGE, GETSET, INCR, INCRBY, INCRBYFLOAT, MGET, MSET, MSETNX, PSETEX, SET, SETBIT, SETEX, SETNX, SETRANGE, STRLEN | BITFIELD
| **Transactions**  | In futrue |

**Notice:** For the unimplemented commands, you can use the Do method, `func (r *Redis) Do(cmd string, args ...interface{}) (reply interface{}, err error)`, to implement it. Welcome to give me a pull request.
