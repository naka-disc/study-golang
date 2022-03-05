# struct
- structの挙動。


## 作成したときの手順
```sh
go mod init naka-disc/struct
mkdir cmd
touch cmd/main.go
```

## ビルド&実行
```sh
go mod tidy
go build -o . cmd/*.go
./main
```
