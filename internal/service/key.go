package service

import (
	"context"
	"redis-desktop-client/internal/define"
	"redis-desktop-client/internal/helper"
)

// KeyList 键列表
func KeyList(req *define.KeyListRequest) ([]string, error) {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	var count = define.DefaultKeyLen
	if req.Keyword != "" {
		count = define.MaxKeyLen
	}
	res, _, err := rdb.Scan(context.Background(), 0, "*"+req.Keyword+"*", count).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetKeyValue 获取键值对
func GetKeyValue(req *define.KeyValueRequest) (*define.KeyValueReply, error) {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	_type, err := rdb.Type(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	var reply = &define.KeyValueReply{
		Type: _type,
	}
	if _type == "string" {
		v, err := rdb.Get(context.Background(), req.Key).Result()
		if err != nil {
			return nil, err
		}
		reply.Value = v
	} else if _type == "hash" {
		keys, _, err := rdb.HScan(context.Background(), req.Key, 0, "", define.MaxHashLen).Result()
		if err != nil {
			return nil, err
		}
		data := make([]*define.KeyValue, 0, len(keys)/2)
		for i := 0; i < len(keys); i += 2 {
			data = append(data, &define.KeyValue{
				Key:   keys[i],
				Value: keys[i+1],
			})
		}
		reply.Value = data
	} else if _type == "list" {
		list, err := rdb.LRange(context.Background(), req.Key, 0, define.MaxListLen-1).Result()
		if err != nil {
			return nil, err
		}
		data := make([]*define.KeyValue, 0, len(list))
		for i := 0; i < len(list); i++ {
			data = append(data, &define.KeyValue{
				Value: list[i],
			})
		}
		reply.Value = data
	} else if _type == "set" {
		sets, _, err := rdb.SScan(context.Background(), req.Key, 0, "", define.MaxSetLen).Result()
		if err != nil {
			return nil, err
		}
		data := make([]*define.KeyValue, 0, len(sets))
		for i := 0; i < len(sets); i++ {
			data = append(data, &define.KeyValue{
				Value: sets[i],
			})
		}
		reply.Value = data
	} else if _type == "zset" {
		z, err := rdb.ZRevRangeWithScores(context.Background(), req.Key, 0, define.MaxZSetLen-1).Result()
		if err != nil {
			return nil, err
		}
		reply.Value = z
	}

	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	reply.TTL = ttl
	return reply, nil
}
