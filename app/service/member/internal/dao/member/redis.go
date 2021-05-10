package member

import (
	"context"
	"encoding/json"
	"fmt"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/cache/redis"
	"go-kartos-study/pkg/log"
)

const (
	_memberKey     = "m:%d"
	_defaultExpire = 600
)

func keyMember(mid int64) string {
	return fmt.Sprintf(_memberKey, mid)
}


func (dao *Dao) cacheGetMember(ctx context.Context, id int64) (m *model.Member, err error) {
	var (
		item []byte
		key = keyMember(id)
		conn = dao.redis.Get(ctx)
	)
	defer conn.Close()
	if item, err = redis.Bytes(conn.Do("GET", key)); err != nil {
		if err == redis.ErrNil {
			err = nil
		}else{
			log.Error("conn.Do(GET %s) error(%v)", key, err)
		}
		return
	}
	if err = json.Unmarshal(item, &m); err != nil {
		log.Error("json.Unmarshal(%v) err(%v)", item, err)
	}
	return
}

func (dao *Dao) cacheSetMember(ctx context.Context, id int64, m *model.Member) (err error) {
	var (
		values []byte
		conn   = dao.redis.Get(ctx)
	)
	defer conn.Close()
	if values, err = json.Marshal(m); err != nil {
		log.Error("json.Marshal(%v) err(%v)", m, err)
		return
	}

	_, err = conn.Do("SET", keyMember(id), values, "EX", _defaultExpire)
	if err != nil {
		log.Error("SET err(%v)", err)
	}
	return
}

func (dao *Dao) cacheDelMember(c context.Context, id int64) (err error) {
	conn := dao.redis.Get(c)
	defer conn.Close()
	_, err = conn.Do("DEL", keyMember(id))
	return
}
