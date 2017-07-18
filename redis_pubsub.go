package redis

import (
	"github.com/garyburd/redigo/redis"
)

// PubSub returns a new PubSubConn object, which use a new Redis connections.
//
// This method is used to handle the subscribed datas.
//
// Notice: The type of the returned value is PubSubConn in the pakcage of
// "github.com/garyburd/redigo/redis".
//
// If you want to publish a message to a channel, please use r.Publish(c, m).
func (r *Redis) PubSub() (redis.PubSubConn, error) {
	conn, err := r.newRedisConn()
	if err != nil {
		return redis.PubSubConn{}, err
	}
	return redis.PubSubConn{Conn: conn.(*redisConn).Conn}, nil
}

// Publish executes the reids command PUBLISH, that's, publishs a message to
// a channel, then returns the number of the clients which receive this message.
//
// New in redis version 2.0.0.
func (r *Redis) Publish(channel, message string) (int64, error) {
	return r.doToInt("PUBLISH", channel, message)
}
