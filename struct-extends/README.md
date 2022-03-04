# struct-extends
- structの継承。


## 作成したときの手順
```sh
go mod init naka-disc/struct-extends
mkdir cmd
touch cmd/main.go
```

## ビルド&実行
```sh
go mod tidy
go build -o . cmd/*.go
./main
```

## 参照／利用サイト
- [Inheritance in GO using struct (Embedding) - Welcome To Golang By Example](https://golangbyexample.com/inheritance-go-struct/)
