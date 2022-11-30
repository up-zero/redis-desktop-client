package service

import (
	"context"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"gitee.com/up-zero/redis-desktop-client/internal/helper"
)

// HashFieldDelete hash 字段删除
func HashFieldDelete(req *define.HashFieldDeleteRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	err = rdb.HDel(context.Background(), req.Key, req.Field...).Err()
	return err
}

// HashAddOrUpdateField hash 字段新增、修改
func HashAddOrUpdateField(req *define.HashAddOrUpdateFieldRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	err = rdb.HSet(context.Background(), req.Key, map[string]interface{}{req.Field: req.Value}).Err()
	return err
}
