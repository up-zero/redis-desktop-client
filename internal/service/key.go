package service

import (
	"context"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"gitee.com/up-zero/redis-desktop-client/internal/helper"
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

func GetKeyValue(req *define.KeyValueRequest) (*define.KeyValueReply, error) {
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
	_type, err := rdb.Type(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	// _type => string
	v, err := rdb.Get(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	return &define.KeyValueReply{
		Value: v,
		TTL:   ttl,
		Type:  _type,
	}, nil
}

func DeleteKeyValue(req *define.KeyValueRequest) error {
	conn, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})
	_, err = rdb.Del(context.Background(), req.Key).Result()
	return err
}
