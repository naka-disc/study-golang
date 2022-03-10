package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// DB接続
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// 接続失敗時
		panic("failed to connect database")
	}

	// マイグレーション
	// 指定した構造体でテーブル作成・変更してくれる
	db.AutoMigrate(&Product{})

	// レコード作成
	db.Create(&Product{Code: "D42", Price: 100})

	// テーブルからデータを引っ張ってくる
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// データ更新
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// データ削除
	db.Delete(&product, 1)
}

// 構造体
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
