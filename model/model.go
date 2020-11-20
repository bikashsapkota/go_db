package model

import (
	"time"
)


type IdentifiedMusics struct {
  Id int
  MusicId int
  DjId int
  PlayedTimestamp time.Time
  CreatedAt time.Time
  UpdatedAt time.Time
  videos string
}

type MusicCampaign struct {
  Id int
  UserId int
  CampaignName string
  Email string
  SpinRate int
  CampaignBalance int
  CreatedAt time.Time
  UpdatedAt time.Time
}

type Notification struct {
  Id int
  UserId int
  Message string
  Href string
  Seen bool
  Type string
  CreatedAt time.Time
  UpdatedAt time.Time
}

type KafkaMessages  struct {
  ID int
  Topic string
  Message string
}
