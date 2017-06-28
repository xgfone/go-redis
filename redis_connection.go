package redis

import (
	"errors"
)

// Auth executes the redis command AUTH.
//
// Return an error if failed, or nil. Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Auth(passwd string) error {
	if v := r.doToString("AUTH", passwd); v != "OK" {
		return errors.New(v)
	}
	return nil
}

// Echo executes the redis command ECHO.
//
// Panic if an error occurs.
//
// New in redis version 1.0.0.
func (r *Redis) Echo(msg string) string {
	return r.doToString("ECHO", msg)
}