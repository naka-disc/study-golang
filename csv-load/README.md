# golang-sample-csvload

## 作成したときの手順
```sh
go mod init naka-disc/csv-load
```

## ビルド
```sh
go mod tidy
go build ./main.go
# go install
```

## 依存関係パッケージ情報
- [golang.org/x/text](https://pkg.go.dev/golang.org/x/text)
  - ライセンスがBSD-3-Clause

## 備考
- CSVは下記から取得。

[国勢調査 時系列データ CSV形式による主要時系列データ | ファイル | 統計データを探す | 政府統計の総合窓口](https://www.e-stat.go.jp/stat-search/files?page=1&layout=datalist&toukei=00200521&tstat=000001011777&cycle=0&tclass1=000001094741&tclass2val=0)
