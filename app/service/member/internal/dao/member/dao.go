package member

import (
	"context"
	"fmt"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/cache/redis"
	"go-kartos-study/pkg/database/sql"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/queue/kafka"
	"go-kartos-study/pkg/stat/prom"
	"go-kartos-study/pkg/sync/pipeline/fanout"
	"golang.org/x/sync/singleflight"
)

// Dao is redis dao.
type Dao struct {
	c                 *conf.Config
	db                *sql.DB
	redis             *redis.Pool
	cacheSingleFlight *singleflight.Group
	cache             *fanout.Fanout
	publisher         kafka.Publisher
}

// New new a dao.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:                 c,
		db:                sql.NewMySQL(c.Mysql),
		redis:             redis.NewPool(c.Redis),
		cacheSingleFlight: &singleflight.Group{},
		cache:             fanout.New("cache", fanout.Worker(1), fanout.Buffer(1024)),
		publisher:         kafka.NewPublisher(c.KafkaPublish),
	}
	return d
}

// Close close dao.
func (dao *Dao) Close() {
	if dao.db != nil {
		dao.db.Close()
	}
}

// Ping ping cpdb
func (dao *Dao) Ping(c context.Context) (err error) {
	return dao.db.Ping(c)
}

func (dao *Dao) InitMember(ctx context.Context, arg *model.Member) (err error) {
	return dao.dbInitMember(ctx, arg)
}

func (dao *Dao) AddMember(ctx context.Context, arg *model.Member) (id int64, err error) {
	mutex := dao.getMemberLock(ctx,1)
	fmt.Println(mutex.Lock(ctx))
	fmt.Println(mutex.Lock(ctx))
	fmt.Println(mutex.Unlock(ctx))
	fmt.Println(mutex.Lock(ctx))
	defer fmt.Println(mutex.Unlock(ctx))
	return dao.dbAddMember(ctx, arg)
}

func (dao *Dao) BatchAddMember(ctx context.Context, args []*model.Member) (affectRow int64, err error) {
	return dao.dbBatchAddMember(ctx, args)
}

func (dao *Dao) GetMemberByID(ctx context.Context, id int64) (res *model.Member, err error) {
	addCache := true
	res, err = dao.cacheGetMember(ctx, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.Id == -1 {
			res = nil
		}
	}()

	if res != nil {
		prom.CacheHit.Incr("member")
		return
	}
	// XXX m == nil RowNotFound
	var rr interface{}
	// 压测即可发现 同时会有很多访问，但是db只有一条，防止击穿
	log.Info("cacheSingleFlight.Do test")
	rr, err, _ = dao.cacheSingleFlight.Do(fmt.Sprintf("m_%d", id), func() (r interface{}, e error) {
		prom.CacheMiss.Incr("member test")
		log.Info("dbGetMemberByID")
		r, e = dao.dbGetMemberByID(ctx, id)
		return
	})

	if err != nil {
		return
	}
	res = rr.(*model.Member)
	miss := res
	if miss == nil {
		miss = &model.Member{Id: -1}
	}
	if !addCache {
		return
	}
	// fanout模式
	dao.cache.Do(ctx, func(ctx context.Context) {
		_ = dao.cacheSetMember(ctx, id, miss)
	})
	return
}

func (dao *Dao) GetMemberByPhone(ctx context.Context, phone string) (m *model.Member, err error) {
	return dao.dbGetMemberByPhone(ctx, phone)
}

func (dao *Dao) GetMemberMaxAge(ctx context.Context) (age int64, err error) {
	return dao.dbGetMemberMaxAge(ctx)
}

func (dao *Dao) GetMemberSumAge(ctx context.Context) (age int64, err error) {
	return dao.dbGetMemberSumAge(ctx)
}

func (dao *Dao) CountMember(ctx context.Context) (count int64, err error) {
	return dao.dbCountMember(ctx)
}

func (dao *Dao) ListMember(ctx context.Context) (res []*model.Member, err error) {
	return dao.dbListMember(ctx)
}

func (dao *Dao) HasMemberByID(ctx context.Context, id int64) (has bool, err error) {
	return dao.dbHasMemberByID(ctx, id)
}

func (dao *Dao) QueryMemberByName(ctx context.Context, name string) (res []*model.Member, err error) {
	return dao.dbQueryMemberByName(ctx, name)
}

func (dao *Dao) QueryMemberByIDs(ctx context.Context, ids []int64) (res map[int64]*model.Member, err error) {
	return dao.dbQueryMemberByIDs(ctx, ids)
}

func (dao *Dao) UpdateMember(ctx context.Context, member *model.Member) (err error) {
	return dao.dbUpdateMember(ctx, member)
}

func (dao *Dao) UpdateMemberAttr(ctx context.Context, id int64, attr int32) (err error) {
	return dao.dbUpdateMemberAttr(ctx, id, attr)
}

func (dao *Dao) SetMember(ctx context.Context, arg *model.Member) (err error) {
	return dao.dbSetMember(ctx, arg)
}

func (dao *Dao) SortMember(ctx context.Context, args model.ArgMemberSort) (err error) {
	return dao.dbSortMember(ctx, args)
}

func (dao *Dao) DeleteMember(ctx context.Context, id int64) (err error) {
	return dao.dbDeleteMember(ctx, id)
}
