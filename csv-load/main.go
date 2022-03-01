package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	// ファイルを展開
	file, err := os.Open("resources/c03.csv")

	// ファイル展開に失敗したらログ出力して終了
	if err != nil {
		log.Fatal(err)
	}

	// CSVのリーダーを作成
	// TODO: Shift-JISのみの対応になっているので要修正（文字コード検知など）
	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	for {
		// リーダーから一行ずつ取得していく
		record, err := reader.Read()
		// errに終端記号が格納されていれば読み込み終了
		if err == io.EOF {
			break
		}
		// その他、何かしらのエラーがあればログ出力して終了
		if err != nil {
			log.Fatal(err)
		}
		// 取得した行を出力
		fmt.Println(record)
	}

}
