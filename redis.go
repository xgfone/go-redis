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
	// ErrInvalidArgs is returned when the arguments of the redis command is not right.
	ErrInvalidArgs = fmt.Errorf("The arguments is invalid.")

	// ErrInvalidResult is returned when the result is not right.
	ErrInvalidResult = fmt.Errorf("The result is invalid.")

	// ErrNotExist is returned when the key or result does not exist.
	ErrNotExist = fmt.Errorf("Not exist")
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

// NewRedis return a new Redis, which uses a connection pool in the underlying
// implementation.
//
// connURL is a URL to connect to the redis server. It should follow the draft
// IANA specification for the scheme (https://www.iana.org/assignments/uri-schemes/prov/redis).
// For example, "redis://password@127.0.0.1:6379/0".
//
// poolSize is the size of the connection pool.
func NewRedis(connURL string, poolSize int) *Redis {
	r := &Redis{
		connURL:  connURL,
		poolSize: poolSize,
		ctx:      context.TODO(),
	}
	r.rp = pools.NewResourcePool(r.newRedisConn, poolSize, poolSize, 0)
	return r
}

func (r *Redis) newRedisConn() (pools.Resource, error) {
	c, err := redis.DialURL(r.connURL, redis.DialConnectTimeout(RedisConnTimeout))
	return &redisConn{c}, err
}

// Close closes the Redis connection.
func (r *Redis) Close() {
	r.rp.Close()
}

func (r *Redis) getConn() (*redisConn, error) {
	c, err := r.rp.Get(r.ctx)
	if err != nil {
		return nil, err
	}
	return c.(*redisConn), nil
}

func (r *Redis) putConn(c *redisConn) {
	r.rp.Put(c)
}

// Do executes the Redis command, then returns the result.
func (r *Redis) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	conn, err := r.getConn()
	if err != nil {
		return nil, err
	}

	defer func() {
		if conn.Err() != nil {
			r.putConn(nil)
		} else {
			r.putConn(conn)
		}
	}()

	return conn.Do(cmd, args...)
}
