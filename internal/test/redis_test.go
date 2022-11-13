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
