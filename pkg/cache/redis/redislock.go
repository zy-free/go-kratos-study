package redis

import (
	"context"
	"go-kartos-study/pkg/log"
	"math/rand"
	"strconv"
	"time"
)

const (
	letters     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lockCommand = `
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])`
	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
	randomLen       = 16
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
)

type RedisLock struct {
	store  *Pool
	expire time.Duration
	key    string
	id     string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// id需要一致时才能释放锁，所以使用姿势不能全局一个锁，这样id都一样，会造成其他goroutine也能释放锁，所以最好具体使用时new
func NewMutex(store *Pool, key string, expire time.Duration) *RedisLock {
	return &RedisLock{
		store:  store,
		key:    key,
		id:     randomStr(randomLen),
		expire: expire,
	}
}

func (rl *RedisLock) Lock(ctx context.Context) (bool, error) {
	script := NewScript(1, lockCommand)
	conn := rl.store.Get(ctx)
	defer conn.Close()
	resp, err := script.Do(conn, rl.key, rl.id, strconv.Itoa(int(rl.expire.Milliseconds())+tolerance))
	if err == ErrNil {
		return false, nil
	} else if err != nil {
		log.Error("Error on acquiring lock for %s, %s", rl.key, err.Error())
		return false, err
	} else if resp == nil {
		return false, nil
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	} else {
		log.Error("Unknown reply when acquiring lock for %s: %v", rl.key, resp)
		return false, nil
	}
}

func (rl *RedisLock) Unlock(ctx context.Context) (bool, error) {
	conn := rl.store.Get(ctx)
	defer conn.Close()
	script := NewScript(1, delCommand)
	resp, err := script.Do(conn, rl.key, rl.id)
	if err != nil {
		return false, err
	}

	if reply, ok := resp.(int64); !ok {
		return false, nil
	} else {
		return reply == 1, nil
	}
}

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
