package service

import (
	"context"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"gitee.com/up-zero/redis-desktop-client/internal/helper"
	"github.com/go-redis/redis/v8"
)

func ZSetValueDelete(req *define.ZSetValueRequest) error {
	conn, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})
	err = rdb.ZRem(context.Background(), req.Key, req.Member).Err()
	return err
}

func ZSetValueCreate(req *define.ZSetValueRequest) error {
	conn, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})
	err = rdb.ZAdd(context.Background(), req.Key, &redis.Z{
		Score:  req.Score,
		Member: req.Member,
	}).Err()
	return err
}
