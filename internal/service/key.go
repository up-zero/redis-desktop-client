package service

import (
	"context"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"gitee.com/up-zero/redis-desktop-client/internal/helper"
	"github.com/go-redis/redis/v8"
)

// KeyList 键列表
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
	var count = define.DefaultKeyLen
	if req.Keyword != "" {
		count = define.MaxKeyLen
	}
	res, _, err := rdb.Scan(context.Background(), 0, "*"+req.Keyword+"*", count).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetKeyValue 获取键值对
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
	var reply = &define.KeyValueReply{
		Type: _type,
	}
	if _type == "string" {
		v, err := rdb.Get(context.Background(), req.Key).Result()
		if err != nil {
			return nil, err
		}
		reply.Value = v
	} else if _type == "hash" {
		keys, _, err := rdb.HScan(context.Background(), req.Key, 0, "", define.MaxHashLen).Result()
		if err != nil {
			return nil, err
		}
		data := make([]*define.KeyValue, 0, len(keys)/2)
		for i := 0; i < len(keys); i += 2 {
			data = append(data, &define.KeyValue{
				Key:   keys[i],
				Value: keys[i+1],
			})
		}
		reply.Value = data
	} else if _type == "list" {
		list, err := rdb.LRange(context.Background(), req.Key, 0, define.MaxListLen-1).Result()
		if err != nil {
			return nil, err
		}
		data := make([]*define.KeyValue, 0, len(list))
		for i := 0; i < len(list); i++ {
			data = append(data, &define.KeyValue{
				Value: list[i],
			})
		}
		reply.Value = data
	}

	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	reply.TTL = ttl
	return reply, nil
}

// DeleteKeyValue 删除键值对
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

// CreateKeyValue 创建键值对
func CreateKeyValue(req *define.CreateKeyValueRequest) error {
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
	if req.Type == "string" {
		err = rdb.Set(context.Background(), req.Key, "value", -1).Err()
	} else if req.Type == "hash" {
		err = rdb.HSet(context.Background(), req.Key, map[string]string{"key": "value"}).Err()
	} else if req.Type == "list" {
		err = rdb.RPush(context.Background(), req.Key, "value").Err()
	}
	return err
}

// UpdateKeyValue 更新键值对数据
func UpdateKeyValue(req *define.UpdateKeyValueRequest) error {
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
	// type ==> string
	err = rdb.Set(context.Background(), req.Key, req.Value, req.TTL).Err()
	return err
}
