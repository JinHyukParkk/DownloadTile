package downloads

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func DownloadTile(site string, lv string, x_start string, y_start string, ln string, vworldTile string) {
	if site == "naver" {
		ConstructNaverUrl(lv, x_start, y_start, ln)
	} else if site == "daum" {
		ConstructDaumUrl(lv, x_start, y_start, ln)
	} else if site == "vworld" {
		if vworldTile == "2d" {
			Construct2dVWorldUrl(lv, x_start, y_start, ln)
		} else {
			Construct3dVWorldUrl(lv, x_start, y_start, ln)
		}

	}
}
func Construct2dVWorldUrl(lv string, x_start string, y_start string, ln string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	for i := -2; i <= 3; i++ {
		for j := -2; j <= 2; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "http://xdworld.vworld.kr:8080/2d/Satellite/201710/" + lv + "/" + x_str + "/" + y_str + ".jpeg"
			fileName := y_str + "_" + x_str
			fileDir := "./tileData/vworld/2D" + ln + "/"
			MakeJPG(url, fileName, fileDir)
		}
	}
}
func Construct3dVWorldUrl(lv string, x_start string, y_start string, ln string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	for i := -2; i <= 3; i++ {
		for j := -2; j <= 2; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "http://xdworld.vworld.kr:8080/XDServer/3DData?Version=2.0.0.0&Request=GetLayer&Layer=tile_mo_HD&Level=" + lv + "&IDX=" + x_str + "&IDY=" + y_str + "&Key=81EC01D7-0327-3868-B85D-67E737396E44"
			fileName := y_str + "_" + x_str
			fileDir := "./tileData/vworld/3D" + ln + "/"
			MakeJPG(url, fileName, fileDir)
		}
	}
}
func ConstructNaverUrl(lv string, x_start string, y_start string, ln string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	for i := -2; i <= 3; i++ {
		for j := -2; j <= 2; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "https://simg.pstatic.net/onetile/get/184/0/1/" + lv + "/" + x_str + "/" + y_str + "/bl_st_bg"
			fileName := y_str + "_" + x_str
			fileDir := "./tileData/naver/" + ln + "/"
			MakeJPG(url, fileName, fileDir)
		}
	}
}
func ConstructDaumUrl(lv string, x_start string, y_start string, ln string) {
	x, _ := strconv.Atoi(x_start)
	y, _ := strconv.Atoi(y_start)
	num := 0
	for i := -2; i <= 3; i++ {
		for j := -2; j <= 2; j++ {
			x_str := strconv.Itoa(x + i)
			y_str := strconv.Itoa(y + j)
			url := "http://map" + strconv.Itoa(num%4) + ".daumcdn.net/map_skyview/L" + lv + "/" + y_str + "/" + x_str + ".jpg?v=160114"
			fileName := y_str + "_" + x_str
			fileDir := "./tileData/daum/" + ln + "/"
			MakeJPG(url, fileName, fileDir)
			num++
		}
	}
}

func MakeJPG(url string, fN string, fileDir string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// os.RemoveAll("./result/*")

	fileName := fileDir + fN + ".jpg"
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
