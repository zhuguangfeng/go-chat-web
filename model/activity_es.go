package model

import (
	_ "embed"
)

const ActivityIndexName = "activity"

var (
	//go:embed activity_index.json
	ActivityIndex string
)

type ActivityEs struct {
	ID                  int64    `json:"id"`
	SponsorID           int64    `json:"sponsorId"`
	GroupID             int64    `json:"groupId"`
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
