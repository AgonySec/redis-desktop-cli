package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"os"
	"redis-desktop-client/internal/define"
)

// ConnectionList 连接列表
func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return []*define.Connection{{
			Addr:     "127.0.0.1",
			Identity: "default",
			Name:     "default",
			Password: "",
			Port:     "6379",
			Username: "",
		}}, nil
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}

// ConnectionCreate 创建连接
func ConnectionCreate(connection *define.Connection) error {
	if connection.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	if connection.Port == "" {
		connection.Port = "6379"
	}
	if connection.Name == "" {
		connection.Name = connection.Addr
	}
	connection.Identity = uuid.New().String()
	conf := new(define.Config)

	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		//配置文件内容
		conf.Connections = []*define.Connection{connection}
		data, _ = json.Marshal(conf)
		// 写入配置文件
		os.MkdirAll(nowPath, 0666)
		os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
		return nil
	}
	json.Unmarshal(data, conf)
	conf.Connections = append(conf.Connections, connection)
	data, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}

func ConnectionEdit(connection *define.Connection) error {
	if connection.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	if connection.Port == "" {
		connection.Port = "6379"
	}
	if connection.Name == "" {
		connection.Name = connection.Addr
	}

	conf := new(define.Config)

	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, v := range conf.Connections {
		if v.Identity == connection.Identity {
			conf.Connections[i] = connection
			break
		}
	}
	data, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}

func ConnectionDelete(identity string) error {
	if identity == "" {
		return errors.New("连接标识不能为空")
	}
	conf := new(define.Config)

	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, v := range conf.Connections {
		if v.Identity == identity {
			conf.Connections = append(conf.Connections[:i], conf.Connections[i+1:]...)
			break
		}
	}
	data, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}
