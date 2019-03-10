package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

var (
	fp  *os.File
	err error
)

func main() {
	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error!! Please input correct file path.")
			return
		}
		defer fp.Close()
	}
	prefix := "http://whisper.wisdom-guild.net/search.php?name="
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		card_name_jp := scanner.Text()
		url := prefix + card_name_jp
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		card_name := doc.Find("div.card a").First().Text()
		fmt.Println(card_name)
	}
}
