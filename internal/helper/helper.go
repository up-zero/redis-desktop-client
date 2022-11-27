package helper

import (
	"encoding/json"
	"errors"
	"gitee.com/up-zero/redis-desktop-client/internal/define"
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
