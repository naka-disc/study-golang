# private-standard-layout
- 個人でGolangProject作るときのディレクトリ構成。


## 作成したときの手順
```sh
go mod init naka-disc/private-standard-layout
mkdir cmd
mkdir -p internal/app/hello
mkdir internal/app/sample
```

## ビルド&実行
```sh
go mod tidy
go build -o . cmd/*.go
./main
```

## 参照／利用サイト
- [ローカル（ソース非公開）プロジェクトでGo言語の推奨ディレクトリ構成（golang-standards/project-layout）を使う - 自分用めも](https://belhb.hateblo.jp/entry/2020/10/31/201803)
- [Goにはディレクトリ構成のスタンダードがあるらしい。 - Qiita](https://qiita.com/sueken/items/87093e5941bfbc09bea8)
- [golang-standards/project-layout: Standard Go Project Layout](https://github.com/golang-standards/project-layout)
