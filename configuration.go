package go_db

import (
	"errors"
	"os"
	"fmt"
)


func ConnString() (connStr string, err error){
	
	if os.Getenv("DB_USER") == "" {
		return connStr, errors.New("Empty user.")
	}
	


	if os.Getenv("DB_PASSWORD") == "" {
		return connStr, errors.New("Empty password.")
	}
	


	if os.Getenv("DB_HOST") == "" {
		return connStr, errors.New("Empty hostname.")
	}
	
	if os.Getenv("DB_PORT") == "" {
		return connStr, errors.New("Empty port number.")
	}
	

	if os.Getenv("DB_NAME") == "" {
		return connStr, errors.New("Empty database name.")
	}
	

	connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))
	
	return connStr, err
}
