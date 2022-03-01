package go_db

import (
	// "github.com/bikashsapkota/godown/model"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"

	//"time"
	"io"
	"log"

	"github.com/bikashsapkota/go_db/model"
)

type MockDatabase struct {
}

func (pg *MockDatabase) InitDb(io.Writer, int, int) {

}

func (pg *MockDatabase) GetAllIdentifiedOfUser(user_id int) (*[]model.IdentifiedMusics, error) {
	identifiedMusic := []model.IdentifiedMusics{}
	return &identifiedMusic, nil
}

//Save Functions
func (pg *MockDatabase) GetTodayIdentifiedOfUser(user_id int) (*[]model.IdentifiedMusics, error) {
	identifiedMusic := []model.IdentifiedMusics{}
	return &identifiedMusic, nil
}

func (pg *MockDatabase) GetTodayIdentifiedOfUserCount(user_id int) (*int, error) {
	identifiedMusic := 0
	return &identifiedMusic, nil
}

func (pg *MockDatabase) AddAllIdentifiedMusic(user_id int, song_id int, played_timestamp string) (bool, error) {
	status := false
	if err := DbConn.Raw("insert into identified_music_alls (music_id, dj_id, played_timestamp) values (?,?,?)", song_id, user_id, played_timestamp).Scan(&status).Error; err != nil {
		log.Printf("DB Error: Get Today IdentifiedOf User Count (%s)", err)
		return false, err
	}

	return true, nil
}

func (pg *MockDatabase) MarkMessageAsConsumed(id int) (bool, error) {
	return true, nil
}

func (pg *MockDatabase) GetUnConsumedMessages() ([]model.KafkaMessages, error) {
	result := []model.KafkaMessages{}
	return result, nil
}

func (pg *MockDatabase) SaveNotification(notification model.Notifications) (bool, error) {
	return true, nil
}

func (pg *MockDatabase) GetAllJobs() ([]model.Job, error) {
	return nil, nil
}

func (pg *MockDatabase) GetKeyerCount() (string, error) {
	return "", nil
}

func (pg *MockDatabase) GetKeyingCount() (int, error) {
	return 0, nil
}

func (pg *MockDatabase) GetDjUserIdWithGenre(string) (*[]int, error) {
	return nil, nil
}

func (pg *MockDatabase) GetWithdrawRequestCount() (int, error) {
	return 0, nil
}
