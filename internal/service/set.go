package service

import (
	"context"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"gitee.com/up-zero/redis-desktop-client/internal/helper"
	"github.com/go-redis/redis/v8"
)

func SetValueDelete(req *define.SetValueRequest) error {
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
	err = rdb.SRem(context.Background(), req.Key, req.Value).Err()
	return err
}

func SetValueCreate(req *define.SetValueRequest) error {
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
	err = rdb.SAdd(context.Background(), req.Key, req.Value).Err()
	return err
}
