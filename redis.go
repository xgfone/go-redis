package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/xgfone/go-tools/pools"
)

// RedisConnTimeout is the timeout to connect to the Redis Server.
var RedisConnTimeout = 5 * time.Second

var (
	// ErrInvalidArgs is panicked when the arguments of the redis command is not right.
	ErrInvalidArgs = fmt.Errorf("The arguments is invalid.")
)

type redisConn struct {
	redis.Conn
}

func (r redisConn) Close() {
	r.Conn.Close()
}

// Redis is a proxy to operate the redis command.
type Redis struct {
	connURL  string
	poolSize int

	ctx context.Context
	rp  *pools.ResourcePool
}

// NewRedis return a new Redis.
func NewRedis(connURL string, poolSize int) *Redis {
	rp := pools.NewResourcePool(func() (pools.Resource, error) {
		r, err := redis.DialURL(connURL, redis.DialConnectTimeout(RedisConnTimeout))
		return &redisConn{r}, err
	}, poolSize, poolSize, 0)

	return &Redis{
		connURL:  connURL,
		poolSize: poolSize,

		ctx: context.TODO(),
		rp:  rp,
	}
}

// Close closes the Redis connection.
func (r *Redis) Close() {
	r.rp.Close()
}

func (r *Redis) getConn() *redisConn {
	c, err := r.rp.Get(r.ctx)
	if err != nil {
		panic(err)
	}
	return c.(*redisConn)
}

func (r *Redis) putConn(c *redisConn) {
	r.rp.Put(c)
}

// Do executes the Redis command, then returns the result.
func (r *Redis) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	conn := r.getConn()

	defer func() {
		if _err := recover(); _err != nil {
			err = fmt.Errorf("%v", _err)
			r.putConn(nil)
			conn.Close()
		} else {
			r.putConn(conn)
		}
	}()

	return conn.Do(cmd, args...)
}
