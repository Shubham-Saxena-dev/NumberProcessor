package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidKey = errors.New("key does not exist")
	ErrRedisGet   = errors.New("redisGet: Failed to get value for the key")
	ErrRedisDel   = errors.New("RedisDel: Failed to delete key")
	ErrDelAll     = errors.New("RedisFlushDB: Failed to flush cache with error")
	ErrRedisSet   = errors.New("redisSetWithLock: Failed to set value")
)

type CacheError struct {
	key     interface{}
	ErrType error
	cause   error
}

func newCacheError(errType error, key interface{}, cause error) *CacheError {
	return &CacheError{
		ErrType: errType,
		key:     key,
		cause:   cause,
	}
}

func (c *CacheError) Error() string {
	return fmt.Sprintf("%s: %v: %v", c.ErrType, c.key, c.cause)
}

func (c *CacheError) Unwrap() error {
	return c.cause
}

func ErrorInvalidKey(key interface{}, err error) *CacheError {
	return newCacheError(ErrInvalidKey, key, err)
}

func ErrorRedisGet(key interface{}, err error) *CacheError {
	return newCacheError(ErrRedisGet, key, err)
}

func ErrorRedisDel(key interface{}, err error) *CacheError {
	return newCacheError(ErrRedisDel, key, err)
}

func ErrorDelAll(err error) *CacheError {
	return newCacheError(ErrDelAll, nil, err)
}

func ErrorRedisSet(key interface{}, err error) *CacheError {
	return newCacheError(ErrRedisSet, key, err)
}
