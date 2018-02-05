# Naver, Daum, Vworld Map에서 위성 Tile을 받아오는 소스입니다.
웹 디버깅 툴인 피들러를 이용하여 url 정보를 알아내었고, 테스트 용도로 작성하였습니다. 단순하게 편리함과 golang 공부를 위해 만들었습니다.
현재 예제는 입력 좌표를 중심으로 타일 30개를 받아오는 소스입니다. 필요에 따라 수정해서 사용하시면 될 것 같습니다.


## 1. Flag 설명
* -site : 'naver', 'daum', 'vworld'   .. 받아올 TileMap의 해당 사이트
* -lv : 해당 Tile의 level
* -x : 해당 Tile을 중심으로 x축 좌표
* -y : 해당 Tile을 중심으로 y축 좌표
* -ln : 해당 Tile들의 위치 이름 .. ex) 독도, 설악산, 북한산, 여의도 ..
* -type : site가 __vwolrd__ 라면 2d or 3d 중 어느 지도에서 가져올 것인지 선택

## 2. 실행 방법  
### 해당 소스 URL
[https://github.com/JinHyukParkk/DownloadTile.git](https://github.com/JinHyukParkk/DownloadTile.git)
### 2.1. 리눅스, Mac OS X  환경
  1. 'DownloadTile' 이라는 실행 파일 다운
  2. 터미널 실행
      * ./DownloadTile -site [site] -lv [Level] -x [x축] -y [y축] -ln [위치 이름]
  3. result 디렉터리가 생성되고, 그 안에 타일 30개를 받아옴

### 2.2. Window 환경
  1. 'DownloadTile.exe' 이라는 실행 파일 다운
  2. CMD 관리자 모드로 실행
      * DownloadTile.exe -site [site] -lv [Level] -x [x축] -y [y축] -ln [위치 이름]
  3. result 디렉터리가 생성되고, 그 안에 타일 30개를 받아옴

## 3. Golang 라이브러리로 사용
```
# go get github.com/JinHyukParkk/DownloadTile
```
* Naver 함수
#### ConstructNaverUrl(level string, x string, y string)
```
import github.com/JinHyukParkk/DownloadTile
함수 : downloads.ConstructNaverUrl([Level],[x축],[y축],[위치])
```
* Daum 함수
#### ConstructDaumUrl(level string, x string, y string)
```
import github.com/JinHyukParkk/DownloadTile
함수 : downloads.ConstructDaumUrl([Level],[x축],[y축],[위치])
```

* Vworld 함수
#### Construct2DVWorldUrl(level string, x string, y string)
#### Construct3DVWorldUrl(level string, x string, y string)
```
import github.com/JinHyukParkk/DownloadTile
함수 : downloads.Construct2DVworldUrl([Level],[x축],[y축],[위치])
함수 : downloads.Construct3DVworldUrl([Level],[x축],[y축],[위치])
```

## 4. git clone 하여 실행 방법
```
# git clone https://github.com/JinHyukParkk/DownloadTile.git
# make
# ./DownloadTile -site naver -lv [Level] -x [x좌표] -y [y좌표] -ln [위치 이름]
예시) ./DownloadTile naver 13 6728 5993   // 13레벨 타일맵 북한산 좌표
```

## 5. 지도 레벨과 좌표 확인
Naver Map tile : [https://navermaps.github.io/maps.js/docs/tutorial-1-maptypes-tilecheck.example.html](https://navermaps.github.io/maps.js/docs/tutorial-1-maptypes-tilecheck.example.html)
 - 링크와 타일맵의 좌표가 맞지 않아서 피들러로 좌표를 확인해야 할 것 같습니다.!

Daum Map tile : [http://apis.map.daum.net/web/sample/getTile/](http://apis.map.daum.net/web/sample/getTile/)

Vworld Map tile : 피들러를 이용하여 찾아야 합니다.

## 결과
'tileData'라는 directory 안에 [site]/[위치 이름]/"y좌표_x좌표.jpg"로 해당 Tile 30개가 저장됩니다.

### 지도서비스 개념 참조
개념 참조 : [http://d2.naver.com/helloworld/1174](http://d2.naver.com/helloworld/1174)
