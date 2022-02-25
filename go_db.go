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
	log.Println(connString)

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

func (pg *PgDatabase) SaveNotification(notification model.Notifications) (bool, error) {
	if err := DbConn.Create(&notification).Error; err != nil {
		log.Printf("DB Error: SaveNotification (%s)", err)
		return false, err
	}

	return true, nil
}

func (pg *PgDatabase) GetAllJobs() ([]model.Job, error) {
	log.Println("GetAllJobsExecuting")
	result := []model.Job{}
	if err := DbConn.Raw("SELECT * FROM job").Scan(&result).Error; err != nil {
		log.Printf("DB Error: GetAllJobs (%s)", err)
		return nil, err
	}

	return result, nil
}

func (pg *PgDatabase) GetKeyingCount() (int, error) {
	type Result struct {
		Count int `json:"count"`
	}

	var result Result

	if err := DbConn.Raw("select count(id) as count from kafka_messages where message->'$.match' = 'no_match'").Scan(&result).Error; err != nil {
		log.Printf("DB Error: Get Keying Count (%s)", err)
		return 0, err
	}

	return result.Count, nil
}

func (pg *PgDatabase) GetKeyerCount() (string, error) {
	type Result struct {
		DjName string `json:"dj_name"`
		Count  int    `json:"count"`
	}

	result := []Result{}

	if err := DbConn.Raw("select message->>'$.payload.dj_name' as dj_name, count(*) as count from kafka_messages where message->'$.match' = 'no_match' group by dj_name;").Scan(&result).Error; err != nil {
		log.Printf("DB Error: Get Keying Count (%s)", err)
		return "", err
	}

	var resp = ""
	for _, res := range result {
		resp += fmt.Sprintf("DjName: %s\tCount: %d\n", res.DjName, res.Count)
	}

	return resp, nil
}

func (pg *PgDatabase) GetDjUserIdWithGenre(genre_id string) (*[]int, error) {
	log.Println("GetDjIdWithGenre")
	var result []int
	if err := DbConn.Raw("SELECT djs.user_id FROM dj__musics, djs where music_type = ? and djs.id = dj__musics.dj_id", genre_id).Scan(&result).Error; err != nil {
		log.Printf("DB Error: GetAllJobs (%s)", err)
		return nil, err
	}

	return &result, nil
}
