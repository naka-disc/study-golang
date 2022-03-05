# scraping
- スクレイピング。


## 作成したときの手順
```sh
go mod init naka-disc/scraping
touch README.md
mkdir cmd
touch cmd/main.go

go get github.com/PuerkitoBio/goquery
```

## ビルド&実行
```sh
go mod tidy
go build -o . cmd/*.go
./main
# go build -o . cmd/*.go && ./main
```

## 依存関係パッケージ情報
- [github.com/PuerkitoBio/goquery](https://pkg.go.dev/github.com/PuerkitoBio/goquery)
  - BSD-3-Clause
