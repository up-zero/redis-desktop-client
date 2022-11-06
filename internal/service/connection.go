package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// ConnectionList 连接列表
func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}

// ConnectionCreate 创建连接
func ConnectionCreate(conn *define.Connection) error {
	if conn.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	// 参数默认值处理
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		// 配置文件的内容初始化
		conf.Connections = []*define.Connection{conn}
		data, _ = json.Marshal(conf)
		// 写入配置内容
		os.MkdirAll(nowPath, 0666)
		ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	}
	json.Unmarshal(data, conf)
	conf.Connections = append(conf.Connections, conn)
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}
