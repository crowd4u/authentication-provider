package database

import "time"

type RedisHandler interface {
	Set(string) (string, error)
	SetWithKey(string, string) error
	Get(string) (string, error)
	Update(string, string) error
	SetWithTime(string, string, time.Duration) error
	Count(string) (int, error)
	AddValue(string, time.Duration) error
	DeclValue(string) (int, error)
	SADD(string, string, time.Duration) (int64, error)
	SREM(string, string) error
	SMEMBERS(string) ([]string, error)
}
