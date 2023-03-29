package zcache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var ctx = context.Background()

type URedisClient struct {
	*redis.Client
	TimeOut time.Duration
}

func NewRedisClient(addr string, password string, db int) *URedisClient {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	return &URedisClient{
		Client: conn,
	}
}

func (redis *URedisClient) GetTableValue(table string, key string) (interface{}, error) {
	var bytes strings.Builder
	bytes.WriteString(table)
	bytes.WriteString(CUnderline)
	bytes.WriteString(key)
	return redis.HGetAll(ctx, bytes.String()).Result()
}
func (redis *URedisClient) SetTableValue(table string, key string, value ...interface{}) (bool, error) {
	var bytes strings.Builder
	bytes.WriteString(table)
	bytes.WriteString(CUnderline)
	bytes.WriteString(key)
	return redis.HMSet(ctx, bytes.String(), value).Result()
}
func (redis *URedisClient) SetTableField(table string, key string, value ...interface{}) (int64, error) {
	var bytes strings.Builder
	bytes.WriteString(table)
	bytes.WriteString(CUnderline)
	bytes.WriteString(key)
	return redis.HSet(ctx, bytes.String(), value).Result()
}
func (redis *URedisClient) GetTableField(table string, key string, field string) (string, error) {
	var bytes strings.Builder
	bytes.WriteString(table)
	bytes.WriteString(CUnderline)
	bytes.WriteString(key)
	return redis.HGet(ctx, bytes.String(), field).Result()
}
func (redis *URedisClient) GetValue(key string) (string, error) {
	return redis.Get(ctx, key).Result()
}

func (redis *URedisClient) SetValue(key string, value interface{}) (string, error) {
	return redis.Set(ctx, key, value, redis.TimeOut).Result()
}
func (redis *URedisClient) GetField(key string, field string) (string, error) {
	return redis.HGet(ctx, key, field).Result()
}

func (redis *URedisClient) SetField(key string, value ...interface{}) (int64, error) {
	return redis.HSet(ctx, key, value, redis.TimeOut).Result()
}
