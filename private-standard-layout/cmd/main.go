package main

import (
	"internal/hello"
	"internal/sample"
	// go.modに記載したrequire書くと、そのファイルを読み込み可能
)

func main() {
	// 別ファイルのfuncを呼び出し
	sample.Say()
	hello.Hello()
}
