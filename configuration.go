package go_db

import (
	"errors"
	"os"
	//"fmt"
)


func ConnString() (connStr string, err error){
	if os.Getenv("DB_USER") == "" {
		return connStr, errors.New("Empty user.")
	}
	//student:student@/musicapp?charset=utf8&parseTime=True
	connStr += os.Getenv("DB_USER") + ":"


	if os.Getenv("DB_PASSWORD") == "" {
		return connStr, errors.New("Empty password.")
	}
	connStr += os.Getenv("DB_PASSWORD") + "@/"


	if os.Getenv("DB_HOST") == "" {
		return connStr, errors.New("Empty hostname.")
	}
	//connStr += " host=" + os.Getenv("DB_HOST")
	if os.Getenv("DB_PORT") == "" {
		return connStr, errors.New("Empty port number.")
	}
	//connStr += " port=" + os.Getenv("DB_PORT")

	if os.Getenv("DB_NAME") == "" {
		return connStr, errors.New("Empty database name.")
	}
	connStr += os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True"
	//connStr += " sslmode=" + os.Getenv("SSL_MODE")

	return connStr, err
}
