package test

import (
	"context"
	"fmt"
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
