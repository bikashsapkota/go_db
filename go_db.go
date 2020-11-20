package go_db

import (
	"fmt"
	"io"
	"log"

	"github.com/bikashsapkota/go_db/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbConn *gorm.DB

type PgDatabase struct {
}

//initialize the  database
func (pg *PgDatabase) InitDb(file io.Writer, maxIdleConns, maxOpenConns int) {

	connString, err := ConnString()
	if err != nil {
		log.Fatalln(err)
	}

	DbConn, err = gorm.Open("mysql", connString)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	} else {
		log.Println("Successful DB Connection!!")
	}

	DbConn.SetLogger(log.New(file, "[godown] ", 0))
	DbConn.LogMode(true)
	DbConn.SingularTable(true)
	DbConn.DB().SetMaxIdleConns(maxIdleConns)
	DbConn.DB().SetMaxOpenConns(maxOpenConns)
}

//close postgres database connection
func (pg *PgDatabase) Close() {
	DbConn.Close()
}

//function to get IdentifiedMusic
func (pg *PgDatabase) GetAllIdentifiedOfUser(user_id int) (*[]model.IdentifiedMusics, error) {
	identifiedMusic := []model.IdentifiedMusics{}

	if err := DbConn.Find(&identifiedMusic, "dj_id = ?", user_id).Error; err != nil {
		log.Printf("DB Error: Get Identified music (%s)", err)
		return nil, err
	}

	return &identifiedMusic, nil
}

func (pg *PgDatabase) GetTodayIdentifiedOfUser(user_id int) (*[]model.IdentifiedMusics, error) {
	identifiedMusic := []model.IdentifiedMusics{}

	if err := DbConn.Raw("SELECT * FROM identified_musics WHERE datediff(CURRENT_TIMESTAMP, played_timestamp)=0 AND identified_musics.dj_id = ?", user_id).Scan(&identifiedMusic).Error; err != nil {
		log.Printf("DB Error: Get Identified music (%s)", err)
		return nil, err
	}

	return &identifiedMusic, nil
}

func (pg *PgDatabase) GetTodayIdentifiedOfUserCount(user_id int) (*int, error) {

	type Result struct {
		Count int `json:"count"`
	}

	var result Result

	if err := DbConn.Raw("SELECT COUNT(played_timestamp) as count FROM identified_musics WHERE datediff(CURRENT_TIMESTAMP, played_timestamp)=0 AND identified_musics.dj_id = ?", user_id).Scan(&result).Error; err != nil {
		log.Printf("DB Error: Get Today IdentifiedOf User Count (%s)", err)
		return nil, err
	}

	return &result.Count, nil
}

func (pg *PgDatabase) AddAllIdentifiedMusic(user_id int, song_id int, played_timestamp string) (bool, error) {

	if err := DbConn.Exec("insert into identified_music_alls (music_id, dj_id, played_timestamp) values (?,?,?)", song_id, user_id, played_timestamp).Error; err != nil {
		log.Printf("DB Error: Get Today IdentifiedOf User Count (%s)", err)
		fmt.Println("Error in Godown")
		return false, err
	}

	return true, nil
}

func (pg *PgDatabase) MarkMessageAsConsumed(id int) (bool, error) {
	if err := DbConn.Exec("update kafka_messages set consumed = 1 where id = ?", id).Error; err != nil {
		log.Printf("DB Error: markMessageAsConsumed (%s)", err)
	}

	return true, nil

}

func (pg *PgDatabase) GetUnConsumedMessages() ([]model.KafkaMessages, error) {
	result := []model.KafkaMessages{}

	if err := DbConn.Raw("SELECT message->>'$.topic' as topic, message, id FROM kafka_messages where consumed = 0 order by id asc").Scan(&result).Error; err != nil {
		log.Printf("DB Error: GetUnConsumedMessages (%s)", err)
		return nil, err
	}

	return result, nil
}
