package test

import (
	"context"
	"fmt"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
	"gitee.com/up-zero/redis-desktop-client/internal/helper"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
	"testing"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
var ctx = context.Background()

func TestInfo(t *testing.T) {
	info, err := rdb.Info(context.Background()).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info)
	res, err := rdb.Info(context.Background(), "keyspace").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	dbs := strings.Split(res, "\n")
	m := make(map[string]int)
	for i := 1; i < len(dbs)-1; i++ {
		v := strings.Split(dbs[i], ":")
		if len(v) < 2 {
			continue
		}
		vv := strings.Split(v[1], ",")
		if len(vv) < 3 {
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
		m[v[0]] = num
	}
	fmt.Println(m)
}

func TestConfigGet(t *testing.T) {
	res, err := rdb.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestType(t *testing.T) {
	res, err := rdb.Type(context.Background(), "hash").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestHash(t *testing.T) {
	intCmd := rdb.HSet(ctx, "hash", map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"})
	fmt.Println(intCmd)
	getOne := rdb.HGet(ctx, "hash", "k1")
	fmt.Println(getOne.Val())
	all := rdb.HGetAll(ctx, "hash")
	for key, value := range all.Val() {
		fmt.Println("key --> ", key, " value --> ", value)
	}
}

func TestHashGet(t *testing.T) {
	resLen, err := rdb.HLen(ctx, "hash").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resLen)
	res, _, err := rdb.HScan(ctx, "hash", 0, "", 200).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(res), res)
}

func TestHashDelete(t *testing.T) {
	err := rdb.HDel(ctx, "hash2", "key").Err()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Success Delete")
}

func TestLRange(t *testing.T) {
	res, err := rdb.LRange(context.Background(), "list", 0, define.MaxListLen-1).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestZRevRange(t *testing.T) {
	rz, err := rdb.ZRevRangeWithScores(context.Background(), "zset2", 0, 199).Result()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range rz {
		fmt.Println(v)
	}
}

func TestGetRedisClient(t *testing.T) {
	rdb, err := helper.GetRedisClient("cc69d2e0-80e4-40ed-96a3-8706403b4c7c", 0)
	if err != nil {
		t.Fatal(err)
	}
	res, err := rdb.Info(context.Background()).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
