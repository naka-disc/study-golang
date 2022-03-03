package main

import (
	"fmt"

	// このSQLパッケージは汎用IFで、別途ドライバーと組み合わせる必要がある
	// https://pkg.go.dev/database/sql
	"database/sql"
	// ドライバーとして使うためのインポート ここで実際に使うわけじゃないので _
	_ "github.com/mattn/go-sqlite3"
)

type Member struct {
	Id   int
	Name string
}

// SELECT
func sel() ([]Member, bool) {
	// 構造体リストを生成しておく
	// SELECTデータをこの中に詰めるため
	var memberList []Member

	// SQLite3と接続
	db, err := sql.Open("sqlite3", "./sample.sqlite3") //接続開始（example.sqlに保存する）
	if err != nil {
		fmt.Println("Database Connection Error.")
		return memberList, false
	}
	defer db.Close()

	// SELECTクエリを実行
	const selSQL = "SELECT * FROM member"
	rows, err := db.Query(selSQL)
	if err != nil {
		fmt.Println("Query Error.")
		return memberList, false
	}
	defer rows.Close()

	// クエリの結果を読み込んでいくため、カーソルを次に動かしていく
	for rows.Next() {
		// SELECT結果のデータを指定ポインタに割り当て、だと思う
		// SELECTの個数と引数は合わせる必要がある
		var member Member
		err := rows.Scan(&member.Id, &member.Name)
		if err != nil {
			fmt.Println("Rows Scan Error.")
		}

		// スキャンに成功しても失敗しても追加
		memberList = append(memberList, member)
	}

	return memberList, true
}

func main() {
	sel()
}
