package models

import "github.com/go-redis/redis/v8"

type User struct {
	Username       string
	ChannelHandler *redis.PubSub
	StopListen     chan struct{}
	Listen         bool
	MessageChannel chan redis.Message
}
