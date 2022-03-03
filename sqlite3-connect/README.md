# sqlite3-connect

## 作成したときの手順
```sh
go mod init naka-disc/sqlite3-connect
```

## ビルド
```sh
sqlite3 sample.sqlite3 < init.sql

go mod tidy
go build
./sqlite3-connect
```

## 依存関係パッケージ情報
- [github.com/mattn/go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)
  - ライセンスがMIT
