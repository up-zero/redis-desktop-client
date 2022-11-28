package define

import "time"

var ConfigName = "redis-client.conf"

type Connection struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	// DefaultKeyLen 键列表的默认查询长度
	DefaultKeyLen int64 = 100
	// MaxKeyLen 键列表的最大查询长度
	MaxKeyLen int64 = 2000
	// MaxHashLen hash列表的最大查询长度
	MaxHashLen int64 = 200
	// MaxListLen list列表的最大查询长度
	MaxListLen int64 = 200
)

type Config struct {
	Connections []*Connection `json:"connections"`
}

type DbItem struct {
	Key    string `json:"key"`    // db0 , db1
	Number int    `json:"number"` // 个数
}

type KeyListRequest struct {
	ConnIdentity string `json:"conn_identity"`
	Db           int    `json:"db"`
	Keyword      string `json:"keyword"`
}

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type KeyValueRequest struct {
	ConnIdentity string `json:"conn_identity"`
	Db           int    `json:"db"`
	Key          string `json:"key"`
}

type KeyValueReply struct {
	Value interface{}   `json:"value"`
	TTL   time.Duration `json:"ttl"`
	Type  string        `json:"type"`
}

type CreateKeyValueRequest struct {
	KeyValueRequest
	Type string `json:"type"`
}

type UpdateKeyValueRequest struct {
	KeyValueRequest
	TTL   time.Duration `json:"ttl"`
	Value string        `json:"value"`
}

type HashFieldDeleteRequest struct {
	KeyValueRequest
	Field []string `json:"field"`
}

type HashAddOrUpdateFieldRequest struct {
	KeyValueRequest
	Field string `json:"field"`
	Value string `json:"value"`
}
