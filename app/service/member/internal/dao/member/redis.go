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
	_memberLockKey     = "mlock:%d"
	// key fb_mid/100000  offset => mid%100000
	// bit value 1 mean unfaved; bit value 0 mean faved
	// 为了避免单个 BITSET 过大或者热点，需要使用 region sharding
	_favedBit = "fb_%d"
	_bucket   = 100000
	_defaultExpire = 600
)

func keyMember(mid int64) string {
	return fmt.Sprintf(_memberKey, mid)
}

func keyMemberLock(mid int64) string {
	return fmt.Sprintf(_memberLockKey, mid)
}

func favedBitKey(mid int64) string {
	return fmt.Sprintf(_favedBit, mid/_bucket)
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

func (dao *Dao) cacheDelMember(ctx context.Context, id int64) (err error) {
	conn := dao.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("DEL", keyMember(id))
	return
}

func (d *Dao) cacheFavedBit(ctx context.Context, mid int64) (err error) {
	key := favedBitKey( mid)
	offset := mid % _bucket
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SETBIT", key, offset, 0); err != nil {
		log.Error("conn.DO(SETBIT) key(%s) offset(%d) err(%v)", key, offset, err)
	}
	return
}

func (d *Dao) cacheUnFavedBit(ctx context.Context, mid int64) (err error) {
	key := favedBitKey(mid)
	offset := mid % _bucket
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SETBIT", key, offset, 1); err != nil {
		log.Error("conn.DO(SETBIT) key(%s) offset(%d) err(%v)", key, offset, err)
	}
	return
}

func (d *Dao) getMemberLock(ctx context.Context, mid int64) (lock *redis.RedisLock) {
	key := keyMemberLock(mid)

	return redis.NewMutex(d.redis,key,30)
}