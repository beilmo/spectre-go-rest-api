package storage

// import (
// 	"github.com/beilmo/spectre-go-rest-api/adapter/gateway"
// 	"github.com/jinzhu/gorm"
// )

// type Database gorm.DB

// var db *Database

// func Connect() *Database {
// 	var err error

// 	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/hoge")

// 	if err != nil {
// 		panic(err)
// 	}
// 	if !db.HasTable(&gateway.Session{}) {
// 		if err := db.Table("users").CreateTable(&gateway.Session{}).Error; err != nil {
// 			panic(err)
// 		}
// 	}
// 	return db
// }

// func CloseConn() {
// 	db.Close()
// }
