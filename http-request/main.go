package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NiconicoAPIのレスポンスデータ
// DataとMetaがあり、Dataの方は複数ある
type NicoNico struct {
	Data []NicoNicoData `json:"data"`
	Meta NicoNicoMeta   `json:"meta"`
}

// 実際のデータ部
type NicoNicoData struct {
	ContentId   string `json:"contentId"`
	Title       string `json:"title"`
	ViewCounter int    `json:"viewCounter"`
}

// レスポンスのメタデータ
type NicoNicoMeta struct {
	Id         string `json:"id"`
	TotalCount int    `json:"totalCount"`
	Status     int    `json:"status"`
}

func main() {
	url := "https://api.search.nicovideo.jp/api/v2/snapshot/video/contents/search?q=%E5%88%9D%E9%9F%B3%E3%83%9F%E3%82%AF&targets=title&fields=contentId,title,viewCounter&filters[viewCounter][gte]=10000&_sort=-viewCounter&_offset=0&_limit=3&_context=apiguide"

	nicoObj, err := executeAPI(url)
	if !err {
		return
	}

	fmt.Println(nicoObj)
}

// APIを実行
func executeAPI(url string) (NicoNico, bool) {
	var nicoObj NicoNico

	// GETリクエスト送信
	resp, err := http.Get(url)
	// エラーがあった場合
	if err != nil {
		return nicoObj, false
	}

	// deferはreturnの時の実行される式
	// https://go.dev/tour/flowcontrol/12
	// レスポンスのクローズ処理は最後にやる
	defer resp.Body.Close()

	// レスポンスボディを読み込み
	body, err := ioutil.ReadAll(resp.Body)
	// エラーがあった場合
	if err != nil {
		return nicoObj, false
	}
	// stringBody := string(body)

	// json.Unmarshalの第二引数は、JSONに対応する構造体のポインタ
	if err := json.Unmarshal(body, &nicoObj); err != nil {
		return nicoObj, false
	}

	return nicoObj, true
}
