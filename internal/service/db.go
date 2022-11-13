package service

import (
	"changeme/internal/define"
	"changeme/internal/helper"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
)

func DbList(identity string) ([]*define.DbItem, error) {
	if identity == "" {
		return nil, errors.New("连接唯一标识不能为空")
	}
	conn, err := helper.GetConnection(identity)
	if err != nil {
		return nil, err
	}
	// 获取数据库的连接对象
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
	})
	// info 获取数据库键的个数
	keySpace, err := rdb.Info(context.Background(), "keyspace").Result()
	if err != nil {
		return nil, err
	}
	// keyspace 数据格式
	// # Keyspace
	// db0:keys=2,avg_ttl...
	//
	m := make(map[string]int)
	v := strings.Split(keySpace, "\n")
	for i := 1; i < len(v)-1; i++ {
		databases := strings.Split(v[i], ":")
		if len(databases) < 2 {
			continue
		}
		vv := strings.Split(databases[1], ",")
		if len(vv) < 1 {
			continue
		}
		keyNumber := strings.Split(vv[0], "=")
		if len(keyNumber) < 2 {
			continue
		}
		num, err := strconv.Atoi(keyNumber[1])
		if err != nil {
			continue
		}
		m[databases[0]] = num
	}
	// config get 获取数据库的个数
	databasesRes, err := rdb.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		return nil, err
	}
	if len(databasesRes) < 2 {
		return nil, errors.New("连接数据异常")
	}
	dbNum, err := strconv.Atoi(databasesRes[1].(string))
	if err != nil {
		return nil, err
	}
	data := make([]*define.DbItem, dbNum)
	for i := 0; i < dbNum; i++ {
		item := &define.DbItem{
			Key: "db" + strconv.Itoa(i),
		}
		if n, ok := m["db"+strconv.Itoa(i)]; ok {
			item.Number = n
		}
		data = append(data, item)
	}
	return data, nil
}
