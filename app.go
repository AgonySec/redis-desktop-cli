package main

import (
	"context"
	"redis-desktop-client/internal/define"
	"redis-desktop-client/internal/service"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
//
//	func (a *App) Greet(p Person) string {
//		return fmt.Sprintf("Hello %s (Age: %d)!", p.Name, p.Age)
//	}

// ConnectionList 获取连接列表
func (a *App) ConnectionList() interface{} {
	conn, err := service.ConnectionList()
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "SUCCESS",
		"data": conn,
	}
}

// DbList 获取数据库列表
func (a *App) DbList(identity string) interface{} {
	dbs, err := service.DbList(identity)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "SUCCESS",
		"data": dbs,
	}
}

// KeyList 键列表
func (a *App) KeyList(req *define.KeyListRequest) interface{} {
	if req.ConnIdentity == "" {
		return map[string]interface{}{
			"code": -1,
			"msg":  "连接的唯一标识不能为空",
		}
	}
	data, err := service.KeyList(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"data": data,
	}
}

// GetKeyValue 键值对查询
func (a *App) GetKeyValue(req *define.KeyValueRequest) interface{} {
	if req.Key == "" || req.ConnIdentity == "" {
		return map[string]interface{}{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	data, err := service.GetKeyValue(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"data": data,
	}
}

// ConnectionCreate 新建连接
func (a *App) ConnectionCreate(connection *define.Connection) interface{} {
	err := service.ConnectionCreate(connection)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "新建成功",
	}
}

// ConnectionEdit 编辑连接
func (a *App) ConnectionEdit(connection *define.Connection) interface{} {
	err := service.ConnectionEdit(connection)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "修改成功",
	}
}

// ConnectionDelete 删除连接
func (a *App) ConnectionDelete(identity string) interface{} {
	err := service.ConnectionDelete(identity)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "删除成功",
	}
}
