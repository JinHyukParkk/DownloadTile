package createDir

import (
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

func CreateDir(site string, lv string, x_start string, y_start string) {
  dirPath := "tileData"
  CreateDirIfNotExist(dirPath)
  dirPath += "/" + site
  CreateDirIfNotExist(dirPath)
  dirPath += "/" + lv
  CreateDirIfNotExist(dirPath)
  y, _ := strconv.Atoi(y_start)
  for j := -2; j <= 2; j++ {
    y_str := strconv.Itoa(y + j)
		ydirPath := dirPath+"/"+y_str
    CreateDirIfNotExist(ydirPath)
  }
}
