package define

var ConfigName = "redis-client.conf"

type Connection struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

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
