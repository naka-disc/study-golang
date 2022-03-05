package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 取得先URL
	url := "https://qiita.com/"

	// URLにGETリクエスト送信して、結果を取得
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// 結果のステータスコードが200 OKじゃないならログ出して終わり
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// HTMLをgoqueryに読み込ませる
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// articleタグを抽出
	selectList := doc.Find("article")
	// その中身を展開していく
	selectList.Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find(".css-32d82q").Attr("href")
		title := s.Find(".css-1wsd998").Text()
		fmt.Println(title)
		fmt.Println(href)
	})
}
