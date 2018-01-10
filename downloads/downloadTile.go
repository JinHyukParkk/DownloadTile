package downloads

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
func ConstructNaverUrl(lv string, x_start string, y_start string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "https://simg.pstatic.net/onetile/get/184/0/1/" + lv + "/" + x_str + "/" + y_str + "/bl_st_bg"
			fileName := "Naver_" + lv + "_x" + x_str + "_y" + y_str
			MakeJPG(url, fileName)
		}
	}
}
func ConstructDaumUrl(lv string, x_start string, y_start string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	num := 0
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "http://map" + strconv.Itoa(num%4) + ".daumcdn.net/map_skyview/L" + lv + "/" + y_str + "/" + x_str + ".jpg?v=160114"
			fileName := "Daum_" + lv + "_x" + x_str + "_y" + y_str
			fmt.Println(url)
			MakeJPG(url, fileName)
			num++
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
