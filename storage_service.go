package go_db

import (
	"io"
	"github.com/bikashsapkota/go_db/model"
)

type DatabaseService interface {
	InitDb(io.Writer, int, int)
	GetAllIdentifiedOfUser(int) (*[]model.IdentifiedMusics, error)
	GetTodayIdentifiedOfUser(int) (*[]model.IdentifiedMusics, error)
	GetTodayIdentifiedOfUserCount(int) (*int, error)
	AddAllIdentifiedMusic( int, int, string) (bool, error)
	MarkMessageAsConsumed(int) (bool, error)
	GetUnConsumedMessages() ([]model.KafkaMessages, error)
	SaveNotification(model.Notifications) (bool, error)
	GetAllJobs() ([]model.Job, error)
	GetKeyingCount()(int, error)
	GetKeyerCount() (string, error)
}
