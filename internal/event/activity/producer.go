package activity

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/zhuguangfeng/go-chat/internal/domain"
)

const TopicSyncActivity = "sync_dynamic_event"

type Producer interface {
	ProducerSyncActivityEvent(ctx context.Context, activity ActivityEvent) error
}

type SaramaProducer struct {
	producer sarama.SyncProducer
}

func NewProducer(producer sarama.SyncProducer) Producer {
	return &SaramaProducer{
		producer: producer,
	}
}

// 发送同步活动消息
func (s *SaramaProducer) ProducerSyncActivityEvent(ctx context.Context, activity ActivityEvent) error {
	val, err := json.Marshal(&activity)
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: TopicSyncActivity,
		Value: sarama.StringEncoder(val),
	})
	return err
}

type ActivityEvent struct {
	ID                  int64    `json:"id"`
	SponsorID           int64    `json:"sponsorId"`
	Title               string   `json:"title"`
	Desc                string   `json:"desc"`
	Media               []string `json:"media"`
	AgeRestrict         uint     `json:"ageRestrict"`
	GenderRestrict      uint     `json:"genderRestrict"`
	CostRestrict        uint     `json:"CostRestrict"`
	Visibility          uint     `json:"visibility"`
	MaxPeopleNumber     int64    `json:"maxPeopleNumber"`
	CurrentPeopleNumber int64    `json:"CurrentPeopleNumber"`
	Address             string   `json:"address"`
	Category            uint     `json:"category"`
	StartTime           uint     `json:"startTime"`
	DeadlineTime        uint     `json:"deadlineTime"`
	Status              uint     `json:"status"`
	CreatedTime         uint     `json:"createdTime"`
	UpdatedTime         uint     `json:"updatedTime"`
}

func ToEvent(activity domain.Activity) ActivityEvent {
	return ActivityEvent{
		ID:                  activity.ID,
		SponsorID:           activity.Sponsor.ID,
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
		Category:            activity.Category.Uint(),
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status.Uint(),
		CreatedTime:         activity.CreatedTime,
		UpdatedTime:         activity.UpdatedTime,
	}
}
