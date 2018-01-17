# Naver map, Daum map에서 위성 사진 tile을 받아오는 소스입니다.
웹 디버깅 툴인 피들러를 이용하여 url 정보를 알아내었고, 테스트 용도로 작성하였습니다. 단순하게 편리함과 golang 공부를 위해 만들었습니다.
현재 예제는 입력 좌표를 중심으로 타일 30개를 받아오는 소스입니다. 필요에 따라 수정해서 사용하시면 될 것 같습니다.

## 1. Naver map에서 받아오기
### 1.1 사용법
```
# go get github.com/JinHyukParkk/DownloadTile
```
* 함수
##### ConstructNaverUrl(level string, x string, y string)
```
import github.com/JinHyukParkk/DownloadTile
함수 : downloads.ConstructNaverUrl([Level],[x축],[y축)
```
* 결과
##### 실행하는 디렉터리에 result란 디렉터리가 생성되고, 그 안에 타일 30개가 저장됩니다.
### 1.2 실행 방법
```
# git clone https://github.com/JinHyukParkk/DownloadTile.git
# make
# ./DownloadTile naver [Level] [x좌표] [y좌표]
예시) ./DownloadTile naver 13 6728 5993   // 13레벨 타일맵 북한산 좌표
```

### 1.3 지도 레벨과 좌표 확인
naver 지도 tile : [https://navermaps.github.io/maps.js/docs/tutorial-1-maptypes-tilecheck.example.html](https://navermaps.github.io/maps.js/docs/tutorial-1-maptypes-tilecheck.example.html)
 - 링크와 타일맵의 좌표가 맞지 않아서 피들러로 좌표를 확인해야 할 것 같습니다.!

## 2. Daum map에서 받아오기
### 2.1 사용법
```
# go get github.com/JinHyukParkk/DownloadTile
```
* 함수
##### ConstructDaumUrl(level string, x string, y string)
```
import github.com/JinHyukParkk/DownloadTile
함수 : downloads.ConstructDaumUrl([Level],[x축],[y축)
```
* 결과
##### 실행하는 디렉터리에 result란 디렉터리가 생성되고, 그 안에 타일 30개가 저장됩니다.

### 2.2 실행 방법 사용법
```
# git clone https://github.com/JinHyukParkk/DownloadTile.git
# make
# ./DownloadTile daum [Level] [x좌표] [y좌표]
```
### 2.3 지도 레벨과 좌표 확인
daum 지도 tile : [http://apis.map.daum.net/web/sample/getTile/](http://apis.map.daum.net/web/sample/getTile/)

개념 참조 : [http://d2.naver.com/helloworld/1174](http://d2.naver.com/helloworld/1174)
