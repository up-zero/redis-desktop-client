package helper

import (
	"encoding/json"
	"errors"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"os"
	"os/user"
)

func GetConnection(identity string) (*define.Connection, error) {
	conf := new(define.Config)
	nowPath := GetConfPath()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	for _, v := range conf.Connections {
		if v.Identity == identity {
			return v, nil
		}
	}
	return nil, errors.New("连接数据不存在")
}

// GetConfPath 获取配置文件路径
func GetConfPath() string {
	current, _ := user.Current()
	return current.HomeDir
}

// GetRedisClient 获取 Redis 客户端对象
//
// connectionIdentity 连接唯一标识
// db 选中的数据库
func GetRedisClient(connectionIdentity string, db int) (*redis.Client, error) {
	conn, err := GetConnection(connectionIdentity)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       db,
	})
	return rdb, err
}
