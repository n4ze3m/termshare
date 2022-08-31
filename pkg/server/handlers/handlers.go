package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	ctx *context.Context
	rdb *redis.Client
}

func NewHandler(ctx *context.Context, rdb *redis.Client) *Handler {
	return &Handler{
		ctx: ctx,
		rdb: rdb,
	}
}
