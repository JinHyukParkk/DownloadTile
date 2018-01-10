package main

import (
	"fmt"
	"os"
	"time"

	"github.com/JinHyukParkk/DownloadTile/downloads"
)

func main() {
	t0 := time.Now()
	//argument
	site := os.Args[1]
	level := os.Args[2]
	firstD := os.Args[3]
	secondD := os.Args[4]

	if site == "naver" {
		downloads.ConstructNaverUrl(level, firstD, secondD)
	} else if site == "daum" {
		downloads.ConstructDaumUrl(level, firstD, secondD)
	}

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
