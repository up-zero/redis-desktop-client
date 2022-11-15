package service

import (
	"changeme/internal/define"
	"changeme/internal/helper"
	"context"
	"github.com/go-redis/redis/v8"
)

func KeyList(req *define.KeyListRequest) ([]string, error) {
	conn, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})
	var count int64 = 100
	if req.Keyword != "" {
		count = 2000
	}
	res, _, err := rdb.Scan(context.Background(), 0, "*"+req.Keyword+"*", count).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}
