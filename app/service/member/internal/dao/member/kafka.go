package member

import (
	"context"
	"encoding/json"
	"go-kartos-study/app/service/member/internal/model"
	"go-kartos-study/pkg/queue/kafka"
	"strconv"
)

func (dao *Dao) KafkaPushMember(ctx context.Context, m *model.Member) (err error) {
	b, _ := json.Marshal(m)
	err = dao.publisher.Publish(ctx, kafka.Event{
		Key:     strconv.Itoa(int(m.ID)),
		Payload: b,
	})
	return
}
