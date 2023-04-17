package resource

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
)

var TableNoteData *gorm.DB

func InitDB() {
	// 用实际的用户名、密码、主机名（或 IP 地址）和数据库名称替换下面的占位符
	username := "root"
	password := "mysecretpassword"
	host := "172.17.0.3"
	database := "note_data"

	// 构建 MySQL 连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	// 连接到 MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: %v", err)
	}

	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatalf("Error connecting to the database: %v", err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("Error pinging the database: %v", err)
	// }

	TableNoteData = db.Table("note-data")

	fmt.Println("Successfully connected to MySQL!")
}
