package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
func ConstructURL(lv string, x_start string, y_start string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "https://simg.pstatic.net/onetile/get/184/0/1/" + lv + "/" + x_str + "/" + y_str + "/bl_st_bg"
			fileName := lv + "_x" + x_str + "_y" + y_str
			MakeJPG(url, fileName)
		}
	}
}

func MakeJPG(url string, fN string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	CreateDirIfNotExist("result")
	os.RemoveAll("./result/*")

	fileName := "./result/" + fN + ".jpg"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	file.Close()
	// fmt.Println("Success")
}

func main() {
	t0 := time.Now()
	//argument
	level := os.Args[1]
	x_start := os.Args[2]
	y_start := os.Args[3]

	ConstructURL(level, x_start, y_start)

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
