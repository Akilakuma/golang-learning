package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 一個在DB叫做User的table結構
type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	PasswordHash string
}

func main() {
	// 建立連線
	// [root] 帳號，
	// [qwe123] 密碼
	// [tcp(127.0.0.1:3306)] ip和port而且一定要外刮一層tcp
	// [Match] DB的名稱
	db, err := gorm.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/Match?charset=utf8&parseTime=True&loc=Local")
	// 選擇要不要打開debug mode 超實用
	// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	db.LogMode(true)

	// 連不連得到，這個超級重要!!
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	// Orm，跟db對起來，gorm用DB必先宣告
	db.AutoMigrate(&User{})

	// 試寫入一筆
	u1 := User{Username: "shen", Email: "abc@hotmail.com", PasswordHash: "werqw123"}
	db.Create(&u1)

	// 讀取剛剛寫入的那筆資料
	var user User
	db.Find(&user)
	// 印出select 結果(正確用法)
	fmt.Print(user)

	// 這樣會印出一串記憶體位置，我不知道這是幹嘛，使用觀念好像跟以前有點不一樣，如上golang是丟了user進去，又將user拿來讀取
	result := db.Find(&user)
	fmt.Print(result)
}
