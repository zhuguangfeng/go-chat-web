package activity

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/zhuguangfeng/go-chat/internal/domain"
	"github.com/zhuguangfeng/go-chat/internal/repository"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"github.com/zhuguangfeng/go-chat/pkg/saramax"
	"time"
)

const SyncActivityConsumerGroupName = "sync_activity"

type ActivityConsumer struct {
	client       sarama.Client
	l            logger.Logger
	activityRepo repository.ActivityRepository
}

func NewActivityConsumer(client sarama.Client, l logger.Logger, activityRepo repository.ActivityRepository) *ActivityConsumer {
	return &ActivityConsumer{
		client:       client,
		l:            l,
		activityRepo: activityRepo,
	}
}

func (a *ActivityConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient(SyncActivityConsumerGroupName, a.client)
	if err != nil {
		return err
	}
	go func() {
		err := cg.Consume(context.Background(), []string{TopicSyncActivity}, saramax.NewHandler[ActivityEvent](a.l, a.Consumer))
		if err != nil {
			a.l.Error("推出了消费 循环异常", logger.String("ConsumerGroupName", SyncActivityConsumerGroupName), logger.Error(err))
		}
	}()
	return err
}

func (a *ActivityConsumer) Consumer(sg *sarama.ConsumerMessage, activity ActivityEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return a.activityRepo.InputActivity(ctx, a.toDomain(activity))
}

func (a *ActivityConsumer) toDomain(activity ActivityEvent) domain.Activity {
	return domain.Activity{
		ID: activity.ID,
		Sponsor: domain.User{
			ID: activity.SponsorID,
		},
		Title:               activity.Title,
		Desc:                activity.Desc,
		Media:               activity.Media,
		AgeRestrict:         activity.AgeRestrict,
		GenderRestrict:      activity.GenderRestrict,
		CostRestrict:        activity.CostRestrict,
		Visibility:          activity.Visibility,
		MaxPeopleNumber:     activity.MaxPeopleNumber,
		CurrentPeopleNumber: activity.CurrentPeopleNumber,
		Address:             activity.Address,
		Category:            domain.ActivityCategory(activity.Category),
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              domain.ActivityStatus(activity.Status),
		CreatedTime:         activity.CreatedTime,
		UpdatedTime:         activity.UpdatedTime,
	}
}
