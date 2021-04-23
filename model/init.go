package model

import (
	"github.com/lexkong/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	//config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	//	username,
	//	password,
	//	addr,
	//	name,
	//	true,
	//	//"Asia/Shanghai"),
	//	"Local")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	//setupDB(db)

	return db

}
