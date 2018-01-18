package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/JinHyukParkk/DownloadTile/downloads"
)

func main() {
	t0 := time.Now()
	//argument

	site := flag.String("site", "naver", "Site Name")
	level := flag.String("lv", "1", "Tile Level")
	firstD := flag.String("x", "1", "Tile x Location")
	secondD := flag.String("y", "1", "Tile y Location")
	flag.Parse()
	if *site == "naver" {
		downloads.ConstructNaverUrl(*level, *firstD, *secondD)
	} else if *site == "daum" {
		downloads.ConstructDaumUrl(*level, *firstD, *secondD)
	}

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
